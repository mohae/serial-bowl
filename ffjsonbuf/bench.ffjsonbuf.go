package ffjsonbuf

import (
	"fmt"
	"testing"

	"github.com/mohae/benchutil"
	"github.com/mohae/serial-bowl/shared"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

var buf fflib.Buffer

var (
	basicMemInfo  [][]byte
	memInfo       [][]byte
	message       [][]byte
	redditAccount [][]byte

	basicMemInfos  []BasicMemInfo
	memInfos       []MemInfo
	messages       []Message
	redditAccounts []RedditAccount
)

// BenchBasicMemInfo runs the BasicMemInfo benches for Marshal/Unmarshal.
func BenchBasicMemInfo() []benchutil.Bench {
	var results []benchutil.Bench
	basicMemInfo = make([][]byte, shared.Len)
	basicMemInfos = PrepareBasicMemInfoData(shared.BasicMemInfoData)

	bench := benchutil.NewBench(shared.FFJSONBuffer.String())
	bench.Iterations = shared.Len
	bench.Group = shared.BasicMemInfo.String()
	bench.Desc = shared.Marshal.String()
	basicMemInfo = make([][]byte, shared.Len)
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
		for j := 0; j < len(shared.BasicMemInfoData); j++ {
			buf.Reset()
			_ = basicMemInfos[j].MarshalJSONBuf(&buf)
			basicMemInfo[j] = buf.Bytes()
		}
	}
}

func basicMemInfoUnmarshal(b *testing.B) {
	var tmp BasicMemInfo
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(basicMemInfo); j++ {
			_ = tmp.UnmarshalJSON(basicMemInfo[j])
		}
	}
	_ = tmp
}

// BenchMemInfo runs the MemInfo benches for Marshal/Unmarshal.
func BenchMemInfo() []benchutil.Bench {
	var results []benchutil.Bench
	memInfo = make([][]byte, shared.Len)
	memInfos = PrepareMemInfoData(shared.MemInfoData)

	bench := benchutil.NewBench(shared.FFJSONBuffer.String())
	bench.Iterations = shared.Len
	bench.Group = shared.MemInfo.String()
	bench.Desc = shared.Marshal.String()
	memInfo = make([][]byte, shared.Len)
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
			buf.Reset()
			_ = memInfos[j].MarshalJSONBuf(&buf)
			memInfo[j] = buf.Bytes()
		}
	}
}

func memInfoUnmarshal(b *testing.B) {
	var tmp MemInfo
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(memInfo); j++ {
			_ = tmp.UnmarshalJSON(memInfo[j])
		}
	}
	_ = tmp
}

// BenchMessage runs the MemInfo benches for Marshal/Unmarshal.
func BenchMessage(l int) []benchutil.Bench {
	var results []benchutil.Bench
	message = make([][]byte, shared.Len)
	messages = PrepareMessageData(shared.MessageData)

	bench := benchutil.NewBench(shared.FFJSONBuffer.String())
	bench.Iterations = shared.Len
	bench.Group = fmt.Sprintf("%s: %d", shared.Message.String(), l)
	bench.Desc = shared.Marshal.String()
	message = make([][]byte, shared.Len)
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
			buf.Reset()
			_ = messages[j].MarshalJSONBuf(&buf)
			message[j] = buf.Bytes()
		}
	}
}

func messageUnmarshal(b *testing.B) {
	var tmp Message
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			_ = tmp.UnmarshalJSON(message[j])
		}
	}
	_ = tmp
}

// BenchRedditAccount runs the MemInfo benches for Marshal/Unmarshal.
func BenchRedditAccount() []benchutil.Bench {
	var results []benchutil.Bench
	redditAccount = make([][]byte, shared.Len)
	redditAccounts = PrepareRedditAccountData(shared.RedditAccountData)

	bench := benchutil.NewBench(shared.FFJSONBuffer.String())
	bench.Iterations = shared.Len
	bench.Group = shared.RedditAccount.String()
	bench.Desc = shared.Marshal.String()
	redditAccount = make([][]byte, shared.Len)
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
			buf.Reset()
			_ = redditAccounts[j].MarshalJSONBuf(&buf)
			redditAccount[j] = buf.Bytes()
		}
	}
}

func redditAccountUnmarshal(b *testing.B) {
	var tmp RedditAccount
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			_ = tmp.UnmarshalJSON(redditAccount[j])
		}
	}
	_ = tmp
}

// PrepareBasicMemInfoData generates the ffjson version of the data.
func PrepareBasicMemInfoData(data []shared.ShBasicMemInfo) []BasicMemInfo {
	tmp := make([]BasicMemInfo, len(data))
	for i := 0; i < len(data); i++ {
		tmp[i] = BasicMemInfo{
			MemTotal:     data[i].MemTotal,
			MemFree:      data[i].MemFree,
			MemAvailable: data[i].MemAvailable,
			Buffers:      data[i].Buffers,
			Cached:       data[i].Cached,
			SwapCached:   data[i].SwapCached,
			SwapTotal:    data[i].SwapTotal,
			SwapFree:     data[i].SwapFree,
		}
	}
	return tmp
}

// PrepareMemInfoData generates the ffjson version of the data.
func PrepareMemInfoData(data []shared.ShMemInfo) []MemInfo {
	tmp := make([]MemInfo, len(data))
	for i := 0; i < len(data); i++ {
		tmp[i] = MemInfo{
			MemTotal:          data[i].MemTotal,
			MemFree:           data[i].MemFree,
			MemAvailable:      data[i].MemAvailable,
			Buffers:           data[i].Buffers,
			Cached:            data[i].Cached,
			SwapCached:        data[i].SwapCached,
			Active:            data[i].Active,
			Inactive:          data[i].Inactive,
			ActiveAnon:        data[i].ActiveAnon,
			InactiveAnon:      data[i].InactiveAnon,
			ActiveFile:        data[i].ActiveFile,
			InactiveFile:      data[i].InactiveFile,
			Unevictable:       data[i].Unevictable,
			Mlocked:           data[i].Mlocked,
			SwapTotal:         data[i].SwapTotal,
			SwapFree:          data[i].SwapFree,
			Dirty:             data[i].Dirty,
			Writeback:         data[i].Writeback,
			AnonPages:         data[i].AnonPages,
			Mapped:            data[i].Mapped,
			Shmem:             data[i].Shmem,
			Slab:              data[i].Slab,
			SReclaimable:      data[i].SReclaimable,
			SUnreclaim:        data[i].SUnreclaim,
			KernelStack:       data[i].KernelStack,
			NFSUnstable:       data[i].NFSUnstable,
			Bounce:            data[i].Bounce,
			WritebackTmp:      data[i].WritebackTmp,
			CommitLimit:       data[i].CommitLimit,
			VmallocTotal:      data[i].VmallocTotal,
			VmallocUsed:       data[i].VmallocUsed,
			VmallocChunk:      data[i].VmallocChunk,
			HardwareCorrupted: data[i].HardwareCorrupted,
			AnonHugePages:     data[i].AnonHugePages,
			HugePagesTotal:    data[i].HugePagesTotal,
			HugePagesFree:     data[i].HugePagesFree,
			HugePagesRsvd:     data[i].HugePagesRsvd,
			Hugepagesize:      data[i].Hugepagesize,
			DirectMap4k:       data[i].DirectMap4k,
			DirectMap2M:       data[i].DirectMap2M,
		}
	}
	return tmp
}

// PrepareMessageData generates the ffjson version of the data.
func PrepareMessageData(data []shared.ShMessage) []Message {
	tmp := make([]Message, len(data))
	for i := 0; i < len(data); i++ {
		tmp[i] = Message{
			ID:     data[i].ID,
			DestID: data[i].DestID,
			Type:   data[i].Type,
			Kind:   data[i].Kind,
			Data:   data[i].Data,
		}
	}
	return tmp
}

// PrepareRedditAccountData generates the ffjson version of the data.
func PrepareRedditAccountData(data []shared.ShRedditAccount) []RedditAccount {
	tmp := make([]RedditAccount, len(data))
	for i := 0; i < len(data); i++ {
		tmp[i] = RedditAccount{
			ID:   data[i].ID,
			Name: data[i].Name,
			Kind: data[i].Kind,
			Data: AccountData{
				CommentKarma: data[i].Data.CommentKarma,
				HasMail:      data[i].Data.HasMail,
				HasModMail:   data[i].Data.HasModMail,
				ID:           data[i].Data.ID,
				InboxCount:   data[i].Data.InboxCount,
				IsFriend:     data[i].Data.IsFriend,
				IsGold:       data[i].Data.IsGold,
				LinkKarma:    data[i].Data.LinkKarma,
				ModHash:      data[i].Data.ModHash,
				Name:         data[i].Data.Name,
				Over18:       data[i].Data.Over18,
			},
		}
	}
	return tmp
}