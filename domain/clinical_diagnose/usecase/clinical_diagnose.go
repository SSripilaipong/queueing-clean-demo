package usecase

import (
	. "queueing-clean-demo/domain/clinical_diagnose/contract"
	"queueing-clean-demo/domain/clinical_diagnose/internal"
	"queueing-clean-demo/domain/contract"
)

func NewClinicalDiagnoseUsecase(visitRepo IVisitRepo, idGenerator IIdGenerator) domain.IClinicalDiagnoseUsecase {
	return &internal.Usecase{
		VisitRepo:   visitRepo,
		IdGenerator: idGenerator,
	}
}
