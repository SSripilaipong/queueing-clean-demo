package manage_doctor_queue

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
