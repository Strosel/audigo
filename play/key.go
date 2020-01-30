package play

import (
	"time"

	"github.com/strosel/audigo/midi"
)

type updateKey struct {
	key   int8
	major bool
}

func (uk updateKey) Duration(time.Duration) time.Duration { return 0 }
func (uk updateKey) TickDuration(uint16) uint16           { return 0 }

func (uk updateKey) ToMIDI(uint16, uint8, uint8) []midi.Event {
	me := midi.MetaKey(uk.key, !uk.major)
	return []midi.Event{me}
}

//UpdateKey updates the key of the song
func UpdateKey(k int8, maj bool) Playable {
	return updateKey{
		key:   k,
		major: maj,
	}
}
