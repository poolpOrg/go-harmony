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
	accidentals int8
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

	var accidentalDelta int8
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

func (note *Note) Equals(target Note) bool {
	return note.natural == target.natural &&
		note.accidentals == target.accidentals &&
		note.octave == target.octave
}

func (note *Note) Name() string {
	if note.accidentals == 0 {
		return note.natural.Name()
	} else if note.accidentals < 0 {
		return note.natural.Name() + strings.Repeat("b", int(-note.accidentals))
	} else {
		return note.natural.Name() + strings.Repeat("#", int(note.accidentals))
	}
}

func (note *Note) OctaveName() string {
	if note.accidentals == 0 {
		return fmt.Sprintf("%s%d", note.natural.Name(), note.Octave())
	} else if note.accidentals < 0 {
		return fmt.Sprintf("%s%d", note.natural.Name()+strings.Repeat("b", int(-note.accidentals)), note.Octave())
	} else {
		return fmt.Sprintf("%s%d", note.natural.Name()+strings.Repeat("#", int(note.accidentals)), note.Octave())
	}
}

func (note *Note) Interval(interval intervals.Interval) *Note {
	// shortcut on an octave multiplier
	if interval.Position()%7 == 0 && interval.Semitones()%12 == 0 {
		targetNote := *note
		targetNote.octave = *targetNote.octave.Add(int8(interval.Semitones() / 12))
		return &targetNote
	}

	// locate target natural note for interval
	targetNatural := naturals.Naturals()[(int(note.natural.Position())+int(interval.Position()))%len(naturals.Naturals())]

	// locate target octave taking into account interval AND natural positions
	targetOctave := *note.octave.Add(int8(interval.Position()) / 7)
	if targetNatural.Position() < note.natural.Position() {
		targetOctave = *targetOctave.Next()
	}

	// build target note with accidentals offset
	targetNote := &Note{
		natural:     targetNatural,
		accidentals: note.accidentals,
		octave:      targetOctave,
	}

	// adjust accidentals to match the interval's semitone distance
	distance := int8(interval.Semitones()) - int8(targetNote.AbsoluteSemitones()-note.AbsoluteSemitones())
	targetNote.accidentals += distance

	return targetNote
}

func (note *Note) Distance(target Note) intervals.Interval {
	var origin Note

	// always compute from lower to higher note regardless of current and target notes
	if note.AbsoluteSemitones() <= target.AbsoluteSemitones() {
		origin = *note
	} else {
		origin = target
		target = *note
	}

	// ie: B4 -> C5 = 1 semitone, <1 octave of distance, wrap
	if target.Position() < origin.Position() {
		target.octave.Previous()
	}

	// because C1->C3, C1->C4, C1->C5 are all fifteenth intervals, wrap target to origin octave+2
	if target.Octave()-note.Octave() > 2 {
		target.octave = *note.octave.Add(2)
	}

	// scan for an interval that would produce target
	for _, interval := range intervals.Intervals() {
		if note.Interval(interval).Equals(target) {
			return interval
		}
	}
	panic("unknown interval")
}

func (note *Note) Position() uint8 {
	return note.natural.Position()
}

func (note *Note) Semitones() uint8 {
	return uint8(int8(note.natural.Semitones()) + note.accidentals)
}

func (note *Note) Frequency() float64 {
	tuner := tunings.NewTuner(tunings.EqualTemperament, tunings.A440)
	return tuner.Frequency(uint8(note.AbsoluteSemitones()) - 4)
}

func (note *Note) Enharmonic(target Note) bool {
	return note.Semitones()%12 == target.Semitones()%12
}

func (note *Note) Octave() int8 {
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
func (note *Note) SetOctave(position int8) error {
	octave, err := octaves.FromPosition(position)
	if err != nil {
		return err
	}
	note.octave = *octave
	return nil
}

func (note *Note) AbsoluteSemitones() uint8 {
	return uint8(note.octave.Position()*12) + note.Semitones()
}

func (note *Note) MIDI() uint8 {
	// go-harmony starts at C0, midi notes start at 0
	return note.AbsoluteSemitones() + 12
}
