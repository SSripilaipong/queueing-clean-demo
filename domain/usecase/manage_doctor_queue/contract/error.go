package contract

type VisitAlreadyExistsError struct {
}

func (v VisitAlreadyExistsError) Error() string {
	return "VisitAlreadyExistsError"
}

type DoctorStillBusyError struct {
}

func (d DoctorStillBusyError) Error() string {
	return "DoctorStillBusyError"
}

type NoVisitInProgressToCompleteError struct {
}

func (n NoVisitInProgressToCompleteError) Error() string {
	return "NoVisitInProgressToCompleteError"
}

type DoctorQueueNotFoundError struct {
}

func (d DoctorQueueNotFoundError) Error() string {
	return "DoctorQueueNotFoundError"
}

type DuplicateDoctorQueueIdError struct {
}

func (d DuplicateDoctorQueueIdError) Error() string {
	return "DuplicateDoctorQueueIdError"
}
