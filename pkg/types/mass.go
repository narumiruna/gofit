package types

type Mass Unit

const (
	ZeroMass Mass = 0.0

	Gram     Mass = 1.0
	KiloGram Mass = Gram * 1e+3

	Pound Mass = KiloGram * 0.45359237
)

func (m Mass) Float64() float64 {
	return float64(m)
}

func (m Mass) MulInt(x int64) Mass {
	return Mass(float64(x) * m.Float64())
}
