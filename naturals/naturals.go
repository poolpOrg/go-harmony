package naturals

import (
	"fmt"
)

type natural struct {
	name      string
	pos       int
	semitones int
}
type Natural natural

var naturals = []Natural{
	{name: "C", pos: 0, semitones: 0},
	{name: "D", pos: 1, semitones: 2},
	{name: "E", pos: 2, semitones: 4},
	{name: "F", pos: 3, semitones: 5},
	{name: "G", pos: 4, semitones: 7},
	{name: "A", pos: 5, semitones: 9},
	{name: "B", pos: 6, semitones: 11},
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
	return uint(natural.pos)
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
