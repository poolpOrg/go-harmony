package intervals

import (
	"testing"
)

var intervalsList = []Interval{
	PerfectUnison,
	AugmentedUnison,
	DiminishedSecond,
	MinorSecond,
	MajorSecond,
	AugmentedSecond,
	DiminishedThird,
	MinorThird,
	MajorThird,
	AugmentedThird,
	DiminishedFourth,
	PerfectFourth,
	AugmentedFourth,
	DiminishedFifth,
	PerfectFifth,
	AugmentedFifth,
	DiminishedSixth,
	MinorSixth,
	MajorSixth,
	AugmentedSixth,
	DiminishedSeventh,
	MinorSeventh,
	MajorSeventh,
	AugmentedSeventh,
	DiminishedOctave,
	Octave,
	AugmentedOctave,
	DiminishedNinth,
	MinorNinth,
	MajorNinth,
	AugmentedNinth,
	DiminishedTenth,
	MinorTenth,
	MajorTenth,
	AugmentedTenth,
	DiminishedEleventh,
	PerfectEleventh,
	AugmentedEleventh,
	DiminishedTwelfth,
	PerfectTwelfth,
	AugmentedTwelfth,
	DiminishedThirteenth,
	MinorThirteenth,
	MajorThirteenth,
	AugmentedThirteenth,
	DiminishedFourteenth,
	MinorFourteenth,
	MajorFourteenth,
	AugmentedFourteenth,
	DiminishedFifteenth,
	PerfectFifteenth,
	AugmentedFifteenth,
}

var positions = map[string]uint{
	"1":    0,
	"1aug": 0,

	"2dim": 1,
	"2min": 1,
	"2maj": 1,
	"2aug": 1,

	"3dim": 2,
	"3min": 2,
	"3maj": 2,
	"3aug": 2,

	"4dim": 3,
	"4":    3,
	"4aug": 3,

	"5dim": 4,
	"5":    4,
	"5aug": 4,

	"6dim": 5,
	"6min": 5,
	"6maj": 5,
	"6aug": 5,

	"7dim": 6,
	"7min": 6,
	"7maj": 6,
	"7aug": 6,

	"8dim": 7,
	"8":    7,
	"8aug": 7,

	"9dim": 8,
	"9min": 8,
	"9maj": 8,
	"9aug": 8,

	"10dim": 9,
	"10min": 9,
	"10maj": 9,
	"10aug": 9,

	"11dim": 10,
	"11":    10,
	"11aug": 10,

	"12dim": 11,
	"12":    11,
	"12aug": 11,

	"13dim": 12,
	"13min": 12,
	"13maj": 12,
	"13aug": 12,

	"14dim": 13,
	"14min": 13,
	"14maj": 13,
	"14aug": 13,

	"15dim": 14,
	"15":    14,
	"15aug": 14,
}

var semitones = map[string]uint{
	"1":    0,
	"1aug": 1,

	"2dim": 0,
	"2min": 1,
	"2maj": 2,
	"2aug": 3,

	"3dim": 2,
	"3min": 3,
	"3maj": 4,
	"3aug": 5,

	"4dim": 4,
	"4":    5,
	"4aug": 6,

	"5dim": 6,
	"5":    7,
	"5aug": 8,

	"6dim": 7,
	"6min": 8,
	"6maj": 9,
	"6aug": 10,

	"7dim": 9,
	"7min": 10,
	"7maj": 11,
	"7aug": 12,

	"8dim": 11,
	"8":    12,
	"8aug": 13,

	"9dim": 12,
	"9min": 13,
	"9maj": 14,
	"9aug": 15,

	"10dim": 14,
	"10min": 15,
	"10maj": 16,
	"10aug": 17,

	"11dim": 16,
	"11":    17,
	"11aug": 18,

	"12dim": 18,
	"12":    19,
	"12aug": 20,

	"13dim": 19,
	"13min": 20,
	"13maj": 21,
	"13aug": 22,

	"14dim": 21,
	"14min": 22,
	"14maj": 23,
	"14aug": 24,

	"15dim": 23,
	"15":    24,
	"15aug": 25,
}

func TestIntervals_length(t *testing.T) {
	intervalsList := Intervals()
	got := len(intervalsList)
	want := 52
	if got != want {
		t.Fatalf(`intervals.Intervals() = %d, want %d`, got, want)
	}
}

func TestIntervals_Intervals(t *testing.T) {
	intervalsList := Intervals()
	for offset, value := range Intervals() {
		if intervalsList[offset] != value {
			t.Fatalf(`intervals.Intervals() has incorrect item at index %d`, offset)
		}
	}
}

func TestIntervals_FromName(t *testing.T) {
	for _, interval := range Intervals() {
		n, err := FromName(interval.Name())
		if err != nil {
			t.Fatalf(`intervals.FromName(%s) = %v, want %v`, interval.Name(), n, interval)
		}
	}
	n, err := FromName("13sup")
	if err == nil {
		t.Fatalf(`intervals.FromName(%s) = %v, want %v`, "13sup", n, nil)
	}
}

func TestIntervals_Position(t *testing.T) {
	for _, interval := range Intervals() {
		if value, exists := positions[interval.Name()]; !exists {
			t.Fatalf(`interval(%s) does not exist`, interval.Name())
		} else {
			if interval.Position() != value {
				t.Fatalf(`interval(%s).Position() = %d, want %d`, interval.Name(), interval.Position(), value)
			}
		}
	}
}

func TestIntervals_Semitones(t *testing.T) {
	for _, interval := range Intervals() {
		if value, exists := semitones[interval.Name()]; !exists {
			t.Fatalf(`interval(%s) does not exist`, interval.Name())
		} else {
			if interval.Semitones() != value {
				t.Fatalf(`interval(%s).Semitones() = %d, want %d`, interval.Name(), interval.Semitones(), value)
			}
		}
	}
}

func TestIntervals_New(t *testing.T) {
	for _, interval := range Intervals() {
		if New(interval.Position(), interval.Semitones()) != interval {
			t.Fatalf(`intervals.New(%d, %d) does not produce correct interval`,
				interval.Position(), interval.Semitones())
		}
	}
}

func TestIntervals_Relative(t *testing.T) {
	for _, interval := range Intervals() {
		if interval.Relative() != New(7-interval.Position(), 12-interval.Semitones()) {
			t.Fatalf(`intervals.Relative(%d, %d) does not produce correct interval`,
				interval.Position(), interval.Semitones())
		}
	}
}
