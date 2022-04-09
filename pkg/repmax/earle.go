package repmax

import (
	"fmt"

	"github.com/narumiruna/gofit/pkg/types"
)

type Welday struct {
	RepMaxEquation

	Weight      types.Mass
	Repetitions int
}

func (r *Welday) String() string {
	return fmt.Sprintf("%v", r.OneRepMax())
}

func (r *Welday) OneRepMax() types.Mass {
	return r.Weight.MulFloat(float64(r.Repetitions)*0.033 + 1)
}

func (r *Welday) RepsMax(repetitions int) types.Mass {
	return r.OneRepMax().DivFloat(float64(repetitions)*0.033 + 1.0)
}
