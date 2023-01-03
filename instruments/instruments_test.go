package instruments

import (
	"testing"

	"github.com/poolpOrg/go-harmony/naturals"
	"github.com/poolpOrg/go-harmony/octaves"
	"github.com/poolpOrg/go-harmony/tunings"
)

func TestInstruments(t *testing.T) {
}

func TestInstruments_NewInstrument(t *testing.T) {
	tuningsList := tunings.Tunings()
	for _, tuning := range tuningsList {
		_ = NewInstrument(tuning)
	}
}

func TestInstruments_Natural(t *testing.T) {
	instrument := NewInstrument(tunings.A440)

	naturalsList := naturals.Naturals()
	for _, natural := range naturalsList {
		got, err := instrument.Natural(natural.Name())
		if err != nil {
			t.Fatalf(`instrument.Natural(%s) = error(%s)`, natural.Name(), err)
		}
		if got.Name() != natural.Name() {
			t.Fatalf(`instrument.Natural(%s) = %s, want %s`, natural.Name(), got.Name(), natural.Name())
		}
	}
}

func TestInstruments_Octave(t *testing.T) {
	instrument := NewInstrument(tunings.A440)

	octavesList := octaves.Octaves()
	for _, octave := range octavesList {
		got, err := instrument.Octave(octave.Name())
		if err != nil {
			t.Fatalf(`instrument.Octave(%s) = error(%s)`, octave.Name(), err)
		}
		if got.Name() != octave.Name() {
			t.Fatalf(`instrument.Octave(%s) = %s, want %s`, octave.Name(), got.Name(), octave.Name())
		}
	}
}

func TestInstruments_Note(t *testing.T) {
	instrument := NewInstrument(tunings.A440)

	naturalsList := naturals.Naturals()
	for _, natural := range naturalsList {
		got, err := instrument.Note(natural.Name())
		if err != nil {
			if err != nil {
				t.Fatalf(`instrument.Note(%s) = error(%s)`, natural.Name(), err)
			}
			if got.Name() != "C" {
				t.Fatalf(`instrument.Note(%s) = %s, want %s`, natural.Name(), got.Name(), natural.Name())
			}
		}
	}
}

func TestInstruments_Notes(t *testing.T) {
}

func TestInstruments_Chord(t *testing.T) {
	instrument := NewInstrument(tunings.A440)

	naturalsList := naturals.Naturals()
	for _, natural := range naturalsList {
		got, err := instrument.Chord(natural.Name())
		if err != nil {
			if err != nil {
				t.Fatalf(`instrument.Chord(%s) = error(%s)`, natural.Name(), err)
			}
			if got.Name() != "C" {
				t.Fatalf(`instrument.Chord(%s) = %s, want %s`, natural.Name(), got.Name(), natural.Name())
			}
		}
	}
}

func TestInstruments_Scale(t *testing.T) {
}

func TestInstruments_Distances(t *testing.T) {
}
