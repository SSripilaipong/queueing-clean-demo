package contract

import (
	"queueing-clean-demo/domain/usecase/clinical_diagnose/contract"
	contract2 "queueing-clean-demo/domain/usecase/manage_doctor_queue/contract"
)

type IClinicalDiagnoseUsecase interface {
	CreateVisit(request contract.CreateVisit) (contract.Visit, error)
	SubmitAssessment(request contract.SubmitAssessment) (contract.Visit, error)
}

type IManageDoctorQueueUsecase interface {
	PushVisit(request contract2.PushVisitToDoctorQueue) (contract2.DoctorQueueResponse, error)
	CallVisit(request contract2.CallVisitFromDoctorQueue) (contract2.DoctorQueueResponse, error)
	CompleteDiagnosis(request contract2.CompleteDiagnosis) (contract2.DoctorQueueResponse, error)
	CheckVisits(request contract2.CheckVisits) (contract2.DoctorQueueResponse, error)
	CreateDoctorQueue(request contract2.CreateDoctorQueue) (contract2.DoctorQueueResponse, error)
}
