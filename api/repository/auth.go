package repository

import (
	"articleproject/api/model/request"
	"articleproject/api/model/response"
	"articleproject/api/redis_internal"
	"articleproject/constants"
	"articleproject/error"
	"articleproject/utils"
	rabbitmq_user "articleproject/utils/rabbitmq/user"
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	amqp "github.com/rabbitmq/amqp091-go"
)

type AuthRepository interface {
	UserRegistration(user request.User) error
	UserLogin(user request.User) (response.User, string, error)
	RefreshToken(string) (int64, bool, error)
}

type authRepository struct {
	pgx *pgx.Conn
	rdb *redis.Client
	amqp *amqp.Connection
}

func NewAuthRepo(pgx *pgx.Conn, rdb *redis.Client, amqp *amqp.Connection) AuthRepository {
	return authRepository{
		pgx: pgx,
		rdb: rdb,
		amqp: amqp,
	}
}

func (a authRepository) UserRegistration(user request.User) error {
	var userID int64
	err := a.pgx.QueryRow(context.Background(), `INSERT INTO users (name, bio, email, password, isadmin) VALUES ($1, $2, $3, $4, $5) RETURNING id`, user.Name, user.Bio, user.Email, user.Password, user.IsAdmin).Scan(&userID)
	if err != nil {
		pgErr, ok := err.(*pgconn.PgError)
        if ok && pgErr.Code == "23505" { 
			fmt.Println(pgErr, ok)
            return errorhandling.DuplicateEmailFound
        }
		return errorhandling.RegistrationFailedError
	}
	redis_internal.SaveUser(user, userID, a.rdb)
	rabbitmq_user.ProduceUserMail(a.amqp, user)
	return nil
}

func (a authRepository) UserLogin(user request.User) (response.User, string, error) {
	var dbUser response.User
	row := a.pgx.QueryRow(context.Background(), `SELECT id, name, bio, email, password, image, isadmin FROM users WHERE email = $1`, user.Email)
	err := row.Scan(&dbUser.ID, &dbUser.Name, &dbUser.Bio, &dbUser.Email, &dbUser.Password, &dbUser.Image, &dbUser.IsAdmin)

	if err == sql.ErrNoRows {
		return response.User{}, constants.EMPTY_STRING, errorhandling.NoUserFound
	}

	passwordMatched := utils.VerifyPassword(user.Password, dbUser.Password)
	if !passwordMatched {
		return response.User{}, "", errorhandling.PasswordNotMatch
	}

	refreshToken, err := utils.CreateAccessToken(time.Now().Add(time.Hour * 24 * 7), dbUser.ID, dbUser.IsAdmin)
	if err != nil {
		return response.User{}, constants.EMPTY_STRING, err
	}

	_, err = a.pgx.Exec(context.Background(), `INSERT INTO refreshtoken (userid, refreshtoken) VALUES ($1, $2)`, dbUser.ID, refreshToken)
	if err != nil {
		return response.User{}, constants.EMPTY_STRING, err
	}

	return dbUser, refreshToken, nil
	//isadmin false, id corrupt
}

func (a authRepository) RefreshToken(token string) (int64, bool, error) {
	row := a.pgx.QueryRow(context.Background(), `SELECT id, isadmin FROM users as u LEFT JOIN refreshtoken as r on u.id = r.userid WHERE r.refreshtoken = $1`, token)
	var isadmin bool
	var id int64
	err := row.Scan(&id, &isadmin)
	if err == sql.ErrNoRows {
		return 0, false, errorhandling.RefreshTokenNotFound
	}

	return id, isadmin, nil
}
