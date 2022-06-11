package _manage_doctor_queue

import (
	"queueing-clean-demo/domain/manage_doctor_queue/contract"
	"time"
)

func DoctorQueueResponseFromDoctorQueue(queue *DoctorQueue) manage_doctor_queue.DoctorQueueResponse {
	var visits []manage_doctor_queue.VisitShortInfoResponse
	visitIterator := queue.IterateVisits()
	for visitIterator.HasNext() {
		visitRaw := visitIterator.Next()
		visits = append(visits, manage_doctor_queue.VisitShortInfoResponse{
			Id:        visitRaw.Id,
			Name:      visitRaw.Name,
			Gender:    visitRaw.Gender,
			Age:       visitRaw.Age,
			EnterTime: visitRaw.EnterTime,
		})
	}
	return manage_doctor_queue.DoctorQueueResponse{Visits: visits}
}

func VisitShortInfoFromPushVisitToDoctorQueueRequest(r manage_doctor_queue.PushVisitToDoctorQueue) VisitShortInfo {
	return VisitShortInfo{
		Id:        r.VisitId,
		Name:      r.PatientName,
		Gender:    r.PatientGender,
		Age:       r.PatientAge,
		EnterTime: time.Time{},
	}
}
