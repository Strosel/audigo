package audigo

import (
	"bytes"
	"encoding/binary"
)

const (
	//DefaultFileType Is the default File type of a WAV
	DefaultFileType = "RIFF"

	//DefaultMediaType Is the default Media type of a WAV
	DefaultMediaType = "WAVE"
)

//WaveHeader describes the header of a .wav file
type WaveHeader struct {
	FileType   string
	FileLength uint32
	MediaType  string
}

//NewWaveHeader Creates a new Generic WaveHeader
func NewWaveHeader() WaveHeader {
	return WaveHeader{
		FileType:   DefaultFileType,
		FileLength: 4,
		MediaType:  DefaultMediaType,
	}
}

//Bytes Returns the binary representation of the WaveHeader
func (wh WaveHeader) Bytes() []byte {
	out := bytes.NewBuffer([]byte{})

	out.WriteString(wh.FileType)
	binary.Write(out, binary.LittleEndian, wh.FileLength)
	out.WriteString(wh.MediaType)

	return out.Bytes()
}
