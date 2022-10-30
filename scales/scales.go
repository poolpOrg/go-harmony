package scales

import (
	"fmt"

	"github.com/poolpOrg/go-harmony/chords"
	"github.com/poolpOrg/go-harmony/intervals"
	"github.com/poolpOrg/go-harmony/notes"
)

type scale struct {
	root      notes.Note
	structure []intervals.Interval
}
type Scale scale

type Degree uint

var (
	Tonic       Degree = 0
	Supertonic  Degree = 1
	Mediant     Degree = 2
	Subdominant Degree = 3
	Dominant    Degree = 4
	Submediant  Degree = 5
	LeadingTone Degree = 6
)

var Ionian = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MajorSecond,
	intervals.MajorThird,
	intervals.PerfectFourth,
	intervals.PerfectFifth,
	intervals.MajorSixth,
	intervals.MajorSeventh,
	intervals.Octave,
}

var Dorian = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MajorSecond,
	intervals.MinorThird,
	intervals.PerfectFourth,
	intervals.PerfectFifth,
	intervals.MajorSixth,
	intervals.MinorSeventh,
	intervals.Octave,
}

var Phrygian = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MinorSecond,
	intervals.MinorThird,
	intervals.PerfectFourth,
	intervals.PerfectFifth,
	intervals.MinorSixth,
	intervals.MinorSeventh,
	intervals.Octave,
}

var Lydian = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MajorSecond,
	intervals.MajorThird,
	intervals.AugmentedFourth,
	intervals.PerfectFifth,
	intervals.MajorSixth,
	intervals.MajorSeventh,
	intervals.Octave,
}

var Mixolydian = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MajorSecond,
	intervals.MajorThird,
	intervals.PerfectFourth,
	intervals.PerfectFifth,
	intervals.MajorSixth,
	intervals.MinorSeventh,
	intervals.Octave,
}

var Aeolian = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MajorSecond,
	intervals.MinorThird,
	intervals.PerfectFourth,
	intervals.PerfectFifth,
	intervals.MinorSixth,
	intervals.MinorSeventh,
	intervals.Octave,
}

var Locrian = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MinorSecond,
	intervals.MinorThird,
	intervals.PerfectFourth,
	intervals.DiminishedFifth,
	intervals.MinorSixth,
	intervals.MinorSeventh,
	intervals.Octave,
}

var HarmonicMinor = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MajorSecond,
	intervals.MinorThird,
	intervals.PerfectFourth,
	intervals.PerfectFifth,
	intervals.MinorSixth,
	intervals.MajorSeventh,
	intervals.Octave,
}

var MelodicMinor = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MajorSecond,
	intervals.MinorThird,
	intervals.PerfectFourth,
	intervals.PerfectFifth,
	intervals.MajorSixth,
	intervals.MajorSeventh,
	intervals.Octave,
	intervals.MinorSeventh,
	intervals.MinorSixth,
	intervals.PerfectFifth,
	intervals.PerfectFourth,
	intervals.MinorThird,
	intervals.MajorSecond,
	intervals.PerfectUnison,
}

var JazzMinor = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MajorSecond,
	intervals.MinorThird,
	intervals.PerfectFourth,
	intervals.PerfectFifth,
	intervals.MajorSixth,
	intervals.MajorSeventh,
	intervals.Octave,
}

var MajorPentatonic = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MajorSecond,
	intervals.MajorThird,
	intervals.PerfectFifth,
	intervals.MajorSixth,
	intervals.Octave,
}

var EgyptianPentatonic = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MajorSecond,
	intervals.PerfectFourth,
	intervals.PerfectFifth,
	intervals.MinorSeventh,
	intervals.Octave,
}

var BluesMinor = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MinorThird,
	intervals.PerfectFourth,
	intervals.MinorSixth,
	intervals.MinorSeventh,
	intervals.Octave,
}

var BluesMajor = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MajorSecond,
	intervals.PerfectFourth,
	intervals.PerfectFifth,
	intervals.MajorSixth,
	intervals.Octave,
}

var MinorPentatonic = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MinorThird,
	intervals.PerfectFourth,
	intervals.PerfectFifth,
	intervals.MinorSeventh,
	intervals.Octave,
}

var scales = map[string][]intervals.Interval{
	"ionian":              Ionian,
	"dorian":              Dorian,
	"phyrigian":           Phrygian,
	"lydian":              Lydian,
	"mixolydian":          Mixolydian,
	"aeolian":             Aeolian,
	"locrian":             Locrian,
	"harmonic minor":      HarmonicMinor,
	"melodic minor":       MelodicMinor,
	"jazz minor":          JazzMinor,
	"major pentatonic":    MajorPentatonic,
	"egyptian pentatonic": EgyptianPentatonic,
	"blues minor":         BluesMinor,
	"blues major":         BluesMajor,
	"minor pentatonic":    MinorPentatonic,
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

	var structure []intervals.Interval
	switch scaleName {
	case "":
		fallthrough
	case "M":
		fallthrough
	case "ionian":
		fallthrough
	case "maj":
		// major scale, ionian mode
		structure = Ionian

	case "dorian":
		// dorian mode
		structure = Dorian

	case "phrygian":
		// phrygian mode
		structure = Phrygian

	case "lydian":
		// lydian mode
		structure = Lydian

	case "mixolydian":
		// mixolydian mode
		structure = Mixolydian

	case "aeolian":
		fallthrough
	case "m":
		fallthrough
	case "min":
		// natural minor scale, or aeolian
		structure = Aeolian

	case "locrian":
		structure = Locrian

	case "harmonicMinor":
		// harmonic minor scale
		structure = HarmonicMinor

	case "melodicMinor":
		// melodic minor scale
		structure = MelodicMinor

	case "jazzMinor":
		// jazz minor scale
		structure = JazzMinor

	case "majorPentatonic":
		// major pentatonic scale
		structure = MajorPentatonic

	case "egyptian":
		// egyptian pentatonic scale
		structure = EgyptianPentatonic

	case "BluesMinor":
		// blues minor pentatonic scale
		structure = BluesMinor

	case "BluesMajor":
		// blues major pentatonic scale
		structure = BluesMajor

	case "minorPentatonic":
		// minor pentatonic scale
		structure = MinorPentatonic

	default:
		return nil, fmt.Errorf("unknown scale name: %s", scaleName)

	}
	return &Scale{
		root:      *n,
		structure: structure,
	}, nil
}

func (scale *Scale) Name() string {
	for name, structure := range scales {
		if len(structure) != len(scale.structure) {
			continue
		}

		for i := 0; i < len(structure); i++ {
			if structure[i] != scale.structure[i] {
				continue
			}
		}
		return name
	}

	return ""
}

func (scale *Scale) Notes() []notes.Note {
	ret := make([]notes.Note, 0)
	for _, interval := range scale.structure {
		ret = append(ret, *scale.root.Interval(interval))
	}
	return ret
}

func (scale *Scale) Triads() []chords.Chord {
	// skip octave in triads construction
	// XXX - fix for melodic minor
	scaleNotes := scale.Notes()[0 : len(scale.Notes())-1]
	ret := make([]chords.Chord, 0)
	for i := 0; i < len(scaleNotes); i++ {
		chord := chords.FromNotes([]notes.Note{scaleNotes[i], scaleNotes[(i+2)%len(scaleNotes)], scaleNotes[(i+4)%len(scaleNotes)]})
		ret = append(ret, chord)
	}
	return ret
}

func (scale *Scale) Sevenths() []chords.Chord {
	// skip octave in triads construction
	// XXX - fix for melodic minor
	scaleNotes := scale.Notes()[0 : len(scale.Notes())-1]
	ret := make([]chords.Chord, 0)
	for i := 0; i < len(scaleNotes); i++ {
		chord := chords.FromNotes([]notes.Note{scaleNotes[i], scaleNotes[(i+2)%len(scaleNotes)], scaleNotes[(i+4)%len(scaleNotes)], scaleNotes[(i+6)%len(scaleNotes)]})
		ret = append(ret, chord)
	}
	return ret
}

func (scale *Scale) Triad(degree Degree) chords.Chord {
	// skip octave in triads construction
	// XXX - fix for melodic minor
	scaleNotes := scale.Notes()[0 : len(scale.Notes())-1]
	return chords.FromNotes([]notes.Note{scaleNotes[degree], scaleNotes[(int(degree)+2)%len(scaleNotes)], scaleNotes[(int(degree)+4)%len(scaleNotes)]})
}

func (scale *Scale) Seventh(degree Degree) chords.Chord {
	// skip octave in seventh construction
	// XXX - fix for melodic minor
	scaleNotes := scale.Notes()[0 : len(scale.Notes())-1]
	return chords.FromNotes([]notes.Note{scaleNotes[degree], scaleNotes[(int(degree)+2)%len(scaleNotes)], scaleNotes[(int(degree)+4)%len(scaleNotes)], scaleNotes[(int(degree)+6)%len(scaleNotes)]})
}

func FromChord(chord chords.Chord) []Scale {
	ret := make([]Scale, 0)
	for scaleName := range scales {
		root := *chord.Root()

		scale := &Scale{
			root:      root,
			structure: scales[scaleName],
		}

		count := 0
		for _, chordNote := range chord.Notes() {
			found := false
			for _, scaleNote := range scale.Notes() {
				if chordNote.Inharmonic(scaleNote) {
					found = true
					count++
					break
				}
			}
			if !found {
				break
			}
		}

		if count > 1 {
			ret = append(ret, *scale)
		}
	}
	return ret
}
