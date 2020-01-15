package play

import (
	"github.com/strosel/audigo/midi"
)

//Song is a song
type Song struct {
	Meter string
	//Tempo is the tempo of the song in quarter notes/min
	//as is usually anotated as <quarter note> = <number>
	Tempo uint32
	//Key is the key in number of flats or sharps, -7<k<7
	//flats are notated with negative numbers and sharps with positive
	Key         int8
	Maj         bool
	Instruments map[string][]Playable
}

//ToMIDI converts the song into a midi file
func (s Song) ToMIDI() midi.MIDI {
	var ticks uint16 = 64 //should be inputed or calculated
	m := midi.MIDI{
		Tracks: []midi.Track{},
	}

	var ch uint8 = 0
	for name, stave := range s.Instruments {
		t := midi.Track{
			midi.MetaChannelPrefix(ch),
			midi.MetaInstrument(name),
		}
		t = append(t, UpdateTempo(s.Tempo).ToMIDI(0, 0, 0)...)
		t = append(t, UpdateMeter(s.Meter).ToMIDI(0, 0, 0)...)
		t = append(t, UpdateKey(s.Key, s.Maj).ToMIDI(0, 0, 0)...)

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
