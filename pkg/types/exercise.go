package types

type Exercise struct {
	Name      string
	Equipment Equipment
	Sets      SetSlice
}

func (e Exercise) Volume(withEquipment bool) Mass {
	v := e.Sets.Volume()

	if withEquipment {
		v += e.Equipment.Weight
	}

	return v
}
