package play

//Instrument describes an instrument staff
type Instrument struct {
	Staff []Playable
	pos   int
}

//HasNext returns if there is another playable in the staff
func (i *Instrument) HasNext() bool {
	return i.pos+1 < len(i.Staff)
}

//Peek returns the next playable in the staff
func (i *Instrument) Peek() Playable {
	return i.Staff[i.pos+1]
}

//Next returns current playable in the staff and increments the internal position
func (i *Instrument) Next() Playable {
	p := i.Staff[i.pos]
	i.pos++
	return p
}
