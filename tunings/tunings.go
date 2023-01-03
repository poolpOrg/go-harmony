package tunings

import (
	"math"
)

type tuningSystem struct {
	name      string
	frequency func(float64, uint8) float64
}
type TuningSystem tuningSystem

var (
	EqualTemperament TuningSystem = TuningSystem{"Equal Temperament", func(reference float64, position uint8) float64 {
		delta := 53 - int8(position) //A4
		if delta == 0 {
			return reference
		} else if delta < 0 {
			return reference * math.Pow(2, float64(-delta)/12)
		} else {
			return reference / math.Pow(2, float64(delta)/12)
		}
	}}
)

var tuningSystems = []TuningSystem{
	EqualTemperament,
}

func Systems() []TuningSystem {
	return tuningSystems
}

type tuning struct {
	name      string
	reference float64
}
type Tuning tuning

func (tuning *Tuning) Name() string {
	return tuning.name
}

func (tuning *Tuning) Reference() float64 {
	return tuning.reference
}

var (
	A432 Tuning = Tuning{name: "A432", reference: 432.}
	A434 Tuning = Tuning{name: "A434", reference: 434.}
	A436 Tuning = Tuning{name: "A436", reference: 436.}
	A438 Tuning = Tuning{name: "A438", reference: 438.}
	A440 Tuning = Tuning{name: "A440", reference: 440.}
	A442 Tuning = Tuning{name: "A442", reference: 442.}
	A444 Tuning = Tuning{name: "A444", reference: 444.}
	A446 Tuning = Tuning{name: "A446", reference: 446.}
)

func Tunings() []Tuning {
	return []Tuning{
		A432, A434, A436, A438, A440, A442, A444, A446,
	}
}

type Tuner struct {
	system TuningSystem
	tuning Tuning
}

func NewTuner(system TuningSystem, tuning Tuning) *Tuner {
	return &Tuner{system: system, tuning: tuning}
}

func (tuner *Tuner) Frequency(position uint8) float64 {
	return tuner.system.frequency(tuner.tuning.reference, position)
}
