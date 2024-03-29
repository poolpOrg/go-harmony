package instruments

import (
	"github.com/poolpOrg/go-harmony/chords"
	"github.com/poolpOrg/go-harmony/intervals"
	"github.com/poolpOrg/go-harmony/naturals"
	"github.com/poolpOrg/go-harmony/notes"
	"github.com/poolpOrg/go-harmony/octaves"
	"github.com/poolpOrg/go-harmony/scales"
	"github.com/poolpOrg/go-harmony/tunings"
)

type Instrument struct{}

func NewInstrument(tuning tunings.Tuning) Instrument {
	return Instrument{}
}

func (instrument *Instrument) Natural(name string) (naturals.Natural, error) {
	natural, err := naturals.FromName(name)
	if err != nil {
		return naturals.Natural{}, err
	}
	return *natural, nil
}

func (instrument *Instrument) Octave(name string) (octaves.Octave, error) {
	octave, err := octaves.FromName(name)
	if err != nil {
		return octaves.Octave{}, err
	}
	return *octave, nil
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

func (instrument *Instrument) Distances(noteNames ...string) ([]intervals.Interval, error) {
	noteSequence := make([]notes.Note, 0)
	for _, noteName := range noteNames {
		note, err := notes.Parse(noteName)
		if err != nil {
			return nil, err
		}
		noteSequence = append(noteSequence, *note)
	}

	intervalsSequence := make([]intervals.Interval, 0)
	for offset, noteName := range noteNames {
		note, err := notes.Parse(noteName)
		if err != nil {
			return nil, err
		}
		if offset == 0 {
			intervalsSequence = append(intervalsSequence, intervals.PerfectUnison)
		} else {
			intervalsSequence = append(intervalsSequence, noteSequence[offset-1].Distance(*note))
		}
	}
	return intervalsSequence, nil
}
