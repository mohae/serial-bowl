package main

import (
	"encoding/json"
	"testing"

	"github.com/mohae/serial-bowl/jsn"
	"github.com/mohae/serial-bowl/shared"
)

var jsonRedditAccount [][]byte
var jsonMemInfo [][]byte
var jsonBasicMemInfo [][]byte
var jsonMessage [][]byte

func init() {
	jsonRedditAccount = make([][]byte, shared.Len)
	jsonMemInfo = make([][]byte, shared.Len)
	jsonBasicMemInfo = make([][]byte, shared.Len)
	jsonMessage = make([][]byte, shared.Len)
}

func BenchRedditAccountJSONMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			jsonRedditAccount[j], _ = json.Marshal(shared.RedditAccountData[j])
		}
	}
}

func BenchRedditAccountJSONUnmarshal(b *testing.B) {
	var tmp jsn.RedditAccount
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			json.Unmarshal(jsonRedditAccount[j], &tmp)
		}
	}
	_ = tmp
}

func BenchMemInfoJSONMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			jsonMemInfo[j], _ = json.Marshal(shared.MemInfoData[j])
		}
	}
}

func BenchMemInfoJSONUnmarshal(b *testing.B) {
	var tmp jsn.MemInfo
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			json.Unmarshal(jsonMemInfo[j], &tmp)
		}
	}
	_ = tmp
}

func BenchBasicMemInfoJSONMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			jsonBasicMemInfo[j], _ = json.Marshal(shared.BasicMemInfoData[j])
		}
	}
}

func BenchBasicMemInfoJSONUnmarshal(b *testing.B) {
	var tmp jsn.BasicMemInfo
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			json.Unmarshal(jsonBasicMemInfo[j], &tmp)
		}
	}
	_ = tmp
}

func BenchMessageJSONMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			jsonMessage[j], _ = json.Marshal(shared.MessageData[j])
		}
	}
}

func BenchMessageJSONUnmarshal(b *testing.B) {
	var tmp jsn.Message
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			json.Unmarshal(jsonMessage[j], &tmp)
		}
	}
	_ = tmp
}
