package common

type VisitNotFoundError struct {
}

func (v VisitNotFoundError) Error() string {
	return "VisitNotFoundError"
}

type InvalidVisitDataError struct {
}

func (i InvalidVisitDataError) Error() string {
	return "InvalidVisitDataError"
}

type DuplicateVisitIdError struct {
}

func (d DuplicateVisitIdError) Error() string {
	return "DuplicateVisitIdError"
}
