package worker

import (
	"go.mongodb.org/mongo-driver/mongo"
	"queueing-clean-demo/domain/contract"
	"queueing-clean-demo/domain/usecase/manage_doctor_queue"
	impl "queueing-clean-demo/implementation"
)

type Deps struct {
	ManageDoctorQueueUsecase contract.IManageDoctorQueueUsecase
}

func createDeps(db *mongo.Database) *Deps {
	return &Deps{
		ManageDoctorQueueUsecase: &manage_doctor_queue.Usecase{
			DoctorQueueRepo: &impl.DoctorQueueRepoInMongo{Collection: db.Collection("DoctorQueueRepo")},
			Clock:           impl.Clock{},
		},
	}
}
