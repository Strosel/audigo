package play

import (
	"time"

	"github.com/strosel/audigo/midi"
)

type updateTempo uint32

func (ut updateTempo) Duration(time.Duration) time.Duration { return 0 }
func (ut updateTempo) TickDuration(uint16) uint16           { return 0 }

func (ut updateTempo) ToMIDI(uint16, uint8, uint8) []midi.Event {
	me := midi.MetaTempo(uint32(ut))
	return []midi.Event{me}
}

//UpdateTempo updates the tempo of the song
func UpdateTempo(t uint32) Playable {
	return updateTempo(uint32(time.Minute.Microseconds()) / t)
}
