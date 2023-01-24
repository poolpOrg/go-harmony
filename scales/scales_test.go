package scales

import (
	"testing"

	"github.com/poolpOrg/go-harmony/intervals"
)

var degrees = map[Degree]uint8{
	Tonic:       0,
	Supertonic:  1,
	Mediant:     2,
	Subdominant: 3,
	Dominant:    4,
	Submediant:  5,
	LeadingTone: 6,
}

func TestScales_Degrees(t *testing.T) {
	for degree, value := range degrees {
		if uint8(degree) != value {
			t.Fatalf(`Degree %v has invalid value, want %d, got %d`, degree.Name(), uint8(degree), value)
		}
	}
}

func TestScales_DegreeName(t *testing.T) {
	testCases := []string{"Tonic", "Supertonic", "Mediant", "Subdominant", "Dominant", "Submediant", "LeadingTone"}
	for degree, value := range degrees {
		if degree.Name() != testCases[value] {
			t.Fatalf(`Degree %v has invalid name, want %s, got %s`, uint8(degree), testCases[value], degree.Name())
		}
	}
}

func TestScales_Parse(t *testing.T) {
	scales := map[string][]intervals.Interval{
		"":                Ionian,
		"M":               Ionian,
		"ionian":          Ionian,
		"maj":             Ionian,
		"dorian":          Dorian,
		"phrygian":        Phrygian,
		"lydian":          Lydian,
		"mixolydian":      Mixolydian,
		"aeolian":         Aeolian,
		"m":               Aeolian,
		"min":             Aeolian,
		"locrian":         Locrian,
		"harmonicMinor":   HarmonicMinor,
		"harmonicMajor":   HarmonicMajor,
		"jazzMinor":       JazzMinor,
		"majorPentatonic": MajorPentatonic,
		"egyptian":        EgyptianPentatonic,
		"BluesMinor":      BluesMinor,
		"BluesMajor":      BluesMajor,
		"minorPentatonic": MinorPentatonic,
	}
	for name, _ := range scales {
		_, err := Parse("C" + name)
		if err != nil {
			t.Fatalf(`could not parse scale name %s`, "C"+name)
		}
		//if s.structure != structure {
		//	t.Fatalf(`could not parse scale name %s`, "C"+name)
		//}
	}
}

func TestScales_Name(t *testing.T) {
}

type scaleNotesTestCase struct {
	input string
	want  []string
}

func TestScales_Notes(t *testing.T) {
	testCases := []scaleNotesTestCase{
		{"Cbb", []string{"Cbb", "Dbb", "Ebb", "Fbb", "Gbb", "Abb", "Bbb"}},
		{"Cb", []string{"Cb", "Db", "Eb", "Fb", "Gb", "Ab", "Bb"}},
		{"C", []string{"C", "D", "E", "F", "G", "A", "B"}},
		{"C#", []string{"C#", "D#", "E#", "F#", "G#", "A#", "B#"}},
		{"C##", []string{"C##", "D##", "E##", "F##", "G##", "A##", "B##"}},

		{"Cbbmin", []string{"Cbb", "Dbb", "Ebbb", "Fbb", "Gbb", "Abbb", "Bbbb"}},
		{"Cbmin", []string{"Cb", "Db", "Ebb", "Fb", "Gb", "Abb", "Bbb"}},
		{"Cmin", []string{"C", "D", "Eb", "F", "G", "Ab", "Bb"}},
		{"C#min", []string{"C#", "D#", "E", "F#", "G#", "A", "B"}},
		{"C##min", []string{"C##", "D##", "E#", "F##", "G##", "A#", "B#"}},
	}

	for _, testCase := range testCases {
		scale, _ := Parse(testCase.input)
		for degree, note := range scale.Notes() {
			got := note.Name()
			if testCase.want[degree] != got {
				t.Fatalf(`Scale(%s).Notes()[%d].Name() = %s, want %s`, testCase.input, degree, got, testCase.want[degree])
			}
		}
	}
}

func TestScales_Triads(t *testing.T) {
}

func TestScales_Sevenths(t *testing.T) {
}

func TestScales_Triad(t *testing.T) {
}

func TestScales_Seventh(t *testing.T) {
}

func TestScales_NotesInChord(t *testing.T) {
}

func TestScales_FromChord(t *testing.T) {
}
