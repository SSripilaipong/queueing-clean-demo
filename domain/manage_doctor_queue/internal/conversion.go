package internal

import (
	. "queueing-clean-demo/domain/manage_doctor_queue"
	"time"
)

func VisitShortInfoFromPushVisitToDoctorQueueRequest(r PushVisitToDoctorQueue) VisitShortInfo {
	return VisitShortInfo{
		Id:        r.VisitId,
		Name:      r.PatientName,
		Gender:    r.PatientGender,
		Age:       r.PatientAge,
		EnterTime: time.Time{},
	}
}
