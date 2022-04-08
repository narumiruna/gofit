package types

type MassSlice []Mass

func (s MassSlice) ToFloats() (o Float64Slice) {
	for _, v := range s {
		o = append(o, v.Float64())
	}
	return o
}
