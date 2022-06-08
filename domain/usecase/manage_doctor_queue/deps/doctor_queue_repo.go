package deps

import (
	"queueing-clean-demo/domain/usecase/manage_doctor_queue"
)

type IDoctorQueueRepo interface {
	Create(queue *manage_doctor_queue.DoctorQueue) (*manage_doctor_queue.DoctorQueue, error)
	FindByDoctorIdAndUpdate(doctorId string, update func(queue *manage_doctor_queue.DoctorQueue) (*manage_doctor_queue.DoctorQueue, error)) (*manage_doctor_queue.DoctorQueue, error)
	FindByDoctorId(doctorId string) (*manage_doctor_queue.DoctorQueue, error)
}
