package midi

import (
	"bytes"
	"encoding/binary"
)

//MetaEvent a midi meta event,
type MetaEvent struct {
	Data     []uint8
	Duration VLQ
}

//Delta return the delta time associated with the event
func (me MetaEvent) Delta() VLQ {
	return me.Duration
}

//Bytes the event as a byte array
func (me MetaEvent) Bytes() []byte {
	out := bytes.NewBuffer([]byte{})

	binary.Write(out, binary.BigEndian, uint8(0xFF))
	binary.Write(out, binary.BigEndian, me.length().Bytes())
	for _, d := range me.Data {
		binary.Write(out, binary.BigEndian, d)
	}

	return out.Bytes()
}

func (me MetaEvent) length() VLQ {
	return VLQ(len(me.Data))
}
