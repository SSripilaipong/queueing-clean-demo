package _manage_doctor_queue

import (
	"queueing-clean-demo/domain/manage_doctor_queue/contract"
	"time"
)

func VisitShortInfoFromPushVisitToDoctorQueueRequest(r manage_doctor_queue.PushVisitToDoctorQueue) VisitShortInfo {
	return VisitShortInfo{
		Id:        r.VisitId,
		Name:      r.PatientName,
		Gender:    r.PatientGender,
		Age:       r.PatientAge,
		EnterTime: time.Time{},
	}
}
