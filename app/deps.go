package app

import (
	"go.mongodb.org/mongo-driver/mongo"
	. "queueing-clean-demo/domain/clinical_diagnose/usecase"
	. "queueing-clean-demo/domain/manage_doctor_queue/usecase"
	impl "queueing-clean-demo/implementation"
	d "queueing-clean-demo/rest/deps"
)

func createRestDeps(database *mongo.Database) d.RestDeps {
	return d.RestDeps{
		ClinicalDiagnoseUsecase: NewClinicalDiagnoseUsecase(
			impl.NewVisitRepoInMongo(database.Collection("VisitRepo")),
			impl.IdGenerator{},
		),
		ManageDoctorQueueUsecase: NewManageDoctorQueueUsecase(
			impl.NewDoctorQueueRepoInMongo(database.Collection("DoctorQueueRepo")),
			impl.Clock{},
		),
	}
}
