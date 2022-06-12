package deps

import (
	"queueing-clean-demo/domain"
)

type IRestDeps interface {
	ClinicalDiagnose() domain.IClinicalDiagnoseUsecase
	ManageDoctorQueue() domain.IManageDoctorQueueUsecase
	Destroy()
}
