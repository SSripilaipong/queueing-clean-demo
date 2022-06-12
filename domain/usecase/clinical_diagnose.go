package usecase

import (
	"queueing-clean-demo/domain/clinical_diagnose"
	. "queueing-clean-demo/domain/clinical_diagnose/contract"
	"queueing-clean-demo/domain/contract"
)

func NewClinicalDiagnoseUsecase(visitRepo IVisitRepo, idGenerator IIdGenerator) domain.IClinicalDiagnoseUsecase {
	return &_clinical_diagnose.Usecase{
		VisitRepo:   visitRepo,
		IdGenerator: idGenerator,
	}
}
