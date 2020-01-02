package play

import (
	"github.com/strosel/audigo/midi"
)

//Playable defines an object that can be played by generating midi events
type Playable interface {
	ToMIDI(ticks uint16, ch, vel uint8) []midi.Event
}
