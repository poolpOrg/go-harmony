package scales

import (
	"fmt"

	"github.com/poolpOrg/go-harmony/chords"
	"github.com/poolpOrg/go-harmony/intervals"
	"github.com/poolpOrg/go-harmony/notes"
)

type scale struct {
	notes  []notes.Note
	chords []chords.Chord
}
type Scale scale

func majorScale(note *notes.Note) *Scale {
	scaleNotes := make([]notes.Note, 0)
	scaleNotes = append(scaleNotes, *note)
	scaleNotes = append(scaleNotes, *note.Interval(intervals.MajorSecond))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.MajorThird))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.PerfectFourth))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.PerfectFifth))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.MajorSixth))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.MajorSeventh))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.Octave))

	scaleChords := make([]chords.Chord, 0)
	chord, _ := chords.Parse(note.Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.MajorSecond).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.MajorThird).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.PerfectFourth).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.PerfectFifth).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.MajorSixth).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.MajorSeventh).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.Octave).Name())
	scaleChords = append(scaleChords, *chord)

	return &Scale{
		notes:  scaleNotes,
		chords: scaleChords,
	}
}

func naturalMinorScale(note *notes.Note) *Scale {
	scaleNotes := make([]notes.Note, 0)
	scaleNotes = append(scaleNotes, *note)
	scaleNotes = append(scaleNotes, *note.Interval(intervals.MajorSecond))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.MinorThird))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.PerfectFourth))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.PerfectFifth))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.MinorSixth))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.MinorSeventh))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.Octave))

	scaleChords := make([]chords.Chord, 0)
	chord, _ := chords.Parse(note.Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.MajorSecond).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.MinorThird).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.PerfectFourth).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.PerfectFifth).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.MinorSixth).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.MinorSeventh).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.Octave).Name())
	scaleChords = append(scaleChords, *chord)

	return &Scale{
		notes:  scaleNotes,
		chords: scaleChords,
	}
}

func harmonicMinorScale(note *notes.Note) *Scale {
	scaleNotes := make([]notes.Note, 0)
	scaleNotes = append(scaleNotes, *note)
	scaleNotes = append(scaleNotes, *note.Interval(intervals.MajorSecond))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.MinorThird))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.PerfectFourth))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.PerfectFifth))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.MinorSixth))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.MajorSeventh))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.Octave))

	scaleChords := make([]chords.Chord, 0)
	chord, _ := chords.Parse(note.Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.MajorSecond).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.MinorThird).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.PerfectFourth).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.PerfectFifth).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.MinorSixth).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.MajorSeventh).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.Octave).Name())
	scaleChords = append(scaleChords, *chord)

	return &Scale{
		notes:  scaleNotes,
		chords: scaleChords,
	}
}

func melodicMinorScale(note *notes.Note) *Scale {
	scaleNotes := make([]notes.Note, 0)
	scaleNotes = append(scaleNotes, *note)
	scaleNotes = append(scaleNotes, *note.Interval(intervals.MajorSecond))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.MinorThird))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.PerfectFourth))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.PerfectFifth))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.MajorSixth))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.MajorSeventh))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.Octave))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.MinorSeventh))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.MinorSixth))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.PerfectFifth))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.PerfectFourth))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.MinorThird))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.MajorSecond))
	scaleNotes = append(scaleNotes, *note)

	scaleChords := make([]chords.Chord, 0)
	chord, _ := chords.Parse(note.Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.MajorSecond).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.MinorThird).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.PerfectFourth).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.PerfectFifth).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.MajorSixth).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.MajorSeventh).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.Octave).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.MinorSeventh).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.MinorSixth).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.PerfectFifth).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.PerfectFourth).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.MinorThird).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.MajorSecond).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Name())
	scaleChords = append(scaleChords, *chord)

	return &Scale{
		notes:  scaleNotes,
		chords: scaleChords,
	}
}

func jazzMinorScale(note *notes.Note) *Scale {
	scaleNotes := make([]notes.Note, 0)
	scaleNotes = append(scaleNotes, *note)
	scaleNotes = append(scaleNotes, *note.Interval(intervals.MajorSecond))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.MinorThird))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.PerfectFourth))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.PerfectFifth))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.MajorSixth))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.MajorSeventh))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.Octave))

	scaleChords := make([]chords.Chord, 0)
	chord, _ := chords.Parse(note.Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.MajorSecond).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.MinorThird).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.PerfectFourth).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.PerfectFifth).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.MajorSixth).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.MajorSeventh).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.Octave).Name())
	scaleChords = append(scaleChords, *chord)

	return &Scale{
		notes:  scaleNotes,
		chords: scaleChords,
	}
}

func majorPentatonicScale(note *notes.Note) *Scale {
	scaleNotes := make([]notes.Note, 0)
	scaleNotes = append(scaleNotes, *note)
	scaleNotes = append(scaleNotes, *note.Interval(intervals.MajorSecond))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.MajorThird))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.PerfectFifth))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.MajorSixth))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.Octave))

	scaleChords := make([]chords.Chord, 0)
	chord, _ := chords.Parse(note.Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.MajorSecond).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.MajorThird).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.PerfectFifth).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.MajorSixth).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.Octave).Name())
	scaleChords = append(scaleChords, *chord)

	return &Scale{
		notes:  scaleNotes,
		chords: scaleChords,
	}
}

func egyptianPentatonicScale(note *notes.Note) *Scale {
	scaleNotes := make([]notes.Note, 0)
	scaleNotes = append(scaleNotes, *note)
	scaleNotes = append(scaleNotes, *note.Interval(intervals.MajorSecond))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.PerfectFourth))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.PerfectFifth))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.MinorSeventh))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.Octave))

	scaleChords := make([]chords.Chord, 0)
	chord, _ := chords.Parse(note.Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.MajorSecond).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.PerfectFourth).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.PerfectFifth).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.MinorSeventh).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.Octave).Name())
	scaleChords = append(scaleChords, *chord)

	return &Scale{
		notes:  scaleNotes,
		chords: scaleChords,
	}
}

func bluesMinorScale(note *notes.Note) *Scale {
	scaleNotes := make([]notes.Note, 0)
	scaleNotes = append(scaleNotes, *note)
	scaleNotes = append(scaleNotes, *note.Interval(intervals.MinorThird))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.PerfectFourth))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.MinorSixth))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.MinorSeventh))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.Octave))

	scaleChords := make([]chords.Chord, 0)
	chord, _ := chords.Parse(note.Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.MinorThird).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.PerfectFourth).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.MinorSixth).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.MinorSeventh).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.Octave).Name())
	scaleChords = append(scaleChords, *chord)

	return &Scale{
		notes:  scaleNotes,
		chords: scaleChords,
	}
}

func bluesMajorScale(note *notes.Note) *Scale {
	scaleNotes := make([]notes.Note, 0)
	scaleNotes = append(scaleNotes, *note)
	scaleNotes = append(scaleNotes, *note.Interval(intervals.MajorSecond))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.PerfectFourth))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.PerfectFifth))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.MajorSixth))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.Octave))

	scaleChords := make([]chords.Chord, 0)
	chord, _ := chords.Parse(note.Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.MajorSecond).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.PerfectFourth).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.PerfectFifth).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.MajorSixth).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.Octave).Name())
	scaleChords = append(scaleChords, *chord)

	return &Scale{
		notes:  scaleNotes,
		chords: scaleChords,
	}
}

func minorPentatonicScale(note *notes.Note) *Scale {
	scaleNotes := make([]notes.Note, 0)
	scaleNotes = append(scaleNotes, *note)
	scaleNotes = append(scaleNotes, *note.Interval(intervals.MinorThird))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.PerfectFourth))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.PerfectFifth))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.MinorSeventh))
	scaleNotes = append(scaleNotes, *note.Interval(intervals.Octave))

	scaleChords := make([]chords.Chord, 0)
	chord, _ := chords.Parse(note.Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.MinorThird).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.PerfectFourth).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.PerfectFifth).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.MinorSeventh).Name())
	scaleChords = append(scaleChords, *chord)
	chord, _ = chords.Parse(note.Interval(intervals.Octave).Name())
	scaleChords = append(scaleChords, *chord)

	return &Scale{
		notes:  scaleNotes,
		chords: scaleChords,
	}
}

func Parse(scale string) (*Scale, error) {
	if len(scale) == 0 {
		return nil, fmt.Errorf("empty scale")
	}

	noteEnd := int(1)
	for _, r := range scale[1:] {
		if r == '#' || r == 'b' {
			noteEnd++
			continue
		}
		break
	}

	noteName := scale[0:noteEnd]
	scaleName := scale[noteEnd:]

	n, err := notes.Parse(noteName)
	if err != nil {
		return nil, err
	}

	switch scaleName {
	case "":
		fallthrough
	case "M":
		fallthrough
	case "maj":
		// major scale
		return majorScale(n), nil

	case "m":
		fallthrough
	case "min":
		// natural minor scale
		return naturalMinorScale(n), nil

	case "harmonicMinor":
		// harmonic minor scale
		return harmonicMinorScale(n), nil

	case "melodicMinor":
		// melodic minor scale
		return melodicMinorScale(n), nil

	case "jazzMinor":
		// jazz minor scale
		return jazzMinorScale(n), nil

	case "majorPentatonic":
		// major pentatonic scale
		return majorPentatonicScale(n), nil

	case "egyptian":
		// egyptian pentatonic scale
		return egyptianPentatonicScale(n), nil

	case "BluesMinor":
		// blues minor pentatonic scale
		return bluesMinorScale(n), nil

	case "BluesMajor":
		// blues major pentatonic scale
		return bluesMajorScale(n), nil

	case "minorPentatonic":
		// minor pentatonic scale
		return minorPentatonicScale(n), nil

	default:
		return nil, fmt.Errorf("unknown scale name: %s", scaleName)

	}
}

func (scale *Scale) Name() string {
	return scale.notes[0].Name()
}

func (scale *Scale) Notes() []notes.Note {
	return scale.notes
}

func (scale *Scale) Chords() []chords.Chord {
	return scale.chords
}
