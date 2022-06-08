package contract

import (
	"time"
)

type DoctorQueueResponse struct {
	Visits          []VisitShortInfoResponse
	VisitInProgress VisitShortInfoResponse
}

type VisitShortInfoResponse struct {
	Id        string
	Name      string
	Gender    string
	Age       int
	EnterTime time.Time
}

type AssessmentResponse struct {
	NursingAssessment string
	PainScore         int
}
