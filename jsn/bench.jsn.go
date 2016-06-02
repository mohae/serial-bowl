package jsn

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/mohae/benchutil"
	"github.com/mohae/serial-bowl/shared"
)

var (
	basicMemInfo  [][]byte
	memInfo       [][]byte
	message       [][]byte
	redditAccount [][]byte
	cpuInfo       [][]byte
)

func newBench(s string) benchutil.Bench {
	bench := benchutil.NewBench("encoding/json")
	bench.Iterations = shared.Len
	bench.Group = s
	bench.SubGroup = shared.JSON.String()
	return bench
}

// BenchBasicMemInfo runs the BasicMemInfo benches for Marshal/Unmarshal.
func BenchBasicMemInfo() []benchutil.Bench {
	var results []benchutil.Bench
	basicMemInfo = make([][]byte, shared.Len)

	bench := newBench(shared.BasicMemInfo.String())
	bench.Desc = shared.Marshal.String()
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(basicMemInfoMarshal))
	results = append(results, bench)
	bench.Desc = shared.Unmarshal.String()
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(basicMemInfoUnmarshal))
	results = append(results, bench)
	basicMemInfo = nil
	return results
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
func BenchMemInfo() []benchutil.Bench {
	var results []benchutil.Bench
	memInfo = make([][]byte, shared.Len)

	bench := newBench(shared.MemInfo.String())
	bench.Desc = shared.Marshal.String()
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(memInfoMarshal))
	results = append(results, bench)
	bench.Desc = shared.Unmarshal.String()
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(memInfoUnmarshal))
	results = append(results, bench)
	memInfo = nil
	return results
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

// BenchMessage runs the Message benches for Marshal/Unmarshal.
func BenchMessage(l int) []benchutil.Bench {
	var results []benchutil.Bench
	message = make([][]byte, shared.Len)

	bench := newBench(fmt.Sprintf("%s: %d", shared.Message.String(), l))
	bench.Desc = shared.Marshal.String()
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(messageMarshal))
	results = append(results, bench)
	bench.Desc = shared.Unmarshal.String()
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(messageUnmarshal))
	results = append(results, bench)
	message = nil
	return results
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

// BenchRedditAccount runs the RedditAccount benches for Marshal/Unmarshal.
func BenchRedditAccount() []benchutil.Bench {
	var results []benchutil.Bench
	redditAccount = make([][]byte, shared.Len)

	bench := newBench(shared.RedditAccount.String())
	bench.Desc = shared.Marshal.String()
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(redditAccountMarshal))
	results = append(results, bench)
	bench.Desc = shared.Unmarshal.String()
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(redditAccountUnmarshal))
	results = append(results, bench)
	redditAccount = nil
	return results
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

// BenchCPUInfo runs the CPUInfo benches for Marshal/Unmarshal.
func BenchCPUInfo(n int) []benchutil.Bench {
	var results []benchutil.Bench
	cpuInfo = make([][]byte, shared.Len)

	bench := newBench(fmt.Sprintf("%s: %d", shared.CPUInfo.String(), n))
	bench.Desc = shared.Marshal.String()
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(cpuInfoMarshal))
	results = append(results, bench)
	bench.Desc = shared.Unmarshal.String()
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(cpuInfoUnmarshal))
	results = append(results, bench)
	redditAccount = nil
	return results
}

func cpuInfoMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			cpuInfo[j], _ = json.Marshal(shared.CPUInfoData[j])
		}
	}
}

func cpuInfoUnmarshal(b *testing.B) {
	var tmp shared.ShCPUInfo
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			json.Unmarshal(cpuInfo[j], &tmp)
		}
	}
	_ = tmp
}
