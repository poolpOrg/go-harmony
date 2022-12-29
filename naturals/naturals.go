package naturals

import (
	"fmt"
)

type natural struct {
	name      string
	position  int
	semitones int
}
type Natural natural

var (
	C Natural = Natural{name: "C", position: 0, semitones: 0}
	D Natural = Natural{name: "D", position: 1, semitones: 2}
	E Natural = Natural{name: "E", position: 2, semitones: 4}
	F Natural = Natural{name: "F", position: 3, semitones: 5}
	G Natural = Natural{name: "G", position: 4, semitones: 7}
	A Natural = Natural{name: "A", position: 5, semitones: 9}
	B Natural = Natural{name: "B", position: 6, semitones: 11}
)

var naturals = []Natural{
	C, D, E, F, G, A, B,
}

func Parse(name string) (*Natural, error) {
	for _, element := range naturals {
		if element.name == name {
			return &element, nil
		}
	}
	return nil, fmt.Errorf("bad name (%s): should be 'C', 'D', 'E', 'F', 'G', 'A' or 'B'", name)
}

func (natural *Natural) Name() string {
	return natural.name
}

func (natural *Natural) Position() uint {
	return uint(natural.position)
}

func (natural *Natural) Semitones() uint {
	return uint(natural.semitones)
}

func (natural *Natural) Previous() *Natural {
	if natural.Position() == 0 {
		return &naturals[len(naturals)-1]
	} else {
		return &naturals[natural.Position()-1]
	}
}

func (natural *Natural) Next() *Natural {
	return &naturals[int(natural.Position()+1)%len(naturals)]
}
