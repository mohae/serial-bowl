package ffjsonbuf

import (
	"fmt"
	"testing"

	"github.com/mohae/serial-bowl/shared"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

var buf fflib.Buffer

var (
	basicMemInfo  [][]byte
	memInfo       [][]byte
	message       [][]byte
	redditAccount [][]byte
)

var (
	shBasicMemInfos  []ShBasicMemInfo
	shMemInfos       []ShMemInfo
	shMessages       []ShMessage
	shRedditAccounts []ShRedditAccount
)

// BenchBasicMemInfo runs the BasicMemInfo benches for Marshal/Unmarshal.
func BenchBasicMemInfo() shared.Bench {
	shBasicMemInfos = PrepareBasicMemInfoData(shared.BasicMemInfoData)
	bench := shared.Bench{Proto: shared.FFJSONBuffer, StructString: shared.BasicMemInfo.String(), Results: map[shared.Op]shared.Result{}}
	basicMemInfo = make([][]byte, shared.Len)
	bench.Results[shared.Marshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(basicMemInfoMarshal))
	bench.Results[shared.Unmarshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(basicMemInfoUnmarshal))
	basicMemInfo = nil
	return bench
}

func basicMemInfoMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(shared.BasicMemInfoData); j++ {
			buf.Reset()
			_ = shBasicMemInfos[j].MarshalJSONBuf(&buf)
			basicMemInfo[j] = buf.Bytes()
		}
	}
}

func basicMemInfoUnmarshal(b *testing.B) {
	var tmp ShBasicMemInfo
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(basicMemInfo); j++ {
			_ = tmp.UnmarshalJSON(basicMemInfo[j])
		}
	}
	_ = tmp
}

// BenchMemInfo runs the MemInfo benches for Marshal/Unmarshal.
func BenchMemInfo() shared.Bench {
	bench := shared.Bench{Proto: shared.FFJSONBuffer, StructString: shared.MemInfo.String(), Results: map[shared.Op]shared.Result{}}
	memInfo = make([][]byte, shared.Len)
	shMemInfos = PrepareMemInfoData(shared.MemInfoData)
	bench.Results[shared.Marshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(memInfoMarshal))
	bench.Results[shared.Unmarshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(memInfoUnmarshal))
	memInfo = nil
	return bench
}

func memInfoMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			buf.Reset()
			_ = shMemInfos[j].MarshalJSONBuf(&buf)
			memInfo[j] = buf.Bytes()
		}
	}
}

func memInfoUnmarshal(b *testing.B) {
	var tmp ShMemInfo
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(memInfo); j++ {
			_ = tmp.UnmarshalJSON(memInfo[j])
		}
	}
	_ = tmp
}

// BenchMessage runs the MemInfo benches for Marshal/Unmarshal.
func BenchMessage(l int) shared.Bench {
	bench := shared.Bench{Proto: shared.FFJSONBuffer, StructString: fmt.Sprintf("%s %dB", shared.Message.String(), l), Results: map[shared.Op]shared.Result{}}
	message = make([][]byte, shared.Len)
	shMessages = PrepareMessageData(shared.MessageData)
	bench.Results[shared.Marshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(messageMarshal))
	bench.Results[shared.Unmarshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(messageUnmarshal))
	message = nil
	return bench
}

func messageMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			buf.Reset()
			_ = shMessages[j].MarshalJSONBuf(&buf)
			message[j] = buf.Bytes()
		}
	}
}

func messageUnmarshal(b *testing.B) {
	var tmp ShMessage
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			_ = tmp.UnmarshalJSON(message[j])
		}
	}
	_ = tmp
}

// BenchRedditAccount runs the MemInfo benches for Marshal/Unmarshal.
func BenchRedditAccount() shared.Bench {
	bench := shared.Bench{Proto: shared.FFJSONBuffer, StructString: shared.RedditAccount.String(), Results: map[shared.Op]shared.Result{}}
	redditAccount = make([][]byte, shared.Len)
	shRedditAccounts = PrepareRedditAccountData(shared.RedditAccountData)
	bench.Results[shared.Marshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(redditAccountMarshal))
	bench.Results[shared.Unmarshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(redditAccountUnmarshal))
	redditAccount = nil
	return bench
}

func redditAccountMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			buf.Reset()
			_ = shRedditAccounts[j].MarshalJSONBuf(&buf)
			redditAccount[j] = buf.Bytes()
		}
	}
}

func redditAccountUnmarshal(b *testing.B) {
	var tmp ShRedditAccount
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			_ = tmp.UnmarshalJSON(redditAccount[j])
		}
	}
	_ = tmp
}

// PrepareBasicMemInfoData generates the protobuf version of the data.
func PrepareBasicMemInfoData(data []shared.ShBasicMemInfo) []ShBasicMemInfo {
	tmp := make([]ShBasicMemInfo, len(data))
	for i := 0; i < len(data); i++ {
		tmp[i] = ShBasicMemInfo{
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

// PrepareMemInfoData generates the protobuf version of the data.
func PrepareMemInfoData(data []shared.ShMemInfo) []ShMemInfo {
	tmp := make([]ShMemInfo, len(data))
	for i := 0; i < len(data); i++ {
		tmp[i] = ShMemInfo{
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

// PrepareMessageData generates the protobuf version of the data.
func PrepareMessageData(data []shared.ShMessage) []ShMessage {
	tmp := make([]ShMessage, len(data))
	for i := 0; i < len(data); i++ {
		tmp[i] = ShMessage{
			ID:     data[i].ID,
			DestID: data[i].DestID,
			Type:   data[i].Type,
			Kind:   data[i].Kind,
			Data:   data[i].Data,
		}
	}
	return tmp
}

// PrepareRedditAccountData generates the protobuf version of the data.
func PrepareRedditAccountData(data []shared.ShRedditAccount) []ShRedditAccount {
	tmp := make([]ShRedditAccount, len(data))
	for i := 0; i < len(data); i++ {
		tmp[i] = ShRedditAccount{
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
