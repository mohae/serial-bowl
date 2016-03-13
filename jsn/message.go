package jsn

type Message struct {
	ID     []byte
	DestID uint32
	Type   int8
	Kind   int16
	Data   []byte
}
