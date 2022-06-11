package _manage_doctor_queue

import (
	common "queueing-clean-demo/domain/common/contract"
	. "queueing-clean-demo/domain/manage_doctor_queue/contract"
	"queueing-clean-demo/domain/manage_doctor_queue/deps"
)

type Usecase struct {
	DoctorQueueRepo _deps.IDoctorQueueRepo
	Clock           _deps.IClock
}

func (u *Usecase) CreateDoctorQueue(request CreateDoctorQueue) (DoctorQueueResponse, error) {

	queue, err := NewDoctorQueue(request.DoctorId)
	if err == nil {
		queue, err = u.DoctorQueueRepo.Create(queue)
	}

	switch err.(type) {
	case DuplicateDoctorQueueIdError:
		return DoctorQueueResponse{}, err
	case nil:
		return DoctorQueueResponseFromDoctorQueue(queue), nil
	}
	panic(err)
}

func (u *Usecase) PushVisit(request PushVisitToDoctorQueue) (DoctorQueueResponse, error) {

	queue, err := u.DoctorQueueRepo.FindByDoctorIdAndUpdate(request.DoctorId, func(queue *DoctorQueue) (*DoctorQueue, error) {

		info := VisitShortInfoFromPushVisitToDoctorQueueRequest(request)
		info.EnterTime = u.Clock.Now()

		if err := queue.PushVisit(info); err != nil {
			return nil, err
		}
		return queue, nil
	})

	switch err.(type) {
	case DoctorQueueNotFoundError:
		return DoctorQueueResponse{}, err
	case VisitAlreadyExistsError:
		return DoctorQueueResponse{}, err
	case nil:
		return DoctorQueueResponseFromDoctorQueue(queue), nil
	}
	panic(err)
}

func (u *Usecase) CallVisit(request CallVisitFromDoctorQueue) (DoctorQueueResponse, error) {

	queue, err := u.DoctorQueueRepo.FindByDoctorIdAndUpdate(request.DoctorId, func(queue *DoctorQueue) (*DoctorQueue, error) {
		if err := queue.CallVisit(request.VisitId); err != nil {
			return nil, err
		}
		return queue, nil
	})

	switch err.(type) {
	case DoctorQueueNotFoundError:
		return DoctorQueueResponse{}, err
	case DoctorStillBusyError:
		return DoctorQueueResponse{}, err
	case common.VisitNotFoundError:
		return DoctorQueueResponse{}, err
	case nil:
		return DoctorQueueResponseFromDoctorQueue(queue), nil
	}
	panic(err)

}

func (u *Usecase) CompleteDiagnosis(request CompleteDiagnosis) (DoctorQueueResponse, error) {

	queue, err := u.DoctorQueueRepo.FindByDoctorIdAndUpdate(request.DoctorId, func(queue *DoctorQueue) (*DoctorQueue, error) {
		if err := queue.CompleteDiagnosis(); err != nil {
			return nil, err
		}
		return queue, nil
	})

	switch err.(type) {
	case DoctorQueueNotFoundError:
		return DoctorQueueResponse{}, err
	case NoVisitInProgressToCompleteError:
		return DoctorQueueResponse{}, err
	case nil:
		return DoctorQueueResponseFromDoctorQueue(queue), nil
	}
	panic(err)

}

func (u *Usecase) CheckVisits(request CheckVisits) (DoctorQueueResponse, error) {

	queue, err := u.DoctorQueueRepo.FindByDoctorId(request.DoctorId)

	switch err.(type) {
	case DoctorQueueNotFoundError:
		return DoctorQueueResponse{}, err
	case nil:
		return DoctorQueueResponseFromDoctorQueue(queue), nil
	}
	panic(err)
}
