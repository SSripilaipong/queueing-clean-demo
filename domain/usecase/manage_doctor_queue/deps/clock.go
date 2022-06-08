package deps

import (
	"time"
)

type IClock interface {
	Now() time.Time
}
