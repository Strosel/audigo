package midi

import (
	"bytes"
	"encoding/binary"
)

//Format one of 3 midi formats
type Format uint16

const (
	//HeaderType is the chink type of a header chunk
	HeaderType = "MThd"

	//Single is midi format 0, a single track
	Single Format = iota
	//Simultaneous is midi format 1, multiple tracks to be played simultaneously
	Simultaneous
	//Independent is midi format 2, multiple tracks to be played independently
	Independent
)

//Header is a midi header chunk
//for now only ticks per quarter note is supported for Division
type Header struct {
	Format   Format
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
