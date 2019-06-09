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

//ReadWaveHeader Reads the header from raw bytes. See Open()
func ReadWaveHeader(file []byte) (WaveHeader, error) {
	wh := WaveHeader{
		FileType:  string(file[:4]),
		MediaType: string(file[8:12]),
	}

	// header buffer
	hb := bytes.NewBuffer(file[4:8])
	err := binary.Read(hb, binary.LittleEndian, &wh.FileLength)
	if err != nil {
		return WaveHeader{}, err
	}

	return wh, nil
}

//Bytes Returns the binary representation of the WaveHeader
func (wh WaveHeader) Bytes() []byte {
	out := bytes.NewBuffer([]byte{})

	out.WriteString(wh.FileType)
	binary.Write(out, binary.LittleEndian, wh.FileLength)
	out.WriteString(wh.MediaType)

	return out.Bytes()
}
