package rabbitmq_user

import (
	"articleproject/api/model/dto"
	"articleproject/api/model/request"
	"context"
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ProduceUserMail(amqp_conn *amqp.Connection, user request.User) {
	ch, err := amqp_conn.Channel()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ch.Close()

	queueName := "user-mail-queue"
	q, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	email := dto.UserMail{
		To:      user.Email,
		Subject: "Registration Completed Successfully.",
		Body:    "Hey " + user.Name + ",\nCongratulations, Your Registration Completed Successfully.",
	}
	body, _ := json.Marshal(email)

	err = ch.PublishWithContext(context.Background(), "", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        body,
	})

	if err != nil {
		fmt.Println(err)
	}

	go ConsumeUserMail(amqp_conn)
}
