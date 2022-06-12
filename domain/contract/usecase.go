package domain

import (
	clinical_diagnose2 "queueing-clean-demo/domain/clinical_diagnose"
	manage_doctor_queue2 "queueing-clean-demo/domain/manage_doctor_queue"
)

type IClinicalDiagnoseUsecase interface {
	CreateVisit(request clinical_diagnose2.CreateVisit) (clinical_diagnose2.VisitResponse, error)
	SubmitAssessment(request clinical_diagnose2.SubmitAssessment) (clinical_diagnose2.VisitResponse, error)
}

type IManageDoctorQueueUsecase interface {
	PushVisit(request manage_doctor_queue2.PushVisitToDoctorQueue) (manage_doctor_queue2.DoctorQueueResponse, error)
	CallVisit(request manage_doctor_queue2.CallVisitFromDoctorQueue) (manage_doctor_queue2.DoctorQueueResponse, error)
	CompleteDiagnosis(request manage_doctor_queue2.CompleteDiagnosis) (manage_doctor_queue2.DoctorQueueResponse, error)
	CheckVisits(request manage_doctor_queue2.CheckVisits) (manage_doctor_queue2.DoctorQueueResponse, error)
	CreateDoctorQueue(request manage_doctor_queue2.CreateDoctorQueue) (manage_doctor_queue2.DoctorQueueResponse, error)
}
