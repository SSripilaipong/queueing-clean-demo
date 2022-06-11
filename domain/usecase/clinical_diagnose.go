package usecase

import (
	"queueing-clean-demo/domain/clinical_diagnose"
	"queueing-clean-demo/domain/clinical_diagnose/deps"
	"queueing-clean-demo/domain/contract"
)

func NewClinicalDiagnoseUsecase(visitRepo _deps.IVisitRepo, idGenerator _deps.IIdGenerator) domain.IClinicalDiagnoseUsecase {
	return &_clinical_diagnose.Usecase{
		VisitRepo:   visitRepo,
		IdGenerator: idGenerator,
	}
}
