package types

type Unit float64

type Mass Unit

const (
	Gram     Mass = 1.0
	KiloGram Mass = Gram * 1e+3

	// US
	Pound Mass = KiloGram * 0.45359237
)
