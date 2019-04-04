package audigo

import (
	"bytes"
	"encoding/binary"
)

const (
	//DCHUNKID Is the deafault DataChunk ID
	DCHUNKID = "data"
)

//Wave Describeas a soundwave as a set of points
type Wave []int16

//DataChunk Descibes a .wav DataChunk
type DataChunk struct {
	ChunkID   string
	ChunkSize uint32
	WaveData  Wave
}

//NewDataChunk Creates a new Generic DataChunk
func NewDataChunk() DataChunk {
	return DataChunk{
		ChunkID:   DCHUNKID,
		ChunkSize: 0,
		WaveData:  Wave{},
	}
}

//Bytes Returns the binary representation of the DataChunk
func (dc DataChunk) Bytes() []byte {
	out := bytes.NewBuffer([]byte{})

	out.WriteString(dc.ChunkID)

	binary.Write(out, binary.LittleEndian, dc.ChunkSize)

	binary.Write(out, binary.LittleEndian, dc.WaveData)

	return out.Bytes()
}

//AddSampleData Adds stereo audio to the DataChunk
func (dc *DataChunk) AddSampleData(leftBuffer, rightBuffer Wave) {
	tmpData := make([]int16, len(leftBuffer)+len(rightBuffer))
	bufferOffset := 0
	for index := 0; index < len(tmpData); index += 2 {
		tmpData[index] = leftBuffer[bufferOffset]
		tmpData[index+1] = rightBuffer[bufferOffset]
		bufferOffset++
	}
	dc.WaveData = append(dc.WaveData, tmpData...)
	dc.ChunkSize = uint32(len(dc.WaveData) * 2)
}

//AddSampleDataMono Adds mono audio to the DataCunk
func (dc *DataChunk) AddSampleDataMono(buffer Wave) {
	tmpData := make([]int16, len(buffer)*2)
	bufferOffset := 0
	for i := 0; i < len(tmpData); i += 2 {
		tmpData[i] = buffer[bufferOffset]
		tmpData[i+1] = buffer[bufferOffset]
		bufferOffset++
	}
	dc.WaveData = append(dc.WaveData, tmpData...)
	dc.ChunkSize = uint32(len(dc.WaveData) * 2)
}
