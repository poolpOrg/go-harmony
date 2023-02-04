package durations

type duration struct {
	division float32
}
type Duration duration

var (
	DoubleWhole  duration = duration{division: 0.5}
	Whole        duration = duration{division: 1}
	Half         duration = duration{division: 2}
	Quarter      duration = duration{division: 4}
	Eighth       duration = duration{division: 8}
	Sixteenth    duration = duration{division: 16}
	ThirtySecond duration = duration{division: 32}
	SixtyFourth  duration = duration{division: 64}
)
