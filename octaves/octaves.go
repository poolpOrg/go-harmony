package octaves

import (
	"fmt"
)

type octave struct {
	position int8
}
type Octave octave

var (
	C0 Octave = Octave{position: 0}
	C1 Octave = Octave{position: 1}
	C2 Octave = Octave{position: 2}
	C3 Octave = Octave{position: 3}
	C4 Octave = Octave{position: 4}
	C5 Octave = Octave{position: 5}
	C6 Octave = Octave{position: 6}
	C7 Octave = Octave{position: 7}
)

var octaves = []Octave{
	C0, C1, C2, C3, C4, C5, C6, C7,
}

func Octaves() []Octave {
	return octaves
}

func FromName(name string) (*Octave, error) {
	for _, element := range octaves {
		if fmt.Sprintf("C%d", element.position) == name {
			return &element, nil
		}
	}
	return nil, fmt.Errorf("bad name (%s): should be 'C0', 'C1', 'C2', 'C3', 'C4', 'C5', 'C6' or 'C7'", name)
}

func FromPosition(position int8) (*Octave, error) {
	for _, element := range octaves {
		if element.position == position {
			return &element, nil
		}
	}
	return nil, fmt.Errorf("bad position (%d): should be '0', '1', '2', '3', '4', '5', '6' or '7'", position)
}

func (octave *Octave) Name() string {
	return fmt.Sprintf("C%d", octave.position)
}

func (octave *Octave) Position() int8 {
	return octave.position
}

func (octave *Octave) Previous() *Octave {
	if octave.position == 0 {
		return nil
	} else {
		return &Octave{position: octave.position - 1}
	}
}

func (octave *Octave) Next() *Octave {
	if octave.position == int8(len(octaves)-1) {
		return nil
	} else {
		return &Octave{position: octave.position + 1}
	}
}

func (octave *Octave) Add(value int8) *Octave {
	if octave.position+value > int8(len(octaves)-1) {
		return nil
	} else {
		return &Octave{position: octave.position + value}
	}
}

func (octave *Octave) Substract(value int8) *Octave {
	if value > octave.position {
		return nil
	} else {
		return &Octave{position: octave.position - value}
	}
}
