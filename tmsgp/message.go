package tmsgp

//go:generate msgp
type Message struct {
	ID     []byte `msg:"id"`
	DestID uint32 `msg:"dest_id"`
	Type   int8   `msg:"type"`
	Kind   int16  `msg:"kind"`
	Data   []byte `msg:"data"`
}
