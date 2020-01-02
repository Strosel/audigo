package play

func main() {
	s := Song{
		Instruments: map[string][]Playable{
			"piano-left": []Playable{},
		},
	}

	for _, v := range []struct {
		n byte
		o int
	}{{'c', 4}, {'d', 4}, {'e', 4}, {'f', 4}, {'g', 4}, {'a', 4}, {'b', 4}, {'C', 5}} {
		s.Instruments["piano-left"] = append(s.Instruments["piano-left"], Note{
			Note:   v.n,
			Octave: v.o,
			Value:  Quarter,
		})
	}

	s.ToMIDI().Save("test")
}
