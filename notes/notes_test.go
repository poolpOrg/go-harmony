package notes

import (
	"fmt"
	"testing"

	"github.com/poolpOrg/go-harmony/intervals"
	"github.com/poolpOrg/go-harmony/naturals"
	"github.com/poolpOrg/go-harmony/octaves"
)

func TestNotes_Parse(t *testing.T) {
	naturalsList := naturals.Naturals()
	octavesList := octaves.Octaves()
	for _, natural := range naturalsList {
		for _, octave := range octavesList {

			name := natural.Name()

			want := name
			got, err := Parse(want)
			if err != nil {
				t.Fatalf(`notes.Parse(%s) = error(%s)`, want, err)
			}
			if got.Name() != want {
				t.Fatalf(`notes.Parse(%s) = %s, want %s`, want, got.Name(), want)
			}

			want = fmt.Sprintf("%s%d", name, octave.Position())
			got, err = Parse(want)
			if err != nil {
				t.Fatalf(`notes.Parse(%s) = error(%s)`, want, err)
			}
			if got.OctaveName() != want {
				t.Fatalf(`notes.Parse(%s) = %s, want %s`, want, got.Name(), want)
			}

			want = name + "#"
			got, err = Parse(want)
			if err != nil {
				t.Fatalf(`notes.Parse(%s) = error(%s)`, want, err)
			}
			if got.Name() != want {
				t.Fatalf(`notes.Parse(%s) = %s, want %s`, want, got.Name(), want)
			}

			want = fmt.Sprintf("%s%s%d", name, "#", octave.Position())
			got, err = Parse(want)
			if err != nil {
				t.Fatalf(`notes.Parse(%s) = error(%s)`, want, err)
			}
			if got.OctaveName() != want {
				t.Fatalf(`notes.Parse(%s) = %s, want %s`, want, got.Name(), want)
			}

			want = name + "##"
			got, err = Parse(want)
			if err != nil {
				t.Fatalf(`notes.Parse(%s) = error(%s)`, want, err)
			}
			if got.Name() != want {
				t.Fatalf(`notes.Parse(%s) = %s, want %s`, want, got.Name(), want)
			}

			want = fmt.Sprintf("%s%s%d", name, "##", octave.Position())
			got, err = Parse(want)
			if err != nil {
				t.Fatalf(`notes.Parse(%s) = error(%s)`, want, err)
			}
			if got.OctaveName() != want {
				t.Fatalf(`notes.Parse(%s) = %s, want %s`, want, got.Name(), want)
			}

			want = name + "b"
			got, err = Parse(want)
			if err != nil {
				t.Fatalf(`notes.Parse(%s) = error(%s)`, want, err)
			}
			if got.Name() != want {
				t.Fatalf(`notes.Parse(%s) = %s, want %s`, want, got.Name(), want)
			}

			want = fmt.Sprintf("%s%s%d", name, "b", octave.Position())
			got, err = Parse(want)
			if err != nil {
				t.Fatalf(`notes.Parse(%s) = error(%s)`, want, err)
			}
			if got.OctaveName() != want {
				t.Fatalf(`notes.Parse(%s) = %s, want %s`, want, got.Name(), want)
			}

			want = name + "bb"
			got, err = Parse(want)
			if err != nil {
				t.Fatalf(`notes.Parse(%s) = error(%s)`, want, err)
			}
			if got.Name() != want {
				t.Fatalf(`notes.Parse(%s) = %s, want %s`, want, got.Name(), want)
			}

			want = fmt.Sprintf("%s%s%d", name, "bb", octave.Position())
			got, err = Parse(want)
			if err != nil {
				t.Fatalf(`notes.Parse(%s) = error(%s)`, want, err)
			}
			if got.OctaveName() != want {
				t.Fatalf(`notes.Parse(%s) = %s, want %s`, want, got.Name(), want)
			}

			want = fmt.Sprintf("%s%s", name, "z")
			_, err = Parse(want)
			if err == nil {
				t.Fatalf(`notes.Parse(%s) = error(%s)`, want, err)
			}

			want = fmt.Sprintf("%s%s%d", name, "z", octave.Position())
			_, err = Parse(want)
			if err == nil {
				t.Fatalf(`notes.Parse(%s) = error(%s)`, want, err)
			}

		}
	}
}

func TestNotes_Name(t *testing.T) {
}

func TestNotes_OctaveName(t *testing.T) {
}

type notesIntervalTestCase struct {
	interval intervals.Interval
	input    string
	want     string
}

func TestNotes_Interval(t *testing.T) {
	testCases := []notesIntervalTestCase{
		{intervals.PerfectUnison, "C", "C4"},
		{intervals.AugmentedUnison, "C", "C#4"},

		{intervals.DiminishedSecond, "C", "Dbb4"},
		{intervals.MinorSecond, "C", "Db4"},
		{intervals.MajorSecond, "C", "D4"},
		{intervals.AugmentedSecond, "C", "D#4"},

		{intervals.DiminishedSeventh, "C", "Bbb4"},
		{intervals.MinorSeventh, "C", "Bb4"},
		{intervals.MajorSeventh, "C", "B4"},
		{intervals.AugmentedSeventh, "C", "B#4"},

		{intervals.DiminishedOctave, "C", "Cb5"},
		{intervals.Octave, "C", "C5"},
		{intervals.AugmentedOctave, "C", "C#5"},

		{intervals.DiminishedFifteenth, "C", "Cb6"},
		{intervals.PerfectFifteenth, "C", "C6"},
		{intervals.AugmentedFifteenth, "C", "C#6"},
	}

	for _, testCase := range testCases {
		note, _ := Parse(testCase.input)
		got := note.Interval(testCase.interval).OctaveName()
		if got != testCase.want {
			t.Fatalf(`notes.Interval(%s) = %s, want %s`, testCase.interval.Name(), got, testCase.want)
		}
	}
}

type notesDistanceTestCase struct {
	root   string
	target string
	want   intervals.Interval
}

func TestNotes_Distance(t *testing.T) {
	testCases := []notesDistanceTestCase{
		{"C", "C", intervals.PerfectUnison},
		{"C", "D", intervals.MajorSecond},
		{"C", "E", intervals.MajorThird},
		{"C", "F", intervals.PerfectFourth},
		{"C", "G", intervals.PerfectFifth},
		{"C", "A", intervals.MajorSixth},
		{"C", "B", intervals.MajorSeventh},
		{"C", "C5", intervals.Octave},
		{"C", "D5", intervals.MajorNinth},
		{"C", "E5", intervals.MajorTenth},
		{"C", "F5", intervals.PerfectEleventh},
		{"C", "G5", intervals.PerfectTwelfth},
		{"C", "A5", intervals.MajorThirteenth},
		{"C", "B5", intervals.MajorFourteenth},
		{"C", "C6", intervals.PerfectFifteenth},

		{"D", "D", intervals.PerfectUnison},
		{"D", "E", intervals.MajorSecond},
		{"D", "F#", intervals.MajorThird},
		{"D", "G", intervals.PerfectFourth},
		{"D", "A", intervals.PerfectFifth},
		{"D", "B", intervals.MajorSixth},
		{"D", "C#5", intervals.MajorSeventh},
		{"D", "D5", intervals.Octave},
		{"D", "E5", intervals.MajorNinth},
		{"D", "F#5", intervals.MajorTenth},
		{"D", "G5", intervals.PerfectEleventh},
		{"D", "A5", intervals.PerfectTwelfth},
		{"D", "B5", intervals.MajorThirteenth},
		{"D", "C#6", intervals.MajorFourteenth},
		{"D", "D6", intervals.PerfectFifteenth},
	}

	for _, testCase := range testCases {
		root, _ := Parse(testCase.root)
		target, _ := Parse(testCase.target)

		got := root.Distance(*target)
		if got != testCase.want {
			t.Fatalf(`note(%s).Distance(%s) = %s, want %s`, testCase.root, testCase.target, got.Name(), testCase.want.Name())
		}
	}
}

func TestNotes_Position(t *testing.T) {
	naturalsList := naturals.Naturals()
	for _, natural := range naturalsList {
		n, _ := Parse(natural.Name())
		if n.Position() != natural.Position() {
			t.Fatalf(`notes.Position(%s) = %s, want %s`, natural.Name(), n.Name(), natural.Name())
		}
	}
}

func TestNotes_Semitones(t *testing.T) {
}

func TestNotes_Frequency(t *testing.T) {
}

func TestNotes_Enharmonic(t *testing.T) {
}

func TestNotes_Octave(t *testing.T) {
}

func TestNotes_Previous(t *testing.T) {
}

func TestNotes_Next(t *testing.T) {
}

func TestNotes_PreviousChromatic(t *testing.T) {
}

func TestNotes_NextChromatic(t *testing.T) {
}

func TestNotes_Lower(t *testing.T) {
}

func TestNotes_Raise(t *testing.T) {
}

func TestNotes_SetOctave(t *testing.T) {
}

func TestNotes_MIDI(t *testing.T) {
}
