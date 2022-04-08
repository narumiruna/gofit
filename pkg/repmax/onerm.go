package repmax

import (
	"github.com/narumiruna/gofit/pkg/types"
)

func CalculateOneRepMax(mass types.Mass, reps int64) types.Mass {
	return mass.MulInt(reps).MulFloat(0.033) + mass
}
