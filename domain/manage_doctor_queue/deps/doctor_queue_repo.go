package _deps

import (
	. "queueing-clean-demo/domain/manage_doctor_queue/contract"
)

type IDoctorQueueRepo interface {
	Create(queue *DoctorQueueRepr) (*DoctorQueueRepr, error)
	FindByDoctorIdAndUpdate(doctorId string, update func(queue *DoctorQueueRepr) (*DoctorQueueRepr, error)) (*DoctorQueueRepr, error)
	FindByDoctorId(doctorId string) (*DoctorQueueRepr, error)
}
