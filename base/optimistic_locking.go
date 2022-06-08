package base

func OptimisticLockingRetry[R any](maxRetries int, f func() (R, error)) (R, error) {
	var r R
	var err error

	for i := 0; i < maxRetries; i++ {
		switch r, err = f(); err.(type) {
		case OptimisticLockFailedError:
			continue
		default:
			return r, err
		}
	}
	return r, OptimisticLockFailedError{"OptimisticLockingRetry exceeds maxRetries"}
}

type OptimisticLockFailedError struct {
	Message string
}

func (o OptimisticLockFailedError) Error() string {
	if o.Message == "" {
		return "OptimisticLockFailedError"
	}
	return o.Message
}
