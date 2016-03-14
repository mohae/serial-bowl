package main

import (
	"fmt"
	"testing"

	"github.com/mohae/serial-bowl/pb"
	"github.com/mohae/serial-bowl/shared"
)

func main() {
	shared.GenData()
	var r shared.Result

	// BasicMemInfo
	// Flatbuffers Serialize
	fmt.Println("")
	br := testing.Benchmark(BenchBasicMemInfoFBSerialize)
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
	// PB v3
	m, u := pb.BenchBasicMemInfo()
	r.SetFromBenchmarkResult(m)
	fmt.Printf("PB v3 Marshal: BasicMemInfo:\t%s\n", r.String())
	r.SetFromBenchmarkResult(u)
	fmt.Printf("PB v3 Unmarshal: BasicMemInfo:\t%s\n", r.String())

	// Marshal/Serialize: MemInfo
	fmt.Println("Marshal: MemInfo")
	// Flatbuffers Serialize
	br = testing.Benchmark(BenchMemInfoFBSerialize)
	r.SetFromBenchmarkResult(br)
	fmt.Printf("Flatbuffers: MemInfo:\t\t%s\n", r.String())
	// JSON Marshal
	br = testing.Benchmark(BenchMemInfoJSONMarshal)
	r.SetFromBenchmarkResult(br)
	fmt.Printf("JSON: MemInfo:\t\t%s\n", r.String())
	// Unmarshal/Deserialize
	fmt.Println("Unmarshal: MemInfo")
	// Flatbuffers Deserialize
	br = testing.Benchmark(BenchMemInfoFBDeserialize)
	r.SetFromBenchmarkResult(br)
	fmt.Printf("Flatbuffers MemInfo:\t\t%s\n", r.String())
	// JSON Unmarshal
	br = testing.Benchmark(BenchMemInfoJSONUnmarshal)
	r.SetFromBenchmarkResult(br)
	fmt.Printf("JSON MemInfo:\t\t%s\n", r.String())
	// PB v3
	m, u = pb.BenchMemInfo()
	r.SetFromBenchmarkResult(m)
	fmt.Printf("PB v3 Marshal: MemInfo:\t%s\n", r.String())
	r.SetFromBenchmarkResult(u)
	fmt.Printf("PB v3 Unmarshal: MemInfo:\t%s\n", r.String())

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
		// Flatbuffers Deserialize
		br = testing.Benchmark(BenchMessageFBDeserialize)
		r.SetFromBenchmarkResult(br)
		fmt.Printf("MessageDeserializeFB:\t\t%s\n", r.String())
		// JSON Unmarshal
		br = testing.Benchmark(BenchMessageJSONUnmarshal)
		r.SetFromBenchmarkResult(br)
		fmt.Printf("MessageJSONUnmarshal:\t\t%s\n", r.String())
		// PB v3
		m, u = pb.BenchMessage()
		r.SetFromBenchmarkResult(m)
		fmt.Printf("PB v3 Marshal: Message:\t%s\n", r.String())
		r.SetFromBenchmarkResult(u)
		fmt.Printf("PB v3 Unmarshal: Message:\t%s\n", r.String())

	}

	// Reddit Account
	// Flatbuffers Serialize
	br = testing.Benchmark(BenchRedditAccountFBSerialize)
	r.SetFromBenchmarkResult(br)
	fmt.Printf("RedditAccountFBSerialize:\t%s\n", r.String())
	// JSON Marshal
	br = testing.Benchmark(BenchRedditAccountJSONMarshal)
	r.SetFromBenchmarkResult(br)
	fmt.Printf("RedditAccountJSONMarshal:\t%s\n", r.String())
	// Flatbuffers Deserialize
	br = testing.Benchmark(BenchRedditAccountFBDeserialize)
	r.SetFromBenchmarkResult(br)
	fmt.Printf("RedditAccountFBDeserialize:\t%s\n", r.String())
	// JSON Unmarshal
	br = testing.Benchmark(BenchRedditAccountJSONUnmarshal)
	r.SetFromBenchmarkResult(br)
	fmt.Printf("RedditAccountJSONUnmarshal:\t%s\n", r.String())
	// PB v3
	m, u = pb.BenchRedditAccount()
	r.SetFromBenchmarkResult(m)
	fmt.Printf("PB v3:  RedditAccountMarshal:\t%s\n", r.String())
	r.SetFromBenchmarkResult(u)
	fmt.Printf("PB v3:  RedditAccountUnmarshal:\t%s\n", r.String())
}
