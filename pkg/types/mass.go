package types

import "fmt"

type Mass Unit

const (
	ZeroMass Mass = 0.0

	Gram     Mass = 1.0
	KiloGram Mass = Gram * 1e+3

	Pound Mass = KiloGram * 0.45359237
)

func (m Mass) String() string {
	return fmt.Sprintf("%.2f kg", m.Float64()/KiloGram.Float64())
}

func (m Mass) Float64() float64 {
	return float64(m)
}
func (m Mass) AddInt(x int64) Mass {
	return Mass(m.Float64() + float64(x))
}

func (m Mass) AddFloat(x float64) Mass {
	return Mass(m.Float64() + float64(x))
}

func (m Mass) MulInt(x int64) Mass {
	return Mass(m.Float64() * float64(x))
}

func (m Mass) MulFloat(x float64) Mass {
	return Mass(m.Float64() * float64(x))
}

func (m Mass) DivInt(x int64) Mass {
	return Mass(m.Float64() / float64(x))
}

func (m Mass) DivFloat(x float64) Mass {
	return Mass(m.Float64() / float64(x))
}
