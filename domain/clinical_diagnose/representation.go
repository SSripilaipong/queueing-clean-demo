package clinical_diagnose

import "queueing-clean-demo/base"

type VisitRepr struct {
	base.AggregateRepr `json:"_aggregate"`

	Id         string
	Name       string
	Gender     string
	Age        int
	Assessment *AssessmentRepr
}

type AssessmentRepr struct {
	NursingAssessment string
	PainScore         int
}
