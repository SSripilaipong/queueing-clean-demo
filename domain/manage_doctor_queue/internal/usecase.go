package _manage_doctor_queue

import (
	"queueing-clean-demo/domain/common/contract"
	. "queueing-clean-demo/domain/manage_doctor_queue"
)

type Usecase struct {
	DoctorQueueRepo IDoctorQueueRepo
	Clock           IClock
}

func (u *Usecase) CreateDoctorQueue(request CreateDoctorQueue) (DoctorQueueResponse, error) {

	queue, err := NewDoctorQueue(request.DoctorId)
	if err == nil {
		_, err = u.DoctorQueueRepo.Create(queue.ToRepr())
	}

	switch err.(type) {
	case DuplicateDoctorQueueIdError:
		return DoctorQueueResponse{}, err
	case nil:
		return queue.ToResponse(), nil
	}
	panic(err)
}

func (u *Usecase) PushVisit(request PushVisitToDoctorQueue) (DoctorQueueResponse, error) {

	repr, err := u.DoctorQueueRepo.FindByDoctorIdAndUpdate(request.DoctorId, func(repr *DoctorQueueRepr) (*DoctorQueueRepr, error) {

		info := VisitShortInfoFromPushVisitToDoctorQueueRequest(request)
		info.EnterTime = u.Clock.Now()

		queue := NewDoctorQueueFromRepr(repr)
		if err := queue.PushVisit(info); err != nil {
			return nil, err
		}
		return queue.ToRepr(), nil
	})

	switch err.(type) {
	case DoctorQueueNotFoundError:
		return DoctorQueueResponse{}, err
	case VisitAlreadyExistsError:
		return DoctorQueueResponse{}, err
	case nil:
		return NewDoctorQueueFromRepr(repr).ToResponse(), nil
	}
	panic(err)
}

func (u *Usecase) CallVisit(request CallVisitFromDoctorQueue) (DoctorQueueResponse, error) {

	repr, err := u.DoctorQueueRepo.FindByDoctorIdAndUpdate(request.DoctorId, func(repr *DoctorQueueRepr) (*DoctorQueueRepr, error) {
		queue := NewDoctorQueueFromRepr(repr)
		if err := queue.CallVisit(request.VisitId); err != nil {
			return nil, err
		}
		return queue.ToRepr(), nil
	})

	switch err.(type) {
	case DoctorQueueNotFoundError:
		return DoctorQueueResponse{}, err
	case DoctorStillBusyError:
		return DoctorQueueResponse{}, err
	case common.VisitNotFoundError:
		return DoctorQueueResponse{}, err
	case nil:
		return NewDoctorQueueFromRepr(repr).ToResponse(), nil
	}
	panic(err)

}

func (u *Usecase) CompleteDiagnosis(request CompleteDiagnosis) (DoctorQueueResponse, error) {

	repr, err := u.DoctorQueueRepo.FindByDoctorIdAndUpdate(request.DoctorId, func(repr *DoctorQueueRepr) (*DoctorQueueRepr, error) {
		queue := NewDoctorQueueFromRepr(repr)
		if err := queue.CompleteDiagnosis(); err != nil {
			return nil, err
		}
		return queue.ToRepr(), nil
	})

	switch err.(type) {
	case DoctorQueueNotFoundError:
		return DoctorQueueResponse{}, err
	case NoVisitInProgressToCompleteError:
		return DoctorQueueResponse{}, err
	case nil:
		return NewDoctorQueueFromRepr(repr).ToResponse(), nil
	}
	panic(err)

}

func (u *Usecase) CheckVisits(request CheckVisits) (DoctorQueueResponse, error) {

	repr, err := u.DoctorQueueRepo.FindByDoctorId(request.DoctorId)

	switch err.(type) {
	case DoctorQueueNotFoundError:
		return DoctorQueueResponse{}, err
	case nil:
		return NewDoctorQueueFromRepr(repr).ToResponse(), nil
	}
	panic(err)
}
