package main

import (
	"articleproject/api/route"
	"articleproject/config"
	"articleproject/constants"
	"articleproject/db"
	"context"
	"log"
	"net/http"
)

func main() {
	config.LoadEnv()
	conn, rdb, amqp, err := db.DBConnection()
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close(context.Background())
	defer amqp.Close()

	r := route.UsersRoutes(conn, rdb, amqp)

	log.Println("Server started on port no. " + constants.PORT_NO)
	log.Fatal(http.ListenAndServe(constants.PORT_NO, r))
}
