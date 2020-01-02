package midi

import (
	"bytes"
	"encoding/binary"
)

//MetaEvent a midi meta event,
type MetaEvent struct {
	Type uint8
	Data []uint8
}

//SetDelta is ignored for meta events
func (me *MetaEvent) SetDelta(v VLQ) {}

//Delta return the delta time associated with the event
func (me MetaEvent) Delta() VLQ {
	return 0
}

//Bytes the event as a byte array
func (me MetaEvent) Bytes() []byte {
	out := bytes.NewBuffer([]byte{})

	binary.Write(out, binary.BigEndian, uint8(0xFF))
	binary.Write(out, binary.BigEndian, me.Type)
	binary.Write(out, binary.BigEndian, me.length().Bytes())
	for _, d := range me.Data {
		binary.Write(out, binary.BigEndian, d)
	}

	return out.Bytes()
}

func (me MetaEvent) length() VLQ {
	return VLQ(len(me.Data))
}

//MetaSequenceNumber an id for the sewuance or track.
//If omitted, the sequences are numbered sequentially in the order the tracks appear.
//For Format 2 MIDI files, this is used to identify each track.
//For Format 1 files, this event should occur on the first track only.
func MetaSequenceNumber(n uint16) MetaEvent {
	d := bytes.NewBuffer([]byte{})
	binary.Write(d, binary.BigEndian, n)
	return MetaEvent{
		Type: 0x00,
		Data: d.Bytes(),
	}
}

//MetaText arbitrary text
func MetaText(txt string) MetaEvent {
	return MetaEvent{
		Type: 0x01,
		Data: []uint8(txt),
	}
}

//MetaCopyright mark copyright
func MetaCopyright(c string) MetaEvent {
	return MetaEvent{
		Type: 0x02,
		Data: []uint8(c),
	}
}

//MetaName name of the sequence or track
func MetaName(name string) MetaEvent {
	return MetaEvent{
		Type: 0x03,
		Data: []uint8(name),
	}
}

//MetaInstrument instrument name
func MetaInstrument(name string) MetaEvent {
	return MetaEvent{
		Type: 0x04,
		Data: []uint8(name),
	}
}

//MetaLyric lyrics, usually timed syllables
func MetaLyric(lyric string) MetaEvent {
	return MetaEvent{
		Type: 0x05,
		Data: []uint8(lyric),
	}
}

//MetaMarker marks a significant point in the sequence (eg "Verse 1")
func MetaMarker(marker string) MetaEvent {
	return MetaEvent{
		Type: 0x06,
		Data: []uint8(marker),
	}
}

//MetaCue cues for events happening on-stage, such as "curtain rises"
func MetaCue(cue string) MetaEvent {
	return MetaEvent{
		Type: 0x07,
		Data: []uint8(cue),
	}
}

//MetaChannelPrefix Associate all following meta-events and sysex-events
//with the specified MIDI channel, until the next <midi_event>
func MetaChannelPrefix(channel uint8) MetaEvent {
	return MetaEvent{
		Type: 0x20,
		Data: []uint8{channel},
	}
}

//MetaEndTrack end the track chunk
func MetaEndTrack() MetaEvent {
	return MetaEvent{
		Type: 0x2F,
	}
}

//MetaTempo sets the tempo in microseconds per quarter note
func MetaTempo(tempo uint32) MetaEvent {
	d := bytes.NewBuffer([]byte{})
	binary.Write(d, binary.BigEndian, tempo)
	return MetaEvent{
		Type: 0x51,
		Data: d.Bytes()[1:],
	}
}

//MetaTimeSignature Time signature of the form: n/2^d
//c is the number of MIDI Clocks per metronome tick. Normally, there are 24 MIDI Clocks per quarter note.
//b defines this in terms of the number of 1/32 notes which make up the usual 24 MIDI Clocks (the 'standard' quarter note).
func MetaTimeSignature(n, d, c, b uint8) MetaEvent {
	return MetaEvent{
		Type: 0x58,
		Data: []uint8{n, d, c, b},
	}
}

//MetaKey the key of the song
//sf 0 represents a key of C, negative numbers represent 'flats', while positive numbers represent 'sharps'
func MetaKey(sf uint8, minor bool) MetaEvent {
	var m uint8 = 0
	if minor {
		m = 1
	}
	return MetaEvent{
		Type: 0x59,
		Data: []uint8{sf, m},
	}
}

//MetaSpecific metadata specific to certain sequencers
func MetaSpecific(id uint8, data []uint8) MetaEvent {
	data = append([]uint8{id}, data...)
	return MetaEvent{
		Type: 0x7F,
		Data: data,
	}
}
