package controller

import (
	"articleproject/api/repository"
	"articleproject/api/service"
	"articleproject/config"
	"articleproject/db"
	"context"
	"log"
	"os"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5"
	amqp "github.com/rabbitmq/amqp091-go"
)

var conn *pgx.Conn
var rdb *redis.Client
var amqpConn *amqp.Connection
var r *chi.Mux
var authService service.AuthService
var articleService service.ArticleService
var topicService service.TopicService
var followerService service.FollowerService

type contextKey string

var (
	ContextKeyID = contextKey("id")
	ContextKeyToken = contextKey("token")
)

func init() {
	config.LoadEnv("../../.config/")
	conn, rdb, amqpConn, _ = db.DBConnection()
	r = chi.NewRouter()

	authRepository := repository.NewAuthRepo(conn, rdb, amqpConn)
	authService = service.NewAuthService(authRepository)

	articleRepository := repository.NewArticleRepo(conn)
	articleService = service.NewArticleService(articleRepository)

	topicRepository := repository.NewTopicRepo(conn)
	topicService = service.NewTopicService(topicRepository)

	followerRepository := repository.NewFollowerRepository(conn)
	followerService = service.NewFollowerService(followerRepository)
}

func TestMain(m *testing.M) {
	err := ClearMockData(conn)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return
	}

	tx, err = InsertMockData(tx)
	if err != nil {
		tx.Rollback(context.Background())
		return
	}
	tx.Commit(context.Background())

	// this is for running test of the controller
	//so from here it will go to actual function.
	exitVal := m.Run()

	err = ClearMockData(conn)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(exitVal)
}

func InsertMockData(tx pgx.Tx) (pgx.Tx, error) {
	_, err := tx.Exec(context.Background(), "INSERT INTO users (name, bio, email, password, isadmin) VALUES ('Niraj Darji', 'Software Engineer', 'nirajdarji@gmail.com', 'Niraj123$', false)")
	if err != nil {
		return tx, err
	}
	return tx, nil
}

func ClearMockData(dbConn *pgx.Conn) error {
	query := "DELETE FROM refreshtoken where userid in (select id from users where email = 'nirajdarji@gmail.com'); " +
	"DELETE FROM users WHERE email = 'nirajdarji@gmail.com'; " + "DELETE FROM users WHERE email = 'chiragmakwana1807@gmail.com';"

	_, err := dbConn.Exec(context.Background(), query)
	if err != nil {
		log.Print(err)
	}
	return nil
}

//generally in production you'll have separate conenction for database in which you will clear all data from table using delete from query.