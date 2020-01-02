package midi

import (
	"bytes"
	"encoding/binary"
)

type MIDIFormat uint16

const (
	//HeaderType is the chink type of a header chunk
	HeaderType = "MThd"

	Single MIDIFormat = iota
	Simultaneous
	Independent
)

//Header is a midi header chunk
//for now only ticks per quarter note is supported for Division
type Header struct {
	Format   MIDIFormat
	Tracks   uint16
	Division uint16
}

//Bytes returns the bytes to be written to the file
func (h Header) Bytes() []byte {
	out := bytes.NewBuffer([]byte{})

	out.WriteString(HeaderType)
	binary.Write(out, binary.BigEndian, int32(6))
	binary.Write(out, binary.BigEndian, h.Format)
	binary.Write(out, binary.BigEndian, h.Tracks)
	binary.Write(out, binary.BigEndian, h.Division)

	return out.Bytes()
}
