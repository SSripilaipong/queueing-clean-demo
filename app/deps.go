package app

import (
	"go.mongodb.org/mongo-driver/mongo"
	"queueing-clean-demo/domain/usecase"
	impl "queueing-clean-demo/implementation"
	d "queueing-clean-demo/rest/deps"
)

func createRestDeps(database *mongo.Database) d.RestDeps {
	return d.RestDeps{
		ClinicalDiagnoseUsecase: usecase.NewClinicalDiagnoseUsecase(
			&impl.VisitRepoInMongo{Collection: database.Collection("VisitRepo")},
			impl.IdGenerator{},
		),
		ManageDoctorQueueUsecase: usecase.NewManageDoctorQueueUsecase(
			&impl.DoctorQueueRepoInMongo{Collection: database.Collection("DoctorQueueRepo")},
			impl.Clock{},
		),
	}
}
