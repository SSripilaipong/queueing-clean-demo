package manage_doctor_queue

import (
	dContract "queueing-clean-demo/domain/contract"
	"queueing-clean-demo/domain/usecase/manage_doctor_queue/contract"
	"queueing-clean-demo/domain/usecase/manage_doctor_queue/deps"
)

type Usecase struct {
	DoctorQueueRepo deps.IDoctorQueueRepo
	Clock           deps.IClock
}

func (u *Usecase) CreateDoctorQueue(request contract.CreateDoctorQueue) (contract.DoctorQueueResponse, error) {

	queue, err := NewDoctorQueue(request.DoctorId)
	if err == nil {
		queue, err = u.DoctorQueueRepo.Create(queue)
	}

	switch err.(type) {
	case contract.DuplicateDoctorQueueIdError:
		return contract.DoctorQueueResponse{}, err
	case nil:
		return DoctorQueueResponseFromDoctorQueue(queue), nil
	}
	panic(err)
}

func (u *Usecase) PushVisit(request contract.PushVisitToDoctorQueue) (contract.DoctorQueueResponse, error) {

	queue, err := u.DoctorQueueRepo.FindByDoctorIdAndUpdate(request.DoctorId, func(queue *DoctorQueue) (*DoctorQueue, error) {

		info := VisitShortInfoFromPushVisitToDoctorQueueRequest(request)
		info.EnterTime = u.Clock.Now()

		if err := queue.PushVisit(info); err != nil {
			return nil, err
		}
		return queue, nil
	})

	switch err.(type) {
	case contract.DoctorQueueNotFoundError:
		return contract.DoctorQueueResponse{}, err
	case contract.VisitAlreadyExistsError:
		return contract.DoctorQueueResponse{}, err
	case nil:
		return DoctorQueueResponseFromDoctorQueue(queue), nil
	}
	panic(err)
}

func (u *Usecase) CallVisit(request contract.CallVisitFromDoctorQueue) (contract.DoctorQueueResponse, error) {

	queue, err := u.DoctorQueueRepo.FindByDoctorIdAndUpdate(request.DoctorId, func(queue *DoctorQueue) (*DoctorQueue, error) {
		if err := queue.CallVisit(request.VisitId); err != nil {
			return nil, err
		}
		return queue, nil
	})

	switch err.(type) {
	case contract.DoctorQueueNotFoundError:
		return contract.DoctorQueueResponse{}, err
	case contract.DoctorStillBusyError:
		return contract.DoctorQueueResponse{}, err
	case dContract.VisitNotFoundError:
		return contract.DoctorQueueResponse{}, err
	case nil:
		return DoctorQueueResponseFromDoctorQueue(queue), nil
	}
	panic(err)

}

func (u *Usecase) CompleteDiagnosis(request contract.CompleteDiagnosis) (contract.DoctorQueueResponse, error) {

	queue, err := u.DoctorQueueRepo.FindByDoctorIdAndUpdate(request.DoctorId, func(queue *DoctorQueue) (*DoctorQueue, error) {
		if err := queue.CompleteDiagnosis(); err != nil {
			return nil, err
		}
		return queue, nil
	})

	switch err.(type) {
	case contract.DoctorQueueNotFoundError:
		return contract.DoctorQueueResponse{}, err
	case contract.NoVisitInProgressToCompleteError:
		return contract.DoctorQueueResponse{}, err
	case nil:
		return DoctorQueueResponseFromDoctorQueue(queue), nil
	}
	panic(err)

}

func (u *Usecase) CheckVisits(request contract.CheckVisits) (contract.DoctorQueueResponse, error) {

	queue, err := u.DoctorQueueRepo.FindByDoctorId(request.DoctorId)

	switch err.(type) {
	case contract.DoctorQueueNotFoundError:
		return contract.DoctorQueueResponse{}, err
	case nil:
		return DoctorQueueResponseFromDoctorQueue(queue), nil
	}
	panic(err)
}
