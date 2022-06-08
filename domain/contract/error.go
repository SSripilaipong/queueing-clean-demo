package contract

type VisitNotFoundError struct {
}

func (v VisitNotFoundError) Error() string {
	return "VisitNotFoundError"
}

type DuplicateVisitIdError struct {
}

func (d DuplicateVisitIdError) Error() string {
	return "DuplicateVisitIdError"
}

type InvalidVisitDataError struct {
}

func (i InvalidVisitDataError) Error() string {
	return "InvalidVisitDataError"
}
