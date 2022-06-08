package app

import (
	"go.mongodb.org/mongo-driver/mongo"
	"queueing-clean-demo/domain/usecase/clinical_diagnose"
	"queueing-clean-demo/domain/usecase/manage_doctor_queue"
	impl "queueing-clean-demo/implementation"
	d "queueing-clean-demo/rest/deps"
)

func createRestDeps(database *mongo.Database) d.RestDeps {
	return d.RestDeps{
		ClinicalDiagnoseUsecase: &clinical_diagnose.Usecase{
			VisitRepo:   &impl.VisitRepoInMongo{Collection: database.Collection("VisitRepo")},
			IdGenerator: impl.IdGenerator{},
		},
		ManageDoctorQueueUsecase: &manage_doctor_queue.Usecase{
			DoctorQueueRepo: &impl.DoctorQueueRepoInMongo{Collection: database.Collection("DoctorQueueRepo")},
			Clock:           impl.Clock{},
		},
	}
}
