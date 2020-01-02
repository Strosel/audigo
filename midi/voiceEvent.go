package midi

import (
	"bytes"
	"encoding/binary"
)

type VoiceEvent struct {
	Status   uint8
	Channel  uint8
	Data     []uint8
	Duration VLQ
}

func (ve VoiceEvent) Delta() VLQ {
	return ve.Duration
}

func (ve VoiceEvent) Bytes() []byte {
	out := bytes.NewBuffer([]byte{})

	ve.Channel %= 16 //only 16 available channels
	if ve.Status < 16 {
		ve.Status *= 16 //move digit "up one step" if neccecary
	}

	binary.Write(out, binary.BigEndian, ve.Status+ve.Channel)
	for _, d := range ve.Data {
		binary.Write(out, binary.BigEndian, d)
	}

	return out.Bytes()
}

func (ve *VoiceEvent) NoteOff(key, velocity uint8) {
	ve.Status = 0x80
	ve.Data = []byte{key, velocity}
}

func (ve *VoiceEvent) NoteOn(key, velocity uint8) {
	ve.Status = 0x90
	ve.Data = []byte{key, velocity}
}

func (ve *VoiceEvent) Aftertouch(key, pressure uint8) {
	ve.Status = 0xA0
	ve.Data = []byte{key, pressure}
}

func (ve *VoiceEvent) ControllerChange(controller, value uint8) {
	ve.Status = 0xB0
	ve.Data = []byte{controller, value}
}

func (ve *VoiceEvent) ProgramChange(program uint8) {
	ve.Status = 0xC0
	ve.Data = []byte{program}
}

func (ve *VoiceEvent) AftertouchAll(pressure uint8) {
	ve.Status = 0xD0
	ve.Data = []byte{pressure}
}

func (ve *VoiceEvent) PitchBend(lsb, msb uint8) {
	ve.Status = 0xE0
	ve.Data = []byte{lsb, msb}
}
