package clinical_diagnose

import (
	v "queueing-clean-demo/domain"
	dContract "queueing-clean-demo/domain/contract"
	"queueing-clean-demo/domain/usecase/clinical_diagnose/contract"
	"queueing-clean-demo/domain/usecase/clinical_diagnose/deps"
)

type Usecase struct {
	VisitRepo   deps.IVisitRepo
	IdGenerator deps.IIdGenerator
}

func (u *Usecase) CreateVisit(request contract.CreateVisit) (contract.Visit, error) {

TryCreateVisit:
	id := u.IdGenerator.GetId()
	visit, err := v.NewVisit(id, request.Name, request.Gender, request.Age)
	if err == nil {
		visit, err = u.VisitRepo.Create(visit)
	}

	switch err.(type) {
	case dContract.DuplicateVisitIdError:
		goto TryCreateVisit
	case dContract.InvalidVisitDataError:
		return contract.Visit{}, err
	case nil:
		return v.VisitResponseFromVisit(visit), nil
	}
	panic(err)
}

func (u *Usecase) SubmitAssessment(request contract.SubmitAssessment) (contract.Visit, error) {

	visit, err := u.VisitRepo.FindByIdAndUpdate(request.VisitId, func(visit *v.Visit) (*v.Visit, error) {
		if err := visit.SubmitAssessment(v.AssessmentFromSubmitAssessmentRequest(request)); err != nil {
			return nil, err
		}
		return visit, nil
	})

	switch err.(type) {
	case dContract.VisitNotFoundError:
		return contract.Visit{}, err
	case contract.AssessmentAlreadyExistError:
		return contract.Visit{}, err
	case nil:
		return v.VisitResponseFromVisit(visit), nil
	}
	panic(err)
}
