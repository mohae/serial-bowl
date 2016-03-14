package main

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/mohae/serial-bowl/pb"
	"github.com/mohae/serial-bowl/shared"
)

// marshal results holder.
var pbRedditAccount [][]byte
var pbMemInfo [][]byte
var pbBasicMemInfo [][]byte
var pbMessage [][]byte

func init() {
	pbMessage = make([][]byte, shared.Len)
}

func BenchMessagePBMarshal(b *testing.B) {
	b.StopTimer()
	tmp := PreparePBMessageData(shared.MessageData)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			pbMessage[j], _ = proto.Marshal(&tmp[j])
		}
	}
}

func BenchMessagePBUnmarshal(b *testing.B) {
	var tmp pb.Message
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			proto.Unmarshal(pbMessage[j], &tmp)
		}
	}
	_ = tmp
}
