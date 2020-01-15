package play

import (
	"time"

	"github.com/strosel/audigo/midi"
)

type updateMeter string

func (um *updateMeter) Duration(time.Duration) time.Duration { return 0 }
func (um *updateMeter) TickDuration(uint16) uint16           { return 0 }

func (um *updateMeter) ToMIDI(uint16, uint8, uint8) []midi.Event {
	//todo use midi.MetaTimeSignature
	return []midi.Event{}
}

//UpdateMeter updates the meter or time signature of the song
func UpdateMeter(m string) Playable {
	um := updateMeter(m)
	return &um
}
