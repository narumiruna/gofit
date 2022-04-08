package types

import "fmt"

type Exercise struct {
	Name          string
	Equipment     Equipment
	WorkedMuscles []Muscle
}

func NewExercise(name string, equipment Equipment) Exercise {
	e := Exercise{
		Name:      name,
		Equipment: equipment,
	}
	return e
}

func BarbellSquat() Exercise {
	return NewExercise("Squat", Barbell())
}

func (e Exercise) String() string {
	return fmt.Sprintf("%v %v", e.Equipment.Name, e.Name)
}
