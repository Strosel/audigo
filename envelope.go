package audigo

import (
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

type nilEnvelope time.Duration

//NewNilEnvelope Creates a nil envelope that doesnt modify the given wave
func NewNilEnvelope(duration time.Duration) Envelope {
	return nilEnvelope(duration)
}

func (ne nilEnvelope) Duration() time.Duration {
	return time.Duration(ne)
}

func (ne nilEnvelope) Amplitude(x float64) float64 {
	return 1.0
}
