package intervals

import (
	"fmt"
)

type interval struct {
	pos      uint
	semitone uint
}

type Interval interval

var (
	PerfectUnison   Interval = Interval{pos: 0, semitone: 0}
	AugmentedUnison Interval = Interval{pos: 0, semitone: 1}

	DiminishedSecond Interval = Interval{pos: 1, semitone: 0}
	MinorSecond      Interval = Interval{pos: 1, semitone: 1}
	MajorSecond      Interval = Interval{pos: 1, semitone: 2}
	AugmentedSecond  Interval = Interval{pos: 1, semitone: 3}

	DiminishedThird Interval = Interval{pos: 2, semitone: 2}
	MinorThird      Interval = Interval{pos: 2, semitone: 3}
	MajorThird      Interval = Interval{pos: 2, semitone: 4}
	AugmentedThird  Interval = Interval{pos: 2, semitone: 5}

	DiminishedFourth Interval = Interval{pos: 3, semitone: 4}
	PerfectFourth    Interval = Interval{pos: 3, semitone: 5}
	AugmentedFourth  Interval = Interval{pos: 3, semitone: 6}

	DiminishedFifth Interval = Interval{pos: 4, semitone: 6}
	PerfectFifth    Interval = Interval{pos: 4, semitone: 7}
	AugmentedFifth  Interval = Interval{pos: 4, semitone: 8}

	DiminishedSixth Interval = Interval{pos: 5, semitone: 7}
	MinorSixth      Interval = Interval{pos: 5, semitone: 8}
	MajorSixth      Interval = Interval{pos: 5, semitone: 9}
	AugmentedSixth  Interval = Interval{pos: 5, semitone: 10}

	DiminishedSeventh Interval = Interval{pos: 6, semitone: 9}
	MinorSeventh      Interval = Interval{pos: 6, semitone: 10}
	MajorSeventh      Interval = Interval{pos: 6, semitone: 11}
	AugmentedSeventh  Interval = Interval{pos: 6, semitone: 12}

	DiminishedOctave Interval = Interval{pos: 7, semitone: 11}
	Octave           Interval = Interval{pos: 7, semitone: 12}
	AugmentedOctave  Interval = Interval{pos: 7, semitone: 13}

	DiminishedNinth Interval = Interval{pos: 8, semitone: 12}
	MinorNinth      Interval = Interval{pos: 8, semitone: 13}
	MajorNinth      Interval = Interval{pos: 8, semitone: 14}
	AugmentedNinth  Interval = Interval{pos: 8, semitone: 15}

	DiminishedTenth Interval = Interval{pos: 9, semitone: 14}
	MinorTenth      Interval = Interval{pos: 9, semitone: 15}
	MajorTenth      Interval = Interval{pos: 9, semitone: 16}
	AugmentedTenth  Interval = Interval{pos: 9, semitone: 17}

	DiminishedEleventh Interval = Interval{pos: 10, semitone: 16}
	PerfectEleventh    Interval = Interval{pos: 10, semitone: 17}
	AugmentedEleventh  Interval = Interval{pos: 10, semitone: 18}

	DiminishedTwelfth Interval = Interval{pos: 11, semitone: 18}
	PerfectTwelfth    Interval = Interval{pos: 11, semitone: 19}
	AugmentedTwelfth  Interval = Interval{pos: 11, semitone: 20}

	DiminishedThirteenth Interval = Interval{pos: 12, semitone: 19}
	MinorThirteenth      Interval = Interval{pos: 12, semitone: 20}
	MajorThirteenth      Interval = Interval{pos: 12, semitone: 21}
	AugmentedThirteenth  Interval = Interval{pos: 12, semitone: 22}

	DiminishedFourteenth Interval = Interval{pos: 13, semitone: 21}
	MinorFourteenth      Interval = Interval{pos: 13, semitone: 22}
	MajorFourteenth      Interval = Interval{pos: 13, semitone: 23}
	AugmentedFourteenth  Interval = Interval{pos: 13, semitone: 24}

	DiminishedFifteenth Interval = Interval{pos: 14, semitone: 23}
	PerfectFifteenth    Interval = Interval{pos: 14, semitone: 24}
	AugmentedFifteenth  Interval = Interval{pos: 14, semitone: 25}
)

func Intervals() []Interval {
	ret := make([]Interval, 0)

	ret = append(ret, PerfectUnison)
	ret = append(ret, AugmentedUnison)

	ret = append(ret, DiminishedSecond)
	ret = append(ret, MinorSecond)
	ret = append(ret, MajorSecond)
	ret = append(ret, AugmentedSecond)

	ret = append(ret, DiminishedThird)
	ret = append(ret, MinorThird)
	ret = append(ret, MajorThird)
	ret = append(ret, AugmentedThird)

	ret = append(ret, DiminishedFourth)
	ret = append(ret, PerfectFourth)
	ret = append(ret, AugmentedFourth)

	ret = append(ret, DiminishedFifth)
	ret = append(ret, PerfectFifth)
	ret = append(ret, AugmentedFifth)

	ret = append(ret, DiminishedSixth)
	ret = append(ret, MinorSixth)
	ret = append(ret, MajorSixth)
	ret = append(ret, AugmentedSixth)

	ret = append(ret, DiminishedSeventh)
	ret = append(ret, MinorSeventh)
	ret = append(ret, MajorSeventh)
	ret = append(ret, AugmentedSeventh)

	ret = append(ret, DiminishedOctave)
	ret = append(ret, Octave)
	ret = append(ret, AugmentedOctave)

	ret = append(ret, DiminishedNinth)
	ret = append(ret, MinorNinth)
	ret = append(ret, MajorNinth)
	ret = append(ret, AugmentedNinth)

	ret = append(ret, DiminishedTenth)
	ret = append(ret, MinorTenth)
	ret = append(ret, MajorTenth)
	ret = append(ret, AugmentedTenth)

	ret = append(ret, DiminishedEleventh)
	ret = append(ret, PerfectEleventh)
	ret = append(ret, AugmentedEleventh)

	ret = append(ret, DiminishedTwelfth)
	ret = append(ret, PerfectTwelfth)
	ret = append(ret, AugmentedTwelfth)

	ret = append(ret, DiminishedThirteenth)
	ret = append(ret, MinorThirteenth)
	ret = append(ret, MajorThirteenth)
	ret = append(ret, AugmentedThirteenth)

	ret = append(ret, DiminishedFourteenth)
	ret = append(ret, MinorFourteenth)
	ret = append(ret, MajorFourteenth)
	ret = append(ret, AugmentedFourteenth)

	ret = append(ret, DiminishedFifteenth)
	ret = append(ret, PerfectFifteenth)
	ret = append(ret, AugmentedFifteenth)

	return ret
}

func New(pos uint, semitone uint) Interval {
	return Interval{pos: pos, semitone: semitone}

}

func (interval Interval) Name() string {
	switch interval {
	case PerfectUnison:
		return "1"
	case AugmentedUnison:
		return "1aug"

	case DiminishedSecond:
		return "2dim"
	case MinorSecond:
		return "2min"
	case MajorSecond:
		return "2maj"
	case AugmentedSecond:
		return "2aug"

	case DiminishedThird:
		return "3dim"
	case MinorThird:
		return "3min"
	case MajorThird:
		return "3maj"
	case AugmentedThird:
		return "3aug"

	case DiminishedFourth:
		return "4dim"
	case PerfectFourth:
		return "4"
	case AugmentedFourth:
		return "4aug"

	case DiminishedFifth:
		return "5dim"
	case PerfectFifth:
		return "5"
	case AugmentedFifth:
		return "5aug"

	case DiminishedSixth:
		return "6dim"
	case MinorSixth:
		return "6min"
	case MajorSixth:
		return "6maj"
	case AugmentedSixth:
		return "6aug"

	case DiminishedSeventh:
		return "7dim"
	case MinorSeventh:
		return "7min"
	case MajorSeventh:
		return "7maj"
	case AugmentedSeventh:
		return "7aug"

	case DiminishedOctave:
		return "8dim"
	case Octave:
		return "8"
	case AugmentedOctave:
		return "8aug"

	case DiminishedNinth:
		return "9dim"
	case MinorNinth:
		return "9min"
	case MajorNinth:
		return "9maj"
	case AugmentedNinth:
		return "9aug"

	case DiminishedTenth:
		return "10dim"
	case MinorTenth:
		return "10min"
	case MajorTenth:
		return "10maj"
	case AugmentedTenth:
		return "10aug"

	case DiminishedEleventh:
		return "11dim"
	case PerfectEleventh:
		return "11"
	case AugmentedEleventh:
		return "11aug"

	case DiminishedTwelfth:
		return "12dim"
	case PerfectTwelfth:
		return "12"
	case AugmentedTwelfth:
		return "12aug"

	case DiminishedThirteenth:
		return "13dim"
	case MinorThirteenth:
		return "13min"
	case MajorThirteenth:
		return "13maj"
	case AugmentedThirteenth:
		return "13aug"

	case DiminishedFourteenth:
		return "14dim"
	case MinorFourteenth:
		return "14min"
	case MajorFourteenth:
		return "14maj"
	case AugmentedFourteenth:
		return "14aug"

	case DiminishedFifteenth:
		return "15dim"
	case PerfectFifteenth:
		return "15"
	case AugmentedFifteenth:
		return "15aug"

	default:
		fmt.Println(interval)
		panic("unknown interval")
	}
}

func (interval Interval) Position() uint {
	return interval.pos
}

func (interval Interval) Semitone() uint {
	return interval.semitone
}

func (interval Interval) Relative() Interval {
	return Interval{pos: 7 - interval.pos, semitone: 12 - interval.semitone}
}

func FromName(intervalName string) (*Interval, error) {
	switch intervalName {
	case "1":
		return &PerfectUnison, nil
	case "1aug":
		return &AugmentedUnison, nil

	case "2dim":
		return &DiminishedSecond, nil
	case "2min":
		return &MinorSecond, nil
	case "2maj":
		return &MajorSecond, nil
	case "2aug":
		return &AugmentedSecond, nil

	case "3dim":
		return &DiminishedThird, nil
	case "3min":
		return &MinorThird, nil
	case "3maj":
		return &MajorThird, nil
	case "3aug":
		return &AugmentedThird, nil

	case "4dim":
		return &DiminishedFourth, nil
	case "4":
		return &PerfectFourth, nil
	case "4aug":
		return &AugmentedFourth, nil

	case "5dim":
		return &DiminishedFifth, nil
	case "5":
		return &PerfectFifth, nil
	case "5aug":
		return &AugmentedFifth, nil

	case "6dim":
		return &DiminishedSixth, nil
	case "6min":
		return &MinorSixth, nil
	case "6maj":
		return &MajorSixth, nil
	case "6aug":
		return &AugmentedSixth, nil

	case "7dim":
		return &DiminishedSeventh, nil
	case "7min":
		return &MinorSeventh, nil
	case "7maj":
		return &MajorSeventh, nil
	case "7aug":
		return &AugmentedSeventh, nil

	case "8dim":
		return &DiminishedOctave, nil
	case "8":
		return &Octave, nil
	case "8aug":
		return &AugmentedOctave, nil

	case "9dim":
		return &DiminishedNinth, nil
	case "9min":
		return &MinorNinth, nil
	case "9maj":
		return &MajorNinth, nil
	case "9aug":
		return &AugmentedNinth, nil

	case "10min":
		return &MinorTenth, nil
	case "10maj":
		return &MajorTenth, nil

	case "11dim":
		return &DiminishedEleventh, nil
	case "11":
		return &PerfectEleventh, nil
	case "11aug":
		return &AugmentedEleventh, nil

	case "12dim":
		return &DiminishedTwelfth, nil
	case "12":
		return &PerfectTwelfth, nil
	case "12aug":
		return &AugmentedTwelfth, nil

	case "13dim":
		return &DiminishedThirteenth, nil
	case "13min":
		return &MinorThirteenth, nil
	case "13maj":
		return &MajorThirteenth, nil
	case "13aug":
		return &AugmentedThirteenth, nil

	case "14dim":
		return &DiminishedFourteenth, nil
	case "14min":
		return &MinorFourteenth, nil
	case "14maj":
		return &MajorFourteenth, nil
	case "14aug":
		return &AugmentedFourteenth, nil

	case "15dim":
		return &DiminishedFifteenth, nil
	case "15":
		return &PerfectFifteenth, nil
	case "15aug":
		return &AugmentedFifteenth, nil

	default:
		return nil, fmt.Errorf("unknown interval name: %s", intervalName)
	}
}
