package gobs

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"testing"

	"github.com/mohae/serial-bowl/shared"
)

var basicMemInfo [][]byte
var memInfo [][]byte
var message [][]byte
var redditAccount [][]byte

var buff bytes.Buffer
var enc = gob.NewEncoder(&buff)
var dec = gob.NewDecoder(&buff)

// BenchBasicMemInfo runs the BasicMemInfo benches for Marshal/Unmarshal.
func BenchBasicMemInfo() shared.Bench {
	basicMemInfo = make([][]byte, shared.Len)
	bench := shared.Bench{Proto: shared.Gob, StructString: shared.BasicMemInfo.String(), Results: map[shared.Op]shared.Result{}}
	bench.Results[shared.Marshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(basicMemInfoMarshal))
	bench.Results[shared.Unmarshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(basicMemInfoUnmarshal))
	basicMemInfo = nil
	return bench
}

func basicMemInfoMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			buff.Reset()
			enc.Encode(shared.BasicMemInfoData[j])
			basicMemInfo[j] = buff.Bytes()
		}
	}
}

func basicMemInfoUnmarshal(b *testing.B) {
	var tmp shared.ShBasicMemInfo
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			buff.Reset()
			buff.Write(basicMemInfo[j])
			dec.Decode(&tmp)
		}
	}
	_ = tmp
}

// BenchMemInfo runs the MemInfo benches for Marshal/Unmarshal.
func BenchMemInfo() shared.Bench {
	memInfo = make([][]byte, shared.Len)
	bench := shared.Bench{Proto: shared.Gob, StructString: shared.MemInfo.String(), Results: map[shared.Op]shared.Result{}}
	bench.Results[shared.Marshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(memInfoMarshal))
	bench.Results[shared.Unmarshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(memInfoUnmarshal))
	memInfo = nil
	return bench
}

func memInfoMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			buff.Reset()
			enc.Encode(shared.MemInfoData[j])
			memInfo[j] = buff.Bytes()
		}
	}
}

func memInfoUnmarshal(b *testing.B) {
	var tmp shared.ShMemInfo
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			buff.Reset()
			buff.Write(memInfo[j])
		}
	}
	_ = tmp
}

// BenchMessage runs the MemInfo benches for Marshal/Unmarshal.
func BenchMessage(l int) shared.Bench {
	message = make([][]byte, shared.Len)
	bench := shared.Bench{Proto: shared.Gob, StructString: fmt.Sprintf("%s %dB", shared.Message.String(), l), Results: map[shared.Op]shared.Result{}}
	bench.Results[shared.Marshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(messageMarshal))
	bench.Results[shared.Unmarshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(messageUnmarshal))
	message = nil
	return bench
}

func messageMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			buff.Reset()
			enc.Encode(shared.MessageData[j])
			message[j] = buff.Bytes()
		}
	}
}

func messageUnmarshal(b *testing.B) {
	var tmp shared.ShMessage
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			buff.Reset()
			buff.Write(message[j])
			dec.Decode(&tmp)
		}
	}
	_ = tmp
}

// BenchRedditAccount runs the MemInfo benches for Marshal/Unmarshal.
func BenchRedditAccount() shared.Bench {
	redditAccount = make([][]byte, shared.Len)
	bench := shared.Bench{Proto: shared.Gob, StructString: shared.RedditAccount.String(), Results: map[shared.Op]shared.Result{}}
	bench.Results[shared.Marshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(redditAccountMarshal))
	bench.Results[shared.Unmarshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(redditAccountUnmarshal))
	redditAccount = nil
	return bench
}

func redditAccountMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			buff.Reset()
			enc.Encode(shared.RedditAccountData[j])
			redditAccount[j] = buff.Bytes()
		}
	}
}

func redditAccountUnmarshal(b *testing.B) {
	var tmp shared.ShRedditAccount
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			buff.Reset()
			buff.Write(redditAccount[j])
			dec.Decode(&tmp)
		}
	}
	_ = tmp
}
