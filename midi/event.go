package midi

type Event interface {
	Delta() VLQ
	Bytes() []byte
}
