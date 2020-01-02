package main

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
	//Eight is a eight note
	Eight
)
