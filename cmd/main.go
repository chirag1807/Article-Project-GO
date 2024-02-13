package main

import (
	"articleproject/api/route"
	"articleproject/config"
	"articleproject/constants"
	"articleproject/db"
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadEnv()
	fmt.Println(config.DatabaseConfig.DATABASE_NAME, config.JWtSecretConfig.SecretKey)
	conn, err := db.DBConnection()

	if err != nil {
		return
	}
	defer conn.Close(context.Background())

	r := route.UsersRoutes(conn)

	fmt.Println("Server started on port no. " + constants.PORT_NO)
	log.Fatal(http.ListenAndServe(constants.PORT_NO, r))
}
