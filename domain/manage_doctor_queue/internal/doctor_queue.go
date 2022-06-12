package internal

import (
	"queueing-clean-demo/base"
	"queueing-clean-demo/domain/common"
	. "queueing-clean-demo/domain/manage_doctor_queue"
	"time"
)

type DoctorQueue struct {
	base.Aggregate
	DoctorId        string
	Visits          VisitCollection
	VisitInProgress *VisitShortInfo
}

type VisitShortInfo struct {
	Id        string
	Name      string
	Gender    string
	Age       int
	EnterTime time.Time
}

func NewDoctorQueue(doctorId string) (*DoctorQueue, error) {
	return &DoctorQueue{DoctorId: doctorId, Visits: VisitCollection{Queue: map[string]VisitShortInfo{}}}, nil
}

func (q *DoctorQueue) PushVisit(visit VisitShortInfo) error {
	if q.Visits.hasVisit(visit.Id) {
		return VisitAlreadyExistsError{}
	}
	q.Visits.add(visit)
	return nil
}

func (q *DoctorQueue) CallVisit(id string) error {
	if q.VisitInProgress != nil {
		return DoctorStillBusyError{}
	}
	if !q.Visits.hasVisit(id) {
		return common.VisitNotFoundError{}
	}
	visit := q.Visits.pop(id)
	q.VisitInProgress = &visit
	return nil
}

func (q *DoctorQueue) CompleteDiagnosis() error {
	if q.VisitInProgress == nil {
		return NoVisitInProgressToCompleteError{}
	}
	q.VisitInProgress = nil
	return nil
}

func (q *DoctorQueue) ToResponse() DoctorQueueResponse {
	visits := make([]VisitShortInfoResponse, len(q.Visits.Queue))
	i := 0
	for _, v := range q.Visits.Queue {
		visits[i] = VisitShortInfoResponse{
			Id:        v.Id,
			Name:      v.Name,
			Gender:    v.Gender,
			Age:       v.Age,
			EnterTime: v.EnterTime,
		}
		i++
	}

	var visitInProgress VisitShortInfoResponse
	if vip := q.VisitInProgress; vip != nil {
		visitInProgress = VisitShortInfoResponse{
			Id:        vip.Id,
			Name:      vip.Name,
			Gender:    vip.Gender,
			Age:       vip.Age,
			EnterTime: vip.EnterTime,
		}
	}

	repr := DoctorQueueResponse{
		Visits:          visits,
		VisitInProgress: visitInProgress,
	}
	return repr
}

func (q *DoctorQueue) ToRepr() *DoctorQueueRepr {
	visits := make(map[string]VisitShortInfoRepr)
	for k, v := range q.Visits.Queue {
		visits[k] = VisitShortInfoRepr{
			Id:        v.Id,
			Name:      v.Name,
			Gender:    v.Gender,
			Age:       v.Age,
			EnterTime: v.EnterTime,
		}
	}

	var visitInProgress *VisitShortInfoRepr
	if vip := q.VisitInProgress; vip != nil {
		visitInProgress = &VisitShortInfoRepr{
			Id:        vip.Id,
			Name:      vip.Name,
			Gender:    vip.Gender,
			Age:       vip.Age,
			EnterTime: vip.EnterTime,
		}
	}

	repr := &DoctorQueueRepr{
		AggregateRepr: base.AggregateRepr{
			Version: q.GetVersion(),
			Events:  q.GetEvents(),
		},
		DoctorId:        q.DoctorId,
		Visits:          visits,
		VisitInProgress: visitInProgress,
	}
	return repr
}

func NewDoctorQueueFromRepr(repr *DoctorQueueRepr) *DoctorQueue {
	visits := make(map[string]VisitShortInfo)
	for k, v := range repr.Visits {
		visits[k] = VisitShortInfo{
			Id:        v.Id,
			Name:      v.Name,
			Gender:    v.Gender,
			Age:       v.Age,
			EnterTime: v.EnterTime,
		}
	}

	var visitInProgress *VisitShortInfo
	if vip := repr.VisitInProgress; vip != nil {
		visitInProgress = &VisitShortInfo{
			Id:        vip.Id,
			Name:      vip.Name,
			Gender:    vip.Gender,
			Age:       vip.Age,
			EnterTime: vip.EnterTime,
		}
	}

	q := &DoctorQueue{
		Aggregate: base.Aggregate{},
		DoctorId:  repr.DoctorId,
		Visits: VisitCollection{
			Queue: visits,
		},
		VisitInProgress: visitInProgress,
	}
	q.SetVersion(repr.GetVersion())
	return q
}
