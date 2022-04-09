package repmax

import (
	"math"

	"github.com/narumiruna/gofit/pkg/types"
)

func CalculateWelday1RM(weight types.Mass, repetitions int) types.Mass {
	w := weight.Float64()
	r := float64(repetitions)

	rm := w * (r*0.0333 + 1.0)

	return types.Mass(rm)
}

func CalculateMayhew1RM(weight types.Mass, repetitions int) types.Mass {
	w := weight.Float64()
	r := float64(repetitions)

	rm := w / (0.522 + 0.419*math.Exp(-0.055*r))

	return types.Mass(rm)
}
func CalculateReynolds1RM(weight types.Mass, repetitions int) types.Mass {
	w := weight.Float64()
	r := float64(repetitions)

	rm := w / (0.4847 + 0.5551*math.Exp(-0.0723*r))

	return types.Mass(rm)
}

func CalculateWathen1RM(weight types.Mass, repetitions int) types.Mass {
	w := weight.Float64()
	r := float64(repetitions)

	rm := w / (0.488 + 0.538*math.Exp(-0.075*r))

	return types.Mass(rm)
}
