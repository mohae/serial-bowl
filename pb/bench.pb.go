package pb

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/mohae/serial-bowl/jsn"
	"github.com/mohae/serial-bowl/shared"
)

// marshal results holder.
var redditAccount [][]byte
var memInfo [][]byte
var basicMemInfo [][]byte
var message [][]byte

// PrepareBasicMemInfoData generates the protobuf version of the data.
func PrepareBasicMemInfoData(data []jsn.BasicMemInfo) []BasicMemInfo {
	tmp := make([]BasicMemInfo, len(data))
	for i := 0; i < len(data); i++ {
		tmp[i] = BasicMemInfo{
			MemTotal:     int64(data[i].MemTotal),
			MemFree:      int64(data[i].MemFree),
			MemAvailable: int64(data[i].MemAvailable),
			Buffers:      int64(data[i].Buffers),
			Cached:       int64(data[i].Cached),
			SwapCached:   int64(data[i].SwapCached),
			SwapTotal:    int64(data[i].SwapTotal),
			SwapFree:     int64(data[i].SwapFree),
		}
	}
	return tmp
}

// PrepareMemInfoData generates the protobuf version of the data.
func PrepareMemInfoData(data []jsn.MemInfo) []MemInfo {
	tmp := make([]MemInfo, len(data))
	for i := 0; i < len(data); i++ {
		tmp[i] = MemInfo{
			MemTotal:          int64(data[i].MemTotal),
			MemFree:           int64(data[i].MemFree),
			MemAvailable:      int64(data[i].MemAvailable),
			Buffers:           int64(data[i].Buffers),
			Cached:            int64(data[i].Cached),
			SwapCached:        int64(data[i].SwapCached),
			Active:            int64(data[i].Active),
			Inactive:          int64(data[i].Inactive),
			ActiveAnon:        int64(data[i].ActiveAnon),
			InactiveAnon:      int64(data[i].InactiveAnon),
			ActiveFile:        int64(data[i].ActiveFile),
			InactiveFile:      int64(data[i].InactiveFile),
			Unevictable:       int64(data[i].Unevictable),
			Mlocked:           int64(data[i].Mlocked),
			SwapTotal:         int64(data[i].SwapTotal),
			SwapFree:          int64(data[i].SwapFree),
			Dirty:             int64(data[i].Dirty),
			Writeback:         int64(data[i].Writeback),
			AnonPages:         int64(data[i].AnonPages),
			Mapped:            int64(data[i].Mapped),
			Shmem:             int64(data[i].Shmem),
			Slab:              int64(data[i].Slab),
			SReclaimable:      int64(data[i].SReclaimable),
			SUnreclaim:        int64(data[i].SUnreclaim),
			KernelStack:       int64(data[i].KernelStack),
			NFSUnstable:       int64(data[i].NFSUnstable),
			Bounce:            int64(data[i].Bounce),
			WritebackTmp:      int64(data[i].WritebackTmp),
			CommitLimit:       int64(data[i].CommitLimit),
			VmallocTotal:      int64(data[i].VmallocTotal),
			VmallocUsed:       int64(data[i].VmallocUsed),
			VmallocChunk:      int64(data[i].VmallocChunk),
			HardwareCorrupted: int64(data[i].HardwareCorrupted),
			AnonHugePages:     int64(data[i].AnonHugePages),
			HugePagesTotal:    int64(data[i].HugePagesTotal),
			HugePagesFree:     int64(data[i].HugePagesFree),
			HugePagesRsvd:     int64(data[i].HugePagesRsvd),
			Hugepagesize:      int64(data[i].Hugepagesize),
			DirectMap4K:       int64(data[i].DirectMap4k),
			DirectMap2M:       int64(data[i].DirectMap2M),
		}
	}
	return tmp
}

// PrepareMessageData generates the protobuf version of the data.
func PrepareMessageData(data []jsn.Message) []Message {
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
func PrepareRedditAccountData(data []jsn.RedditAccount) []RedditAccount {
	tmp := make([]RedditAccount, len(data))
	for i := 0; i < len(data); i++ {
		tmp[i] = RedditAccount{
			ID:   data[i].ID,
			Name: data[i].Name,
			Kind: data[i].Kind,
			Data: &AccountData{
				CommentKarma: int64(data[i].Data.CommentKarma),
				HasMail:      data[i].Data.HasMail,
				HasModMail:   data[i].Data.HasModMail,
				ID:           data[i].Data.ID,
				InboxCount:   int64(data[i].Data.InboxCount),
				IsFriend:     data[i].Data.IsFriend,
				IsGold:       data[i].Data.IsGold,
				LinkKarma:    int64(data[i].Data.LinkKarma),
				ModHash:      data[i].Data.ModHash,
				Name:         data[i].Data.Name,
				Over18:       data[i].Data.Over18,
			},
		}
	}
	return tmp
}

// BenchBasicMemInfo runs the MemInfo benches for Marshal/Unmarshal.
func BenchBasicMemInfo() (m, u testing.BenchmarkResult) {
	basicMemInfo = make([][]byte, shared.Len)
	m = testing.Benchmark(basicMemInfoMarshal)
	u = testing.Benchmark(basicMemInfoUnmarshal)
	basicMemInfo = nil
	return m, u
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
func BenchMemInfo() (m, u testing.BenchmarkResult) {
	memInfo = make([][]byte, shared.Len)
	m = testing.Benchmark(memInfoMarshal)
	u = testing.Benchmark(memInfoUnmarshal)
	memInfo = nil
	return m, u
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
func BenchMessage() (m, u testing.BenchmarkResult) {
	message = make([][]byte, shared.Len)
	m = testing.Benchmark(messageMarshal)
	u = testing.Benchmark(messageUnmarshal)
	message = nil
	return m, u
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
func BenchRedditAccount() (m, u testing.BenchmarkResult) {
	redditAccount = make([][]byte, shared.Len)
	m = testing.Benchmark(redditAccountMarshal)
	u = testing.Benchmark(redditAccountUnmarshal)
	redditAccount = nil
	return m, u
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
