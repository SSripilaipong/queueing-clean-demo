package usecase

import (
	"queueing-clean-demo/domain/contract"
	"queueing-clean-demo/domain/manage_doctor_queue"
	"queueing-clean-demo/domain/manage_doctor_queue/deps"
)

func NewManageDoctorQueueUsecase(doctorQueueRepo _deps.IDoctorQueueRepo, clock _deps.IClock) domain.IManageDoctorQueueUsecase {
	return &_manage_doctor_queue.Usecase{
		DoctorQueueRepo: doctorQueueRepo,
		Clock:           clock,
	}
}
