package play

import (
	"math"
	"time"

	"github.com/strosel/audigo/midi"
)

//MultiChord describes a chord wehere one or several notes changes during the duration
//index of each note should be the same in each chord and a None Duration is used to show no change
type MultiChord []Chord

//Duration calculates and returns the duration of the tie based on
//how long the measure takes, aka the length of a whole note
func (mc MultiChord) Duration(measure time.Duration) time.Duration {
	m := time.Duration(0)
	for _, n := range mc {
		if n.Duration(measure) > m {
			m = n.Duration(measure)
		}
	}
	return m
}

//TickDuration calculates and returns the duration of the tie in ticks
//based on how many ticks a quarter note is
func (mc MultiChord) TickDuration(quarter uint16) uint16 {
	m := uint16(0)
	for _, n := range mc {
		if n.TickDuration(quarter) > m {
			m = n.TickDuration(quarter)
		}
	}
	return m
}

//ToMIDI converst the chord into an array of midi events
func (mc MultiChord) ToMIDI(ticks uint16, ch, vel uint8) []midi.Event {
	out := []midi.Event{}
	e := &midi.VoiceEvent{
		Channel: ch,
	}
	//Start the first chord
	for i, n := range mc[0] {
		if i != 0 {
			e = &midi.VoiceEvent{
				Channel: ch,
			}
		}
		e.NoteOn(0x3C+uint8(n.Dist()), vel)
		out = append(out, e)
	}

	depth := []int{}
	//passed since last start on row
	dt := []uint16{}
	for i := 0; i < len(mc[0]); i++ {
		depth = append(depth, 0)
		dt = append(dt, 0)
	}

	order := mc.order()
	mat := mc.matrixDuration(ticks)
	for _, i := range order {
		//duration = time this note should take - time since this row last changed
		dur := mat[depth[i]][i] - dt[i]

		//stop the current note
		stop := mc[depth[i]][i].ToMIDI(ticks, ch, vel)[1]
		stop.SetDelta(midi.VLQ(dur))
		out = append(out, stop)

		//start the next note if there is one
		depth[i]++
		if depth[i] < len(mc) && mc[depth[i]][i].Value != None {
			start := mc[depth[i]][i].ToMIDI(ticks, ch, vel)[0]
			start.SetDelta(0)
			out = append(out, start)
		}

		//increse all times since change by duration
		for j := range dt {
			dt[j] += dur
		}
		//this row just changed so time since last is 0
		dt[i] = 0
	}

	return out
}

//matrixDuration returns a matrix of the note durations
//this function is only implemented for ticks as it is only used for durations relative to eaher
func (mc MultiChord) matrixDuration(quarter uint16) [][]uint16 {
	out := [][]uint16{}
	for i, c := range mc {
		out = append(out, []uint16{})
		for _, n := range c {
			out[i] = append(out[i], n.TickDuration(quarter))
		}
	}

	return out
}

//order is the order of termination for the notes
func (mc MultiChord) order() []int {
	out := []int{}
	durs := mc.matrixDuration(64) // constant 64 to support down to 256th note
	depth := []int{}
	//summed durations of chord parts "top" to "bottom"
	d := []uint16{}
	//setup
	for i := 0; i < len(durs[0]); i++ {
		d = append(d, durs[0][i])
		depth = append(depth, 1)
	}

	for {
		i := mini(d)
		if d[i] == math.MaxUint16 {
			break
		}
		out = append(out, i)
		if depth[i] < len(durs) {
			d[i] += durs[depth[i]][i]
			depth[i]++
		} else {
			d[i] = math.MaxUint16
		}
	}

	return out
}

//mini index of min in a
func mini(a []uint16) int {
	m := 0
	for i := range a {
		if a[i] < a[m] {
			m = i
		}
	}
	return m
}
