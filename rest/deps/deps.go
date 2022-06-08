package deps

import (
	"queueing-clean-demo/domain/contract"
)

type RestDeps struct {
	ClinicalDiagnoseUsecase  contract.IClinicalDiagnoseUsecase
	ManageDoctorQueueUsecase contract.IManageDoctorQueueUsecase
}
