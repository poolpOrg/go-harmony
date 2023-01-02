package tunings

import (
	"testing"
)

func TestTuning432Frequencies(t *testing.T) {
	tuner := NewTuner(EqualTemperament, A432)

	freq := tuner.Frequency(53)
	if freq != 432. {
		t.Fatalf(`tuner.Frequency(0) = %f, want %f`, freq, 432.)
	}
}

func TestTuning434Frequencies(t *testing.T) {
	tuner := NewTuner(EqualTemperament, A434)

	freq := tuner.Frequency(53)
	if freq != 434. {
		t.Fatalf(`tuner.Frequency(0) = %f, want %f`, freq, 434.)
	}
}

func TestTuning436Frequencies(t *testing.T) {
	tuner := NewTuner(EqualTemperament, A436)

	freq := tuner.Frequency(53)
	if freq != 436. {
		t.Fatalf(`tuner.Frequency(0) = %f, want %f`, freq, 436.)
	}
}

func TestTuning438Frequencies(t *testing.T) {
	tuner := NewTuner(EqualTemperament, A438)

	freq := tuner.Frequency(53)
	if freq != 438. {
		t.Fatalf(`tuner.Frequency(0) = %f, want %f`, freq, 438.)
	}
}

func TestTuning440Frequencies(t *testing.T) {
	tuner := NewTuner(EqualTemperament, A440)

	freq := tuner.Frequency(53)
	if freq != 440. {
		t.Fatalf(`tuner.Frequency(0) = %f, want %f`, freq, 440.)
	}
}

func TestTuning442Frequencies(t *testing.T) {
	tuner := NewTuner(EqualTemperament, A442)

	freq := tuner.Frequency(53)
	if freq != 442. {
		t.Fatalf(`tuner.Frequency(0) = %f, want %f`, freq, 442.)
	}
}

func TestTuning444Frequencies(t *testing.T) {
	tuner := NewTuner(EqualTemperament, A444)

	freq := tuner.Frequency(53)
	if freq != 444. {
		t.Fatalf(`tuner.Frequency(0) = %f, want %f`, freq, 444.)
	}
}

func TestTuning446Frequencies(t *testing.T) {
	tuner := NewTuner(EqualTemperament, A446)

	freq := tuner.Frequency(53)
	if freq != 446. {
		t.Fatalf(`tuner.Frequency(0) = %f, want %f`, freq, 446.)
	}
}
