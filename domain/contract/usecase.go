package domain

import (
	clinical_diagnose2 "queueing-clean-demo/domain/clinical_diagnose"
	"queueing-clean-demo/domain/manage_doctor_queue/contract"
)

type IClinicalDiagnoseUsecase interface {
	CreateVisit(request clinical_diagnose2.CreateVisit) (clinical_diagnose2.VisitResponse, error)
	SubmitAssessment(request clinical_diagnose2.SubmitAssessment) (clinical_diagnose2.VisitResponse, error)
}

type IManageDoctorQueueUsecase interface {
	PushVisit(request manage_doctor_queue.PushVisitToDoctorQueue) (manage_doctor_queue.DoctorQueueResponse, error)
	CallVisit(request manage_doctor_queue.CallVisitFromDoctorQueue) (manage_doctor_queue.DoctorQueueResponse, error)
	CompleteDiagnosis(request manage_doctor_queue.CompleteDiagnosis) (manage_doctor_queue.DoctorQueueResponse, error)
	CheckVisits(request manage_doctor_queue.CheckVisits) (manage_doctor_queue.DoctorQueueResponse, error)
	CreateDoctorQueue(request manage_doctor_queue.CreateDoctorQueue) (manage_doctor_queue.DoctorQueueResponse, error)
}
