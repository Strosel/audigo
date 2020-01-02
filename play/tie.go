package play

import (
	"github.com/strosel/audigo/midi"
)

//Tie is a set of notes with a tie
type Tie struct {
	Note   Note
	Values []Duration
	Dots   []int
}

//NewTie creates a tie from a set sof notes
func NewTie(notes ...Note) Tie {
	v := []Duration{}
	d := []int{}
	for _, n := range notes {
		v = append(v, n.Value)
		d = append(d, n.Dots)
	}
	return Tie{
		Note:   notes[0],
		Values: v,
		Dots:   d,
	}
}

//ToMIDI converst the tied set into an array of midi events
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
