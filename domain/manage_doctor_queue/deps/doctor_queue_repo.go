package _deps

import (
	. "queueing-clean-demo/domain/manage_doctor_queue"
)

type IDoctorQueueRepo interface {
	Create(queue *DoctorQueue) (*DoctorQueue, error)
	FindByDoctorIdAndUpdate(doctorId string, update func(queue *DoctorQueue) (*DoctorQueue, error)) (*DoctorQueue, error)
	FindByDoctorId(doctorId string) (*DoctorQueue, error)
}
