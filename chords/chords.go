package chords

import (
	"fmt"
	"strings"

	"github.com/poolpOrg/go-harmony/intervals"
	"github.com/poolpOrg/go-harmony/notes"
)

type Structure []intervals.Interval

type chord struct {
	root      notes.Note
	structure Structure
	bass      intervals.Interval
}
type Chord chord

var (
	MajorTriad Structure = Structure{
		intervals.PerfectUnison,
		intervals.MajorThird,
		intervals.PerfectFifth,
	}

	MinorTriad Structure = Structure{
		intervals.PerfectUnison,
		intervals.MinorThird,
		intervals.PerfectFifth,
	}

	AugmentedTriad Structure = Structure{
		intervals.PerfectUnison,
		intervals.MajorThird,
		intervals.AugmentedFifth,
	}

	DiminishedTriad Structure = Structure{
		intervals.PerfectUnison,
		intervals.MinorThird,
		intervals.DiminishedFifth,
	}

	PowerChord Structure = Structure{
		intervals.PerfectUnison,
		intervals.PerfectFifth,
	}

	Sixth Structure = Structure{
		intervals.PerfectUnison,
		intervals.MajorThird,
		intervals.PerfectFifth,
		intervals.MajorSixth,
	}

	MinorSixth Structure = Structure{
		intervals.PerfectUnison,
		intervals.MinorThird,
		intervals.PerfectFifth,
		intervals.MajorSixth,
	}

	DominantSeventh Structure = Structure{
		intervals.PerfectUnison,
		intervals.MajorThird,
		intervals.PerfectFifth,
		intervals.MinorSeventh,
	}

	MajorSeventh Structure = Structure{
		intervals.PerfectUnison,
		intervals.MajorThird,
		intervals.PerfectFifth,
		intervals.MajorSeventh,
	}

	MinorMajorSeventh Structure = Structure{
		intervals.PerfectUnison,
		intervals.MinorThird,
		intervals.PerfectFifth,
		intervals.MajorSeventh,
	}

	MinorSeventh Structure = Structure{
		intervals.PerfectUnison,
		intervals.MinorThird,
		intervals.PerfectFifth,
		intervals.MinorSeventh,
	}

	AugmentedMajorSeventh Structure = Structure{
		intervals.PerfectUnison,
		intervals.MajorThird,
		intervals.AugmentedFifth,
		intervals.MajorSeventh,
	}

	AugmentedSeventh Structure = Structure{
		intervals.PerfectUnison,
		intervals.MajorThird,
		intervals.AugmentedFifth,
		intervals.MinorSeventh,
	}

	HalfDiminishedSeventh Structure = Structure{
		intervals.PerfectUnison,
		intervals.MinorThird,
		intervals.DiminishedFifth,
		intervals.MinorSeventh,
	}

	DiminishedSeventh Structure = Structure{
		intervals.PerfectUnison,
		intervals.MinorThird,
		intervals.DiminishedFifth,
		intervals.DiminishedSeventh,
	}

	DominantSeventhFlatFive Structure = Structure{
		intervals.PerfectUnison,
		intervals.MajorThird,
		intervals.DiminishedFifth,
		intervals.MinorSeventh,
	}

	MajorNinth               Structure = append(MajorSeventh, intervals.MajorNinth)
	DominantNinth            Structure = append(DominantSeventh, intervals.MajorNinth)
	DominantMinorNinth       Structure = append(DominantSeventh, intervals.MinorNinth)
	MinorMajorNinth          Structure = append(MinorMajorSeventh, intervals.MajorNinth)
	MinorNinth               Structure = append(MinorSeventh, intervals.MajorNinth)
	AugmentedMajorNinth      Structure = append(AugmentedMajorSeventh, intervals.MajorNinth)
	AugmentedDominantNinth   Structure = append(AugmentedSeventh, intervals.MajorNinth)
	HalfDiminishedNinth      Structure = append(HalfDiminishedSeventh, intervals.MajorNinth)
	HalfDiminishedMinorNinth Structure = append(HalfDiminishedSeventh, intervals.MinorNinth)
	DiminishedNinth          Structure = append(DiminishedSeventh, intervals.MajorNinth)
	DiminishedMinorNinth     Structure = append(DiminishedSeventh, intervals.MinorNinth)

	Eleventh           Structure = append(DominantNinth, intervals.PerfectEleventh)
	MajorEleventh      Structure = append(MajorNinth, intervals.PerfectEleventh)
	MinorMajorEleventh Structure = append(MinorMajorNinth, intervals.PerfectEleventh)
	MinorEleventh      Structure = append(MinorNinth, intervals.PerfectEleventh)

	AugmentedMajorEleventh Structure = append(AugmentedMajorNinth, intervals.PerfectEleventh)
	AugmentedEleventh      Structure = append(AugmentedDominantNinth, intervals.PerfectEleventh)
	HalfDiminishedEleventh Structure = append(HalfDiminishedNinth, intervals.PerfectEleventh)
	DiminishedEleventh     Structure = append(DiminishedNinth, intervals.PerfectEleventh)

	MajorThirteenth          Structure = append(MajorEleventh, intervals.MajorThirteenth)
	Thirteenth               Structure = append(Eleventh, intervals.MajorThirteenth)
	MinorMajorThirteenth     Structure = append(MinorMajorEleventh, intervals.MajorThirteenth)
	MinorThirteenth          Structure = append(MinorEleventh, intervals.MajorThirteenth)
	AugmentedMajorThirteenth Structure = append(AugmentedMajorEleventh, intervals.MajorThirteenth)
	AugmentedThirteenth      Structure = append(AugmentedEleventh, intervals.MajorThirteenth)
	HalfDiminishedThirteenth Structure = append(HalfDiminishedEleventh, intervals.MajorThirteenth)

	AddNinth      Structure = append(MajorTriad, intervals.MajorNinth)
	AddEleventh   Structure = append(MajorTriad, intervals.PerfectEleventh)
	AddThirteenth Structure = append(MajorTriad, intervals.MajorThirteenth)

	SusSecond Structure = Structure{
		intervals.PerfectUnison,
		intervals.MajorSecond,
		intervals.PerfectFifth,
	}

	SusFourth Structure = Structure{
		intervals.PerfectUnison,
		intervals.PerfectFourth,
		intervals.PerfectFifth,
	}
)

func (structure Structure) Equals(target Structure) bool {
	if len(structure) != len(target) {
		return false
	}
	for i := 0; i < len(structure); i++ {
		if structure[i] != target[i] {
			return false
		}
	}
	return true
}

func (structure Structure) Name() string {
	if structure.Equals(MajorTriad) {
		return "maj"
	}
	if structure.Equals(MinorTriad) {
		return "min"
	}
	if structure.Equals(AugmentedTriad) {
		return "aug"
	}
	if structure.Equals(DiminishedTriad) {
		return "dim"
	}
	if structure.Equals(PowerChord) {
		return "5"
	}
	if structure.Equals(Sixth) {
		return "6"
	}
	if structure.Equals(MinorSixth) {
		return "min6"
	}
	if structure.Equals(DominantSeventh) {
		return "7"
	}
	if structure.Equals(MajorSeventh) {
		return "maj7"
	}
	if structure.Equals(MinorMajorSeventh) {
		return "-M7"
	}
	if structure.Equals(MinorSeventh) {
		return "min7"
	}
	if structure.Equals(AugmentedMajorSeventh) {
		return "+M7"
	}
	if structure.Equals(AugmentedSeventh) {
		return "aug7"
	}
	if structure.Equals(HalfDiminishedSeventh) {
		return "m7b5"
	}
	if structure.Equals(DiminishedSeventh) {
		return "dim7"
	}
	if structure.Equals(DominantSeventhFlatFive) {
		return "7dim5"
	}

	if structure.Equals(MajorNinth) {
		return "maj9"
	}
	if structure.Equals(DominantNinth) {
		return "9"
	}
	if structure.Equals(DominantMinorNinth) {
		return "7b9"
	}
	if structure.Equals(MinorMajorNinth) {
		return "-M9"
	}
	if structure.Equals(MinorNinth) {
		return "min9"
	}
	if structure.Equals(AugmentedMajorNinth) {
		return "+M9"
	}
	if structure.Equals(AugmentedDominantNinth) {
		return "aug9"
	}
	if structure.Equals(HalfDiminishedNinth) {
		return "øb9"
	}
	if structure.Equals(DiminishedNinth) {
		return "dim9"
	}
	if structure.Equals(DiminishedMinorNinth) {
		return "dimb9"
	}

	if structure.Equals(Eleventh) {
		return "11"
	}
	if structure.Equals(MajorEleventh) {
		return "maj11"
	}
	if structure.Equals(MinorMajorEleventh) {
		return "-M11"
	}
	if structure.Equals(MinorEleventh) {
		return "min11"
	}
	if structure.Equals(AugmentedMajorEleventh) {
		return "+M11"
	}
	if structure.Equals(AugmentedEleventh) {
		return "aug11"
	}
	if structure.Equals(HalfDiminishedEleventh) {
		return "ø11"
	}
	if structure.Equals(DiminishedEleventh) {
		return "dim11"
	}

	if structure.Equals(MajorThirteenth) {
		return "maj13"
	}
	if structure.Equals(Thirteenth) {
		return "13"
	}
	if structure.Equals(MinorMajorThirteenth) {
		return "-M13"
	}
	if structure.Equals(AugmentedMajorThirteenth) {
		return "+M13"
	}
	if structure.Equals(AugmentedThirteenth) {
		return "aug13"
	}
	if structure.Equals(HalfDiminishedThirteenth) {
		return "ø13"
	}

	if structure.Equals(AddNinth) {
		return "add9"
	}
	if structure.Equals(AddEleventh) {
		return "add11"
	}
	if structure.Equals(AddThirteenth) {
		return "add13"
	}

	if structure.Equals(SusSecond) {
		return "sus2"
	}
	if structure.Equals(SusFourth) {
		return "sus4"
	}

	return ""
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

	inversion := ""
	inversionIndex := strings.Index(chordName, "/")
	if inversionIndex != -1 {
		inversion = chordName[inversionIndex+1:]
		chordName = chordName[:inversionIndex]
	}

	var structure Structure
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

	c := &Chord{
		root:      *n,
		structure: structure,
		bass:      structure[0],
	}

	if inversion != "" {
		if inversion == n.Name() {
			return nil, fmt.Errorf("inversion note %s is root of chord: %s", inversion, c.Name())
		}
		n2, err := notes.Parse(inversion)
		if err != nil {
			return nil, err
		}

		found := false
		for offset, interval := range c.structure {
			if n.Interval(interval).Name() == n2.Name() {
				c.bass = structure[offset]
				found = true
			}
		}
		if !found {
			return nil, fmt.Errorf("inversion note %s not found in chord: %s", inversion, c.Name())
		}
	}

	return c, nil
}

func (chord *Chord) Name() string {
	// TODO: construct chord name by analyzing intervals
	name := chord.root.Name()
	structureName := chord.structure.Name()
	if structureName == "" {
		panic("unknown structure name")
	}
	name += structureName
	if chord.bass != intervals.PerfectUnison {
		name += "/" + chord.root.Interval(chord.bass).Name()
	}
	return name
}

func (chord *Chord) Root() *notes.Note {
	return &chord.root
}

func (chord *Chord) Notes() []notes.Note {
	ret := make([]notes.Note, 0)
	ret = append(ret, *chord.root.Interval(chord.bass))
	for _, interval := range chord.structure {
		if chord.bass != interval {
			ret = append(ret, *chord.root.Interval(interval))
		}
	}
	return ret
}

func FromNotes(notes []notes.Note) Chord {
	structure := Structure{}

	root := notes[0]
	structure = append(structure, intervals.PerfectUnison)
	for _, note := range notes[1:] {
		targetPosition := note.Position()
		if targetPosition < root.Position() {
			targetPosition += 7
		}

		targetSemitone := note.Semitone()
		if targetSemitone < root.Semitone() {
			targetSemitone += 12
		}

		structure = append(structure, intervals.New(targetPosition-root.Position(), targetSemitone-root.Semitone()))
	}

	return Chord{
		root:      notes[0],
		structure: structure,
	}
}
