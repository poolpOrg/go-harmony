package instruments

import (
	"github.com/poolpOrg/go-harmony/chords"
	"github.com/poolpOrg/go-harmony/notes"
	"github.com/poolpOrg/go-harmony/tunings"
)

type Instrument struct{}

func NewInstrument(tuning tunings.Tuning) Instrument {
	return Instrument{}
}

func (instrument *Instrument) Note(name string) (notes.Note, error) {
	note, err := notes.Parse(name)
	if err != nil {
		return notes.Note{}, err
	}
	return *note, nil
}

func (instrument *Instrument) Chord(name string) (chords.Chord, error) {
	chord, err := chords.Parse(name)
	if err != nil {
		return chords.Chord{}, err
	}
	return *chord, nil
}
