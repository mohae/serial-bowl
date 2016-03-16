package capnp

import (
	"fmt"
	"os"
	"testing"

	"github.com/mohae/serial-bowl/shared"
	capnp "zombiezen.com/go/capnproto2"
)

var basicMemInfo [][]byte
var memInfo [][]byte
var message [][]byte
var redditAccount [][]byte

// BenchBasicMemInfo runs the BasicMemInfo benches for Marshal/Unmarshal.
func BenchBasicMemInfo() shared.Bench {
	bench := shared.Bench{Proto: shared.CapnProto2, StructString: shared.BasicMemInfo.String(), Results: map[shared.Op]shared.Result{}}
	basicMemInfo = make([][]byte, shared.Len)
	bench.Results[shared.Marshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(basicMemInfoMarshal))
	bench.Results[shared.Unmarshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(basicMemInfoUnmarshal))
	basicMemInfo = nil
	return bench
}

func basicMemInfoMarshal(b *testing.B) {
	arena := capnp.SingleSegment(nil)
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			msg, seg, _ := capnp.NewMessage(arena)
			m, _ := NewRootBasicMemInfo(seg)
			m.SetMemTotal(shared.BasicMemInfoData[j].MemTotal)
			m.SetMemFree(shared.BasicMemInfoData[j].MemFree)
			m.SetMemAvailable(shared.BasicMemInfoData[j].MemAvailable)
			m.SetBuffers(shared.BasicMemInfoData[j].Buffers)
			m.SetCached(shared.BasicMemInfoData[j].Cached)
			m.SetSwapCached(shared.BasicMemInfoData[j].SwapCached)
			m.SetSwapTotal(shared.BasicMemInfoData[j].SwapTotal)
			m.SetSwapFree(shared.BasicMemInfoData[j].SwapFree)
			basicMemInfo[j], _ = msg.Marshal()
		}
	}
}

func basicMemInfoUnmarshal(b *testing.B) {
	var tmp BasicMemInfo
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			msg, _ := capnp.Unmarshal(basicMemInfo[j])
			tmp, _ = ReadRootBasicMemInfo(msg)
		}
	}
	_ = tmp
}

// BenchMemInfo runs the BasicMemInfo benches for Marshal/Unmarshal.
func BenchMemInfo() shared.Bench {
	bench := shared.Bench{Proto: shared.CapnProto2, StructString: shared.MemInfo.String(), Results: map[shared.Op]shared.Result{}}
	memInfo = make([][]byte, shared.Len)
	bench.Results[shared.Marshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(memInfoMarshal))
	bench.Results[shared.Unmarshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(memInfoUnmarshal))
	basicMemInfo = nil
	return bench
}

func memInfoMarshal(b *testing.B) {
	arena := capnp.SingleSegment(nil)
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			msg, seg, _ := capnp.NewMessage(arena)
			m, _ := NewRootMemInfo(seg)
			m.SetMemTotal(shared.MemInfoData[j].MemTotal)
			m.SetMemFree(shared.MemInfoData[j].MemFree)
			m.SetMemAvailable(shared.MemInfoData[j].MemAvailable)
			m.SetBuffers(shared.MemInfoData[j].Buffers)
			m.SetCached(shared.MemInfoData[j].Cached)
			m.SetSwapCached(shared.MemInfoData[j].SwapCached)
			m.SetActive(shared.MemInfoData[j].Active)
			m.SetInactive(shared.MemInfoData[j].Inactive)
			m.SetActiveAnon(shared.MemInfoData[j].ActiveAnon)
			m.SetInactiveAnon(shared.MemInfoData[j].InactiveAnon)
			m.SetActiveFile(shared.MemInfoData[j].ActiveFile)
			m.SetInactiveFile(shared.MemInfoData[j].InactiveFile)
			m.SetMlocked(shared.MemInfoData[j].Mlocked)
			m.SetSwapTotal(shared.MemInfoData[j].SwapTotal)
			m.SetSwapFree(shared.MemInfoData[j].SwapFree)
			m.SetDirty(shared.MemInfoData[j].Dirty)
			m.SetWriteback(shared.MemInfoData[j].Writeback)
			m.SetAnonPages(shared.MemInfoData[j].AnonPages)
			m.SetMapped(shared.MemInfoData[j].Mapped)
			m.SetShmem(shared.MemInfoData[j].Shmem)
			m.SetSlab(shared.MemInfoData[j].Slab)
			m.SetSReclaimable(shared.MemInfoData[j].SReclaimable)
			m.SetSUnreclaim(shared.MemInfoData[j].SUnreclaim)
			m.SetKernelStack(shared.MemInfoData[j].KernelStack)
			m.SetNFSUnstable(shared.MemInfoData[j].NFSUnstable)
			m.SetWritebackTmp(shared.MemInfoData[j].WritebackTmp)
			m.SetCommitLimit(shared.MemInfoData[j].CommitLimit)
			m.SetVmallocTotal(shared.MemInfoData[j].VmallocTotal)
			m.SetVmallocUsed(shared.MemInfoData[j].VmallocUsed)
			m.SetVmallocChunk(shared.MemInfoData[j].VmallocChunk)
			m.SetHardwareCorrupted(shared.MemInfoData[j].HardwareCorrupted)
			m.SetAnonHugePages(shared.MemInfoData[j].AnonHugePages)
			m.SetHugePagesTotal(shared.MemInfoData[j].HugePagesTotal)
			m.SetHugePagesFree(shared.MemInfoData[j].HugePagesFree)
			m.SetHugePagesRsvd(shared.MemInfoData[j].HugePagesRsvd)
			m.SetDirectMap4k(shared.MemInfoData[j].DirectMap4k)
			m.SetDirectMap2M(shared.MemInfoData[j].DirectMap2M)
			memInfo[j], _ = msg.Marshal()
		}
	}
}

func memInfoUnmarshal(b *testing.B) {
	var tmp MemInfo
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			msg, _ := capnp.Unmarshal(memInfo[j])
			tmp, _ = ReadRootMemInfo(msg)
		}
	}
	_ = tmp
}

// BenchMessage runs the Message benches for Marshal/Unmarshal.
func BenchMessage(l int) shared.Bench {
	bench := shared.Bench{Proto: shared.CapnProto2, StructString: fmt.Sprintf("%s %dB", shared.Message.String(), l), Results: map[shared.Op]shared.Result{}}
	message = make([][]byte, shared.Len)
	bench.Results[shared.Marshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(messageMarshal))
	bench.Results[shared.Unmarshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(messageUnmarshal))
	message = nil
	return bench
}

func messageMarshal(b *testing.B) {
	arena := capnp.SingleSegment(nil)
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			msg, seg, err := capnp.NewMessage(arena)
			if err != nil {
				fmt.Printf("%d: %s\n", j, err.Error())
				os.Exit(1)
			}
			m, _ := NewRootMessage(seg)
			m.SetID(shared.MessageData[j].ID)
			m.SetDestID(shared.MessageData[j].DestID)
			m.SetType(shared.MessageData[j].Type)
			m.SetKind(shared.MessageData[j].Kind)
			m.SetData(shared.MessageData[j].Data)
			message[j], _ = msg.Marshal()
		}
	}
}

func messageUnmarshal(b *testing.B) {
	var tmp Message
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			msg, _ := capnp.Unmarshal(message[j])
			tmp, _ = ReadRootMessage(msg)
		}
	}
	_ = tmp
}

// BenchRedditAccount runs the RedditAccount benches for Marshal/Unmarshal.
func BenchRedditAccount() shared.Bench {
	bench := shared.Bench{Proto: shared.CapnProto2, StructString: shared.RedditAccount.String(), Results: map[shared.Op]shared.Result{}}
	redditAccount = make([][]byte, shared.Len)
	bench.Results[shared.Marshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(redditAccountMarshal))
	bench.Results[shared.Unmarshal] = shared.ResultFromBenchmarkResult(testing.Benchmark(redditAccountUnmarshal))
	redditAccount = nil
	return bench
}

func redditAccountMarshal(b *testing.B) {
	arena := capnp.SingleSegment(nil)
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			msg, seg, _ := capnp.NewMessage(arena)
			ra, _ := NewRootRedditAccount(seg)
			ra.SetID(shared.RedditAccountData[j].ID)
			ra.SetName(shared.RedditAccountData[j].Name)
			ra.SetKind(shared.RedditAccountData[j].Kind)
			d, _ := ra.NewData()
			d.SetCommentKarma(shared.RedditAccountData[j].Data.CommentKarma)
			d.SetHasMail(shared.RedditAccountData[j].Data.HasMail)
			d.SetHasModMail(shared.RedditAccountData[j].Data.HasModMail)
			d.SetID(shared.RedditAccountData[j].Data.ID)
			d.SetInboxCount(shared.RedditAccountData[j].Data.InboxCount)
			d.SetIsFriend(shared.RedditAccountData[j].Data.IsFriend)
			d.SetIsGold(shared.RedditAccountData[j].Data.IsGold)
			d.SetLinkKarma(shared.RedditAccountData[j].Data.LinkKarma)
			d.SetModHash(shared.RedditAccountData[j].Data.ModHash)
			d.SetName(shared.RedditAccountData[j].Data.Name)
			d.SetOver18(shared.RedditAccountData[j].Data.Over18)
			redditAccount[j], _ = msg.Marshal()
		}
	}
}

func redditAccountUnmarshal(b *testing.B) {
	var tmp RedditAccount
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			msg, _ := capnp.Unmarshal(redditAccount[j])
			tmp, _ = ReadRootRedditAccount(msg)
		}
	}
	_ = tmp
}
