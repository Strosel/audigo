package play

//Duration is the duration of a note relative to the meter
//as notes are based in powers of two Duration is stored as the exponent
//and therefore easily expanded upon
type Duration int

const (
	//Whole is a whole note
	Whole Duration = iota
	//Half is a half note
	Half
	//Quarter is a quarter note
	Quarter
	//Eighth is a eighth note
	Eighth
	//Sixteenth is a sixteenth note
	Sixteenth
)
