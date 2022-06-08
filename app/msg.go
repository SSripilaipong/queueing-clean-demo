package app

import (
	"github.com/streadway/amqp"
)

func SetupMessageBroker() {
	var err error

	var conn *amqp.Connection
	if conn, err = amqp.Dial("amqp://root:admin@rabbitmq:5672"); err != nil {
		panic(err)
	}
	defer conn.Close()

	var ch *amqp.Channel
	if ch, err = conn.Channel(); err != nil {
		defer conn.Close()
		panic(err)
	}
	defer ch.Close()

	if _, err = ch.QueueDeclare(
		"allEvents",
		false,
		false,
		false,
		false,
		nil,
	); err != nil {
		panic(err)
	}
}
