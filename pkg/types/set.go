package types

type Set struct {
	Exercise    Exercise
	Weight      Mass
	Repetitions int64
}

func (s Set) Volume() Mass {
	return s.Weight.MulInt(s.Repetitions) + s.Exercise.Equipment.Mass
}
