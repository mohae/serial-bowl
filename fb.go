package main

import (
	"testing"

	"github.com/mohae/serial-bowl/fb"
)

var fbRedditAccount [][]byte
var fbMemInfo [][]byte
var fbBasicMemInfo [][]byte

func init() {
	fbRedditAccount = make([][]byte, Len)
	fbMemInfo = make([][]byte, Len)
	fbBasicMemInfo = make([][]byte, Len)
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

func BenchMemInfoFBDeserialize(b *testing.B) {
	var tmp *fb.MemInfo
	for i := 0; i < b.N; i++ {
		for j := 0; j < Len; j++ {
			tmp = fb.GetRootAsMemInfo(fbMemInfo[j], 0)
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

func BenchBasicMemInfoFBDeserialize(b *testing.B) {
	var tmp *fb.BasicMemInfo
	for i := 0; i < b.N; i++ {
		for j := 0; j < Len; j++ {
			tmp = fb.GetRootAsBasicMemInfo(fbBasicMemInfo[j], 0)
		}
	}
	_ = tmp
}
