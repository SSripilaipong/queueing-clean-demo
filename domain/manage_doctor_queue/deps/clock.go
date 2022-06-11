package _deps

import (
	"time"
)

type IClock interface {
	Now() time.Time
}
