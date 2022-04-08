package types

type SetSlice []Set

func (s SetSlice) Volumes() (o MassSlice) {
	for _, set := range s {
		o = append(o, set.Volume())
	}
	return o
}

func (s SetSlice) Volume() Mass {
	return Mass(s.Volumes().ToFloats().Sum())
}
