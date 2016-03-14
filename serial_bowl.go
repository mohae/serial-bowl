package main

import (
	"fmt"

	"github.com/mohae/serial-bowl/fb"
	"github.com/mohae/serial-bowl/jsn"
	"github.com/mohae/serial-bowl/pb"
	"github.com/mohae/serial-bowl/shared"
)

func main() {
	shared.GenData()
	var r shared.Result

	// BasicMemInfo
	// Flatbuffers
	m, u := fb.BenchBasicMemInfo()
	r.SetFromBenchmarkResult(m)
	fmt.Printf("Flatbuffers Serialize: BasicMemInfo:\t%s\n", r.String())
	r.SetFromBenchmarkResult(u)
	fmt.Printf("Flatbuffers Deserialize: BasicMemInfo:\t%s\n", r.String())
	// JSON
	m, u = jsn.BenchBasicMemInfo()
	r.SetFromBenchmarkResult(m)
	fmt.Printf("JSON Marshal: BasicMemInfo:\t%s\n", r.String())
	r.SetFromBenchmarkResult(u)
	fmt.Printf("JSONUnmarshal: BasicMemInfo:\t%s\n", r.String())
	// PBv3
	m, u = pb.BenchBasicMemInfo()
	r.SetFromBenchmarkResult(m)
	fmt.Printf("PB v3 Marshal: BasicMemInfo:\t%s\n", r.String())
	r.SetFromBenchmarkResult(u)
	fmt.Printf("PB v3 Unmarshal: BasicMemInfo:\t%s\n", r.String())

	// MemInfo
	fmt.Println("")
	// Flatbuffers
	m, u = fb.BenchMemInfo()
	r.SetFromBenchmarkResult(m)
	fmt.Printf("Flatbuffers Serialize: MemInfo:\t%s\n", r.String())
	r.SetFromBenchmarkResult(u)
	fmt.Printf("Flatbuffers Deserialize: MemInfo:\t%s\n", r.String())
	// JSON
	m, u = jsn.BenchMemInfo()
	r.SetFromBenchmarkResult(m)
	fmt.Printf("JSON Marshal: Meminfo:\t%s\n", r.String())
	r.SetFromBenchmarkResult(u)
	fmt.Printf("JSONUnmarshal: Meminfo:\t%s\n", r.String())
	// PBv3
	m, u = pb.BenchMemInfo()
	r.SetFromBenchmarkResult(m)
	fmt.Printf("PB v3 Marshal: MemInfo:\t%s\n", r.String())
	r.SetFromBenchmarkResult(u)
	fmt.Printf("PB v3 Unmarshal: MemInfo:\t%s\n", r.String())

	// Message Data
	dataLen := []int{16, 256, 1024, 4096}
	for _, v := range dataLen {
		fmt.Println("")
		shared.GenMessageData(v, shared.Len)
		fmt.Printf("Data: %d bytes\n", v)
		// Flatbuffers
		m, u = fb.BenchMessage()
		r.SetFromBenchmarkResult(m)
		fmt.Printf("Flatbuffers Serialize: Message:\t%s\n", r.String())
		r.SetFromBenchmarkResult(u)
		fmt.Printf("Flatbuffers Deserialize: Message:\t%s\n", r.String())
		// JSON Marshal
		m, u = jsn.BenchMessage()
		r.SetFromBenchmarkResult(m)
		fmt.Printf("JSON Marshal: Message:\t%s\n", r.String())
		r.SetFromBenchmarkResult(u)
		fmt.Printf("JSONUnmarshal: Message:\t%s\n", r.String())
		// PBv3
		m, u = pb.BenchMessage()
		r.SetFromBenchmarkResult(m)
		fmt.Printf("PB v3 Marshal: Message:\t%s\n", r.String())
		r.SetFromBenchmarkResult(u)
		fmt.Printf("PB v3 Unmarshal: Message:\t%s\n", r.String())

	}

	// Reddit Account
	fmt.Println("")
	// Flatbuffers
	m, u = fb.BenchRedditAccount()
	r.SetFromBenchmarkResult(m)
	fmt.Printf("Flatbuffers Serialize: RedditAccount:\t%s\n", r.String())
	r.SetFromBenchmarkResult(u)
	fmt.Printf("Flatbuffers Deserialize: RedditAccount:\t%s\n", r.String())
	// JSON Marshal
	m, u = jsn.BenchRedditAccount()
	r.SetFromBenchmarkResult(m)
	fmt.Printf("JSON Marshal: RedditAccount:\t%s\n", r.String())
	r.SetFromBenchmarkResult(u)
	fmt.Printf("JSONUnmarshal: RedditAccount:\t%s\n", r.String())
	// PB v3
	m, u = pb.BenchRedditAccount()
	r.SetFromBenchmarkResult(m)
	fmt.Printf("PBv3 Marshal:  RedditAccount:\t%s\n", r.String())
	r.SetFromBenchmarkResult(u)
	fmt.Printf("PBv3 Unmarshal:  RedditAccount:\t%s\n", r.String())
}
