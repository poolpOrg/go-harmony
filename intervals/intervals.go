package intervals

type interval struct {
	name     string
	degree   uint
	distance uint
}

type Interval interval

var (
	PerfectUnison     Interval = Interval{name: "1", degree: 0, distance: 0}
	AugmentedUnison   Interval = Interval{name: "1 aug", degree: 0, distance: 1}
	MinorSecond       Interval = Interval{name: "2 min", degree: 1, distance: 1}
	MajorSecond       Interval = Interval{name: "2 maj", degree: 1, distance: 2}
	AugmentedSecond   Interval = Interval{name: "2 aug", degree: 1, distance: 3}
	MinorThird        Interval = Interval{name: "3 min", degree: 2, distance: 3}
	MajorThird        Interval = Interval{name: "3 maj", degree: 2, distance: 4}
	PerfectFourth     Interval = Interval{name: "4", degree: 3, distance: 5}
	AugmentedFourth   Interval = Interval{name: "4 aug", degree: 3, distance: 6}
	DiminishedFifth   Interval = Interval{name: "5 dim", degree: 4, distance: 6}
	PerfectFifth      Interval = Interval{name: "5", degree: 4, distance: 7}
	AugmentedFifth    Interval = Interval{name: "5 aug", degree: 4, distance: 8}
	MinorSixth        Interval = Interval{name: "6 min", degree: 5, distance: 8}
	MajorSixth        Interval = Interval{name: "6 maj", degree: 5, distance: 9}
	DiminishedSeventh Interval = Interval{name: "7 dim", degree: 6, distance: 9}
	MinorSeventh      Interval = Interval{name: "7 min", degree: 6, distance: 10}
	MajorSeventh      Interval = Interval{name: "7 maj", degree: 6, distance: 11}
	Octave            Interval = Interval{name: "8", degree: 7, distance: 12}
)

func (interval *Interval) Name() string {
	return interval.name
}

func (interval *Interval) Degree() uint {
	return interval.degree
}

func (interval *Interval) Distance() uint {
	return interval.distance
}
