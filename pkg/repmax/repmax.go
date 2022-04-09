package repmax

import (
	"github.com/narumiruna/gofit/pkg/types"
)

type RepMaxEquation interface {
	OneRepMax() types.Mass
	RepsMax(reps int) types.Mass
}
