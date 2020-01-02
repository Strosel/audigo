package midi

import (
	"bytes"
	"encoding/binary"
)

type SysexEvent struct {
	Type     uint8
	Data     []uint8
	Duration VLQ
}

func (se SysexEvent) Delta() VLQ {
	return se.Duration
}

func (se SysexEvent) Bytes() []byte {
	out := bytes.NewBuffer([]byte{})

	binary.Write(out, binary.BigEndian, se.Type)
	binary.Write(out, binary.BigEndian, se.Length().Bytes())
	for _, d := range se.Data {
		binary.Write(out, binary.BigEndian, d)
	}

	return out.Bytes()
}

func (se SysexEvent) Length() VLQ {
	return VLQ(len(se.Data))
}
