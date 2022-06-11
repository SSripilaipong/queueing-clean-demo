package deps

import (
	"queueing-clean-demo/domain/contract"
)

type RestDeps struct {
	ClinicalDiagnoseUsecase  domain.IClinicalDiagnoseUsecase
	ManageDoctorQueueUsecase domain.IManageDoctorQueueUsecase
}
