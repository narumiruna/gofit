package repmax

import (
	"fmt"
	"math"

	"github.com/narumiruna/gofit/pkg/types"
)

type Reynolds struct {
	RepMaxEquation

	Weight      types.Mass
	Repetitions int
}

func (rm *Reynolds) String() string {
	return fmt.Sprintf("%v", rm.OneRepMax())
}

func (rm Reynolds) OneRepMax() types.Mass {
	return rm.Weight.DivFloat(0.5551 * math.Exp(-0.0723*float64(rm.Repetitions)+0.4847))
}
