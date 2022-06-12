package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

func makeConnectionAndChannel(username string, password string, host string, port string) (*amqp.Connection, *amqp.Channel) {
	var err error

	url := fmt.Sprintf("amqp://%s:%s@%s:%s", username, password, host, port)
	var conn *amqp.Connection
	if conn, err = amqp.Dial(url); err != nil {
		panic(err)
	}

	var ch *amqp.Channel
	if ch, err = conn.Channel(); err != nil {
		defer func(conn *amqp.Connection) {
			if err := conn.Close(); err != nil {
				panic(err)
			}
		}(conn)
		panic(err)
	}
	return conn, ch
}
