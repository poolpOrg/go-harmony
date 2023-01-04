package notes

import (
	"fmt"
	"strings"

	"github.com/poolpOrg/go-harmony/intervals"
	"github.com/poolpOrg/go-harmony/naturals"
	"github.com/poolpOrg/go-harmony/octaves"
	"github.com/poolpOrg/go-harmony/tunings"
)

type note struct {
	natural     naturals.Natural
	accidentals int
	octave      octaves.Octave
}
type Note note

func Parse(note string) (*Note, error) {
	if len(note) == 0 {
		return nil, fmt.Errorf("empty note")
	}

	var octave octaves.Octave = octaves.C4
	var accidentals = note[1:]

	r := rune(note[len(note)-1])
	if r >= '0' && r <= '9' {
		t, err := octaves.FromName("C" + note[len(note)-1:])
		if err != nil {
			return nil, err
		}
		octave = *t
		accidentals = note[1 : len(note)-1]
	}

	var accidentalDelta int
	for _, accidental := range accidentals {
		switch accidental {
		case 'b':
			accidentalDelta -= 1
		case '#':
			accidentalDelta += 1
		default:
			return nil, fmt.Errorf("bad accidental (%c): should be 'b' or '#'", accidental)
		}
	}

	natural, err := naturals.FromName(note[0:1])
	if err != nil {
		return nil, err
	}

	return &Note{
		natural:     *natural,
		accidentals: accidentalDelta,
		octave:      octave,
	}, nil
}

func (note *Note) Name() string {
	if note.accidentals == 0 {
		return note.natural.Name()
	} else if note.accidentals < 0 {
		return note.natural.Name() + strings.Repeat("b", -note.accidentals)
	} else {
		return note.natural.Name() + strings.Repeat("#", note.accidentals)
	}
}

func (note *Note) OctaveName() string {
	if note.accidentals == 0 {
		return fmt.Sprintf("%s%d", note.natural.Name(), note.Octave())
	} else if note.accidentals < 0 {
		return fmt.Sprintf("%s%d", note.natural.Name()+strings.Repeat("b", -note.accidentals), note.Octave())
	} else {
		return fmt.Sprintf("%s%d", note.natural.Name()+strings.Repeat("#", note.accidentals), note.Octave())
	}
}

func (note *Note) Interval(interval intervals.Interval) *Note {
	if interval.Position()%7 == 0 && interval.Semitones()%12 == 0 {
		targetNote := *note
		targetNote.octave = *targetNote.octave.Add(uint8(interval.Semitones() / 12))
		return &targetNote
	}

	target := naturals.Naturals()[(int(note.natural.Position())+int(interval.Position()))%len(naturals.Naturals())]
	sourceSemitone := int(note.natural.Semitones()) + note.accidentals

	targetOctave := note.octave.Add(uint8(interval.Semitones() / 12))
	if target.Position() < note.natural.Position() {
		targetOctave = targetOctave.Next()
	}

	targetSemitone := target.Semitones() + uint(note.accidentals)
	targetAccidentals := note.accidentals
	if targetSemitone < note.natural.Semitones() {
		targetSemitone += 12
	}

	distance := int(targetSemitone) - (int(interval.Semitones()%12) + sourceSemitone)
	targetAccidentals += -distance
	//	targetOctave += uint8(targetAccidentals / 12)
	targetAccidentals = targetAccidentals % 12

	if interval == intervals.AugmentedSeventh || interval == intervals.AugmentedFourteenth {
		targetOctave = targetOctave.Previous()
		targetAccidentals = targetAccidentals + 12
	}
	if interval == intervals.DiminishedOctave || interval == intervals.DiminishedFifteenth {
		targetOctave = targetOctave.Next()
		targetAccidentals = targetAccidentals - 12
	}

	if targetAccidentals < 0 {
		if -targetAccidentals%12 == 0 {
			targetAccidentals = 0
		}
	} else {
		if targetAccidentals%12 == 0 {
			targetAccidentals = 0
		}
	}
	return &Note{
		natural:     target,
		accidentals: targetAccidentals,
		octave:      *targetOctave,
	}
}

func (note *Note) Distance(target Note) intervals.Interval {
	var origin Note
	var destination Note

	offset1 := uint8(note.Semitones())
	offset2 := uint8(target.Semitones())

	if offset1 < offset2 {
		origin = *note
		destination = target
	} else {
		origin = target
		destination = *note
	}

	targetPosition := destination.Position()
	targetSemitone := destination.Semitones()

	targetPosition -= origin.Position()
	targetSemitone -= origin.Semitones()

	if target.Octave() > note.Octave() {
		targetPosition += 7
		targetSemitone += 12
	}
	return *intervals.New(targetPosition, uint(targetSemitone))
}

func (note *Note) Position() uint {
	return note.natural.Position()
}

func (note *Note) Semitones() int {
	return int(note.natural.Semitones()) + note.accidentals
}

func (note *Note) Frequency() float64 {
	tuner := tunings.NewTuner(tunings.EqualTemperament, tunings.A440)
	return tuner.Frequency(uint8(note.Position()))
}

func (note *Note) Enharmonic(target Note) bool {
	return note.Semitones()%12 == target.Semitones()%12
}

func (note *Note) Octave() uint8 {
	return note.octave.Position()
}

func (note *Note) Previous() *Note {
	n := *note.Interval(intervals.MajorSeventh)
	if note.natural.Name() == "C" {
		n.octave = *n.octave.Previous()
	}
	return &n
}

func (note *Note) Next() *Note {
	n := *note.Interval(intervals.MinorSecond)
	if note.natural.Name() == "B" {
		n.octave = *n.octave.Next()
	}
	return &n
}

func (note *Note) NextChromatic() *Note {
	var err error
	var n *Note
	switch note.Name() {
	case "C":
		n, err = Parse("C#")
	case "C#":
		n, err = Parse("D")
	case "D":
		n, err = Parse("D#")
	case "D#":
		n, err = Parse("E")
	case "E":
		n, err = Parse("F")
	case "F":
		n, err = Parse("F#")
	case "F#":
		n, err = Parse("G")
	case "G":
		n, err = Parse("G#")
	case "G#":
		n, err = Parse("A")
	case "A":
		n, err = Parse("A#")
	case "A#":
		n, err = Parse("B")
	case "B":
		n, err = Parse("C")
	}
	if err != nil {
		panic(err)
	}
	return n
}

func (note *Note) Lower() *Note {
	n := *note
	n.accidentals -= 1
	return &n
}

func (note *Note) Raise() *Note {
	n := *note
	n.accidentals += 1
	return &n
}

// temporarily until a more generic method is devised
func (note *Note) SetOctave(position uint8) error {
	octave, err := octaves.FromPosition(position)
	if err != nil {
		return err
	}
	note.octave = *octave
	return nil
}

func (note *Note) AbsoluteSemitones() uint8 {
	return note.octave.Position()*12 + uint8(note.Semitones())
}

func (note *Note) MIDI() uint8 {
	return uint8(note.Semitones()) + ((note.Octave() + 1) * 12)
}
