package jsn

import (
	"encoding/json"
	"testing"

	"github.com/mohae/serial-bowl/shared"
)

var basicMemInfo [][]byte
var memInfo [][]byte
var message [][]byte
var redditAccount [][]byte

// BenchBasicMemInfo runs the BasicMemInfo benches for Marshal/Unmarshal.
func BenchBasicMemInfo() (m, u testing.BenchmarkResult) {
	basicMemInfo = make([][]byte, shared.Len)
	m = testing.Benchmark(basicMemInfoMarshal)
	u = testing.Benchmark(basicMemInfoUnmarshal)
	basicMemInfo = nil
	return m, u
}

func basicMemInfoMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			basicMemInfo[j], _ = json.Marshal(shared.BasicMemInfoData[j])
		}
	}
}

func basicMemInfoUnmarshal(b *testing.B) {
	var tmp shared.BasicMemInfo
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			json.Unmarshal(basicMemInfo[j], &tmp)
		}
	}
	_ = tmp
}

// BenchMemInfo runs the MemInfo benches for Marshal/Unmarshal.
func BenchMemInfo() (m, u testing.BenchmarkResult) {
	memInfo = make([][]byte, shared.Len)
	m = testing.Benchmark(memInfoMarshal)
	u = testing.Benchmark(memInfoUnmarshal)
	memInfo = nil
	return m, u
}

func memInfoMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			memInfo[j], _ = json.Marshal(shared.MemInfoData[j])
		}
	}
}

func memInfoUnmarshal(b *testing.B) {
	var tmp shared.MemInfo
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			json.Unmarshal(memInfo[j], &tmp)
		}
	}
	_ = tmp
}

// BenchMessage runs the MemInfo benches for Marshal/Unmarshal.
func BenchMessage() (m, u testing.BenchmarkResult) {
	message = make([][]byte, shared.Len)
	m = testing.Benchmark(messageMarshal)
	u = testing.Benchmark(messageUnmarshal)
	message = nil
	return m, u
}

func messageMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			message[j], _ = json.Marshal(shared.MessageData[j])
		}
	}
}

func messageUnmarshal(b *testing.B) {
	var tmp shared.Message
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			json.Unmarshal(message[j], &tmp)
		}
	}
	_ = tmp
}

// BenchRedditAccount runs the MemInfo benches for Marshal/Unmarshal.
func BenchRedditAccount() (m, u testing.BenchmarkResult) {
	redditAccount = make([][]byte, shared.Len)
	m = testing.Benchmark(redditAccountMarshal)
	u = testing.Benchmark(redditAccountUnmarshal)
	redditAccount = nil
	return m, u
}

func redditAccountMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			redditAccount[j], _ = json.Marshal(shared.RedditAccountData[j])
		}
	}
}

func redditAccountUnmarshal(b *testing.B) {
	var tmp shared.RedditAccount
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			json.Unmarshal(redditAccount[j], &tmp)
		}
	}
	_ = tmp
}
