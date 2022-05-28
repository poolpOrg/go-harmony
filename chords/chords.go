package chords

import (
	"fmt"

	"github.com/poolpOrg/go-harmony/intervals"
	"github.com/poolpOrg/go-harmony/notes"
)

type chord struct {
	name  string
	notes []notes.Note
}
type Chord chord

func majorTriad(note *notes.Note) []notes.Note {
	chordNotes := make([]notes.Note, 0)
	chordNotes = append(chordNotes, *note)
	chordNotes = append(chordNotes, *note.Interval(intervals.MajorThird))
	chordNotes = append(chordNotes, *note.Interval(intervals.PerfectFifth))
	return chordNotes
}

func minorTriad(note *notes.Note) []notes.Note {
	chordNotes := make([]notes.Note, 0)
	chordNotes = append(chordNotes, *note)
	chordNotes = append(chordNotes, *note.Interval(intervals.MinorThird))
	chordNotes = append(chordNotes, *note.Interval(intervals.PerfectFifth))
	return chordNotes
}

func augmentedTriad(note *notes.Note) []notes.Note {
	chordNotes := make([]notes.Note, 0)
	chordNotes = append(chordNotes, *note)
	chordNotes = append(chordNotes, *note.Interval(intervals.MajorThird))
	chordNotes = append(chordNotes, *note.Interval(intervals.AugmentedFifth))
	return chordNotes
}

func diminishedTriad(note *notes.Note) []notes.Note {
	chordNotes := make([]notes.Note, 0)
	chordNotes = append(chordNotes, *note)
	chordNotes = append(chordNotes, *note.Interval(intervals.MinorThird))
	chordNotes = append(chordNotes, *note.Interval(intervals.DiminishedFifth))
	return chordNotes
}

func dominantSeventh(note *notes.Note) []notes.Note {
	chordNotes := make([]notes.Note, 0)
	chordNotes = append(chordNotes, *note)
	chordNotes = append(chordNotes, *note.Interval(intervals.MajorThird))
	chordNotes = append(chordNotes, *note.Interval(intervals.PerfectFifth))
	chordNotes = append(chordNotes, *note.Interval(intervals.MinorSeventh))
	return chordNotes
}

func majorSeventh(note *notes.Note) []notes.Note {
	chordNotes := make([]notes.Note, 0)
	chordNotes = append(chordNotes, *note)
	chordNotes = append(chordNotes, *note.Interval(intervals.MajorThird))
	chordNotes = append(chordNotes, *note.Interval(intervals.PerfectFifth))
	chordNotes = append(chordNotes, *note.Interval(intervals.MajorSeventh))
	return chordNotes
}

func minorMajorSeventh(note *notes.Note) []notes.Note {
	chordNotes := make([]notes.Note, 0)
	chordNotes = append(chordNotes, *note)
	chordNotes = append(chordNotes, *note.Interval(intervals.MinorThird))
	chordNotes = append(chordNotes, *note.Interval(intervals.PerfectFifth))
	chordNotes = append(chordNotes, *note.Interval(intervals.MajorSeventh))
	return chordNotes
}

func minorSeventh(note *notes.Note) []notes.Note {
	chordNotes := make([]notes.Note, 0)
	chordNotes = append(chordNotes, *note)
	chordNotes = append(chordNotes, *note.Interval(intervals.MinorThird))
	chordNotes = append(chordNotes, *note.Interval(intervals.PerfectFifth))
	chordNotes = append(chordNotes, *note.Interval(intervals.MinorSeventh))
	return chordNotes
}

func augmentedMajorSeventh(note *notes.Note) []notes.Note {
	chordNotes := make([]notes.Note, 0)
	chordNotes = append(chordNotes, *note)
	chordNotes = append(chordNotes, *note.Interval(intervals.MajorThird))
	chordNotes = append(chordNotes, *note.Interval(intervals.AugmentedFifth))
	chordNotes = append(chordNotes, *note.Interval(intervals.MajorSeventh))
	return chordNotes
}

func augmentedSeventh(note *notes.Note) []notes.Note {
	chordNotes := make([]notes.Note, 0)
	chordNotes = append(chordNotes, *note)
	chordNotes = append(chordNotes, *note.Interval(intervals.MajorThird))
	chordNotes = append(chordNotes, *note.Interval(intervals.AugmentedFifth))
	chordNotes = append(chordNotes, *note.Interval(intervals.MinorSeventh))
	return chordNotes
}

func halfDiminishedSeventh(note *notes.Note) []notes.Note {
	chordNotes := make([]notes.Note, 0)
	chordNotes = append(chordNotes, *note)
	chordNotes = append(chordNotes, *note.Interval(intervals.MinorThird))
	chordNotes = append(chordNotes, *note.Interval(intervals.DiminishedFifth))
	chordNotes = append(chordNotes, *note.Interval(intervals.MinorSeventh))
	return chordNotes
}

func diminishedSeventh(note *notes.Note) []notes.Note {
	chordNotes := make([]notes.Note, 0)
	chordNotes = append(chordNotes, *note)
	chordNotes = append(chordNotes, *note.Interval(intervals.MinorThird))
	chordNotes = append(chordNotes, *note.Interval(intervals.DiminishedFifth))
	chordNotes = append(chordNotes, *note.Interval(intervals.DiminishedSeventh))
	return chordNotes
}

func dominantSeventhFlatFive(note *notes.Note) []notes.Note {
	chordNotes := make([]notes.Note, 0)
	chordNotes = append(chordNotes, *note)
	chordNotes = append(chordNotes, *note.Interval(intervals.MajorThird))
	chordNotes = append(chordNotes, *note.Interval(intervals.DiminishedFifth))
	chordNotes = append(chordNotes, *note.Interval(intervals.MinorSeventh))
	return chordNotes
}

func Parse(chord string) (*Chord, error) {
	if len(chord) == 0 {
		return nil, fmt.Errorf("empty chord")
	}

	noteEnd := int(1)
	for _, r := range chord[1:] {
		if r == '#' || r == 'b' {
			noteEnd++
			continue
		}
		break
	}

	noteName := chord[0:noteEnd]
	chordName := chord[noteEnd:]

	n, err := notes.Parse(noteName)
	if err != nil {
		return nil, err
	}

	var chordNotes []notes.Note
	switch chordName {
	case "":
		fallthrough
	case "M":
		fallthrough
	case "maj":
		// major triad
		chordNotes = majorTriad(n)

	case "-":
		fallthrough
	case "m":
		fallthrough
	case "min":
		// minor triad
		chordNotes = minorTriad(n)

	case "+":
		fallthrough
	case "aug":
		fallthrough
	case "M#5":
		fallthrough
	case "M+5":
		// augmented triad
		chordNotes = augmentedTriad(n)

	case "o":
		fallthrough
	case "dim":
		fallthrough
	case "mb5":
		fallthrough
	case "mo5":
		// diminished triad
		chordNotes = diminishedTriad(n)

	case "7":
		// dominant seventh
		chordNotes = dominantSeventh(n)

	case "M7":
		fallthrough
	case "Ma7":
		fallthrough
	case "maj7":
		// major seventh
		chordNotes = majorSeventh(n)

	case "mM7":
		fallthrough
	case "m#7":
		fallthrough
	case "-M7":
		fallthrough
	case "minmaj7":
		// minor-major seventh
		chordNotes = minorMajorSeventh(n)

	case "m7":
		fallthrough
	case "-7":
		fallthrough
	case "min7":
		// minor seventh
		chordNotes = minorSeventh(n)

	case "+M7":
		fallthrough
	case "augmaj7":
		fallthrough
	case "M7#5":
		fallthrough
	case "M7+5":
		// augmented-major seventh
		chordNotes = augmentedMajorSeventh(n)

	case "+7":
		fallthrough
	case "aug7":
		fallthrough
	case "7#5":
		fallthrough
	case "7+5":
		// augmented seventh
		chordNotes = augmentedSeventh(n)

	case "ø":
		fallthrough
	case "ø7":
		fallthrough
	case "min7dim5":
		fallthrough
	case "m7b5":
		fallthrough
	case "m7o5":
		fallthrough
	case "-7b5":
		fallthrough
	case "-7o5":
		// half-diminished seventh
		chordNotes = halfDiminishedSeventh(n)

	case "o7":
		fallthrough
	case "dim7":
		// diminished seventh
		chordNotes = diminishedSeventh(n)

	case "7b5":
		fallthrough
	case "7dim5":
		// dominant seventh flat five
		chordNotes = dominantSeventhFlatFive(n)

	default:
		return nil, fmt.Errorf("unknown chord name: %s", chordName)
	}

	return &Chord{
		name:  chord,
		notes: chordNotes,
	}, nil
}

func (chord *Chord) Name() string {
	return chord.notes[0].Name()
}

func (chord *Chord) Notes() []notes.Note {
	return chord.notes
}
