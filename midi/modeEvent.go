package midi

import (
	"bytes"
	"encoding/binary"
)

type ModeEvent struct {
	Channel  uint8
	Data     []uint8
	Duration VLQ
}

func (me ModeEvent) Delta() VLQ {
	return me.Duration
}

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

func (me *ModeEvent) AllSoundOff() {
	me.Data = []byte{0x78}
}

func (me *ModeEvent) ResetAllControllers() {
	me.Data = []byte{0x79}
}

func (me *ModeEvent) LocalControll(off bool) {
	if off {
		me.Data = []byte{0x7A}
	} else {
		me.Data = []byte{0x7A, 0x7F}
	}
}

func (me *ModeEvent) AllNotesOff() {
	me.Data = []byte{0x7B}
}

func (me *ModeEvent) OmniModeOff() {
	me.Data = []byte{0x7C}
}

func (me *ModeEvent) OmniModeOn() {
	me.Data = []byte{0x7D}
}

func (me *ModeEvent) MonoModeOn(channels uint8) {
	me.Data = []byte{0x7E, channels}
}

func (me *ModeEvent) PolyModeOn() {
	me.Data = []byte{0x7F}
}
