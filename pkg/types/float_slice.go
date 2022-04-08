package types

import "gonum.org/v1/gonum/floats"

type Float64Slice []float64

func (s Float64Slice) Sum() float64 {
	return floats.Sum(s)
}
