package chords

import (
	"fmt"

	"github.com/poolpOrg/go-harmony/intervals"
	"github.com/poolpOrg/go-harmony/notes"
)

type chord struct {
	name      string
	root      notes.Note
	structure []intervals.Interval
}
type Chord chord

var MajorTriad = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MajorThird,
	intervals.PerfectFifth,
}

var MinorTriad = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MinorThird,
	intervals.PerfectFifth,
}

var AugmentedTriad = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MajorThird,
	intervals.AugmentedFifth,
}

var DiminishedTriad = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MinorThird,
	intervals.DiminishedFifth,
}

var PowerChord = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.PerfectFifth,
}

var Sixth = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MajorThird,
	intervals.PerfectFifth,
	intervals.MajorSixth,
}

var MinorSixth = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MinorThird,
	intervals.PerfectFifth,
	intervals.MajorSixth,
}

var DominantSeventh = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MajorThird,
	intervals.PerfectFifth,
	intervals.MinorSeventh,
}

var MajorSeventh = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MajorThird,
	intervals.PerfectFifth,
	intervals.MajorSeventh,
}

var MinorMajorSeventh = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MinorThird,
	intervals.PerfectFifth,
	intervals.MajorSeventh,
}

var MinorSeventh = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MinorThird,
	intervals.PerfectFifth,
	intervals.MinorSeventh,
}

var AugmentedMajorSeventh = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MajorThird,
	intervals.AugmentedFifth,
	intervals.MajorSeventh,
}

var AugmentedSeventh = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MajorThird,
	intervals.AugmentedFifth,
	intervals.MinorSeventh,
}

var HalfDiminishedSeventh = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MinorThird,
	intervals.DiminishedFifth,
	intervals.MinorSeventh,
}

var DiminishedSeventh = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MinorThird,
	intervals.DiminishedFifth,
	intervals.DiminishedSeventh,
}

var DominantSeventhFlatFive = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MajorThird,
	intervals.DiminishedFifth,
	intervals.MinorSeventh,
}

var MajorNinth = append(MajorSeventh, intervals.MajorNinth)
var DominantNinth = append(DominantSeventh, intervals.MajorNinth)
var DominantMinorNinth = append(DominantSeventh, intervals.MinorNinth)
var MinorMajorNinth = append(MinorMajorSeventh, intervals.MajorNinth)
var MinorNinth = append(MinorSeventh, intervals.MajorNinth)
var AugmentedMajorNinth = append(AugmentedMajorSeventh, intervals.MajorNinth)
var AugmentedDominantNinth = append(AugmentedSeventh, intervals.MajorNinth)
var HalfDiminishedNinth = append(HalfDiminishedSeventh, intervals.MajorNinth)
var HalfDiminishedMinorNinth = append(HalfDiminishedSeventh, intervals.MinorNinth)
var DiminishedNinth = append(DiminishedSeventh, intervals.MajorNinth)
var DiminishedMinorNinth = append(DiminishedSeventh, intervals.MinorNinth)

var Eleventh = append(DominantNinth, intervals.PerfectEleventh)
var MajorEleventh = append(MajorNinth, intervals.PerfectEleventh)
var MinorMajorEleventh = append(MinorMajorNinth, intervals.PerfectEleventh)
var MinorEleventh = append(MinorNinth, intervals.PerfectEleventh)

var AugmentedMajorEleventh = append(AugmentedMajorNinth, intervals.PerfectEleventh)
var AugmentedEleventh = append(AugmentedDominantNinth, intervals.PerfectEleventh)
var HalfDiminishedEleventh = append(HalfDiminishedNinth, intervals.PerfectEleventh)
var DiminishedEleventh = append(DiminishedNinth, intervals.PerfectEleventh)

var MajorThirteenth = append(MajorEleventh, intervals.MajorThirteenth)
var Thirteenth = append(Eleventh, intervals.MajorThirteenth)
var MinorMajorThirteenth = append(MinorMajorEleventh, intervals.MajorThirteenth)
var MinorThirteenth = append(MinorEleventh, intervals.MajorThirteenth)
var AugmentedMajorThirteenth = append(AugmentedMajorEleventh, intervals.MajorThirteenth)
var AugmentedThirteenth = append(AugmentedEleventh, intervals.MajorThirteenth)
var HalfDiminishedThirteenth = append(HalfDiminishedEleventh, intervals.MajorThirteenth)

var AddNinth = append(MajorTriad, intervals.MajorNinth)
var AddEleventh = append(MajorTriad, intervals.PerfectEleventh)
var AddThirteenth = append(MajorTriad, intervals.MajorThirteenth)

var SusSecond = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.MajorSecond,
	intervals.PerfectFifth,
}

var SusFourth = []intervals.Interval{
	intervals.PerfectUnison,
	intervals.PerfectFourth,
	intervals.PerfectFifth,
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

	var structure []intervals.Interval
	switch chordName {
	case "":
		fallthrough
	case "M":
		fallthrough
	case "maj":
		// major triad
		structure = MajorTriad

	case "-":
		fallthrough
	case "m":
		fallthrough
	case "min":
		// minor triad
		structure = MinorTriad

	case "+":
		fallthrough
	case "aug":
		fallthrough
	case "M#5":
		fallthrough
	case "M+5":
		// augmented triad
		structure = AugmentedTriad

	case "o":
		fallthrough
	case "dim":
		fallthrough
	case "mb5":
		fallthrough
	case "mo5":
		// diminished triad
		structure = DiminishedTriad

	case "5":
		// power chords, though not chords stricly-speaking
		structure = PowerChord

	case "M6":
		fallthrough
	case "6":
		// sixth chords
		structure = Sixth

	case "m6":
		fallthrough
	case "-6":
		fallthrough
	case "min6":
		// sixth chords
		structure = MinorSixth

	case "7":
		// dominant seventh
		structure = DominantSeventh

	case "M7":
		fallthrough
	case "Ma7":
		fallthrough
	case "maj7":
		// major seventh
		structure = MajorSeventh

	case "mM7":
		fallthrough
	case "m#7":
		fallthrough
	case "-M7":
		fallthrough
	case "minmaj7":
		// minor-major seventh
		structure = MinorMajorSeventh

	case "m7":
		fallthrough
	case "-7":
		fallthrough
	case "min7":
		// minor seventh
		structure = MinorSeventh

	case "+M7":
		fallthrough
	case "augmaj7":
		fallthrough
	case "M7#5":
		fallthrough
	case "M7+5":
		// augmented-major seventh
		structure = AugmentedMajorSeventh

	case "+7":
		fallthrough
	case "aug7":
		fallthrough
	case "7#5":
		fallthrough
	case "7+5":
		// augmented seventh
		structure = AugmentedSeventh

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
		structure = HalfDiminishedSeventh

	case "o7":
		fallthrough
	case "dim7":
		// diminished seventh
		structure = DiminishedSeventh

	case "7b5":
		fallthrough
	case "7dim5":
		// dominant seventh flat five
		structure = DominantSeventhFlatFive

	case "M9":
		fallthrough
	case "maj9":
		// major ninth
		structure = MajorNinth

	case "9":
		// dominant ninth
		structure = DominantNinth

	case "7b9":
		// dominant minor ninth
		structure = DominantMinorNinth

	case "mM9":
		fallthrough
	case "-M9":
		fallthrough
	case "minmaj9":
		// minor-major ninth
		structure = MinorMajorNinth

	case "m9":
		fallthrough
	case "-9":
		fallthrough
	case "min9":
		// minor ninth
		structure = MinorNinth

	case "+M9":
		fallthrough
	case "augmaj9":
		// augmented major ninth
		structure = AugmentedMajorNinth

	case "9#5":
		fallthrough
	case "aug9":
		// augmented dominant ninth
		structure = AugmentedDominantNinth

	case "ø9":
		// half-diminished ninth
		structure = HalfDiminishedNinth

	case "øb9":
		// half-diminished ninth
		structure = HalfDiminishedMinorNinth

	case "o9":
		fallthrough
	case "dim9":
		// diminished ninth
		structure = DiminishedNinth

	case "ob9":
		fallthrough
	case "dimb9":
		// diminished ninth
		structure = DiminishedMinorNinth

	case "11":
		// eleventh
		structure = Eleventh

	case "M11":
		fallthrough
	case "maj11":
		// major eleventh
		structure = MajorEleventh

	case "mM11":
		fallthrough
	case "-M11":
		fallthrough
	case "minmaj11":
		// minor-major eleventh
		structure = MinorMajorEleventh

	case "m11":
		fallthrough
	case "-11":
		fallthrough
	case "min11":
		// minor eleventh
		structure = MinorEleventh

	case "+M11":
		fallthrough
	case "augmaj11":
		// augmented major eleventh
		structure = AugmentedMajorEleventh

	case "11#5":
		fallthrough
	case "aug11":
		// augmented eleventh
		structure = AugmentedEleventh

	case "ø11":
		// half-diminished eleventh
		structure = HalfDiminishedEleventh

	case "o11":
		fallthrough
	case "dim11":
		// diminished ninth
		structure = DiminishedEleventh

	case "M13":
		fallthrough
	case "maj13":
		// major thirteenth
		structure = MajorThirteenth

	case "13":
		// thirteenth
		structure = Thirteenth

	case "mM13":
		fallthrough
	case "-M13":
		fallthrough
	case "minmaj13":
		// minor-major thirteenth
		structure = MinorMajorThirteenth

	case "m13":
		fallthrough
	case "-13":
		fallthrough
	case "min13":
		// minor-major thirteenth
		structure = MinorThirteenth

	case "+M13":
		fallthrough
	case "augmajM13":
		// augmented major thirteenth
		structure = AugmentedMajorThirteenth

	case "13#5":
		fallthrough
	case "aug13":
		// augmented thirteenth
		structure = AugmentedThirteenth

	case "ø13":
		// half-diminished thirteenth
		structure = HalfDiminishedThirteenth

		// TODO: add slash chord inversions

	case "add9":
		fallthrough
	case "+9":
		structure = AddNinth

	case "add11":
		fallthrough
	case "+11":
		structure = AddEleventh

	case "add13":
		fallthrough
	case "+13":
		structure = AddThirteenth

	case "sus2":
		structure = SusSecond

	case "sus4":
		structure = SusFourth

	default:
		return nil, fmt.Errorf("unknown chord name: %s", chordName)
	}

	return &Chord{
		root:      *n,
		structure: structure,
	}, nil
}

func (chord *Chord) Name() string {
	// TODO: construct chord name by analyzing intervals
	return chord.root.Name()
}

func (chord *Chord) Notes() []notes.Note {
	ret := make([]notes.Note, 0)
	for _, interval := range chord.structure {
		ret = append(ret, *chord.root.Interval(interval))
	}
	return ret
}
