package internal

import (
	. "queueing-clean-demo/domain/clinical_diagnose/contract"
	"queueing-clean-demo/domain/common/contract"
)

type Usecase struct {
	VisitRepo   IVisitRepo
	IdGenerator IIdGenerator
}

func (u *Usecase) CreateVisit(request CreateVisit) (VisitResponse, error) {

TryCreateVisit:
	id := u.IdGenerator.GetId()
	visit, err := NewVisit(id, request.Name, request.Gender, request.Age)
	if err == nil {
		repr := visit.ToRepr()
		_, err = u.VisitRepo.Create(&repr)
	}

	switch err.(type) {
	case common.DuplicateVisitIdError:
		goto TryCreateVisit
	case common.InvalidVisitDataError:
		return VisitResponse{}, err
	case nil:
		return VisitResponseFromVisit(visit), nil
	}
	panic(err)
}

func (u *Usecase) SubmitAssessment(request SubmitAssessment) (VisitResponse, error) {

	repr, err := u.VisitRepo.FindByIdAndUpdate(request.VisitId, func(repr *VisitRepr) (*VisitRepr, error) {
		visit := NewVisitFromRepr(*repr)
		if err := visit.SubmitAssessment(AssessmentFromSubmitAssessmentRequest(request)); err != nil {
			return nil, err
		}
		result := visit.ToRepr()
		return &result, nil
	})

	switch err.(type) {
	case common.VisitNotFoundError:
		return VisitResponse{}, err
	case AssessmentAlreadyExistError:
		return VisitResponse{}, err
	case nil:
		return VisitResponseFromVisit(NewVisitFromRepr(*repr)), nil
	}
	panic(err)
}
