package play

import (
	"time"

	"github.com/strosel/audigo/midi"
)

//Chord desctibes a chord or group of notes played simoultaniously
type Chord []Note

//Duration calculates and returns the duration of the tie based on
//how long the measure takes, aka the length of a whole note
func (c Chord) Duration(measure time.Duration) time.Duration {
	return c[0].Duration(measure)
}

//TickDuration calculates and returns the duration of the tie in ticks
//based on how many ticks a quarter note is
func (c Chord) TickDuration(quarter uint16) uint16 {
	return c[0].TickDuration(quarter)
}

//RestDuration calculates and returns the duration of the rest preceding the
//tie based on how long the measure takes, aka the length of a whole note
func (c Chord) RestDuration(measure time.Duration) time.Duration {
	return c[0].RestDuration(measure)
}

//RestTickDuration calculates and returns the duration of the rest preceding the
//tie in ticks based on how many ticks a quarter note is
func (c Chord) RestTickDuration(quarter uint16) uint16 {
	return c[0].RestTickDuration(quarter)
}

//ToMIDI converst the chord into an array of midi events
func (c Chord) ToMIDI(ticks uint16, ch, vel uint8) []midi.Event {
	out := []midi.Event{}
	e := &midi.VoiceEvent{
		Channel:  ch,
		Duration: midi.VLQ(c.RestTickDuration(ticks)),
	}
	//Generate the NoteOn events 0 time from eachother
	for i, n := range c {
		if i != 0 {
			e.Duration = 0
		}
		e.NoteOn(0x3C+uint8(n.Dist()), vel)
		out = append(out, e)
	}

	//set the duration of the chord aka the time between last NoteOn and first NoteOff
	e.Duration = midi.VLQ(c.TickDuration(ticks))
	//Generate the Noteff events 0 time from eachother
	for i, n := range c {
		if i != 0 {
			e.Duration = 0
		}
		e.NoteOff(0x3C+uint8(n.Dist()), vel)
		out = append(out, e)
	}

	return out
}
