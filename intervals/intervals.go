package intervals

type interval struct {
	name     string
	pos      uint
	semitone uint
}

type Interval interval

var (
	PerfectUnison     Interval = Interval{name: "1", pos: 0, semitone: 0}
	AugmentedUnison   Interval = Interval{name: "1 aug", pos: 0, semitone: 1}
	MinorSecond       Interval = Interval{name: "2 min", pos: 1, semitone: 1}
	MajorSecond       Interval = Interval{name: "2 maj", pos: 1, semitone: 2}
	AugmentedSecond   Interval = Interval{name: "2 aug", pos: 1, semitone: 3}
	MinorThird        Interval = Interval{name: "3 min", pos: 2, semitone: 3}
	MajorThird        Interval = Interval{name: "3 maj", pos: 2, semitone: 4}
	PerfectFourth     Interval = Interval{name: "4", pos: 3, semitone: 5}
	AugmentedFourth   Interval = Interval{name: "4 aug", pos: 3, semitone: 6}
	DiminishedFifth   Interval = Interval{name: "5 dim", pos: 4, semitone: 6}
	PerfectFifth      Interval = Interval{name: "5", pos: 4, semitone: 7}
	AugmentedFifth    Interval = Interval{name: "5 aug", pos: 4, semitone: 8}
	MinorSixth        Interval = Interval{name: "6 min", pos: 5, semitone: 8}
	MajorSixth        Interval = Interval{name: "6 maj", pos: 5, semitone: 9}
	DiminishedSeventh Interval = Interval{name: "7 dim", pos: 6, semitone: 9}
	MinorSeventh      Interval = Interval{name: "7 min", pos: 6, semitone: 10}
	MajorSeventh      Interval = Interval{name: "7 maj", pos: 6, semitone: 11}
	Octave            Interval = Interval{name: "8", pos: 7, semitone: 12}
	MinorNinth        Interval = Interval{name: "9 min", pos: 8, semitone: 13}
	MajorNinth        Interval = Interval{name: "9 maj", pos: 8, semitone: 14}
	AugmentedNinth    Interval = Interval{name: "9 aug", pos: 8, semitone: 15}
)

func (interval *Interval) Name() string {
	return interval.name
}

func (interval *Interval) Position() uint {
	return interval.pos
}

func (interval *Interval) Semitone() uint {
	return interval.semitone
}
