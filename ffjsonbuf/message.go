package ffjsonbuf

type Message struct {
	ID     []byte `json:"id"`
	DestID uint32 `json:"dest_id"`
	Type   int8   `json:"type"`
	Kind   int16  `json:"kind"`
	Data   []byte `json:"data"`
}
