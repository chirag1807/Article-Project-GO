package rabbitmq_user

import (
	"articleproject/api/model/dto"
	"articleproject/utils"
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ConsumeUserMail(amqp_conn *amqp.Connection) {
	ch, err := amqp_conn.Channel()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ch.Close()

	queueName := "user-mail-queue"
	msgs, err := ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	var forever = make(chan bool)

	go func() {
		for m := range msgs {
			var email dto.UserMail
			_ = json.Unmarshal(m.Body, &email)
			fmt.Println(email.To, email.Subject, email.Body)
			utils.SendMail(email)
		}
	}()

	<-forever
	// <-forever which blocks our main function from completing until the channel is satisfied.
}
