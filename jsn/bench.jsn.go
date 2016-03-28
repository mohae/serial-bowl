package jsn

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/mohae/serial-bowl/shared"
)

var basicMemInfo [][]byte
var memInfo [][]byte
var message [][]byte
var redditAccount [][]byte

// BenchBasicMemInfo runs the BasicMemInfo benches for Marshal/Unmarshal.
func BenchBasicMemInfo() shared.Bench {
	basicMemInfo = make([][]byte, shared.Len)
	bench := shared.Bench{Proto: shared.JSON, StructString: shared.BasicMemInfo.String(), Results: map[shared.Op]shared.Result{}}
	bench.Results[shared.Marshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(basicMemInfoMarshal))
	bench.Results[shared.Unmarshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(basicMemInfoUnmarshal))
	basicMemInfo = nil
	return bench
}

func basicMemInfoMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			basicMemInfo[j], _ = json.Marshal(shared.BasicMemInfoData[j])
		}
	}
}

func basicMemInfoUnmarshal(b *testing.B) {
	var tmp shared.ShBasicMemInfo
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			json.Unmarshal(basicMemInfo[j], &tmp)
		}
	}
	_ = tmp
}

// BenchMemInfo runs the MemInfo benches for Marshal/Unmarshal.
func BenchMemInfo() shared.Bench {
	memInfo = make([][]byte, shared.Len)
	bench := shared.Bench{Proto: shared.JSON, StructString: shared.MemInfo.String(), Results: map[shared.Op]shared.Result{}}
	bench.Results[shared.Marshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(memInfoMarshal))
	bench.Results[shared.Unmarshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(memInfoUnmarshal))
	memInfo = nil
	return bench
}

func memInfoMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			memInfo[j], _ = json.Marshal(shared.MemInfoData[j])
		}
	}
}

func memInfoUnmarshal(b *testing.B) {
	var tmp shared.ShMemInfo
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			json.Unmarshal(memInfo[j], &tmp)
		}
	}
	_ = tmp
}

// BenchMessage runs the MemInfo benches for Marshal/Unmarshal.
func BenchMessage(l int) shared.Bench {
	message = make([][]byte, shared.Len)
	bench := shared.Bench{Proto: shared.JSON, StructString: fmt.Sprintf("%s %dB", shared.Message.String(), l), Results: map[shared.Op]shared.Result{}}
	bench.Results[shared.Marshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(messageMarshal))
	bench.Results[shared.Unmarshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(messageUnmarshal))
	message = nil
	return bench
}

func messageMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			message[j], _ = json.Marshal(shared.MessageData[j])
		}
	}
}

func messageUnmarshal(b *testing.B) {
	var tmp shared.ShMessage
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			json.Unmarshal(message[j], &tmp)
		}
	}
	_ = tmp
}

// BenchRedditAccount runs the MemInfo benches for Marshal/Unmarshal.
func BenchRedditAccount() shared.Bench {
	redditAccount = make([][]byte, shared.Len)
	bench := shared.Bench{Proto: shared.JSON, StructString: shared.RedditAccount.String(), Results: map[shared.Op]shared.Result{}}
	bench.Results[shared.Marshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(redditAccountMarshal))
	bench.Results[shared.Unmarshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(redditAccountUnmarshal))
	redditAccount = nil
	return bench
}

func redditAccountMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			redditAccount[j], _ = json.Marshal(shared.RedditAccountData[j])
		}
	}
}

func redditAccountUnmarshal(b *testing.B) {
	var tmp shared.ShRedditAccount
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			json.Unmarshal(redditAccount[j], &tmp)
		}
	}
	_ = tmp
}
