package audigo

import (
	"io/ioutil"
)

//WAV Describes a .wav file
type WAV struct {
	Header WaveHeader
	Format FormatChunk
	Data   DataChunk
}

//NewWAV Create a new Generic WAV
func NewWAV() WAV {
	return WAV{
		Header: NewWaveHeader(),
		Format: NewFormatChunk(),
		Data:   NewDataChunk(),
	}
}

//AddTrack Adds a new audiotrack to the WAV
func (w *WAV) AddTrack() {
	w.Data.firstTrack = false
	w.Data.index = 0
}

//Bytes Returns the binary representation of the WAV
func (w WAV) Bytes() []byte {
	w.Header.FileLength = uint32(4 + len(w.Format.Bytes()) + len(w.Data.Bytes()))

	tmpBytes := []byte{}
	tmpBytes = append(tmpBytes, w.Header.Bytes()...)
	tmpBytes = append(tmpBytes, w.Format.Bytes()...)
	tmpBytes = append(tmpBytes, w.Data.Bytes()...)

	return tmpBytes
}

//Save saves the WAV as a new .wav file
func (w WAV) Save(filename string) error {
	if len(filename) < 4 || filename[len(filename)-4:] != ".wav" {
		filename += ".wav"
	}

	tmpBytes := w.Bytes()

	return ioutil.WriteFile(filename, tmpBytes, 0666)
}
