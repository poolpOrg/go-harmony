package sound

type sound struct {
	frequency float64
	amplitude float64
}
type Sound sound

func NewSound(frequency float64, amplitude float64) *Sound {
	return &Sound{frequency: frequency, amplitude: amplitude}
}

func (snd *Sound) Frequency() float64 {
	return snd.frequency
}

func (snd *Sound) Amplitude() float64 {
	return snd.amplitude
}
