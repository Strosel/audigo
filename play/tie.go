package play

import (
	"time"

	"github.com/strosel/audigo/midi"
)

//Tie is a set of notes with a tie
type Tie []Playable

//NewTie creates a tie from a set sof notes
func NewTie(plays ...Playable) Tie {
	return Tie(plays)
}

//Duration calculates and returns the duration of the tie based on
//how long the measure takes, aka the length of a whole note
func (t Tie) Duration(measure time.Duration) time.Duration {
	out := time.Duration(0)
	for _, p := range t {
		out += p.Duration(measure)
	}
	return out
}

//TickDuration calculates and returns the duration of the tie in ticks
//based on how many ticks a quarter note is
func (t Tie) TickDuration(quarter uint16) uint16 {
	out := uint16(0)
	for _, p := range t {
		out += p.TickDuration(quarter)
	}
	return out
}

//RestDuration calculates and returns the duration of the rest preceding the
//tie based on how long the measure takes, aka the length of a whole note
func (t Tie) RestDuration(measure time.Duration) time.Duration {
	return t[0].RestDuration(measure)
}

//RestTickDuration calculates and returns the duration of the rest preceding the
//tie in ticks based on how many ticks a quarter note is
func (t Tie) RestTickDuration(quarter uint16) uint16 {
	return t[0].RestTickDuration(quarter)
}

//ToMIDI converst the tied set into an array of midi events
func (t Tie) ToMIDI(ticks uint16, ch, vel uint8) []midi.Event {
	out := t[0].ToMIDI(ticks, ch, vel)

	out[len(out)/2].SetDelta(midi.VLQ(t.TickDuration(ticks)))
	return out
}
