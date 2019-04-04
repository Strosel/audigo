package audigo

import (
	"bytes"
	"encoding/binary"
)

const (
	//FCHUNKID Is the default FormatChunk ID
	FCHUNKID = "fmt "
)

//FormatChunk Descibes a .wav FormatChunk
type FormatChunk struct {
	ChunkID            string
	ChunkSize          uint32
	FormatTag          uint16
	AverageBytesPerSec uint32
	BlockAlign         uint16
	Channels           uint16
	Frequency          uint32
	BitsPerSample      uint16
}

//NewFormatChunk Creates a new Generic FormatChunk
func NewFormatChunk() FormatChunk {
	fc := FormatChunk{
		ChunkID:       FCHUNKID,
		ChunkSize:     16,
		FormatTag:     1,
		Channels:      2,
		Frequency:     44100,
		BitsPerSample: 16,
	}
	fc.recalcBlockSizes()
	return fc
}

func (fc *FormatChunk) recalcBlockSizes() {
	fc.BlockAlign = fc.Channels * (fc.BitsPerSample / 8)
	fc.AverageBytesPerSec = fc.Frequency * uint32(fc.BlockAlign)
}

//Bytes Returns the binary representation of the FormatChunk
func (fc FormatChunk) Bytes() []byte {
	fc.recalcBlockSizes()
	out := bytes.NewBuffer([]byte{})

	out.WriteString(fc.ChunkID)
	binary.Write(out, binary.LittleEndian, fc.ChunkSize)
	binary.Write(out, binary.LittleEndian, fc.FormatTag)
	binary.Write(out, binary.LittleEndian, fc.Channels)
	binary.Write(out, binary.LittleEndian, fc.Frequency)
	binary.Write(out, binary.LittleEndian, fc.AverageBytesPerSec)
	binary.Write(out, binary.LittleEndian, fc.BlockAlign)
	binary.Write(out, binary.LittleEndian, fc.BitsPerSample)

	return out.Bytes()
}
