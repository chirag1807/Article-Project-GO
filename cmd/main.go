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
	config.LoadEnv("../.config/")
	conn, rdb, amqp, err := db.DBConnection()
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close(context.Background())
	defer amqp.Close()

	r := route.UsersRoutes(conn, rdb, amqp)

	// utils.SendSMS("+919327963104", "Hey Chirag, This is Testing Purpose SMS From ZURU TECH.")
	// utils.MakeVoiceVall("+919327963104", "http://demo.twilio.com/docs/voice.xml")

	log.Println("Server started on port no. " + constants.PORT_NO)
	log.Fatal(http.ListenAndServe(constants.PORT_NO, r))
}
