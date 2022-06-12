package rabbitmq

import (
	"context"
	"encoding/json"
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

func (c *Client) Subscribe(ctx context.Context, key string, handle func(amqp.Delivery)) {
	var err error
	var delivery <-chan amqp.Delivery

	if delivery, err = c.ch.Consume(
		key,
		"",
		false,
		false,
		false,
		false,
		nil,
	); err != nil {
		panic(err)
	}

	running := true
	for running {
		select {
		case msg := <-delivery:
			handle(msg)

		case <-ctx.Done():
			running = false
		}
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
