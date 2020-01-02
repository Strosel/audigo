package midi

import (
	"bytes"
	"encoding/binary"
)

type MetaEvent struct {
	Data     []uint8
	Duration VLQ
}

func (me MetaEvent) Delta() VLQ {
	return me.Duration
}

func (me MetaEvent) Bytes() []byte {
	out := bytes.NewBuffer([]byte{})

	binary.Write(out, binary.BigEndian, uint8(0xFF))
	binary.Write(out, binary.BigEndian, me.Length().Bytes())
	for _, d := range me.Data {
		binary.Write(out, binary.BigEndian, d)
	}

	return out.Bytes()
}

func (me MetaEvent) Length() VLQ {
	return VLQ(len(me.Data))
}
