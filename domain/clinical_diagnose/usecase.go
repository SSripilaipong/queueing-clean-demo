package _clinical_diagnose

import (
	"queueing-clean-demo/domain/clinical_diagnose/contract"
	"queueing-clean-demo/domain/clinical_diagnose/deps"
	"queueing-clean-demo/domain/common/contract"
)

type Usecase struct {
	VisitRepo   _deps.IVisitRepo
	IdGenerator _deps.IIdGenerator
}

func (u *Usecase) CreateVisit(request clinical_diagnose.CreateVisit) (clinical_diagnose.VisitResponse, error) {

TryCreateVisit:
	id := u.IdGenerator.GetId()
	visit, err := NewVisit(id, request.Name, request.Gender, request.Age)
	if err == nil {
		visit, err = u.VisitRepo.Create(visit)
	}

	switch err.(type) {
	case common.DuplicateVisitIdError:
		goto TryCreateVisit
	case common.InvalidVisitDataError:
		return clinical_diagnose.VisitResponse{}, err
	case nil:
		return VisitResponseFromVisit(visit), nil
	}
	panic(err)
}

func (u *Usecase) SubmitAssessment(request clinical_diagnose.SubmitAssessment) (clinical_diagnose.VisitResponse, error) {

	visit, err := u.VisitRepo.FindByIdAndUpdate(request.VisitId, func(visit *Visit) (*Visit, error) {
		if err := visit.SubmitAssessment(AssessmentFromSubmitAssessmentRequest(request)); err != nil {
			return nil, err
		}
		return visit, nil
	})

	switch err.(type) {
	case common.VisitNotFoundError:
		return clinical_diagnose.VisitResponse{}, err
	case clinical_diagnose.AssessmentAlreadyExistError:
		return clinical_diagnose.VisitResponse{}, err
	case nil:
		return VisitResponseFromVisit(visit), nil
	}
	panic(err)
}
