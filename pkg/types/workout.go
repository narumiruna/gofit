package types

import "time"

type Workout struct {
	Since time.Time
	Until time.Time
	Sets  SetSlice
}
