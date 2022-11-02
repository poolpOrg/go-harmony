package notes

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/poolpOrg/go-harmony/intervals"
	"github.com/poolpOrg/go-harmony/tunings"
)

type natural struct {
	name     string
	pos      int
	semitone int
}

var naturals = []natural{
	{name: "C", pos: 0, semitone: 0},
	{name: "D", pos: 1, semitone: 2},
	{name: "E", pos: 2, semitone: 4},
	{name: "F", pos: 3, semitone: 5},
	{name: "G", pos: 4, semitone: 7},
	{name: "A", pos: 5, semitone: 9},
	{name: "B", pos: 6, semitone: 11},
}

type note struct {
	name        string
	pos         int
	semitone    int
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

	name := note[0:1]
	var pos int
	var semitone int
	switch name {
	case "C":
		pos = 0
		semitone = 0
	case "D":
		pos = 1
		semitone = 2
	case "E":
		pos = 2
		semitone = 4
	case "F":
		pos = 3
		semitone = 5
	case "G":
		pos = 4
		semitone = 7
	case "A":
		pos = 5
		semitone = 9
	case "B":
		pos = 6
		semitone = 11
	default:
		return nil, fmt.Errorf("bad note (%s): should be 'C', 'D', 'E', 'F', 'G', 'A' or 'B'", name)
	}

	return &Note{
		name:        name,
		pos:         pos,
		semitone:    semitone,
		accidentals: accidentalDelta,
		octave:      octave,
	}, nil
}

func (note *Note) Name() string {
	if note.accidentals == 0 {
		return note.name
	} else if note.accidentals < 0 {
		return note.name + strings.Repeat("b", -note.accidentals)
	} else {
		return note.name + strings.Repeat("#", note.accidentals)
	}
}

func (note *Note) OctaveName() string {
	if note.accidentals == 0 {
		return fmt.Sprintf("%s%d", note.name, note.Octave())
	} else if note.accidentals < 0 {
		return fmt.Sprintf("%s%d", note.name+strings.Repeat("b", -note.accidentals), note.Octave())
	} else {
		return fmt.Sprintf("%s%d", note.name+strings.Repeat("#", note.accidentals), note.Octave())
	}
}

func (note *Note) Interval(interval intervals.Interval) *Note {

	// interval aligned on an octave
	if interval.Semitone()%12 == 0 {
		targetNote := *note
		targetNote.octave += uint8(interval.Semitone() / 12)
		return &targetNote
	}

	target := naturals[(note.pos+int(interval.Position()))%len(naturals)]
	sourceSemitone := note.semitone + note.accidentals

	targetOctave := note.octave + uint8(interval.Semitone()/12)
	if target.pos < note.pos {
		targetOctave++
	}

	targetSemitone := target.semitone + note.accidentals
	targetAccidentals := note.accidentals
	if targetSemitone < note.semitone {
		targetSemitone += 12
	}

	distance := targetSemitone - (int(interval.Semitone()%12) + sourceSemitone)
	targetAccidentals += -distance

	return &Note{
		name:        target.name,
		pos:         target.pos,
		semitone:    target.semitone,
		accidentals: targetAccidentals,
		octave:      targetOctave,
	}
}

func (note *Note) Distance(target Note) intervals.Interval {
	targetPosition := target.Position()
	if targetPosition < note.Position() {
		targetPosition += 7
	}

	targetSemitone := target.Semitone()
	if targetSemitone < note.Semitone() {
		targetSemitone += 12
	}

	targetPosition -= note.Position()
	targetSemitone -= note.Semitone()

	//targetPosition += (uint(target.octave) - uint(note.octave)) * 7
	//targetSemitone += (int(target.octave) - int(note.octave)) * 12

	//if target.pos < note.pos {
	//	targetSemitone += 12
	//}

	return intervals.New(targetPosition, uint(targetSemitone))
}

func (note *Note) Position() uint {
	return uint(note.pos)
}

func (note *Note) Semitone() int {
	return note.semitone + note.accidentals
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
	if note.name == "C" {
		n.octave -= 1
	}
	return &n
}

func (note *Note) Next() *Note {
	n := *note.Interval(intervals.MinorSecond)
	if note.name == "B" {
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
