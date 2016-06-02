package capnp

import (
	"fmt"
	"os"
	"testing"

	"github.com/mohae/benchutil"
	"github.com/mohae/serial-bowl/shared"
	capnp "zombiezen.com/go/capnproto2"
)

var (
	basicMemInfo  [][]byte
	memInfo       [][]byte
	message       [][]byte
	redditAccount [][]byte
	cpuInfo       [][]byte
)

func newBench(s string) benchutil.Bench {
	bench := benchutil.NewBench("zombiezen.com/go/capnproto2")
	bench.Iterations = shared.Len
	bench.Group = s
	bench.SubGroup = shared.CapnProto2.String()
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
func BenchRedditAccount() []benchutil.Bench {
	var results []benchutil.Bench
	redditAccount = make([][]byte, shared.Len)

	bench := newBench(shared.RedditAccount.String())
	bench.Iterations = shared.Len
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

// BenchCPUInfo runs the CPUInfo benches for Marshal/Unmarshal.
func BenchCPUInfo(n int) []benchutil.Bench {
	var results []benchutil.Bench
	cpuInfo = make([][]byte, shared.Len)
	shared.GenCPUInfoData(n, shared.Len)

	bench := newBench(fmt.Sprintf("%s: %d", shared.CPUInfo, n))
	bench.Desc = shared.Marshal.String()
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(cpuInfoMarshal))
	results = append(results, bench)
	bench.Desc = shared.Unmarshal.String()
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(cpuInfoUnmarshal))
	results = append(results, bench)
	message = nil
	return results
}

func cpuInfoMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			msg, seg, _ := capnp.NewMessage(capnp.SingleSegment(nil))
			info, _ := NewRootCPUInfo(seg)
			cpus, _ := NewCPU_List(seg, int32(len(shared.CPUInfoData[j].CPUs)))
			for k := 0; k < len(shared.CPUInfoData[j].CPUs); k++ {
				cpu, _ := NewCPU(seg)
				cpu.SetProcessor(shared.CPUInfoData[j].CPUs[k].Processor)
				cpu.SetVendorID(shared.CPUInfoData[j].CPUs[k].VendorID)
				cpu.SetCPUFamily(shared.CPUInfoData[j].CPUs[k].CPUFamily)
				cpu.SetModel(shared.CPUInfoData[j].CPUs[k].Model)
				cpu.SetStepping(shared.CPUInfoData[j].CPUs[k].Stepping)
				cpu.SetMicrocode(shared.CPUInfoData[j].CPUs[k].Microcode)
				cpu.SetCPUMHz(shared.CPUInfoData[j].CPUs[k].CPUMHz)
				cpu.SetCacheSize(shared.CPUInfoData[j].CPUs[k].CacheSize)
				cpu.SetPhysicalID(shared.CPUInfoData[j].CPUs[k].PhysicalID)
				cpu.SetSiblings(shared.CPUInfoData[j].CPUs[k].Siblings)
				cpu.SetCoreID(shared.CPUInfoData[j].CPUs[k].CoreID)
				cpu.SetCPUCores(shared.CPUInfoData[j].CPUs[k].CPUCores)
				cpu.SetApicID(shared.CPUInfoData[j].CPUs[k].ApicID)
				cpu.SetInitialApicID(shared.CPUInfoData[j].CPUs[k].InitialApicID)
				cpu.SetFPU(shared.CPUInfoData[j].CPUs[k].FPU)
				cpu.SetFPUException(shared.CPUInfoData[j].CPUs[k].FPUException)
				cpu.SetCPUIDLevel(shared.CPUInfoData[j].CPUs[k].CPUIDLevel)
				cpu.SetWP(shared.CPUInfoData[j].CPUs[k].WP)
				cpu.SetFlags(shared.CPUInfoData[j].CPUs[k].Flags)
				cpu.SetBogoMIPS(shared.CPUInfoData[j].CPUs[k].BogoMIPS)
				cpu.SetCLFlushSize(shared.CPUInfoData[j].CPUs[k].CLFlushSize)
				cpu.SetCacheAlignment(shared.CPUInfoData[j].CPUs[k].CacheAlignment)
				cpu.SetAddressSizes(shared.CPUInfoData[j].CPUs[k].AddressSizes)
				cpu.SetPowerManagement(shared.CPUInfoData[j].CPUs[k].PowerManagement)
				_ = cpus.Set(k, cpu)
			}
			_ = info.SetCpus(cpus)
			cpuInfo[j], _ = msg.Marshal()
		}
	}
}

func cpuInfoUnmarshal(b *testing.B) {
	var tmp CPUInfo
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			inf, _ := capnp.Unmarshal(cpuInfo[j])
			tmp, _ = ReadRootCPUInfo(inf)
		}
	}
	_ = tmp
}
