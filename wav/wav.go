package audigo

import (
	"fmt"
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

//Open Opens the given .wav file. Assumes formatting according to this spec: http://www.topherlee.com/software/pcm-tut-wavformat.html [2019-06-09]
func Open(filename string) (WAV, error) {
	if len(filename) < 4 || filename[len(filename)-4:] != ".wav" {
		filename += ".wav"
	}

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return WAV{}, fmt.Errorf("ReadFile: %v", err)
	}

	header, err := ReadWaveHeader(file)
	if err != nil {
		return WAV{}, fmt.Errorf("ReadWaveHeader: %v", err)
	}

	FormatChunk, err := ReadFormatChunk(file)
	if err != nil {
		return WAV{}, fmt.Errorf("ReadFormatChunk: %v", err)
	}

	DataChunk, err := ReadDataChunk(file)
	if err != nil {
		return WAV{}, fmt.Errorf("ReadDataChunk: %v", err)
	}

	return WAV{
		Header: header,
		Format: FormatChunk,
		Data:   DataChunk,
	}, nil
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
