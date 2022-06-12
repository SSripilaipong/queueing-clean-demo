package mongodb

type AggregateNotFoundError struct {
}

func (e AggregateNotFoundError) Error() string {
	return "AggregateNotFoundError"
}

type DuplicateIdError struct {
}

func (e DuplicateIdError) Error() string {
	return "DuplicateIdError"
}
