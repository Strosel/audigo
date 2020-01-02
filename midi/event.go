package midi

//Event a midi track chunk component
type Event interface {
	Delta() VLQ
	Bytes() []byte
}
