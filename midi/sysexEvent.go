package midi

import (
	"bytes"
	"encoding/binary"
)

//SysexEvent a midi system exclusive event
type SysexEvent struct {
	Type     uint8
	Data     []uint8
	Duration VLQ
}

//SetDelta set the delta time associated with the event
func (se *SysexEvent) SetDelta(v VLQ) {
	se.Duration = v
}

//Delta return the delta time associated with the event
func (se *SysexEvent) Delta() VLQ {
	return se.Duration
}

//Bytes the event as a byte array
func (se *SysexEvent) Bytes() []byte {
	out := bytes.NewBuffer([]byte{})

	binary.Write(out, binary.BigEndian, se.Type)
	binary.Write(out, binary.BigEndian, se.length().Bytes())
	for _, d := range se.Data {
		binary.Write(out, binary.BigEndian, d)
	}

	return out.Bytes()
}

func (se SysexEvent) length() VLQ {
	return VLQ(len(se.Data))
}
