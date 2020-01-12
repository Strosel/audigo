package play

import (
	"math"
	"time"

	"github.com/strosel/audigo/midi"
)

type Rest Duration

//Duration calculates and returns the duration of the note based on
//how long the measure takes, aka the length of a whole note
func (r Rest) Duration(measure time.Duration) time.Duration {
	if Duration(r) == None {
		return 0
	}
	//the fraction of the measure the note takes
	fraq := 1. / math.Pow(2., float64(r))

	return time.Duration(float64(measure) * fraq)
}

//TickDuration calculates and returns the duration of the note in ticks
//based on how many ticks a quarter note is
func (r Rest) TickDuration(quarter uint16) uint16 {
	if Duration(r) == None {
		return 0
	}
	//the fraction of the measure the note takes
	fraq := 1. / math.Pow(2., float64(r))

	return uint16(float64(4*quarter) * fraq)
}

//ToMIDI converts the note to a midi NoteOn and NoteOff event
//with the given ticks for a quarter note, channel and velocity
func (r Rest) ToMIDI(ticks uint16, ch, vel uint8) []midi.Event {
	out := []midi.Event{}
	//ignore vel for an always silent rest
	vel = 0
	e := &midi.VoiceEvent{
		Channel: ch,
	}
	e.NoteOn(0x00, vel)
	out = append(out, e)

	e = &midi.VoiceEvent{
		Channel:  ch,
		Duration: midi.VLQ(r.TickDuration(ticks)),
	}
	e.NoteOff(0x00, vel)
	out = append(out, e)
	return out
}
