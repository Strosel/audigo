package play

import "github.com/strosel/audigo/midi"

//Slur is a sequence of notes with a slur line
type Slur []Note

//ToMIDI converst the sluured sequence into an array of midi events
func (s Slur) ToMIDI(ticks uint16, ch, vel uint8) []midi.Event {
	out := []midi.Event{}

	for _, n := range s {
		out = append(out, n.ToMIDI(ticks, ch, vel)...)
	}

	return out
}
