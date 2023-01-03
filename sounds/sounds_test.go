package sounds

import (
	"testing"
)

func TestSound(t *testing.T) {
	sound := NewSound(440.0, .1)

	freq := sound.Frequency()
	if freq != 440. {
		t.Fatalf(`sound.Frequency() = %f, want %f`, freq, 440.)
	}

	amplitude := sound.Amplitude()
	if amplitude != .1 {
		t.Fatalf(`sound.Amplitude() = %f, want %f`, amplitude, .1)
	}

}
