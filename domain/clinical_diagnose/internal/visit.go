package internal

import (
	"queueing-clean-demo/base"
	clinical_diagnose2 "queueing-clean-demo/domain/clinical_diagnose"
	"queueing-clean-demo/domain/common"
)

type Visit struct {
	base.Aggregate
	Id         string
	Name       string
	Gender     string
	Age        int
	Assessment *Assessment
}

type Assessment struct {
	NursingAssessment string
	PainScore         int
}

func NewVisit(id string, name string, gender string, age int) (*Visit, error) {
	if age < 0 {
		return nil, common.InvalidVisitDataError{}
	}

	visit := &Visit{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
	}
	return visit, nil
}

func (v *Visit) SubmitAssessment(assessment Assessment) error {
	if v.Assessment != nil {
		return clinical_diagnose2.AssessmentAlreadyExistError{}
	}
	v.Assessment = &assessment
	v.AppendEvent(clinical_diagnose2.VisitAssessedEvent{
		VisitId: v.Id,
		Name:    v.Name,
		Gender:  v.Gender,
		Age:     v.Age,
	})
	return nil
}

func (v *Visit) ToRepr() clinical_diagnose2.VisitRepr {
	var assessment *clinical_diagnose2.AssessmentRepr
	if v.Assessment != nil {
		assessment = &clinical_diagnose2.AssessmentRepr{
			NursingAssessment: v.Assessment.NursingAssessment,
			PainScore:         v.Assessment.PainScore,
		}
	}

	return clinical_diagnose2.VisitRepr{
		Id:         v.Id,
		Name:       v.Name,
		Gender:     v.Gender,
		Age:        v.Age,
		Assessment: assessment,
		AggregateRepr: base.AggregateRepr{
			Version: v.GetVersion(),
			Events:  v.GetEvents(),
		},
	}
}

func NewVisitFromRepr(repr clinical_diagnose2.VisitRepr) *Visit {
	var assessment *Assessment
	if repr.Assessment != nil {
		assessment = &Assessment{
			NursingAssessment: repr.Assessment.NursingAssessment,
			PainScore:         repr.Assessment.PainScore,
		}
	}
	visit := &Visit{
		Aggregate:  base.Aggregate{},
		Id:         repr.Id,
		Name:       repr.Name,
		Gender:     repr.Gender,
		Age:        repr.Age,
		Assessment: assessment,
	}
	visit.SetVersion(repr.Version)
	return visit
}
