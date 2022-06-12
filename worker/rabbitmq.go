package worker

import (
	"context"
	"github.com/streadway/amqp"
	"queueing-clean-demo/toolbox/mongodb"
	"queueing-clean-demo/worker/deps"
)

func workerLoop(ctx context.Context, handle func(*worker_deps.Deps, amqp.Delivery)) {
	var err error

	var mgConn *mongodb.Connection
	if mgConn, err = mongodb.CreateConnection(ctx, "root", "admin", "mongodb", "27017"); err != nil {
		panic(err)
	}
	defer mgConn.Disconnect(ctx)

	deps := worker_deps.CreateDeps(mgConn.Client.Database("OPD"))

	rbConn, rbCh := makeChannel()
	defer func(rbConn *amqp.Connection) {
		_ = rbConn.Close()
	}(rbConn)
	defer func(rbCh *amqp.Channel) {
		_ = rbCh.Close()
	}(rbCh)

	var delivery <-chan amqp.Delivery
	if delivery, err = rbCh.Consume(
		"allEvents",
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
			handle(deps, msg)
			if err := msg.Ack(false); err != nil {
				panic(err)
			}

		case <-ctx.Done():
			running = false
		}
	}
}

func makeChannel() (*amqp.Connection, *amqp.Channel) {
	var err error

	var conn *amqp.Connection
	if conn, err = amqp.Dial("amqp://root:admin@rabbitmq:5672"); err != nil {
		panic(err)
	}

	var ch *amqp.Channel
	if ch, err = conn.Channel(); err != nil {
		defer func(conn *amqp.Connection) {
			_ = conn.Close()
		}(conn)
		panic(err)
	}
	return conn, ch
}
