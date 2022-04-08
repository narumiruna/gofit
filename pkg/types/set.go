package types

type Set struct {
	Weight      Mass
	Repetitions int64
}

func (s Set) Volume() Mass {
	return s.Weight.MulInt(s.Repetitions)
}

type SetSlice []Set

func (s SetSlice) Volume() Mass {
	sum := ZeroMass
	for _, set := range s {
		sum += set.Volume()
	}
	return sum
}
