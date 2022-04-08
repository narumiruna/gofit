package types

type SetSlice []Set

func (s SetSlice) Volume() Mass {
	sum := ZeroMass
	for _, set := range s {
		sum += set.Volume()
	}
	return sum
}
