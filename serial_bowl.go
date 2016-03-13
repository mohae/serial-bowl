package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/mohae/serial-bowl/jsn"
)

const alphanum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const maxInt = 1<<63 - 1
const Len = 1000

// Data for the tests
var MemData []jsn.MemInfo
var BasicMemData []jsn.BasicMemInfo
var AccountData []jsn.RedditAccount

type Result struct {
	Ops      int64
	NsOp     int64
	MBSec    float64
	BytesOp  int64
	AllocsOp int64
}

func (r *Result) SetFromBenchmarkResult(br testing.BenchmarkResult) {
	r.Ops = int64(br.N) * int64(Len)
	r.NsOp = br.T.Nanoseconds() / r.Ops
	r.MBSec = (float64(br.Bytes) * float64(r.Ops) / 1e6) / br.T.Seconds()
	r.BytesOp = int64(br.MemBytes) / r.Ops
	r.AllocsOp = int64(br.MemAllocs) / r.Ops
}

func (r Result) OpsString() string {
	return fmt.Sprintf("%d ops", r.Ops)
}

func (r Result) NsOpString() string {
	return fmt.Sprintf("%d ns/Op", r.NsOp)
}

func (r Result) MBSecString() string {
	return fmt.Sprintf("%9.2f MB/s", r.MBSec)
}

func (r Result) BytesOpString() string {
	return fmt.Sprintf("%d bytes/Op", r.BytesOp)
}

func (r Result) AllocsOpString() string {
	return fmt.Sprintf("%d allocs/Op", r.AllocsOp)
}

func (r Result) String() string {
	return fmt.Sprintf("%s%s%s%s", column(15, r.OpsString()), column(15, r.NsOpString()), column(18, r.BytesOpString()), column(16, r.AllocsOpString()))
}

func column(w int, s string) string {
	pad := w - len(s)
	padding := make([]byte, pad)
	for i := 0; i < pad; i++ {
		padding[i] = 0x20
	}
	return fmt.Sprintf("%s%s", string(padding), s)
}

func init() {
	rand.Seed(time.Now().UnixNano())
	GenMemData(Len)
	GenBasicMemData(Len)
	GenAccountData(Len)
}

// RandString returns a randomly generated string of length n.
func RandString(l int) string {
	str := make([]byte, l)
	for i := 0; i < l; i++ {
		str[i] = alphanum[rand.Intn(len(alphanum))]
	}
	return string(str)
}

func RandBool() bool {
	if rand.Intn(100)%2 == 0 {
		return false
	}
	return true
}

func GenAccountData(n int) {
	AccountData = make([]jsn.RedditAccount, n)
	for i := 0; i < n; i++ {
		AccountData[i] = jsn.RedditAccount{
			ID:   RandString(20),
			Name: RandString(rand.Intn(30)),
			Kind: RandString(5),
			Data: jsn.AccountData{
				CommentKarma: rand.Intn(maxInt),
				HasMail:      RandBool(),
				HasModMail:   RandBool(),
				ID:           RandString(20),
				InboxCount:   rand.Intn(maxInt),
				IsFriend:     RandBool(),
				IsGold:       RandBool(),
				LinkKarma:    rand.Intn(maxInt),
				ModHash:      RandString(88),
				Name:         RandString(rand.Intn(30)),
				Over18:       RandBool(),
			},
		}
	}
}

func GenMemData(n int) {
	MemData = make([]jsn.MemInfo, n)
	for i := 0; i < n; i++ {
		MemData[i] = jsn.MemInfo{
			MemTotal:          rand.Intn(maxInt),
			MemFree:           rand.Intn(maxInt),
			MemAvailable:      rand.Intn(maxInt),
			Buffers:           rand.Intn(maxInt),
			Cached:            rand.Intn(maxInt),
			SwapCached:        rand.Intn(maxInt),
			Active:            rand.Intn(maxInt),
			Inactive:          rand.Intn(maxInt),
			ActiveAnon:        rand.Intn(maxInt),
			InactiveAnon:      rand.Intn(maxInt),
			ActiveFile:        rand.Intn(maxInt),
			InactiveFile:      rand.Intn(maxInt),
			Unevictable:       rand.Intn(maxInt),
			Mlocked:           rand.Intn(maxInt),
			SwapTotal:         rand.Intn(maxInt),
			SwapFree:          rand.Intn(maxInt),
			Dirty:             rand.Intn(maxInt),
			Writeback:         rand.Intn(maxInt),
			AnonPages:         rand.Intn(maxInt),
			Mapped:            rand.Intn(maxInt),
			Shmem:             rand.Intn(maxInt),
			Slab:              rand.Intn(maxInt),
			SReclaimable:      rand.Intn(maxInt),
			SUnreclaim:        rand.Intn(maxInt),
			KernelStack:       rand.Intn(maxInt),
			NFSUnstable:       rand.Intn(maxInt),
			Bounce:            rand.Intn(maxInt),
			WritebackTmp:      rand.Intn(maxInt),
			CommitLimit:       rand.Intn(maxInt),
			VmallocTotal:      rand.Intn(maxInt),
			VmallocUsed:       rand.Intn(maxInt),
			VmallocChunk:      rand.Intn(maxInt),
			HardwareCorrupted: rand.Intn(maxInt),
			AnonHugePages:     rand.Intn(maxInt),
			HugePagesTotal:    rand.Intn(maxInt),
			HugePagesFree:     rand.Intn(maxInt),
			HugePagesRsvd:     rand.Intn(maxInt),
			Hugepagesize:      rand.Intn(maxInt),
			DirectMap4k:       rand.Intn(maxInt),
			DirectMap2M:       rand.Intn(maxInt),
		}
	}
}

func GenBasicMemData(n int) {
	BasicMemData = make([]jsn.BasicMemInfo, n)
	for i := 0; i < n; i++ {
		BasicMemData[i] = jsn.BasicMemInfo{
			MemTotal:     rand.Intn(maxInt),
			MemFree:      rand.Intn(maxInt),
			MemAvailable: rand.Intn(maxInt),
			Buffers:      rand.Intn(maxInt),
			Cached:       rand.Intn(maxInt),
			SwapCached:   rand.Intn(maxInt),
			SwapTotal:    rand.Intn(maxInt),
			SwapFree:     rand.Intn(maxInt),
		}
	}
}

func main() {
	var r Result
	// Reddit Account
	// Flatbuffers Serialize
	br := testing.Benchmark(BenchRedditAccountFBSerialize)
	r.SetFromBenchmarkResult(br)
	fmt.Printf("RedditAccountSerializeFB:\t%s\n", r.String())
	// JSON Marshal
	br = testing.Benchmark(BenchRedditAccountJSONMarshal)
	r.SetFromBenchmarkResult(br)
	fmt.Printf("RedditAccountJSONMarshal:\t%s\n", r.String())
	fmt.Println("")
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
	fmt.Println("")
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
	fmt.Println("")
	// Flatbuffers Deserialize
	br = testing.Benchmark(BenchMemInfoFBDeserialize)
	r.SetFromBenchmarkResult(br)
	fmt.Printf("MemInfoDeserializeFB:\t\t%s\n", r.String())
	// JSON Unmarshal
	br = testing.Benchmark(BenchMemInfoJSONUnmarshal)
	r.SetFromBenchmarkResult(br)
	fmt.Printf("MemInfoJSONUnmarshal:\t\t%s\n", r.String())
}
