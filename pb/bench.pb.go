package pb

import (
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/mohae/serial-bowl/shared"
)

// marshal results holder.
var redditAccount [][]byte
var memInfo [][]byte
var basicMemInfo [][]byte
var message [][]byte

// BenchBasicMemInfo runs the MemInfo benches for Marshal/Unmarshal.
func BenchBasicMemInfo() shared.Bench {
	bench := shared.Bench{Proto: shared.ProtobufV3, StructString: shared.BasicMemInfo.String(), Results: map[shared.Op]shared.Result{}}
	basicMemInfo = make([][]byte, shared.Len)
	bench.Results[shared.Marshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(basicMemInfoMarshal))
	bench.Results[shared.Unmarshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(basicMemInfoUnmarshal))
	basicMemInfo = nil
	return bench
}

func basicMemInfoMarshal(b *testing.B) {
	b.StopTimer()
	tmp := PrepareBasicMemInfoData(shared.BasicMemInfoData)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			basicMemInfo[j], _ = proto.Marshal(&tmp[j])
		}
	}
}

func basicMemInfoUnmarshal(b *testing.B) {
	var tmp BasicMemInfo
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			proto.Unmarshal(basicMemInfo[j], &tmp)
		}
	}
	_ = tmp
}

// BenchMemInfo runs the MemInfo benches for Marshal/Unmarshal.
func BenchMemInfo() shared.Bench {
	bench := shared.Bench{Proto: shared.ProtobufV3, StructString: shared.MemInfo.String(), Results: map[shared.Op]shared.Result{}}
	memInfo = make([][]byte, shared.Len)
	bench.Results[shared.Marshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(memInfoMarshal))
	bench.Results[shared.Unmarshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(memInfoUnmarshal))
	memInfo = nil
	return bench
}

func memInfoMarshal(b *testing.B) {
	b.StopTimer()
	tmp := PrepareMemInfoData(shared.MemInfoData)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			memInfo[j], _ = proto.Marshal(&tmp[j])
		}
	}
}

func memInfoUnmarshal(b *testing.B) {
	var tmp MemInfo
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			proto.Unmarshal(memInfo[j], &tmp)
		}
	}
	_ = tmp
}

// BenchMessage runs the Message benches for Marshal/Unmarshal.
func BenchMessage(l int) shared.Bench {
	bench := shared.Bench{Proto: shared.ProtobufV3, StructString: fmt.Sprintf("%s %dB", shared.Message.String(), l), Results: map[shared.Op]shared.Result{}}
	message = make([][]byte, shared.Len)
	bench.Results[shared.Marshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(messageMarshal))
	bench.Results[shared.Unmarshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(messageUnmarshal))
	message = nil
	return bench
}

func messageMarshal(b *testing.B) {
	b.StopTimer()
	tmp := PrepareMessageData(shared.MessageData)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			message[j], _ = proto.Marshal(&tmp[j])
		}
	}
}

func messageUnmarshal(b *testing.B) {
	var tmp Message
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			proto.Unmarshal(message[j], &tmp)
		}
	}
	_ = tmp
}

// BenchmarkRedditAccount runs the RedditAccount benches for Marshal/Unmarshal.
func BenchRedditAccount() shared.Bench {
	bench := shared.Bench{Proto: shared.ProtobufV3, StructString: shared.RedditAccount.String(), Results: map[shared.Op]shared.Result{}}
	redditAccount = make([][]byte, shared.Len)
	bench.Results[shared.Marshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(redditAccountMarshal))
	bench.Results[shared.Unmarshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(redditAccountUnmarshal))
	redditAccount = nil
	return bench
}

func redditAccountMarshal(b *testing.B) {
	b.StopTimer()
	tmp := PrepareRedditAccountData(shared.RedditAccountData)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			redditAccount[j], _ = proto.Marshal(&tmp[j])
		}
	}
}

func redditAccountUnmarshal(b *testing.B) {
	var tmp RedditAccount
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			proto.Unmarshal(redditAccount[j], &tmp)
		}
	}
	_ = tmp
}

// PrepareBasicMemInfoData generates the protobuf version of the data.
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

// PrepareMemInfoData generates the protobuf version of the data.
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
			DirectMap4K:       data[i].DirectMap4k,
			DirectMap2M:       data[i].DirectMap2M,
		}
	}
	return tmp
}

// PrepareMessageData generates the protobuf version of the data.
func PrepareMessageData(data []shared.ShMessage) []Message {
	tmp := make([]Message, len(data))
	for i := 0; i < len(data); i++ {
		tmp[i] = Message{
			ID:     data[i].ID,
			DestID: data[i].DestID,
			Type:   int32(data[i].Type),
			Kind:   int32(data[i].Kind),
			Data:   data[i].Data,
		}
	}
	return tmp
}

// PrepareRedditAccountData generates the protobuf version of the data.
func PrepareRedditAccountData(data []shared.ShRedditAccount) []RedditAccount {
	tmp := make([]RedditAccount, len(data))
	for i := 0; i < len(data); i++ {
		tmp[i] = RedditAccount{
			ID:   data[i].ID,
			Name: data[i].Name,
			Kind: data[i].Kind,
			Data: &AccountData{
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
