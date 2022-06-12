package deps

import (
	"queueing-clean-demo/app/connection"
	"queueing-clean-demo/outbox/deps"
	"queueing-clean-demo/toolbox/mongo_watcher"
	"queueing-clean-demo/toolbox/rabbitmq"
)

type outboxDeps struct {
	watcher *mongo_watcher.MongoWatcher
	rabbit  *rabbitmq.Client
}

func NewOutboxDeps() deps.IOutboxDeps {
	watcher := connection.MakeMongoWatcher()
	rabbit := connection.MakeRabbitMQClient()

	return &outboxDeps{
		rabbit:  rabbit,
		watcher: watcher,
	}
}

func (d *outboxDeps) Broker() deps.IPublisher {
	return d.rabbit
}

func (d *outboxDeps) Stream() deps.IStream {
	return d.watcher
}

func (d *outboxDeps) Destroy() {
	d.watcher.Disconnect()
	d.rabbit.Disconnect()
}
