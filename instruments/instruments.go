package instruments

import (
	"github.com/poolpOrg/go-harmony/chords"
	"github.com/poolpOrg/go-harmony/notes"
	"github.com/poolpOrg/go-harmony/scales"
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

func (instrument *Instrument) Notes(noteNames ...string) (chords.Chord, error) {
	noteSequence := make([]notes.Note, 0)
	for _, noteName := range noteNames {
		note, err := notes.Parse(noteName)
		if err != nil {
			return chords.Chord{}, err
		}
		noteSequence = append(noteSequence, *note)
	}
	return chords.FromNotes(noteSequence), nil
}

func (instrument *Instrument) Chord(name string) (chords.Chord, error) {
	chord, err := chords.Parse(name)
	if err != nil {
		return chords.Chord{}, err
	}
	return *chord, nil
}

func (instrument *Instrument) Scale(name string) (scales.Scale, error) {
	scale, err := scales.Parse(name)
	if err != nil {
		return scales.Scale{}, err
	}
	return *scale, nil
}
