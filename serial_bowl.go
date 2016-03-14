package main

import (
	"fmt"
	"testing"

	"github.com/mohae/serial-bowl/shared"
)

func main() {
	shared.GenData()
	var r shared.Result
	// Reddit Account
	// Flatbuffers Serialize
	br := testing.Benchmark(BenchRedditAccountFBSerialize)
	r.SetFromBenchmarkResult(br)
	fmt.Printf("RedditAccountSerializeFB:\t%s\n", r.String())
	// JSON Marshal
	br = testing.Benchmark(BenchRedditAccountJSONMarshal)
	r.SetFromBenchmarkResult(br)
	fmt.Printf("RedditAccountJSONMarshal:\t%s\n", r.String())
	// Flatbuffers Deserialize
	br = testing.Benchmark(BenchRedditAccountFBDeserialize)
	r.SetFromBenchmarkResult(br)
	fmt.Printf("RedditAccountDeserializeFB:\t%s\n", r.String())
	// JSON Unmarshal
	br = testing.Benchmark(BenchRedditAccountJSONUnmarshal)
	r.SetFromBenchmarkResult(br)
	fmt.Printf("RedditAccountJSONUnmarshal:\t%s\n", r.String())

	// BasicMemInfo
	// Flatbuffers Serialize
	fmt.Println("")
	br = testing.Benchmark(BenchBasicMemInfoFBSerialize)
	r.SetFromBenchmarkResult(br)
	fmt.Printf("BasicMemInfoSerializeFB:\t%s\n", r.String())
	// JSON Marshal
	br = testing.Benchmark(BenchBasicMemInfoJSONMarshal)
	r.SetFromBenchmarkResult(br)
	fmt.Printf("BasicMemInfoJSONMarshal:\t%s\n", r.String())
	// Flatbuffers Deserialize
	br = testing.Benchmark(BenchBasicMemInfoFBDeserialize)
	r.SetFromBenchmarkResult(br)
	fmt.Printf("BasicMemInfoDeserializeFB:\t%s\n", r.String())
	// JSON Unmarshal
	br = testing.Benchmark(BenchBasicMemInfoJSONUnmarshal)
	r.SetFromBenchmarkResult(br)
	fmt.Printf("BasicMemInfoJSONUnmarshal:\t%s\n", r.String())

	// MemInfo
	// Flatbuffers Serialize
	fmt.Println("")
	br = testing.Benchmark(BenchMemInfoFBSerialize)
	r.SetFromBenchmarkResult(br)
	fmt.Printf("MemInfoSerializeFB:\t\t%s\n", r.String())
	// JSON Marshal
	br = testing.Benchmark(BenchMemInfoJSONMarshal)
	r.SetFromBenchmarkResult(br)
	fmt.Printf("MemInfoJSONMarshal:\t\t%s\n", r.String())
	// Flatbuffers Deserialize
	br = testing.Benchmark(BenchMemInfoFBDeserialize)
	r.SetFromBenchmarkResult(br)
	fmt.Printf("MemInfoDeserializeFB:\t\t%s\n", r.String())
	// JSON Unmarshal
	br = testing.Benchmark(BenchMemInfoJSONUnmarshal)
	r.SetFromBenchmarkResult(br)
	fmt.Printf("MemInfoJSONUnmarshal:\t\t%s\n", r.String())

	// Message Data
	dataLen := []int{256, 1024, 4096}
	for _, v := range dataLen {
		fmt.Println("")
		shared.GenMessageData(v, shared.Len)
		fmt.Printf("Data: %d bytes\n", v)
		br = testing.Benchmark(BenchMessageFBSerialize)
		r.SetFromBenchmarkResult(br)
		fmt.Printf("MessageSerializeFB:\t\t%s\n", r.String())
		// JSON Marshal
		br = testing.Benchmark(BenchMessageJSONMarshal)
		r.SetFromBenchmarkResult(br)
		fmt.Printf("MessageJSONMarshal:\t\t%s\n", r.String())
		// Protbuf Marshal
		br = testing.Benchmark(BenchMessagePBMarshal)
		r.SetFromBenchmarkResult(br)
		fmt.Printf("MessagePBMarshal:\t\t%s\n", r.String())
		// Flatbuffers Deserialize
		br = testing.Benchmark(BenchMessageFBDeserialize)
		r.SetFromBenchmarkResult(br)
		fmt.Printf("MessageDeserializeFB:\t\t%s\n", r.String())
		// JSON Unmarshal
		br = testing.Benchmark(BenchMessageJSONUnmarshal)
		r.SetFromBenchmarkResult(br)
		fmt.Printf("MessageJSONUnmarshal:\t\t%s\n", r.String())
		// Protobuf Unmarshal
		br = testing.Benchmark(BenchMessagePBUnmarshal)
		r.SetFromBenchmarkResult(br)
		fmt.Printf("MessagePBUnmarshal:\t\t%s\n", r.String())

	}
}
