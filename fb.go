package main

import (
	"testing"

	"github.com/mohae/serial-bowl/fb"
)

var fbRedditAccount [][]byte
var fbMemInfo [][]byte

func init() {
	fbRedditAccount = make([][]byte, Len)
	fbMemInfo = make([][]byte, Len)
}

func BenchRedditAccountFBSerialize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < Len; j++ {
			fbRedditAccount[j] = fb.SerializeRedditAccount(AccountData[j])
		}
	}
}

func BenchRedditAccountFBDeserialize(b *testing.B) {
	b.N /= Len
	var tmp *fb.RedditAccount
	for i := 0; i < b.N; i++ {
		for j := 0; j < Len; j++ {
			tmp = fb.GetRootAsRedditAccount(fbRedditAccount[j], 0)
		}
	}
	_ = tmp
}

func BenchMemInfoFBSerialize(b *testing.B) {
	b.N /= Len
	for i := 0; i < b.N; i++ {
		for j := 0; j < Len; j++ {
			fbMemInfo[j] = fb.SerializeMemInfo(MemData[j])
		}
	}
}

func BenchMemInfoFBDeserialize(b *testing.B) {
	b.N /= Len
	var tmp *fb.MemInfo
	for i := 0; i < b.N; i++ {
		for j := 0; j < Len; j++ {
			tmp = fb.GetRootAsMemInfo(fbMemInfo[j], 0)
		}
	}
	_ = tmp
}
