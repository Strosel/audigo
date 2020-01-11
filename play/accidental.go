package play

//Accidental represents the different musical accidentals
type Accidental int

const (
	//Flat is a flat (b) note
	Flat Accidental = iota - 1
	//Natural is a unmodified note (♮)
	Natural
	//Sharp is a sharp (#) note
	Sharp
)
