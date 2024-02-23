package db

import (
	"articleproject/config"
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5"
	amqp "github.com/rabbitmq/amqp091-go"
)

func DBConnection() (conn *pgx.Conn, rdb *redis.Client, amqpConn *amqp.Connection, err error) {
	DATABASE_URL := "postgresql://" + config.Config.Database.Username + ":" + config.Config.Database.Password + "@127.0.0.1:" + config.Config.Database.Port + "/" + config.Config.Database.Name + "?sslmode=" + config.Config.Database.SSLMode
	connConfig, err := pgx.ParseConfig(DATABASE_URL)
	if err != nil {
		fmt.Println(err)
		return nil, nil, nil, err
	}
	conn, err = pgx.ConnectConfig(context.Background(), connConfig)
	if err != nil {
		fmt.Println(err)
		return nil, nil, nil, err
	}
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Port,
		Password: config.Config.Redis.Password,
		DB:       config.Config.Redis.DB,
	})
	_, err = rdb.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err)
		return nil, nil, nil, err
	}
	rabbitmqConn, err := amqp.Dial("amqp://" + config.Config.RabbitMQ.Username + ":" + config.Config.RabbitMQ.Password + "@localhost:" + config.Config.RabbitMQ.Port + "/")
	if err != nil {
		fmt.Println(err)
		return nil, nil, nil, err
	}
	return conn, rdb, rabbitmqConn, nil
}
