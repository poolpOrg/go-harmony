package scales

import (
	"testing"
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

func TestScales_Parse(t *testing.T) {
}

func TestScales_Name(t *testing.T) {
}

func TestScales_Notes(t *testing.T) {
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
