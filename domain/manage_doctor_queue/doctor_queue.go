package _manage_doctor_queue

import (
	"queueing-clean-demo/base"
	"queueing-clean-demo/domain/common/contract"
	"queueing-clean-demo/domain/manage_doctor_queue/contract"
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
		return manage_doctor_queue.VisitAlreadyExistsError{}
	}
	q.Visits.add(visit)
	return nil
}

func (q *DoctorQueue) CallVisit(id string) error {
	if q.VisitInProgress != nil {
		return manage_doctor_queue.DoctorStillBusyError{}
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
		return manage_doctor_queue.NoVisitInProgressToCompleteError{}
	}
	q.VisitInProgress = nil
	return nil
}

func (q *DoctorQueue) ToResponse() manage_doctor_queue.DoctorQueueResponse {
	visits := make([]manage_doctor_queue.VisitShortInfoResponse, len(q.Visits.Queue))
	i := 0
	for _, v := range q.Visits.Queue {
		visits[i] = manage_doctor_queue.VisitShortInfoResponse{
			Id:        v.Id,
			Name:      v.Name,
			Gender:    v.Gender,
			Age:       v.Age,
			EnterTime: v.EnterTime,
		}
		i++
	}

	var visitInProgress manage_doctor_queue.VisitShortInfoResponse
	if vip := q.VisitInProgress; vip != nil {
		visitInProgress = manage_doctor_queue.VisitShortInfoResponse{
			Id:        vip.Id,
			Name:      vip.Name,
			Gender:    vip.Gender,
			Age:       vip.Age,
			EnterTime: vip.EnterTime,
		}
	}

	repr := manage_doctor_queue.DoctorQueueResponse{
		Visits:          visits,
		VisitInProgress: visitInProgress,
	}
	return repr
}

func (q *DoctorQueue) ToRepr() *manage_doctor_queue.DoctorQueueRepr {
	visits := make(map[string]manage_doctor_queue.VisitShortInfoRepr)
	for k, v := range q.Visits.Queue {
		visits[k] = manage_doctor_queue.VisitShortInfoRepr{
			Id:        v.Id,
			Name:      v.Name,
			Gender:    v.Gender,
			Age:       v.Age,
			EnterTime: v.EnterTime,
		}
	}

	var visitInProgress *manage_doctor_queue.VisitShortInfoRepr
	if vip := q.VisitInProgress; vip != nil {
		visitInProgress = &manage_doctor_queue.VisitShortInfoRepr{
			Id:        vip.Id,
			Name:      vip.Name,
			Gender:    vip.Gender,
			Age:       vip.Age,
			EnterTime: vip.EnterTime,
		}
	}

	repr := &manage_doctor_queue.DoctorQueueRepr{
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

func NewDoctorQueueFromRepr(repr *manage_doctor_queue.DoctorQueueRepr) *DoctorQueue {
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
