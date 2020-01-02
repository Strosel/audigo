package midi

import (
	"bytes"
	"encoding/binary"
)

//ModeEvent a midi mode event
type ModeEvent struct {
	Channel  uint8
	Data     []uint8
	Duration VLQ
}

//Delta return the delta time associated with the event
func (me ModeEvent) Delta() VLQ {
	return me.Duration
}

//Bytes the event as a byte array
func (me ModeEvent) Bytes() []byte {
	out := bytes.NewBuffer([]byte{})

	me.Channel %= 16 //only 16 available channels

	binary.Write(out, binary.BigEndian, 0xB0+me.Channel)

	binary.Write(out, binary.BigEndian, me.Data[0])
	if len(me.Data) == 1 {
		binary.Write(out, binary.BigEndian, 0x00)
	} else {
		binary.Write(out, binary.BigEndian, me.Data[1])
	}

	return out.Bytes()
}

//AllSoundOff transforms the event into a AllSoundOff event
func (me *ModeEvent) AllSoundOff() {
	me.Data = []byte{0x78}
}

//ResetAllControllers transforms the event into a ResetAllControllers event
func (me *ModeEvent) ResetAllControllers() {
	me.Data = []byte{0x79}
}

//LocalControll transforms the event into a LocalControll event with the given parameters
func (me *ModeEvent) LocalControll(off bool) {
	if off {
		me.Data = []byte{0x7A}
	} else {
		me.Data = []byte{0x7A, 0x7F}
	}
}

//AllNotesOff transforms the event into a AllNotesOff event
func (me *ModeEvent) AllNotesOff() {
	me.Data = []byte{0x7B}
}

//OmniModeOff transforms the event into a OmniModeOff event
func (me *ModeEvent) OmniModeOff() {
	me.Data = []byte{0x7C}
}

//OmniModeOn transforms the event into a OmniModeOn event
func (me *ModeEvent) OmniModeOn() {
	me.Data = []byte{0x7D}
}

//MonoModeOn transforms the event into a MonoModeOn event with the given parameters
func (me *ModeEvent) MonoModeOn(channels uint8) {
	me.Data = []byte{0x7E, channels}
}

//PolyModeOn transforms the event into a PolyModeOn event
func (me *ModeEvent) PolyModeOn() {
	me.Data = []byte{0x7F}
}
