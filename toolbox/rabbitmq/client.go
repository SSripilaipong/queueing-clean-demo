package rabbitmq

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
)

type Client struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewClient(username string, password string, host string, port string) *Client {
	conn, ch := makeConnectionAndChannel(username, password, host, port)

	return &Client{
		conn: conn,
		ch:   ch,
	}
}

func (c *Client) Publish(key string, event any) {
	var err error

	var j []byte
	if j, err = json.Marshal(event); err != nil {
		panic(err)
	}

	if err = c.ch.Publish("", key, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        j,
	}); err != nil {
		panic(err)
	}
}

func (c *Client) Disconnect() {
	if err := c.ch.Close(); err != nil {
		panic(err)
	}
	if err := c.conn.Close(); err != nil {
		panic(err)
	}
}

func (c *Client) Ch() *amqp.Channel {
	return c.ch
}

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
