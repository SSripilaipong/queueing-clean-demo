package worker_deps

import (
	"go.mongodb.org/mongo-driver/mongo"
	"queueing-clean-demo/domain"
	"queueing-clean-demo/domain/manage_doctor_queue/usecase"
	impl "queueing-clean-demo/implementation"
)

type Deps struct {
	ManageDoctorQueueUsecase domain.IManageDoctorQueueUsecase
}

func CreateDeps(db *mongo.Database) *Deps {
	return &Deps{
		ManageDoctorQueueUsecase: usecase.NewManageDoctorQueueUsecase(
			impl.NewDoctorQueueRepoInMongo(db.Collection("DoctorQueueRepo")),
			impl.Clock{},
		),
	}
}
