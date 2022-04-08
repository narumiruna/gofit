package types

import "fmt"

type Equipment struct {
	Name string
	Mass Mass
}

func (e Equipment) String() string {
	return fmt.Sprintf("%v (%v)", e.Name, e.Mass)
}

func Barbell() Equipment {
	return Equipment{Name: "Barbell", Mass: 20 * KiloGram}
}
