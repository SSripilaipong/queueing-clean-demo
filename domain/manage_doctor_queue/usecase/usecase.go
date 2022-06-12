package usecase

import (
	"queueing-clean-demo/domain"
	. "queueing-clean-demo/domain/manage_doctor_queue"
	"queueing-clean-demo/domain/manage_doctor_queue/internal"
)

func NewManageDoctorQueueUsecase(doctorQueueRepo IDoctorQueueRepo, clock IClock) domain.IManageDoctorQueueUsecase {
	return &_manage_doctor_queue.Usecase{
		DoctorQueueRepo: doctorQueueRepo,
		Clock:           clock,
	}
}
