package chords

import (
	"fmt"
	"sort"
	"strings"

	"github.com/poolpOrg/go-harmony/intervals"
	"github.com/poolpOrg/go-harmony/notes"
)

type Structure []intervals.Interval

type chord struct {
	root      notes.Note
	structure Structure
	inversion intervals.Interval
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

	FlatFifthTriad Structure = Structure{
		intervals.PerfectUnison,
		intervals.MajorThird,
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

	DominantSeventhSusSecond Structure = Structure{
		intervals.PerfectUnison,
		intervals.MajorSecond,
		intervals.PerfectFifth,
		intervals.MinorSeventh,
	}

	DominantSeventhSusFourth Structure = Structure{
		intervals.PerfectUnison,
		intervals.PerfectFourth,
		intervals.PerfectFifth,
		intervals.MinorSeventh,
	}

	MajorSeventhSusSecond Structure = Structure{
		intervals.PerfectUnison,
		intervals.MajorSecond,
		intervals.PerfectFifth,
		intervals.MajorSeventh,
	}

	MajorSeventhSusFourth Structure = Structure{
		intervals.PerfectUnison,
		intervals.PerfectFourth,
		intervals.PerfectFifth,
		intervals.MajorSeventh,
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

	SusSecondSusFourth Structure = Structure{
		intervals.PerfectUnison,
		intervals.MajorSecond,
		intervals.PerfectFourth,
		intervals.PerfectFifth,
	}
)

func Structures() []Structure {
	ret := make([]Structure, 0)

	ret = append(ret, MajorTriad)
	ret = append(ret, MinorTriad)
	ret = append(ret, AugmentedTriad)
	ret = append(ret, DiminishedTriad)
	ret = append(ret, FlatFifthTriad)

	ret = append(ret, PowerChord)

	ret = append(ret, Sixth)
	ret = append(ret, MinorSixth)

	ret = append(ret, DominantSeventh)
	ret = append(ret, MajorSeventh)
	ret = append(ret, MinorMajorSeventh)
	ret = append(ret, MinorSeventh)
	ret = append(ret, AugmentedMajorSeventh)
	ret = append(ret, AugmentedSeventh)
	ret = append(ret, HalfDiminishedSeventh)
	ret = append(ret, DiminishedSeventh)
	ret = append(ret, DominantSeventhFlatFive)
	ret = append(ret, DominantSeventhSusSecond)
	ret = append(ret, DominantSeventhSusFourth)
	ret = append(ret, MajorSeventhSusSecond)
	ret = append(ret, MajorSeventhSusFourth)

	ret = append(ret, MajorNinth)
	ret = append(ret, DominantNinth)
	ret = append(ret, DominantMinorNinth)
	ret = append(ret, MinorMajorNinth)
	ret = append(ret, MinorNinth)
	ret = append(ret, AugmentedMajorNinth)
	ret = append(ret, AugmentedDominantNinth)
	ret = append(ret, HalfDiminishedNinth)
	ret = append(ret, HalfDiminishedMinorNinth)
	ret = append(ret, DiminishedNinth)
	ret = append(ret, DiminishedMinorNinth)

	ret = append(ret, MajorEleventh)
	ret = append(ret, Eleventh)
	ret = append(ret, MinorMajorEleventh)
	ret = append(ret, MinorEleventh)
	ret = append(ret, AugmentedMajorEleventh)
	ret = append(ret, AugmentedEleventh)
	ret = append(ret, HalfDiminishedEleventh)
	ret = append(ret, DiminishedEleventh)

	ret = append(ret, MajorThirteenth)
	ret = append(ret, Thirteenth)
	ret = append(ret, MinorMajorThirteenth)
	ret = append(ret, MinorThirteenth)
	ret = append(ret, AugmentedMajorThirteenth)
	ret = append(ret, AugmentedThirteenth)
	ret = append(ret, HalfDiminishedThirteenth)

	ret = append(ret, AddNinth)
	ret = append(ret, AddEleventh)
	ret = append(ret, AddThirteenth)

	ret = append(ret, SusSecond)
	ret = append(ret, SusFourth)
	ret = append(ret, SusSecondSusFourth)

	return ret
}

func (structure Structure) Semitones() []uint8 {
	semitones := make([]uint8, 0)
	for index, interval := range structure {
		if index == 0 {
			semitones = append(semitones, 0)
		} else {
			semitones = append(semitones, uint8(interval.Semitones())-uint8(structure[0].Semitones()))
		}
	}
	return semitones
}

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

func (structure Structure) Equivalent(target Structure) bool {
	originSemitones := make(map[uint8]bool)
	targetSemitones := make(map[uint8]bool)
	originSemitonesStructure := make([]uint8, 0)
	targetSemitonesStructure := make([]uint8, 0)

	for _, interval := range structure {
		wrapped := uint8(interval.Semitones()) % 12
		if _, exists := originSemitones[wrapped]; !exists {
			originSemitones[wrapped] = true
			originSemitonesStructure = append(originSemitonesStructure, wrapped)
		}
	}

	for _, interval := range target {
		wrapped := uint8(interval.Semitones()) % 12
		if _, exists := targetSemitones[wrapped]; !exists {
			targetSemitones[wrapped] = true
			targetSemitonesStructure = append(targetSemitonesStructure, wrapped)
		}
	}

	if len(originSemitonesStructure) != len(targetSemitonesStructure) {
		return false
	}

	sort.SliceStable(originSemitonesStructure, func(i, j int) bool {
		return originSemitonesStructure[i] < originSemitonesStructure[j]
	})

	sort.SliceStable(targetSemitonesStructure, func(i, j int) bool {
		return targetSemitonesStructure[i] < targetSemitonesStructure[j]
	})

	for i, _ := range originSemitonesStructure {
		if originSemitonesStructure[i] != targetSemitonesStructure[i] {
			return false
		}
	}

	return true
}

func (structure Structure) Inversion(target Structure) *intervals.Interval {
	if len(structure) != len(target) {
		return nil
	}

	if structure.Equals(target) {
		return &intervals.PerfectUnison
	}

	c, err := Parse("C")
	if err != nil {
		panic(err)
	}
	c.structure = structure

	for i := 1; i < len(c.Notes()); i++ {
		inversionChord, err := Parse(fmt.Sprintf("%s/%s", c.Name(), c.Notes()[i].OctaveName()))
		if err != nil {
			panic(err)
		}
		substructure := Structure{}
		inversionNote := inversionChord.root.Interval(inversionChord.inversion)
		for j := 0; j < len(inversionChord.Notes()); j++ {
			substructure = append(substructure, inversionNote.Distance(inversionChord.Notes()[j]))
		}
		if substructure.Equivalent(target) {
			return &c.structure[i]
		}
	}

	return nil
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
	if structure.Equals(FlatFifthTriad) {
		return "(b5)"
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
	if structure.Equals(DominantSeventhSusSecond) {
		return "7sus2"
	}
	if structure.Equals(DominantSeventhSusFourth) {
		return "7sus4"
	}

	if structure.Equals(MajorSeventhSusSecond) {
		return "maj7sus2"
	}
	if structure.Equals(MajorSeventhSusFourth) {
		return "maj7sus4"
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
		return "ø9"
	}
	if structure.Equals(HalfDiminishedMinorNinth) {
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
	if structure.Equals(SusSecondSusFourth) {
		return "sus2sus4"
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

	case "(b5)":
		// major third, flat fifth
		structure = FlatFifthTriad

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
	case "m7(b5)":
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

	case "7sus2":
		// dominant seventh sus 2
		structure = DominantSeventhSusSecond

	case "7sus4":
		// dominant seventh sus 4
		structure = DominantSeventhSusFourth

	case "maj7sus2":
		// dominant seventh sus 2
		structure = MajorSeventhSusSecond

	case "maj7sus4":
		// dominant seventh sus 4
		structure = MajorSeventhSusFourth

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

	case "sus2sus4":
		structure = SusSecondSusFourth

	default:
		return nil, fmt.Errorf("unknown chord name: %s", chordName)
	}

	c := &Chord{
		root:      *n,
		structure: structure,
		inversion: intervals.PerfectUnison,
	}

	if inversion != "" {
		if inversion == n.Name() {
			return nil, fmt.Errorf("inversion note %s is root of chord: %s", inversion, c.Name())
		}
		n2, err := notes.Parse(inversion)
		if err != nil {
			return nil, err
		}

		if inversion != n2.OctaveName() {
			n2.SetOctave(n.Octave() - 1)
		}

		found := false
		for _, interval := range c.structure {
			if n.Interval(interval).Name() == n2.Name() {
				found = true
				c.inversion = interval
			}
		}
		if !found {
			return nil, fmt.Errorf("inversion note %s not found in chord: %s", inversion, c.Name())
		}
		c.root.SetOctave(n2.Octave())
	}

	return c, nil
}

func (chord *Chord) Name() string {
	// TODO: construct chord name by analyzing intervals
	name := chord.root.Name()
	structureName := chord.structure.Name()
	if structureName == "" {
		structureName = " ("
		for _, interval := range chord.structure[1:] {
			structureName += interval.Name() + ":" + chord.root.Interval(interval).Name() + ","
		}
		structureName += ")"
	}
	name += structureName
	if chord.inversion != intervals.PerfectUnison {
		name += "/" + chord.root.Interval(chord.inversion).Name()
	}
	return name
}

func (chord *Chord) Root() *notes.Note {
	return &chord.root
}

func (chord *Chord) Notes() []notes.Note {

	ret := make([]notes.Note, 0)
	ret = append(ret, *chord.root.Interval(chord.inversion))

	prev := &ret[0]

	if chord.inversion != intervals.PerfectUnison && prev.Octave() > chord.root.Octave() {
		prev.SetOctave(chord.root.Octave())
	}

	var begin int
	for begin = 0; begin < len(chord.structure); begin++ {
		if chord.structure[begin] == chord.inversion {
			break
		}
	}

	for offset := begin + 1; offset%len(chord.structure) != begin; offset++ {
		interval := chord.structure[offset%len(chord.structure)]
		n := chord.root.Interval(interval)
		if n.Name() != chord.root.Interval(chord.inversion).Name() {
			if n.Octave() < prev.Octave() {
				n = n.Interval(intervals.Octave)
			} else if n.Position() < prev.Position() && n.Octave() <= prev.Octave() {
				n = n.Interval(intervals.Octave)
			}
			ret = append(ret, *n)
			prev = n
		}
	}

	return ret
}

func (chord *Chord) Structure() []intervals.Interval {
	return chord.structure
}

func (chord *Chord) SetRoot(root notes.Note) {
	chord.root = root
}

func (chord *Chord) Semitones() []uint8 {
	return chord.structure.Semitones()
}

func FromNotes(notes []notes.Note) Chord {
	// sort notes by ascending order
	sort.SliceStable(notes, func(i, j int) bool {
		return notes[i].AbsoluteSemitones() < notes[j].AbsoluteSemitones()
	})

	root := notes[0]

	chordStructure := Structure{}
	chordStructure = append(chordStructure, intervals.PerfectUnison)
	chordInversion := intervals.PerfectUnison
	for _, note := range notes[1:] {
		chordStructure = append(chordStructure, *intervals.New(root.Distance(note).Position(), root.Distance(note).Semitones()))
	}

	// first try to match a general structure
	structures := Structures()
	for _, refStructure := range structures {
		if chordStructure.Equivalent(refStructure) {
			return Chord{
				root:      root,
				structure: refStructure,
				inversion: chordInversion,
			}
		}
	}

	// none found, try matching inversions
	for _, refStructure := range structures {
		inversionInterval := refStructure.Inversion(chordStructure)
		if inversionInterval == nil {
			continue
		}
		root = *root.Interval(inversionInterval.Relative())
		chordInversion = *inversionInterval
		chordStructure = refStructure
		break
	}

	return Chord{
		root:      root,
		structure: chordStructure,
		inversion: chordInversion,
	}
}

func (chord *Chord) Relative() *Chord {
	structure := make([]intervals.Interval, len(chord.structure))
	copy(structure, chord.structure)
	if structure[1] == intervals.MinorThird {
		structure[1] = intervals.MajorThird
		root := *chord.root.Interval(intervals.MinorThird)
		root.SetOctave(root.Interval(chord.inversion).Octave())
		ret := &Chord{
			root:      root,
			structure: structure,
			inversion: intervals.PerfectUnison,
		}
		return ret

	} else if structure[1] == intervals.MajorThird {
		structure[1] = intervals.MinorThird
		root := *chord.root.Interval(intervals.MajorSixth)
		root.SetOctave(root.Interval(chord.inversion).Octave())
		ret := &Chord{
			root:      root,
			structure: structure,
			inversion: intervals.PerfectUnison,
		}
		return ret
	}
	return nil
}
