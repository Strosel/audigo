package midi

import (
	"bytes"
	"encoding/binary"
)

//TrackType is the chink type of a track chunk
const TrackType = "MTrk"

type Track []Event

func (t Track) Size() uint32 {
	l := 0
	for _, e := range t {
		l += len(e.Delta().Bytes())
		l += len(e.Bytes())
	}
	return uint32(l)
}

//Bytes returns the bytes to be written to the file
func (t Track) Bytes() []byte {
	out := bytes.NewBuffer([]byte{})

	out.WriteString(TrackType)
	binary.Write(out, binary.BigEndian, t.Size())

	for _, e := range t {
		binary.Write(out, binary.BigEndian, e.Delta().Bytes())
		binary.Write(out, binary.BigEndian, e.Bytes())
	}

	return out.Bytes()
}
