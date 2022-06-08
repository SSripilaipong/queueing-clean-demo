package contract

type AssessmentAlreadyExistError struct {
}

func (a AssessmentAlreadyExistError) Error() string {
	return "AssessmentAlreadyExistError"
}
