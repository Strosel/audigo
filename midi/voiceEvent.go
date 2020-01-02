package midi

import (
	"bytes"
	"encoding/binary"
)

//VoiceEvent is a midi voice event
type VoiceEvent struct {
	Status   uint8
	Channel  uint8
	Data     []uint8
	Duration VLQ
}

//Delta return the delta time associated with the event
func (ve VoiceEvent) Delta() VLQ {
	return ve.Duration
}

//Bytes the event as a byte array
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

//NoteOff transform the event into a NoteOff with the given parameters
func (ve *VoiceEvent) NoteOff(key, velocity uint8) {
	ve.Status = 0x80
	ve.Data = []byte{key, velocity}
}

//NoteOn transform the event into a NoteOn with the given parameters
func (ve *VoiceEvent) NoteOn(key, velocity uint8) {
	ve.Status = 0x90
	ve.Data = []byte{key, velocity}
}

//Aftertouch transform the event into a Aftertouch with the given parameters
func (ve *VoiceEvent) Aftertouch(key, pressure uint8) {
	ve.Status = 0xA0
	ve.Data = []byte{key, pressure}
}

//ControllerChange transform the event into a ControllerChange with the given parameters
func (ve *VoiceEvent) ControllerChange(controller, value uint8) {
	ve.Status = 0xB0
	ve.Data = []byte{controller, value}
}

//ProgramChange transform the event into a ProgramChange with the given parameters
func (ve *VoiceEvent) ProgramChange(program uint8) {
	ve.Status = 0xC0
	ve.Data = []byte{program}
}

//AftertouchAll transform the event into a AftertouchAll with the given parameters
func (ve *VoiceEvent) AftertouchAll(pressure uint8) {
	ve.Status = 0xD0
	ve.Data = []byte{pressure}
}

//PitchBend transform the event into a PitchBend with the given parameters
func (ve *VoiceEvent) PitchBend(lsb, msb uint8) {
	ve.Status = 0xE0
	ve.Data = []byte{lsb, msb}
}
