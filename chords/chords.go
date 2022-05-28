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

func majorNinth(note *notes.Note) []notes.Note {
	chordNotes := majorSeventh(note)
	chordNotes = append(chordNotes, *note.Interval(intervals.MajorNinth))
	return chordNotes
}

func dominantNinth(note *notes.Note) []notes.Note {
	chordNotes := dominantSeventh(note)
	chordNotes = append(chordNotes, *note.Interval(intervals.MajorNinth))
	return chordNotes
}

func dominantMinorNinth(note *notes.Note) []notes.Note {
	chordNotes := dominantSeventh(note)
	chordNotes = append(chordNotes, *note.Interval(intervals.MinorNinth))
	return chordNotes
}

func minorMajorNinth(note *notes.Note) []notes.Note {
	chordNotes := minorMajorSeventh(note)
	chordNotes = append(chordNotes, *note.Interval(intervals.MajorNinth))
	return chordNotes
}

func minorNinth(note *notes.Note) []notes.Note {
	chordNotes := minorSeventh(note)
	chordNotes = append(chordNotes, *note.Interval(intervals.MajorNinth))
	return chordNotes
}

func augmentedMajorNinth(note *notes.Note) []notes.Note {
	chordNotes := augmentedMajorSeventh(note)
	chordNotes = append(chordNotes, *note.Interval(intervals.MajorNinth))
	return chordNotes
}

func augmentedDominantNinth(note *notes.Note) []notes.Note {
	chordNotes := augmentedSeventh(note)
	chordNotes = append(chordNotes, *note.Interval(intervals.MajorNinth))
	return chordNotes
}

func halfDiminishedNinth(note *notes.Note) []notes.Note {
	chordNotes := halfDiminishedSeventh(note)
	chordNotes = append(chordNotes, *note.Interval(intervals.MajorNinth))
	return chordNotes
}

func halfDiminishedMinorNinth(note *notes.Note) []notes.Note {
	chordNotes := halfDiminishedSeventh(note)
	chordNotes = append(chordNotes, *note.Interval(intervals.MinorNinth))
	return chordNotes
}

func diminishedNinth(note *notes.Note) []notes.Note {
	chordNotes := diminishedSeventh(note)
	chordNotes = append(chordNotes, *note.Interval(intervals.MajorNinth))
	return chordNotes
}

func diminishedMinorNinth(note *notes.Note) []notes.Note {
	chordNotes := diminishedSeventh(note)
	chordNotes = append(chordNotes, *note.Interval(intervals.MinorNinth))
	return chordNotes
}

func eleventh(note *notes.Note) []notes.Note {
	chordNotes := dominantNinth(note)
	chordNotes = append(chordNotes, *note.Interval(intervals.PerfectEleventh))
	return chordNotes
}

func majorEleventh(note *notes.Note) []notes.Note {
	chordNotes := majorNinth(note)
	chordNotes = append(chordNotes, *note.Interval(intervals.PerfectEleventh))
	return chordNotes
}

func minorMajorEleventh(note *notes.Note) []notes.Note {
	chordNotes := minorMajorNinth(note)
	chordNotes = append(chordNotes, *note.Interval(intervals.PerfectEleventh))
	return chordNotes
}

func minorEleventh(note *notes.Note) []notes.Note {
	chordNotes := minorNinth(note)
	chordNotes = append(chordNotes, *note.Interval(intervals.PerfectEleventh))
	return chordNotes
}

func augmentedMajorEleventh(note *notes.Note) []notes.Note {
	chordNotes := augmentedMajorNinth(note)
	chordNotes = append(chordNotes, *note.Interval(intervals.PerfectEleventh))
	return chordNotes
}

func augmentedEleventh(note *notes.Note) []notes.Note {
	chordNotes := augmentedDominantNinth(note)
	chordNotes = append(chordNotes, *note.Interval(intervals.PerfectEleventh))
	return chordNotes
}

func halfDiminishedEleventh(note *notes.Note) []notes.Note {
	chordNotes := halfDiminishedNinth(note)
	chordNotes = append(chordNotes, *note.Interval(intervals.PerfectEleventh))
	return chordNotes
}

func diminishedEleventh(note *notes.Note) []notes.Note {
	chordNotes := diminishedNinth(note)
	chordNotes = append(chordNotes, *note.Interval(intervals.PerfectEleventh))
	return chordNotes
}

func majorThirteenth(note *notes.Note) []notes.Note {
	chordNotes := majorEleventh(note)
	chordNotes = append(chordNotes, *note.Interval(intervals.MajorThirteenth))
	return chordNotes
}

func thirteenth(note *notes.Note) []notes.Note {
	chordNotes := eleventh(note)
	chordNotes = append(chordNotes, *note.Interval(intervals.MajorThirteenth))
	return chordNotes
}

func minorMajorThirteenth(note *notes.Note) []notes.Note {
	chordNotes := minorMajorEleventh(note)
	chordNotes = append(chordNotes, *note.Interval(intervals.MajorThirteenth))
	return chordNotes
}

func minorThirteenth(note *notes.Note) []notes.Note {
	chordNotes := minorEleventh(note)
	chordNotes = append(chordNotes, *note.Interval(intervals.MajorThirteenth))
	return chordNotes
}

func augmentedMajorThirteenth(note *notes.Note) []notes.Note {
	chordNotes := augmentedMajorEleventh(note)
	chordNotes = append(chordNotes, *note.Interval(intervals.MajorThirteenth))
	return chordNotes
}

func augmentedThirteenth(note *notes.Note) []notes.Note {
	chordNotes := augmentedEleventh(note)
	chordNotes = append(chordNotes, *note.Interval(intervals.MajorThirteenth))
	return chordNotes
}

func halfDiminishedThirteenth(note *notes.Note) []notes.Note {
	chordNotes := halfDiminishedEleventh(note)
	chordNotes = append(chordNotes, *note.Interval(intervals.MajorThirteenth))
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

	case "M9":
		fallthrough
	case "maj9":
		// major ninth
		chordNotes = majorNinth(n)

	case "9":
		// dominant ninth
		chordNotes = dominantNinth(n)

	case "7b9":
		// dominant minor ninth
		chordNotes = dominantMinorNinth(n)

	case "mM9":
		fallthrough
	case "-M9":
		fallthrough
	case "minmaj9":
		// minor-major ninth
		chordNotes = minorMajorNinth(n)

	case "m9":
		fallthrough
	case "-9":
		fallthrough
	case "min9":
		// minor ninth
		chordNotes = minorNinth(n)

	case "+M9":
		fallthrough
	case "augmaj9":
		// augmented major ninth
		chordNotes = augmentedMajorNinth(n)

	case "+9":
		fallthrough
	case "9#5":
		fallthrough
	case "aug9":
		// augmented dominant ninth
		chordNotes = augmentedDominantNinth(n)

	case "ø9":
		// half-diminished ninth
		chordNotes = halfDiminishedNinth(n)

	case "øb9":
		// half-diminished ninth
		chordNotes = halfDiminishedMinorNinth(n)

	case "o9":
		fallthrough
	case "dim9":
		// diminished ninth
		chordNotes = diminishedNinth(n)

	case "ob9":
		fallthrough
	case "dimb9":
		// diminished ninth
		chordNotes = diminishedMinorNinth(n)

	case "11":
		// eleventh
		chordNotes = eleventh(n)

	case "M11":
		fallthrough
	case "maj11":
		// major eleventh
		chordNotes = majorEleventh(n)

	case "mM11":
		fallthrough
	case "-M11":
		fallthrough
	case "minmaj11":
		// minor-major eleventh
		chordNotes = minorMajorEleventh(n)

	case "m11":
		fallthrough
	case "-11":
		fallthrough
	case "min11":
		// minor eleventh
		chordNotes = minorEleventh(n)

	case "+M11":
		fallthrough
	case "augmaj11":
		// augmented major eleventh
		chordNotes = augmentedMajorEleventh(n)

	case "+11":
		fallthrough
	case "11#5":
		fallthrough
	case "aug11":
		// augmented eleventh
		chordNotes = augmentedEleventh(n)

	case "ø11":
		// half-diminished eleventh
		chordNotes = halfDiminishedEleventh(n)

	case "o11":
		fallthrough
	case "dim11":
		// diminished ninth
		chordNotes = diminishedEleventh(n)

	case "M13":
		fallthrough
	case "maj13":
		// major thirteenth
		chordNotes = majorThirteenth(n)

	case "13":
		// thirteenth
		chordNotes = thirteenth(n)

	case "mM13":
		fallthrough
	case "-M13":
		fallthrough
	case "minmaj13":
		// minor-major thirteenth
		chordNotes = minorMajorThirteenth(n)

	case "m13":
		fallthrough
	case "-13":
		fallthrough
	case "min13":
		// minor-major thirteenth
		chordNotes = minorThirteenth(n)

	case "+M13":
		fallthrough
	case "augmajM13":
		// augmented major thirteenth
		chordNotes = augmentedMajorThirteenth(n)

	case "+13":
		fallthrough
	case "13#5":
		fallthrough
	case "aug13":
		// augmented thirteenth
		chordNotes = augmentedThirteenth(n)

	case "ø13":
		// half-diminished thirteenth
		chordNotes = halfDiminishedThirteenth(n)

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
