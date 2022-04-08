package types

import "time"

type Workout struct {
	Exercises []Exercise
	Since     time.Time
	Until     time.Time
}
