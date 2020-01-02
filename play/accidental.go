package main

//Accidental represents the different musical accidentals
type Accidental int

const (
	//Flat is a flat (b) note
	Flat Accidental = iota - 2
	//Natural is a unmodified note (♮)
	Natural
	//Sharp is a sharp (#) note
	Sharp
)
