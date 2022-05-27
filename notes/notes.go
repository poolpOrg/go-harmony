package notes

import (
	"fmt"

	"github.com/poolpOrg/go-harmony/intervals"
)

type note struct {
	name        rune
	degree      uint
	semitone    int
	accidentals string
	octave      uint
}
type Note note

type natural struct {
	degree   uint
	semitone int
}
type Natural natural

var naturals = map[rune]Natural{
	'C': Natural{degree: 0, semitone: 0},
	'D': Natural{degree: 1, semitone: 2},
	'E': Natural{degree: 2, semitone: 4},
	'F': Natural{degree: 3, semitone: 5},
	'G': Natural{degree: 4, semitone: 7},
	'A': Natural{degree: 5, semitone: 9},
	'B': Natural{degree: 6, semitone: 11},
}

func getNaturals() []rune {
	return []rune{'C', 'D', 'E', 'F', 'G', 'A', 'B'}
}

func NoteByName(name string) (*Note, error) {
	if len(name) < 1 {
		return nil, fmt.Errorf("no such note")
	}

	if natural, exists := naturals[rune(name[0])]; !exists {
		return nil, fmt.Errorf("no such note")
	} else {
		accidentals := ""
		if len(name) > 1 {
			accidentals = name[1:]
			switch accidentals {
			case "bb":
				if natural.semitone > 2 {
					natural.semitone = (natural.semitone - 2) % 12
				} else {
					natural.semitone = (12 + natural.semitone) - 2
				}
			case "b":
				if natural.semitone > 1 {
					natural.semitone = (natural.semitone - 1) % 12
				} else {
					natural.semitone = (12 + natural.semitone) - 1
				}
			case "#":
				natural.semitone = (natural.semitone + 1) % 12
			case "##":
				natural.semitone = (natural.semitone + 2) % 12
			default:
				return nil, fmt.Errorf("no such accidental")
			}
		}
		return &Note{
			name:        rune(name[0]),
			degree:      natural.degree,
			semitone:    natural.semitone,
			accidentals: accidentals,
			octave:      4,
		}, nil
	}
}

func (note Note) Name() string {
	return fmt.Sprintf("%c%s", note.name, note.accidentals)
}

func (note *Note) Interval(interval intervals.Interval) Note {
	targetName := getNaturals()[(note.degree+interval.Degree())%7]
	targetNatural := naturals[getNaturals()[(note.degree+interval.Degree())%7]]

	fmt.Println("note", note.degree, note.semitone)
	fmt.Println("interval distance", interval.Distance())

	var distance int
	if note.semitone == targetNatural.semitone {
		distance = 0
	} else if note.semitone < targetNatural.semitone {
		fmt.Println("target", targetNatural.degree, targetNatural.semitone)
		distance = int(interval.Distance()) - (targetNatural.semitone - note.semitone)
	} else {
		// XXX THIS IS BROKEN
		// need to compute distance accounting for the wrap around C
		targetSemitone := targetNatural.semitone + 12
		fmt.Println("target", targetNatural.degree, targetNatural.semitone, "/ wrapped", targetSemitone)
		distance = int(interval.Distance()) - (targetSemitone - note.semitone)
		fmt.Println("actual distance", distance)
	}
	fmt.Println("note.semitone", note.semitone, "natural.semitone", targetNatural.semitone, "interval.distance", interval.Distance(), "distance", distance)

	distance = distance % 12
	var name string
	switch distance {
	case -2:
		name = fmt.Sprintf("%cbb", targetName)
	case -1:
		name = fmt.Sprintf("%cb", targetName)
	case 0:
		name = fmt.Sprintf("%c", targetName)
	case 1:
		name = fmt.Sprintf("%c#", targetName)
	case 2:
		name = fmt.Sprintf("%c##", targetName)
	}

	//fmt.Println(name)
	target, _ := NoteByName(name)
	return *target
}
