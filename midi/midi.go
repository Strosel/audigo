package midi

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"
	"regexp"
	"strings"
)

type MIDI struct {
	Header Header
	Tracks []Track
}

//Bytes the MIDI file byte array
func (m MIDI) Bytes() []byte {
	out := bytes.NewBuffer([]byte{})

	//if there is only one track format should always be single
	if len(m.Tracks) == 1 {
		m.Header.Format = Single
	}

	binary.Write(out, binary.BigEndian, m.Header.Bytes())
	for _, t := range m.Tracks {
		binary.Write(out, binary.BigEndian, t.Bytes())
	}

	return out.Bytes()
}

//Save saves the MIDI file as a new .midi/.mid file
func (m MIDI) Save(filename string) error {
	re, _ := regexp.Compile(`\.midi?$`)
	if !re.MatchString(strings.ToLower(filename)) {
		filename += ".midi"
	}

	return ioutil.WriteFile(filename, m.Bytes(), 0666)
}
