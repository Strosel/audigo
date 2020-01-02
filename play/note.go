package play

import (
	"math"
	"regexp"
	"strconv"
	"time"

	"github.com/strosel/audigo/midi"
)

const a4 = 440.0

//Note is a musical note
type Note struct {
	Note       byte
	Octave     int
	Accidental Accidental
	Value      Duration
	Dots       int
	//Rest is a rest before the note
	Rest Duration
}

//Parse parses a string of format c4b. into a note
//uses regex ([a-gA-G])(\d+)([b#]?)(\.*) and panics on error
func Parse(note string, d Duration) Note {
	if len(note) < 2 {
		panic("Invalid note")
	}
	re, err := regexp.Compile(`([a-gA-G])(\d+)([b#]?)(\.*)`)
	if err != nil {
		panic("Parser Error")
	}

	all := re.FindAllStringSubmatch(note, -1)
	if all == nil || len(all) == 0 {
		panic("Invalid note")
	}
	sNote := all[0][1:]

	oct, err := strconv.Atoi(sNote[1])
	if err != nil {
		panic("Invalid octave")
	}

	acc := Natural
	if sNote[2] == "b" {
		acc = Flat
	} else if sNote[2] == "#" {
		acc = Sharp
	}

	n := Note{
		Note:       sNote[0][0],
		Octave:     oct,
		Accidental: acc,
		Value:      d,
		Dots:       len(sNote[3]),
		Rest:       None,
	}

	return n
}

//ParseR Like Parse but with the rest r before the note
func ParseR(note string, d Duration, r Duration) Note {
	n := Parse(note, d)
	n.Rest = r
	return n
}

//Dist calculates the dustance from c4 in half steps
func (n Note) Dist() float64 {
	//number of half steps from c within the same octave
	hs := map[byte]float64{
		'a': 9,
		'b': 11,
		'c': 0,
		'd': 2,
		'e': 4,
		'f': 5,
		'g': 7,
	}

	//bitwise or with 96 for upper to lower case conversion
	n.Note |= 96
	//wrap around after g
	n.Note %= 'h'

	//calc half steps
	return hs[n.Note] + float64(n.Octave-4)*12. + float64(n.Accidental)
}

//Frequency calculates and returns the frequency of the note
//based on https://pages.mtu.edu/~suits/NoteFreqCalcs.html
func (n Note) Frequency() float64 {
	//calc frequency, -9 half steps to base on a4
	return a4 * math.Pow(math.Pow(2, 1./12.), n.Dist()-9)
}

//Duration calculates and returns the duration of the note based on
//how long the measure takes, aka the length of a whole note
func (n Note) Duration(measure time.Duration) time.Duration {
	//the fraction of the measure the note takes
	fraq := 1. / math.Pow(2., float64(n.Value))

	//add dots
	if n.Dots > 0 {
		//dots depend on the previous value/length, keep that
		prev := fraq
		for i := 0; i < n.Dots; i++ {
			prev /= 2.
			fraq += prev
		}
	}

	return time.Duration(float64(measure) * fraq)
}

//TickDuration calculates and returns the duration of the note in ticks
//based on how many ticks a quarter note is
func (n Note) TickDuration(quarter uint16) uint16 {
	//the fraction of the measure the note takes
	fraq := 1. / math.Pow(2., float64(n.Value))

	//add dots
	if n.Dots > 0 {
		//dots depend on the previous value/length, keep that
		prev := fraq
		for i := 0; i < n.Dots; i++ {
			prev /= 2.
			fraq += prev
		}
	}

	return uint16(float64(4*quarter) * fraq)
}

//RestDuration calculates and returns the duration of the rest preceding the
//note based on how long the measure takes, aka the length of a whole note
func (n Note) RestDuration(measure time.Duration) time.Duration {
	if n.Rest == None {
		return 0
	}
	//the fraction of the measure the rest takes
	fraq := 1. / math.Pow(2., float64(n.Rest))
	return time.Duration(float64(measure) * fraq)
}

//RestTickDuration calculates and returns the duration of the rest preceding the
//note in ticks based on how many ticks a quarter note is
func (n Note) RestTickDuration(quarter uint16) uint16 {
	if n.Rest == None {
		return 0
	}
	//the fraction of the measure the rest takes
	fraq := 1. / math.Pow(2., float64(n.Rest))
	return uint16(float64(4*quarter) * fraq)
}

//ToMIDI converts the note to a midi NoteOn and NoteOff event
//with the given ticks for a quarter note, channel and velocity
func (n Note) ToMIDI(ticks uint16, ch, vel uint8) []midi.Event {
	out := []midi.Event{}
	e := &midi.VoiceEvent{
		Channel:  ch,
		Duration: midi.VLQ(n.RestTickDuration(ticks)),
	}
	e.NoteOn(0x3C+uint8(n.Dist()), vel)
	out = append(out, e)

	e = &midi.VoiceEvent{
		Channel:  ch,
		Duration: midi.VLQ(n.TickDuration(ticks)),
	}
	e.NoteOff(0x3C+uint8(n.Dist()), vel)
	out = append(out, e)
	return out
}
