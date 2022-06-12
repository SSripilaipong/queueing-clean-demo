package deps

import (
	"queueing-clean-demo/domain"
)

type RestDeps struct {
	ClinicalDiagnoseUsecase  domain.IClinicalDiagnoseUsecase
	ManageDoctorQueueUsecase domain.IManageDoctorQueueUsecase
}
