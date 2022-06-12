package app

import "queueing-clean-demo/toolbox/rabbitmq"

func makeRabbitMQClient() *rabbitmq.Client {
	return rabbitmq.NewClient("root", "admin", "rabbitmq", "5672")
}
