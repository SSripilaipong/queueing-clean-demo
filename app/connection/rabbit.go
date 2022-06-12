package connection

import "queueing-clean-demo/toolbox/rabbitmq"

func MakeRabbitMQClient() *rabbitmq.Client {
	return rabbitmq.NewClient("root", "admin", "rabbitmq", "5672")
}
