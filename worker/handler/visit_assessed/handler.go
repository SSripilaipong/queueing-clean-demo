package visit_assessed

import (
	"queueing-clean-demo/domain/clinical_diagnose"
	"queueing-clean-demo/domain/manage_doctor_queue"
)

func Handler(deps deps.IWorkerDeps, e clinical_diagnose.VisitAssessedEvent) {
	_, err := deps.ManageDoctorQueue().PushVisit(manage_doctor_queue.PushVisitToDoctorQueue{
		DoctorId:      "629c93cae6509bc3a7b1aaf7", // fixed for simplicity
		VisitId:       e.VisitId,
		PatientName:   e.Name,
		PatientGender: e.Gender,
		PatientAge:    e.Age,
	})

	switch err.(type) {
	case manage_doctor_queue.DoctorQueueNotFoundError:
		return
	case manage_doctor_queue.VisitAlreadyExistsError:
		return
	case nil:
		return
	}
	panic(err)
}
