package main

import (
	"testing"

	"github.com/mohae/serial-bowl/fb"
	"github.com/mohae/serial-bowl/shared"
)

var fbRedditAccount [][]byte
var fbMemInfo [][]byte
var fbBasicMemInfo [][]byte
var fbMessage [][]byte

func init() {
	fbRedditAccount = make([][]byte, shared.Len)
	fbMemInfo = make([][]byte, shared.Len)
	fbBasicMemInfo = make([][]byte, shared.Len)
	fbMessage = make([][]byte, shared.Len)
}

func BenchRedditAccountFBSerialize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			fbRedditAccount[j] = fb.SerializeRedditAccount(shared.RedditAccountData[j])
		}
	}
}

func BenchRedditAccountFBDeserialize(b *testing.B) {
	var tmp *fb.RedditAccount
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			tmp = fb.GetRootAsRedditAccount(fbRedditAccount[j], 0)
		}
	}
	_ = tmp
}

func BenchBasicMemInfoFBSerialize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			fbBasicMemInfo[j] = fb.SerializeBasicMemInfo(shared.BasicMemInfoData[j])
		}
	}
}

func BenchBasicMemInfoFBDeserialize(b *testing.B) {
	var tmp *fb.BasicMemInfo
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			tmp = fb.GetRootAsBasicMemInfo(fbBasicMemInfo[j], 0)
		}
	}
	_ = tmp
}

func BenchMemInfoFBSerialize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			fbMemInfo[j] = fb.SerializeMemInfo(shared.MemInfoData[j])
		}
	}
}

func BenchMemInfoFBDeserialize(b *testing.B) {
	var tmp *fb.MemInfo
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			tmp = fb.GetRootAsMemInfo(fbMemInfo[j], 0)
		}
	}
	_ = tmp
}

func BenchMessageFBSerialize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			fbMessage[j] = fb.SerializeMessage(shared.MessageData[j])
		}
	}
}

func BenchMessageFBDeserialize(b *testing.B) {
	var tmp *fb.Message
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			tmp = fb.GetRootAsMessage(fbMessage[j], 0)
		}
	}
	_ = tmp
}
