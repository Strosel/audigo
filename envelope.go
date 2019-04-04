package audigo

import (
	"math"
	"time"
)

const (
	//SAMPLERATE Is a generic sampleRate
	SAMPLERATE = 44100
)

//Envelope Describes a generic sound envelope
type Envelope interface {
	Duration() time.Duration
	Amplitude(float64) float64
	GenerateSine(uint32, ...float64) Wave
}

type linearASDR struct {
	attackTime  time.Duration
	decayTime   time.Duration
	sustainTime time.Duration
	releaseTime time.Duration

	maxAmplitude float64
	sustainLevel float64
}

//NewLinearASDREnvelope Creates an ADSR envelope with "linear, harch edges"
func NewLinearASDREnvelope(attackTime, decayTime, sustainTime, releaseTime time.Duration, maxAmplitude, sustainLevel float64) Envelope {
	ble := linearASDR{
		attackTime:  attackTime,
		decayTime:   attackTime + decayTime,
		sustainTime: attackTime + decayTime + sustainTime,
		releaseTime: attackTime + decayTime + sustainTime + releaseTime,

		maxAmplitude: maxAmplitude,
		sustainLevel: sustainLevel,
	}

	return ble
}

func (ble linearASDR) Duration() time.Duration {
	return ble.releaseTime
}

func (ble linearASDR) Amplitude(x float64) float64 {
	k := 0.0
	m := 0.0

	if x < float64(ble.attackTime) { //attack stage
		k = ble.maxAmplitude / float64(ble.attackTime)
	} else if float64(ble.attackTime) <= x && x < float64(ble.decayTime) { //decay stage
		k = (ble.sustainLevel - ble.maxAmplitude) / float64(ble.decayTime-ble.attackTime)
		m = ble.maxAmplitude - k*float64(ble.attackTime)
	} else if float64(ble.decayTime) <= x && x < float64(ble.sustainTime) { //sustain stage
		return ble.sustainLevel
	} else if float64(ble.sustainTime) <= x && x < float64(ble.releaseTime) { //release stage
		k = -ble.sustainLevel / float64(ble.releaseTime-ble.sustainTime)
		m = -k * float64(ble.releaseTime)
	}

	return x*k + m
}

func (ble linearASDR) GenerateSine(sampleRate uint32, frequency ...float64) Wave {
	bufferSize := uint32(float64(sampleRate) * ble.Duration().Seconds())

	timePeriod := make([]float64, len(frequency))
	for i, freq := range frequency {
		timePeriod[i] = (math.Pi * 2 * freq) / float64(sampleRate)
	}

	sinesum := make([]float64, bufferSize)
	sinemax := 0.0
	for i := range sinesum {
		sum := 0.0
		for _, tp := range timePeriod {
			sum += math.Sin(tp * float64(i))
		}

		if sum > sinemax {
			sinemax = sum
		}

		sinesum[i] = sum
	}

	data := make(Wave, bufferSize)
	for i, sine := range sinesum {
		sec := (float64(i) / float64(bufferSize)) * float64(ble.Duration())
		amp := ble.Amplitude(sec)

		data[i] = int16(amp * sine * (1 / sinemax))
	}

	return data
}
