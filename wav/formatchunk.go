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

//ReadFormatChunk Reads the FormatChunk from raw bytes. See Open()
func ReadFormatChunk(file []byte) (FormatChunk, error) {
	fc := FormatChunk{}

	var err error
	locs := []int{12, 16, 20, 22, 24, 28, 32, 34, 36, 40}
	for i, v := range locs[:9] {
		buff := bytes.NewBuffer(file[v:locs[i+1]])

		switch i {
		case 0:
			fc.ChunkID = string(file[v:locs[i+1]])
		case 1:
			err = binary.Read(buff, binary.LittleEndian, &fc.ChunkSize)
		case 2:
			err = binary.Read(buff, binary.LittleEndian, &fc.FormatTag)
		case 3:
			err = binary.Read(buff, binary.LittleEndian, &fc.Channels)
		case 4:
			err = binary.Read(buff, binary.LittleEndian, &fc.Frequency)
		case 5:
			err = binary.Read(buff, binary.LittleEndian, &fc.AverageBytesPerSec)
		case 6:
			err = binary.Read(buff, binary.LittleEndian, &fc.BlockAlign)
		case 7:
			err = binary.Read(buff, binary.LittleEndian, &fc.BitsPerSample)
		}

		if err != nil {
			return FormatChunk{}, err
		}
	}

	return fc, nil
}

func (fc *FormatChunk) recalcBlockSizes() {
	fc.BlockAlign = fc.Channels * (fc.BitsPerSample / 8)
	fc.AverageBytesPerSec = fc.Frequency * uint32(fc.BlockAlign)
}

//Bytes Returns the binary representation of the FormatChunk
func (fc *FormatChunk) Bytes() []byte {
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
