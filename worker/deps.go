package worker

import (
	"go.mongodb.org/mongo-driver/mongo"
	"queueing-clean-demo/domain/contract"
	"queueing-clean-demo/domain/manage_doctor_queue/contract"
	impl "queueing-clean-demo/implementation"
)

type Deps struct {
	ManageDoctorQueueUsecase domain.IManageDoctorQueueUsecase
}

func createDeps(db *mongo.Database) *Deps {
	return &Deps{
		ManageDoctorQueueUsecase: manage_doctor_queue.manage_doctor_queue.NewUsecase(
			&impl.DoctorQueueRepoInMongo{Collection: db.Collection("DoctorQueueRepo")},
			impl.Clock{},
		),
	}
}
