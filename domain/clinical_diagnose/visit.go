package _clinical_diagnose

import (
	"queueing-clean-demo/base"
	"queueing-clean-demo/domain/clinical_diagnose/contract"
	"queueing-clean-demo/domain/common/contract"
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
		return clinical_diagnose.AssessmentAlreadyExistError{}
	}
	v.Assessment = &assessment
	v.AppendEvent(clinical_diagnose.VisitAssessedEvent{
		VisitId: v.Id,
		Name:    v.Name,
		Gender:  v.Gender,
		Age:     v.Age,
	})
	return nil
}
