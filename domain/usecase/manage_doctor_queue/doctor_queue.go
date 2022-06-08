package manage_doctor_queue

import (
	"queueing-clean-demo/base"
	derr "queueing-clean-demo/domain/contract"
	"queueing-clean-demo/domain/usecase/manage_doctor_queue/contract"
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
		return contract.VisitAlreadyExistsError{}
	}
	q.Visits.add(visit)
	return nil
}

func (q *DoctorQueue) CallVisit(id string) error {
	if q.VisitInProgress != nil {
		return contract.DoctorStillBusyError{}
	}
	if !q.Visits.hasVisit(id) {
		return derr.VisitNotFoundError{}
	}
	visit := q.Visits.pop(id)
	q.VisitInProgress = &visit
	return nil
}

func (q *DoctorQueue) CompleteDiagnosis() error {
	if q.VisitInProgress == nil {
		return contract.NoVisitInProgressToCompleteError{}
	}
	q.VisitInProgress = nil
	return nil
}

func (q *DoctorQueue) IterateVisits() *visitIterator {
	return q.Visits.iterateVisits()
}

type VisitCollection struct {
	Queue map[string]VisitShortInfo `json:"queue"`
}

func (q *VisitCollection) hasVisit(visitId string) bool {
	_, exists := q.Queue[visitId]
	return exists
}

func (q *VisitCollection) add(shortInfo VisitShortInfo) {
	q.Queue[shortInfo.Id] = shortInfo
}

func (q *VisitCollection) pop(id string) VisitShortInfo {
	visit, exists := q.Queue[id]
	if exists {
		delete(q.Queue, id)
	}
	return visit
}

func (q *VisitCollection) iterateVisits() *visitIterator {
	visits := make([]VisitShortInfo, len(q.Queue))
	for _, v := range q.Queue {
		visits = append(visits, v)
	}
	return &visitIterator{visits: visits}
}

type visitIterator struct {
	index  int
	visits []VisitShortInfo
}

func (i *visitIterator) HasNext() bool {
	return i.index < len(i.visits)
}

func (i *visitIterator) Next() VisitShortInfo {
	index := i.index
	i.index++
	return i.visits[index]
}
