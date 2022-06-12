package deps

import (
	"context"
	"github.com/streadway/amqp"
	"queueing-clean-demo/domain"
)

type IWorkerDeps interface {
	ManageDoctorQueue() domain.IManageDoctorQueueUsecase
	Broker() ISubscriber
	Destroy()
}

type ISubscriber interface {
	Subscribe(ctx context.Context, key string, handle func(amqp.Delivery))
}
