package notes

import (
	"fmt"
	"testing"

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

func TestNotes_Interval(t *testing.T) {
}

func TestNotes_Distance(t *testing.T) {
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
