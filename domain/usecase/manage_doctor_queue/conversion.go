package manage_doctor_queue

import (
	"queueing-clean-demo/domain/usecase/manage_doctor_queue/contract"
	"time"
)

func DoctorQueueResponseFromDoctorQueue(queue *DoctorQueue) contract.DoctorQueueResponse {
	var visits []contract.VisitShortInfoResponse
	visitIterator := queue.IterateVisits()
	for visitIterator.HasNext() {
		visitRaw := visitIterator.Next()
		visits = append(visits, contract.VisitShortInfoResponse{
			Id:        visitRaw.Id,
			Name:      visitRaw.Name,
			Gender:    visitRaw.Gender,
			Age:       visitRaw.Age,
			EnterTime: visitRaw.EnterTime,
		})
	}
	return contract.DoctorQueueResponse{Visits: visits}
}

func VisitShortInfoFromPushVisitToDoctorQueueRequest(r contract.PushVisitToDoctorQueue) VisitShortInfo {
	return VisitShortInfo{
		Id:        r.VisitId,
		Name:      r.PatientName,
		Gender:    r.PatientGender,
		Age:       r.PatientAge,
		EnterTime: time.Time{},
	}
}
