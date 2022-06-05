package intervals

type interval struct {
	pos      uint
	semitone uint
}

type Interval interval

var (
	PerfectUnison     Interval = Interval{pos: 0, semitone: 0}
	AugmentedUnison   Interval = Interval{pos: 0, semitone: 1}
	MinorSecond       Interval = Interval{pos: 1, semitone: 1}
	MajorSecond       Interval = Interval{pos: 1, semitone: 2}
	AugmentedSecond   Interval = Interval{pos: 1, semitone: 3}
	MinorThird        Interval = Interval{pos: 2, semitone: 3}
	MajorThird        Interval = Interval{pos: 2, semitone: 4}
	PerfectFourth     Interval = Interval{pos: 3, semitone: 5}
	AugmentedFourth   Interval = Interval{pos: 3, semitone: 6}
	DiminishedFifth   Interval = Interval{pos: 4, semitone: 6}
	PerfectFifth      Interval = Interval{pos: 4, semitone: 7}
	AugmentedFifth    Interval = Interval{pos: 4, semitone: 8}
	MinorSixth        Interval = Interval{pos: 5, semitone: 8}
	MajorSixth        Interval = Interval{pos: 5, semitone: 9}
	DiminishedSeventh Interval = Interval{pos: 6, semitone: 9}
	MinorSeventh      Interval = Interval{pos: 6, semitone: 10}
	MajorSeventh      Interval = Interval{pos: 6, semitone: 11}
	Octave            Interval = Interval{pos: 7, semitone: 12}
	AugmentedOctave   Interval = Interval{pos: 7, semitone: 13}
	MinorNinth        Interval = Interval{pos: 8, semitone: 13}
	MajorNinth        Interval = Interval{pos: 8, semitone: 14}
	AugmentedNinth    Interval = Interval{pos: 8, semitone: 15}
	MinorTenth        Interval = Interval{pos: 9, semitone: 15}
	MajorTenth        Interval = Interval{pos: 9, semitone: 16}
	PerfectEleventh   Interval = Interval{pos: 10, semitone: 17}
	AugmentedEleventh Interval = Interval{pos: 10, semitone: 18}
	DiminishedTwelfth Interval = Interval{pos: 11, semitone: 18}
	PerfectTwelfth    Interval = Interval{pos: 11, semitone: 19}
	AugmentedTwelfth  Interval = Interval{pos: 11, semitone: 20}
	MinorThirteenth   Interval = Interval{pos: 12, semitone: 20}
	MajorThirteenth   Interval = Interval{pos: 12, semitone: 21}
)

func New(pos uint, semitone uint) Interval {
	return Interval{pos: pos, semitone: semitone}

}

func (interval Interval) Name() string {
	switch interval {
	case PerfectUnison:
		return "1"
	case AugmentedUnison:
		return "1aug"
	case MinorSecond:
		return "2min"
	case MajorSecond:
		return "2maj"
	case AugmentedSecond:
		return "2aug"
	case MinorThird:
		return "3min"
	case MajorThird:
		return "3maj"
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
	case MinorSixth:
		return "6min"
	case MajorSixth:
		return "6maj"
	case DiminishedSeventh:
		return "7dim"
	case MinorSeventh:
		return "7min"
	case MajorSeventh:
		return "7maj"
	case Octave:
		return "8"
	case AugmentedOctave:
		return "8aug"
	case MinorNinth:
		return "9min"
	case MajorNinth:
		return "9maj"
	case AugmentedNinth:
		return "9aug"
	case MinorTenth:
		return "10min"
	case MajorTenth:
		return "10maj"
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
	case MinorThirteenth:
		return "13min"
	case MajorThirteenth:
		return "13maj"
	default:
		panic("unknown interval")
	}
}

func (interval Interval) Position() uint {
	return interval.pos
}

func (interval Interval) Semitone() uint {
	return interval.semitone
}
