package deps

import (
	"context"
	connection2 "queueing-clean-demo/app/connection"
	"queueing-clean-demo/domain"
	. "queueing-clean-demo/domain/clinical_diagnose/usecase"
	. "queueing-clean-demo/domain/manage_doctor_queue/usecase"
	impl "queueing-clean-demo/implementation"
	d "queueing-clean-demo/rest/deps"
	"queueing-clean-demo/toolbox/mongodb"
)

type restDeps struct {
	connection        *mongodb.Connection
	clinicalDiagnose  domain.IClinicalDiagnoseUsecase
	manageDoctorQueue domain.IManageDoctorQueueUsecase
}

func NewRestDeps() d.IRestDeps {
	connection := connection2.MakeMongoDbConnection()

	database := connection.Client.Database("OPD")

	return &restDeps{
		clinicalDiagnose: NewClinicalDiagnoseUsecase(
			impl.NewVisitRepoInMongo(database.Collection("VisitRepo")),
			impl.IdGenerator{},
		),
		manageDoctorQueue: NewManageDoctorQueueUsecase(
			impl.NewDoctorQueueRepoInMongo(database.Collection("DoctorQueueRepo")),
			impl.Clock{},
		),
		connection: connection,
	}
}

func (d *restDeps) ClinicalDiagnose() domain.IClinicalDiagnoseUsecase {
	return d.clinicalDiagnose
}

func (d *restDeps) ManageDoctorQueue() domain.IManageDoctorQueueUsecase {
	return d.manageDoctorQueue
}

func (d *restDeps) Destroy() {
	d.connection.Disconnect(context.Background())
}
