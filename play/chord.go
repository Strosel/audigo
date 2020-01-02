package main

import "github.com/strosel/audigo/midi"

//Chord desctibes a chord or group of notes played simoultaniously
type Chord []Note

//ToMIDI converst the chord into an array of midi events
func (c Chord) ToMIDI(ticks uint16, ch, vel uint8) []midi.Event {
	out := []midi.Event{}
	e := midi.VoiceEvent{
		Channel: ch,
	}
	//Generate the NoteOn events 0 time from eachother
	for _, n := range c {
		e.NoteOn(0x3C+uint8(n.Dist()), vel)
		out = append(out, e)
	}

	//set the duration of the chord aka the time between last NoteOn and first NoteOff
	e.Duration = midi.VLQ(c[0].TickDuration(ticks))
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
