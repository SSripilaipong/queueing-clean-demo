package usecase

import (
	"queueing-clean-demo/domain"
	. "queueing-clean-demo/domain/clinical_diagnose"
	"queueing-clean-demo/domain/clinical_diagnose/internal"
)

func NewClinicalDiagnoseUsecase(visitRepo IVisitRepo, idGenerator IIdGenerator) domain.IClinicalDiagnoseUsecase {
	return &internal.Usecase{
		VisitRepo:   visitRepo,
		IdGenerator: idGenerator,
	}
}
