package main

import (
	"time"

	"github.com/strosel/audigo/midi"
)

//Song is a song
type Song struct {
	Meter       string
	Measure     time.Duration
	BPM         int
	Instruments map[string][]Playable
}

//ToMIDI converts the song into a midi file
func (s Song) ToMIDI() midi.MIDI {
	var ticks uint16 = 50 //should be inputed or calculated
	m := midi.MIDI{
		Tracks: []midi.Track{},
	}

	var ch uint8 = 0
	for name, stave := range s.Instruments {
		t := midi.Track{
			midi.MetaChannelPrefix(ch),
			midi.MetaInstrument(name),
		}

		for _, n := range stave {
			t = append(t, n.ToMIDI(ticks, ch, 40)...)
		}

		t = append(t, midi.MetaEndTrack())

		m.Tracks = append(m.Tracks, t)
		ch++
	}

	m.Header = midi.Header{
		Tracks:   uint16(len(m.Tracks)),
		Format:   midi.Single,
		Division: ticks,
	}
	return m
}
