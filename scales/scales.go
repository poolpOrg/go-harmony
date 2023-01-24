package scales

import (
	"fmt"
	"sort"

	"github.com/poolpOrg/go-harmony/chords"
	"github.com/poolpOrg/go-harmony/intervals"
	"github.com/poolpOrg/go-harmony/notes"
)

type scale struct {
	root      notes.Note
	structure []intervals.Interval
}
type Scale scale

type Degree uint8

var (
	Tonic       Degree = 0
	Supertonic  Degree = 1
	Mediant     Degree = 2
	Subdominant Degree = 3
	Dominant    Degree = 4
	Submediant  Degree = 5
	LeadingTone Degree = 6
)

func (degree *Degree) Name() string {
	switch *degree {
	case Tonic:
		return "Tonic"
	case Supertonic:
		return "Supertonic"
	case Mediant:
		return "Mediant"
	case Subdominant:
		return "Subdominant"
	case Dominant:
		return "Dominant"
	case Submediant:
		return "Submediant"
	case LeadingTone:
		return "LeadingTone"
	}
	panic("unsupported degree")
}

func Degrees() []Degree {
	return []Degree{Tonic, Supertonic, Mediant, Subdominant, Dominant, Submediant, LeadingTone}
}

var Ionian = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MajorSecond,
	intervals.MajorThird,
	intervals.PerfectFourth,
	intervals.PerfectFifth,
	intervals.MajorSixth,
	intervals.MajorSeventh,
}

var Dorian = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MajorSecond,
	intervals.MinorThird,
	intervals.PerfectFourth,
	intervals.PerfectFifth,
	intervals.MajorSixth,
	intervals.MinorSeventh,
}

var Phrygian = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MinorSecond,
	intervals.MinorThird,
	intervals.PerfectFourth,
	intervals.PerfectFifth,
	intervals.MinorSixth,
	intervals.MinorSeventh,
}

var Lydian = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MajorSecond,
	intervals.MajorThird,
	intervals.AugmentedFourth,
	intervals.PerfectFifth,
	intervals.MajorSixth,
	intervals.MajorSeventh,
}

var Mixolydian = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MajorSecond,
	intervals.MajorThird,
	intervals.PerfectFourth,
	intervals.PerfectFifth,
	intervals.MajorSixth,
	intervals.MinorSeventh,
}

var Aeolian = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MajorSecond,
	intervals.MinorThird,
	intervals.PerfectFourth,
	intervals.PerfectFifth,
	intervals.MinorSixth,
	intervals.MinorSeventh,
}

var Locrian = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MinorSecond,
	intervals.MinorThird,
	intervals.PerfectFourth,
	intervals.DiminishedFifth,
	intervals.MinorSixth,
	intervals.MinorSeventh,
}

var HarmonicMinor = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MajorSecond,
	intervals.MinorThird,
	intervals.PerfectFourth,
	intervals.PerfectFifth,
	intervals.MinorSixth,
	intervals.MajorSeventh,
}

var HarmonicMajor = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MajorSecond,
	intervals.MajorThird,
	intervals.PerfectFourth,
	intervals.PerfectFifth,
	intervals.MinorSixth,
	intervals.MajorSeventh,
}

var JazzMinor = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MajorSecond,
	intervals.MinorThird,
	intervals.PerfectFourth,
	intervals.PerfectFifth,
	intervals.MajorSixth,
	intervals.MajorSeventh,
}

var MajorPentatonic = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MajorSecond,
	intervals.MajorThird,
	intervals.None,
	intervals.PerfectFifth,
	intervals.MajorSixth,
	intervals.None,
}

var EgyptianPentatonic = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MajorSecond,
	intervals.None,
	intervals.PerfectFourth,
	intervals.PerfectFifth,
	intervals.None,
	intervals.MinorSeventh,
}

var BluesMinor = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.None,
	intervals.MinorThird,
	intervals.PerfectFourth,
	intervals.None,
	intervals.MinorSixth,
	intervals.MinorSeventh,
}

var BluesMajor = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MajorSecond,
	intervals.None,
	intervals.PerfectFourth,
	intervals.PerfectFifth,
	intervals.MajorSixth,
	intervals.None,
}

var MinorPentatonic = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.None,
	intervals.MinorThird,
	intervals.PerfectFourth,
	intervals.PerfectFifth,
	intervals.None,
	intervals.MinorSeventh,
}

var scales = map[string][]intervals.Interval{
	"ionian":         Ionian,
	"dorian":         Dorian,
	"phyrigian":      Phrygian,
	"lydian":         Lydian,
	"mixolydian":     Mixolydian,
	"aeolian":        Aeolian,
	"locrian":        Locrian,
	"harmonic minor": HarmonicMinor,
	"harmonic major": HarmonicMajor,
	//"melodic minor":       MelodicMinor,
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

	case "harmonicMajor":
		// harmonic major scale
		structure = HarmonicMajor

		/*
			case "melodicMinor":
				// melodic minor scale
				structure = MelodicMinor
		*/
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

		found := true
		for i := 0; i < len(structure); i++ {
			if structure[i] != scale.structure[i] {
				found = false
				break
			}
		}
		if !found {
			continue
		}
		return name
	}

	return ""
}

func (scale *Scale) Notes() []*notes.Note {
	ret := make([]*notes.Note, 0)
	for degree, interval := range scale.structure {
		if scale.structure[degree] != intervals.None {
			ret = append(ret, scale.root.Interval(interval))
		} else {
			ret = append(ret, nil)
		}
	}
	return ret
}

func (scale *Scale) Triads() []*chords.Chord {
	ret := make([]*chords.Chord, 0)
	for _, degree := range Degrees() {
		if scale.structure[degree] != intervals.None {
			chord := scale.Triad(degree)
			ret = append(ret, &chord)
		} else {
			ret = append(ret, nil)
		}
	}
	return ret

}

func (scale *Scale) Sevenths() []*chords.Chord {
	ret := make([]*chords.Chord, 0)
	for _, degree := range Degrees() {
		if scale.structure[degree] != intervals.None {
			chord := scale.Seventh(degree)
			ret = append(ret, &chord)
		} else {
			ret = append(ret, nil)
		}
	}
	return ret
}

func (scale *Scale) Triad(degree Degree) chords.Chord {
	// skip octave in triads construction
	// XXX - fix for melodic minor
	scaleNotes := scale.Notes()
	root := scaleNotes[int(degree)]
	third := scaleNotes[(int(degree)+2)%len(scaleNotes)]
	fifth := scaleNotes[(int(degree)+4)%len(scaleNotes)]

	if third.Position() < root.Position() || int(degree)+2 >= len(scaleNotes) {
		third = third.Interval(intervals.Octave)
	}
	if fifth.Position() < root.Position() || int(degree)+4 >= len(scaleNotes) {
		fifth = fifth.Interval(intervals.Octave)
	}

	return chords.FromNotes([]notes.Note{*root, *third, *fifth})
}

func (scale *Scale) Seventh(degree Degree) chords.Chord {
	// skip octave in seventh construction
	// XXX - fix for melodic minor
	scaleNotes := scale.Notes()
	root := scaleNotes[int(degree)]
	third := scaleNotes[(int(degree)+2)%len(scaleNotes)]
	fifth := scaleNotes[(int(degree)+4)%len(scaleNotes)]
	seventh := scaleNotes[(int(degree)+6)%len(scaleNotes)]

	if third.Position() < root.Position() || int(degree)+2 >= len(scaleNotes) {
		third = third.Interval(intervals.Octave)
	}
	if fifth.Position() < root.Position() || int(degree)+4 >= len(scaleNotes) {
		fifth = fifth.Interval(intervals.Octave)
	}
	if seventh.Position() < root.Position() || int(degree)+6 >= len(scaleNotes) {
		seventh = seventh.Interval(intervals.Octave)
	}

	return chords.FromNotes([]notes.Note{*root, *third, *fifth, *seventh})
}

func (scale *Scale) NotesInChord(chord chords.Chord) int {
	count := 0
	for _, chordNote := range chord.Notes() {
		for _, scaleNote := range scale.Notes() {
			if scaleNote != nil {
				if chordNote.Enharmonic(*scaleNote) {
					count++
					break
				}
			}
		}
	}
	return count
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
			for _, scaleNote := range scale.Notes() {
				if scaleNote != nil {
					if chordNote.Enharmonic(*scaleNote) {
						count++
						break
					}
				}
			}
		}

		if count > 1 {
			ret = append(ret, *scale)
		}
	}

	sort.SliceStable(ret, func(i, j int) bool {
		if ret[i].NotesInChord(chord) == ret[j].NotesInChord(chord) {
			return ret[i].Name() < ret[j].Name()
		} else {
			return ret[i].NotesInChord(chord) > ret[j].NotesInChord(chord)
		}
	})

	return ret
}
