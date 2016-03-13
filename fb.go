package main

import (
	"testing"

	"github.com/mohae/serial-bowl/fb"
)

var fbRedditAccount [][]byte
var fbMemInfo [][]byte
var fbBasicMemInfo [][]byte
var fbMessage [][]byte

func init() {
	fbRedditAccount = make([][]byte, Len)
	fbMemInfo = make([][]byte, Len)
	fbBasicMemInfo = make([][]byte, Len)
	fbMessage = make([][]byte, Len)
}

func BenchRedditAccountFBSerialize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < Len; j++ {
			fbRedditAccount[j] = fb.SerializeRedditAccount(AccountData[j])
		}
	}
}

func BenchRedditAccountFBDeserialize(b *testing.B) {
	var tmp *fb.RedditAccount
	for i := 0; i < b.N; i++ {
		for j := 0; j < Len; j++ {
			tmp = fb.GetRootAsRedditAccount(fbRedditAccount[j], 0)
		}
	}
	_ = tmp
}

func BenchBasicMemInfoFBSerialize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < Len; j++ {
			fbBasicMemInfo[j] = fb.SerializeBasicMemInfo(BasicMemData[j])
		}
	}
}

func BenchBasicMemInfoFBDeserialize(b *testing.B) {
	var tmp *fb.BasicMemInfo
	for i := 0; i < b.N; i++ {
		for j := 0; j < Len; j++ {
			tmp = fb.GetRootAsBasicMemInfo(fbBasicMemInfo[j], 0)
		}
	}
	_ = tmp
}

func BenchMemInfoFBSerialize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < Len; j++ {
			fbMemInfo[j] = fb.SerializeMemInfo(MemData[j])
		}
	}
}

func BenchMemInfoFBDeserialize(b *testing.B) {
	var tmp *fb.MemInfo
	for i := 0; i < b.N; i++ {
		for j := 0; j < Len; j++ {
			tmp = fb.GetRootAsMemInfo(fbMemInfo[j], 0)
		}
	}
	_ = tmp
}

func BenchMessageFBSerialize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < Len; j++ {
			fbMessage[j] = fb.SerializeMessage(MessageData[j])
		}
	}
}

func BenchMessageFBDeserialize(b *testing.B) {
	var tmp *fb.Message
	for i := 0; i < b.N; i++ {
		for j := 0; j < Len; j++ {
			tmp = fb.GetRootAsMessage(fbMessage[j], 0)
		}
	}
	_ = tmp
}
