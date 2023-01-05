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

		{intervals.PerfectUnison, "Cb", "Cb4"},
		{intervals.AugmentedUnison, "Cb", "C4"},
		{intervals.DiminishedSecond, "Cb", "Dbbb4"},
		{intervals.MinorSecond, "Cb", "Dbb4"},
		{intervals.MajorSecond, "Cb", "Db4"},
		{intervals.AugmentedSecond, "Cb", "D4"},
		{intervals.DiminishedSeventh, "Cb", "Bbbb4"},
		{intervals.MinorSeventh, "Cb", "Bbb4"},
		{intervals.MajorSeventh, "Cb", "Bb4"},
		{intervals.AugmentedSeventh, "Cb", "B4"},
		{intervals.DiminishedOctave, "Cb", "Cbb5"},
		{intervals.Octave, "Cb", "Cb5"},
		{intervals.AugmentedOctave, "Cb", "C5"},
		{intervals.DiminishedFifteenth, "Cb", "Cbb6"},
		{intervals.PerfectFifteenth, "Cb", "Cb6"},
		{intervals.AugmentedFifteenth, "Cb", "C6"},

		{intervals.PerfectUnison, "C#", "C#4"},
		{intervals.AugmentedUnison, "C#", "C##4"},
		{intervals.DiminishedSecond, "C#", "Db4"},
		{intervals.MinorSecond, "C#", "D4"},
		{intervals.MajorSecond, "C#", "D#4"},
		{intervals.AugmentedSecond, "C#", "D##4"},
		{intervals.DiminishedSeventh, "C#", "Bb4"},
		{intervals.MinorSeventh, "C#", "B4"},
		{intervals.MajorSeventh, "C#", "B#4"},
		{intervals.AugmentedSeventh, "C#", "B##4"},
		{intervals.DiminishedOctave, "C#", "C5"},
		{intervals.Octave, "C#", "C#5"},
		{intervals.AugmentedOctave, "C#", "C##5"},
		{intervals.DiminishedFifteenth, "C#", "C6"},
		{intervals.PerfectFifteenth, "C#", "C#6"},
		{intervals.AugmentedFifteenth, "C#", "C##6"},

		{intervals.PerfectUnison, "B", "B4"},
		{intervals.AugmentedUnison, "B", "B#4"},
		{intervals.DiminishedSecond, "B", "Cb5"},
		{intervals.MinorSecond, "B", "C5"},
		{intervals.MajorSecond, "B", "C#5"},
		{intervals.AugmentedSecond, "B", "C##5"},
		{intervals.DiminishedSeventh, "B", "Ab5"},
		{intervals.MinorSeventh, "B", "A5"},
		{intervals.MajorSeventh, "B", "A#5"},
		{intervals.AugmentedSeventh, "B", "A##5"},
		{intervals.DiminishedOctave, "B", "Bb5"},
		{intervals.Octave, "B", "B5"},
		{intervals.AugmentedOctave, "B", "B#5"},
		{intervals.DiminishedFifteenth, "B", "Bb6"},
		{intervals.PerfectFifteenth, "B", "B6"},
		{intervals.AugmentedFifteenth, "B", "B#6"},

		{intervals.PerfectUnison, "Bb", "Bb4"},
		{intervals.AugmentedUnison, "Bb", "B4"},
		{intervals.DiminishedSecond, "Bb", "Cbb5"},
		{intervals.MinorSecond, "Bb", "Cb5"},
		{intervals.MajorSecond, "Bb", "C5"},
		{intervals.AugmentedSecond, "Bb", "C#5"},
		{intervals.DiminishedSeventh, "Bb", "Abb5"},
		{intervals.MinorSeventh, "Bb", "Ab5"},
		{intervals.MajorSeventh, "Bb", "A5"},
		{intervals.AugmentedSeventh, "Bb", "A#5"},
		{intervals.DiminishedOctave, "Bb", "Bbb5"},
		{intervals.Octave, "Bb", "Bb5"},
		{intervals.AugmentedOctave, "Bb", "B5"},
		{intervals.DiminishedFifteenth, "Bb", "Bbb6"},
		{intervals.PerfectFifteenth, "Bb", "Bb6"},
		{intervals.AugmentedFifteenth, "Bb", "B6"},

		{intervals.PerfectUnison, "B#", "B#4"},
		{intervals.AugmentedUnison, "B#", "B##4"},
		{intervals.DiminishedSecond, "B#", "C5"},
		{intervals.MinorSecond, "B#", "C#5"},
		{intervals.MajorSecond, "B#", "C##5"},
		{intervals.AugmentedSecond, "B#", "C###5"},
		{intervals.DiminishedSeventh, "B#", "A5"},
		{intervals.MinorSeventh, "B#", "A#5"},
		{intervals.MajorSeventh, "B#", "A##5"},
		{intervals.AugmentedSeventh, "B#", "A###5"},
		{intervals.DiminishedOctave, "B#", "B5"},
		{intervals.Octave, "B#", "B#5"},
		{intervals.AugmentedOctave, "B#", "B##5"},
		{intervals.DiminishedFifteenth, "B#", "B6"},
		{intervals.PerfectFifteenth, "B#", "B#6"},
		{intervals.AugmentedFifteenth, "B#", "B##6"},
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
