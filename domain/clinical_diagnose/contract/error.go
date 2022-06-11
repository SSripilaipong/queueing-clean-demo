package clinical_diagnose

type AssessmentAlreadyExistError struct {
}

func (a AssessmentAlreadyExistError) Error() string {
	return "AssessmentAlreadyExistError"
}
