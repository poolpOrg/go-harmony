package tunings

import (
	"testing"
)

type tuningNameTestCase struct {
	tuning Tuning
	want   string
}

type tuningReferenceTestCase struct {
	tuning Tuning
	want   float64
}

type tunerFrequencyTestCase struct {
	tuning Tuning
	input  uint8
	want   float64
}

func TestTuning_Tunings(t *testing.T) {
	var testCases = []Tuning{
		A432, A434, A436, A438, A440, A442, A444, A446,
	}

	tunings := Tunings()
	if len(testCases) != len(tunings) {
		t.Fatalf(`testCases lenght != tunings length, got %d, want %d`, len(tunings), len(testCases))
	}
	for offset, testCase := range testCases {
		if tunings[offset] != testCase {
			t.Fatalf(`Tuning<%s>.Name() = %s, want %s`, testCase.Name(), testCase.Name(), tunings[offset].Name())
		}
	}
}

func TestTuning_Tuning_Name(t *testing.T) {
	var testCases = []tuningNameTestCase{
		{A432, "A432"},
		{A434, "A434"},
		{A436, "A436"},
		{A438, "A438"},
		{A440, "A440"},
		{A442, "A442"},
		{A444, "A444"},
		{A446, "A446"},
	}

	for _, testCase := range testCases {
		if testCase.tuning.Name() != testCase.want {
			t.Fatalf(`Tuning<%s>.Name() = %s, want %s`, testCase.tuning.Name(), testCase.tuning.Name(), testCase.want)
		}
	}
}

func TestTuning_Tuning_Reference(t *testing.T) {
	var testCases = []tuningReferenceTestCase{
		{A432, 432.},
		{A434, 434.},
		{A436, 436.},
		{A438, 438.},
		{A440, 440.},
		{A442, 442.},
		{A444, 444.},
		{A446, 446.},
	}

	for _, testCase := range testCases {
		if testCase.tuning.Reference() != testCase.want {
			t.Fatalf(`Tuning<%s>.Name() = %f, want %f`, testCase.tuning.Name(), testCase.tuning.Reference(), testCase.want)
		}
	}
}

func TestTuning_Systems(t *testing.T) {
	var testCases = []TuningSystem{
		EqualTemperament,
	}

	tuningSystems := Systems()
	if len(testCases) != len(tuningSystems) {
		t.Fatalf(`testCases length != tunings length, got %d, want %d`, len(tuningSystems), len(testCases))
	}
	for offset, testCase := range testCases {
		if tuningSystems[offset].Name() != testCase.Name() {
			t.Fatalf(`Tuning<%s>.Name() = %s, want %s`, testCase.Name(), testCase.Name(), tuningSystems[offset].Name())
		}
	}
}

func TestTuning_NewTuner(t *testing.T) {
	tunings := Tunings()
	for _, tuning := range tunings {
		n := NewTuner(EqualTemperament, tuning)
		if n == nil {
			t.Fatalf(`could not create tuner for tuning %s`, tuning.Name())
		}
	}
}

func TestTuning_TunerFrequency(t *testing.T) {
	var testCases = []tunerFrequencyTestCase{
		{A432, 53, 432.},
		{A432, 53 - 12, 216.},
		{A432, 53 + 12, 864.},

		{A434, 53, 434.},
		{A434, 53 - 12, 217.},
		{A434, 53 + 12, 868.},

		{A436, 53, 436.},
		{A436, 53 - 12, 218.},
		{A436, 53 + 12, 872.},

		{A438, 53, 438.},
		{A438, 53 - 12, 219.},
		{A438, 53 + 12, 876.},

		{A440, 53, 440.},
		{A440, 53 - 12, 220.},
		{A440, 53 + 12, 880.},

		{A442, 53, 442.},
		{A442, 53 - 12, 221.},
		{A442, 53 + 12, 884.},

		{A444, 53, 444.},
		{A444, 53 - 12, 222.},
		{A444, 53 + 12, 888.},

		{A446, 53, 446.},
		{A446, 53 - 12, 223.},
		{A446, 53 + 12, 892.},
	}

	for _, testCase := range testCases {
		tuner := NewTuner(EqualTemperament, testCase.tuning)
		got := tuner.Frequency(testCase.input)
		if got != testCase.want {
			t.Fatalf(`tuner(%s).Frequency(%d) = %f, want %f`, testCase.tuning.Name(), testCase.input, got, testCase.want)
		}
	}
}
