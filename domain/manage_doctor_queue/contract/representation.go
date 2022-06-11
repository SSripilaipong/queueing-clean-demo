package manage_doctor_queue

import (
	"queueing-clean-demo/base"
	"time"
)

type DoctorQueueRepr struct {
	base.AggregateRepr `json:"_aggregate"`
	DoctorId           string
	Visits             map[string]VisitShortInfoRepr
	VisitInProgress    *VisitShortInfoRepr
}

type VisitShortInfoRepr struct {
	Id        string
	Name      string
	Gender    string
	Age       int
	EnterTime time.Time
}
