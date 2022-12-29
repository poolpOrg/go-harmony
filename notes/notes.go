package notes

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/poolpOrg/go-harmony/intervals"
	"github.com/poolpOrg/go-harmony/naturals"
	"github.com/poolpOrg/go-harmony/tunings"
)

type note struct {
	natural     naturals.Natural
	accidentals int
	octave      uint8
}
type Note note

func Parse(note string) (*Note, error) {
	if len(note) == 0 {
		return nil, fmt.Errorf("empty note")
	}

	var octave uint8
	var accidentals string

	r := rune(note[len(note)-1])
	if r >= '0' && r <= '9' {
		value, err := strconv.ParseUint(note[len(note)-1:], 10, 8)
		if err != nil {
			return nil, err
		}
		if value > 8 {
			return nil, fmt.Errorf("bad octave (%d): should be 0 <= n <= 8)", value)
		}
		octave = uint8(value)
		accidentals = note[1 : len(note)-1]
	} else {
		octave = uint8(4)
		accidentals = note[1:]
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

	natural, err := naturals.Parse(note[0:1])
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
	// interval aligned on an octave
	if interval.Semitone()%12 == 0 {
		targetNote := *note
		targetNote.octave += uint8(interval.Semitone() / 12)
		return &targetNote
	}

	target := naturals.Naturals()[(int(note.natural.Position())+int(interval.Position()))%len(naturals.Naturals())]
	sourceSemitone := int(note.natural.Semitones()) + note.accidentals

	targetOctave := note.octave + uint8(interval.Semitone()/12)
	if target.Position() < note.natural.Position() {
		targetOctave++
	}

	targetSemitone := target.Semitones() + uint(note.accidentals)
	targetAccidentals := note.accidentals
	if targetSemitone < note.natural.Semitones() {
		targetSemitone += 12
	}

	distance := int(targetSemitone) - (int(interval.Semitone()%12) + sourceSemitone)
	targetAccidentals += -distance
	//	targetOctave += uint8(targetAccidentals / 12)
	targetAccidentals = targetAccidentals % 12

	return &Note{
		natural:     target,
		accidentals: targetAccidentals,
		octave:      targetOctave,
	}
}

func (note *Note) Distance(target Note) intervals.Interval {
	var origin Note
	var destination Note

	offset1 := note.octave*12 + uint8(note.Semitone())
	offset2 := target.octave*12 + uint8(target.Semitone())
	if offset1 < offset2 {
		origin = *note
		destination = target
	} else {
		origin = target
		destination = *note
	}

	targetPosition := destination.Position()
	targetSemitone := destination.Semitone()

	targetPosition -= origin.Position()
	targetSemitone -= origin.Semitone()

	if target.Octave() > note.Octave() {
		targetPosition += 7
		targetSemitone += 12
	}

	return intervals.New(targetPosition, uint(targetSemitone))
}

func (note *Note) Position() uint {
	return note.natural.Position()
}

func (note *Note) Semitone() int {
	return int(note.natural.Semitones()) + note.accidentals
}

func (note *Note) Frequency() float64 {
	tuning := tunings.A440
	octave := note.octave
	semitone := note.Semitone()
	if semitone < 0 {
		semitone = (12 + semitone) % 12
		if octave != 0 {
			octave -= 1
		}
	}
	return tuning.Frequency(uint(semitone), uint(octave))
}

func (note *Note) Inharmonic(target Note) bool {
	return note.Semitone()%12 == target.Semitone()%12
}

func (note *Note) Octave() uint8 {
	return note.octave
}

func (note *Note) Previous() *Note {
	n := *note.Interval(intervals.MajorSeventh)
	if note.natural.Name() == "C" {
		n.octave -= 1
	}
	return &n
}

func (note *Note) Next() *Note {
	n := *note.Interval(intervals.MinorSecond)
	if note.natural.Name() == "B" {
		n.octave += 1
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
func (note *Note) SetOctave(octave uint8) {
	note.octave = octave
}
