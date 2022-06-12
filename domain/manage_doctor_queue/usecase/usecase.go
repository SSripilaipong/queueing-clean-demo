package usecase

import (
	"queueing-clean-demo/domain"
	. "queueing-clean-demo/domain/manage_doctor_queue"
	"queueing-clean-demo/domain/manage_doctor_queue/internal"
)

func NewManageDoctorQueueUsecase(doctorQueueRepo IDoctorQueueRepo, clock IClock) domain.IManageDoctorQueueUsecase {
	return &internal.Usecase{
		DoctorQueueRepo: doctorQueueRepo,
		Clock:           clock,
	}
}
