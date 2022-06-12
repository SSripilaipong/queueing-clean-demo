package app

import (
	"context"
	"queueing-clean-demo/domain"
	"queueing-clean-demo/domain/manage_doctor_queue/usecase"
	"queueing-clean-demo/implementation"
	"queueing-clean-demo/toolbox/mongodb"
	"queueing-clean-demo/toolbox/rabbitmq"
	d "queueing-clean-demo/worker/deps"
)

type workerDeps struct {
	mongo             *mongodb.Connection
	rabbit            *rabbitmq.Client
	manageDoctorQueue domain.IManageDoctorQueueUsecase
}

func newWorkerDeps() d.IWorkerDeps {
	connection := makeMongoDbConnection()
	rabbit := makeRabbitMQClient()

	db := connection.Client.Database("OPD")

	return &workerDeps{
		rabbit: rabbit,
		mongo:  connection,
		manageDoctorQueue: usecase.NewManageDoctorQueueUsecase(
			implementation.NewDoctorQueueRepoInMongo(db.Collection("DoctorQueueRepo")),
			implementation.Clock{},
		),
	}
}

func (d *workerDeps) ManageDoctorQueue() domain.IManageDoctorQueueUsecase {
	return d.manageDoctorQueue
}

func (d *workerDeps) Broker() d.IBroker {
	return d.rabbit
}

func (d *workerDeps) Destroy() {
	d.mongo.Disconnect(context.Background())
	d.rabbit.Disconnect()
}
