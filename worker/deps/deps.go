package deps

import (
	"context"
	"github.com/streadway/amqp"
	"queueing-clean-demo/domain"
)

type IWorkerDeps interface {
	ManageDoctorQueue() domain.IManageDoctorQueueUsecase
	Broker() IBroker
	Destroy()
}

type IBroker interface {
	Subscribe(ctx context.Context, key string, handle func(amqp.Delivery))
}
