package usecase

import (
	"queueing-clean-demo/domain/contract"
	"queueing-clean-demo/domain/manage_doctor_queue"
	. "queueing-clean-demo/domain/manage_doctor_queue/contract"
)

func NewManageDoctorQueueUsecase(doctorQueueRepo IDoctorQueueRepo, clock IClock) domain.IManageDoctorQueueUsecase {
	return &_manage_doctor_queue.Usecase{
		DoctorQueueRepo: doctorQueueRepo,
		Clock:           clock,
	}
}
