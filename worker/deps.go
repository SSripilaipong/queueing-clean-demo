package worker

import (
	"go.mongodb.org/mongo-driver/mongo"
	"queueing-clean-demo/domain/contract"
	"queueing-clean-demo/domain/manage_doctor_queue/usecase"
	impl "queueing-clean-demo/implementation"
)

type Deps struct {
	ManageDoctorQueueUsecase domain.IManageDoctorQueueUsecase
}

func createDeps(db *mongo.Database) *Deps {
	return &Deps{
		ManageDoctorQueueUsecase: usecase.NewManageDoctorQueueUsecase(
			&impl.DoctorQueueRepoInMongo{Collection: db.Collection("DoctorQueueRepo")},
			impl.Clock{},
		),
	}
}
