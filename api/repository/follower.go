package repository

import (
	"articleproject/api/model/response"
	"context"

	"github.com/jackc/pgx/v5"
)

type FollowerRepository interface {
	Follow(int64, int64) error
	UnFollow(int64, int64) error
	FetchFollowers(int64) ([]response.Follower, error)
	FetchFollowings(int64) ([]response.Follower, error)
}

type followerRepository struct {
	pgx *pgx.Conn
}

func NewFollowerRepository(pgx *pgx.Conn) FollowerRepository {
	return followerRepository{
		pgx: pgx,
	}
}

func (f followerRepository) Follow(userid int64, followersid int64) error {
	_, err := f.pgx.Exec(context.Background(), `INSERT INTO followers (userid, followersid) VALUES ($1, $2)`, userid, followersid)
	if err != nil {
		return err
	}
	return nil
}

func (f followerRepository) UnFollow(userid int64, follwersid int64) error {
	_, err := f.pgx.Exec(context.Background(), `DELETE FROM followers WHERE userid = $1 AND followersid = $2`, userid, follwersid)
	if err != nil {
		return err
	}
	return nil
}

func (f followerRepository) FetchFollowers(userid int64) ([]response.Follower, error) {
	followers, err := f.pgx.Query(context.Background(), `SELECT f.followersid, u.name, u.bio, u.image FROM followers f JOIN users u ON f.followersid = u.id WHERE f.userid = $1`, userid)
	followersSlice := make([]response.Follower, 0)

	if err != nil {
		return followersSlice, err
	}
	defer followers.Close()

	var follower response.Follower
	for followers.Next() {
		if err := followers.Scan(&follower.ID, &follower.Name, &follower.Bio, &follower.Image); err != nil {
			return followersSlice, err
		}
		followersSlice = append(followersSlice, follower)
	}

	return followersSlice, nil
}

func (f followerRepository) FetchFollowings(followersid int64) ([]response.Follower, error) {
	followings, err := f.pgx.Query(context.Background(), `SELECT f.userid, u.name, u.bio, u.image FROM followers f JOIN users u ON f.userid = u.id WHERE f.followersid = $1`, followersid)
	followingsSlice := make([]response.Follower, 0)

	if err != nil {
		return followingsSlice, err
	}
	defer followings.Close()

	var following response.Follower
	for followings.Next() {
		if err := followings.Scan(&following.ID, &following.Name, &following.Bio, &following.Image); err != nil {
			return followingsSlice, err
		}
		followingsSlice = append(followingsSlice, following)
	}

	return followingsSlice, nil
}
