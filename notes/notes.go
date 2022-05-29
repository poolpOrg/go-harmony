package notes

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/poolpOrg/go-harmony/intervals"
	"github.com/poolpOrg/go-harmony/tuning"
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
	indexOctave := strings.Index(note, "^")
	if indexOctave != -1 {
		value, err := strconv.ParseUint(note[indexOctave+1:], 10, 8)
		if err != nil {
			return nil, err
		}
		if value > 8 {
			return nil, fmt.Errorf("bad octave (%d): should be 0 <= n <= 8)", value)
		}
		octave = uint8(value)
		accidentals = note[1:indexOctave]
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

func (note *Note) Interval(interval intervals.Interval) *Note {
	if interval.Semitone()%12 == 0 {
		targetNote := *&note
		targetNote.octave += uint8(interval.Semitone() / 12)
		return targetNote
	}

	target := naturals[(note.pos+int(interval.Position()))%len(naturals)]
	sourceSemitone := note.semitone + note.accidentals
	if sourceSemitone < 0 {
		sourceSemitone = (12 + sourceSemitone) % 12
	}

	targetOctave := note.octave
	targetSemitone := target.semitone + note.accidentals
	targetAccidentals := note.accidentals
	if targetSemitone < note.semitone {
		targetSemitone += 12
	}

	//fmt.Println("distance:", interval.Semitone(), "source:", sourceSemitone, "target:", targetSemitone)
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

func (note *Note) Frequency() float64 {
	tuning := tuning.NewA440()
	semitone := note.semitone + note.accidentals
	if semitone < 0 {
		semitone = (12 + semitone) % 12
	}
	return tuning.Frequency(uint(semitone), uint(note.octave))
}
