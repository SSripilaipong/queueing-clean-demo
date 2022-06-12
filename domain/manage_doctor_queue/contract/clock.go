package manage_doctor_queue

import (
	"time"
)

type IClock interface {
	Now() time.Time
}
