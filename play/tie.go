package play

import "github.com/strosel/audigo/midi"

//Tie is a pair of notes with a tie
type Tie struct {
	Note   Note
	Values [2]Duration
	Dots   [2]int
}

//ToMIDI converst the tied pair into an array of midi events
func (t Tie) ToMIDI(ticks uint16, ch, vel uint8) []midi.Event {
	out := t.Note.ToMIDI(ticks, ch, vel)

	e := midi.VoiceEvent{
		Channel: ch,
	}
	var d uint16 = 0
	for i := range t.Values {
		t.Note.Value = t.Values[i]
		t.Note.Dots = t.Dots[i]
		d += t.Note.TickDuration(ticks)
	}
	e.Duration = midi.VLQ(d)
	e.NoteOff(0x3C+uint8(t.Note.Dist()), vel)
	out[1] = e
	return out
}
