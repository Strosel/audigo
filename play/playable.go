package play

import (
	"time"

	"github.com/strosel/audigo/midi"
)

//Playable defines an object that can be played by generating midi events
type Playable interface {
	Duration(measure time.Duration) time.Duration
	TickDuration(quarter uint16) uint16
	ToMIDI(ticks uint16, ch, vel uint8) []midi.Event
}
