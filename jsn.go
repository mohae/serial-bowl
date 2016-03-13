package main

import (
	"encoding/json"
	"testing"

	"github.com/mohae/serial-bowl/jsn"
)

var jsonRedditAccount [][]byte
var jsonMemInfo [][]byte

func init() {
	jsonRedditAccount = make([][]byte, Len)
	jsonMemInfo = make([][]byte, Len)
}

func BenchRedditAccountJSONMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < Len; j++ {
			jsonRedditAccount[j], _ = json.Marshal(AccountData[j])
		}
	}
}

func BenchRedditAccountJSONUnmarshal(b *testing.B) {
	b.N /= Len
	var tmp jsn.RedditAccount
	for i := 0; i < b.N; i++ {
		for j := 0; j < Len; j++ {
			json.Unmarshal(jsonRedditAccount[j], &tmp)
		}
	}
	_ = tmp
}

func BenchMemInfoJSONMarshal(b *testing.B) {
	b.N /= Len
	for i := 0; i < b.N; i++ {
		for j := 0; j < Len; j++ {
			jsonMemInfo[j], _ = json.Marshal(MemData[j])
		}
	}
}

func BenchMemInfoJSONUnmarshal(b *testing.B) {
	b.N /= Len
	var tmp jsn.MemInfo
	for i := 0; i < b.N; i++ {
		for j := 0; j < Len; j++ {
			json.Unmarshal(jsonMemInfo[j], &tmp)
		}
	}
	_ = tmp
}
