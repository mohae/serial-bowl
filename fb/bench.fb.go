package fb

import (
	"fmt"
	"testing"

	"github.com/google/flatbuffers/go"
	"github.com/mohae/benchutil"
	"github.com/mohae/serial-bowl/shared"
)

var (
	basicMemInfo  [][]byte
	memInfo       [][]byte
	message       [][]byte
	redditAccount [][]byte
	cpuInfo       [][]byte

	bldr = flatbuffers.NewBuilder(0)
)

func newBench(s string) benchutil.Bench {
	bench := benchutil.NewBench("google/flatbuffers/go")
	bench.Iterations = shared.Len
	bench.Group = s
	bench.SubGroup = shared.Flatbuffers.String()
	return bench
}

// BenchBasicMemInfo runs the BasicMemInfo benches for Serialize/Deserialize.
func BenchBasicMemInfo() []benchutil.Bench {
	var results []benchutil.Bench
	basicMemInfo = make([][]byte, shared.Len)

	bench := newBench(shared.BasicMemInfo.String())
	bench.Desc = shared.Marshal.String()
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(basicMemInfoSerialize))
	results = append(results, bench)
	bench.Desc = shared.Unmarshal.String()
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(basicMemInfoDeserialize))
	results = append(results, bench)
	basicMemInfo = nil
	return results
}

func basicMemInfoSerialize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			basicMemInfo[j] = serializeBasicMemInfo(shared.BasicMemInfoData[j])
		}
	}
}

func basicMemInfoDeserialize(b *testing.B) {
	var tmp *BasicMemInfo
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			tmp = GetRootAsBasicMemInfo(basicMemInfo[j], 0)
		}
	}
	_ = tmp
}

func serializeBasicMemInfo(sh shared.ShBasicMemInfo) []byte {
	bldr.Reset()
	BasicMemInfoStart(bldr)
	BasicMemInfoAddMemTotal(bldr, sh.MemTotal)
	BasicMemInfoAddMemFree(bldr, sh.MemFree)
	BasicMemInfoAddMemAvailable(bldr, sh.MemAvailable)
	BasicMemInfoAddBuffers(bldr, sh.Buffers)
	BasicMemInfoAddCached(bldr, sh.Cached)
	BasicMemInfoAddSwapCached(bldr, sh.SwapCached)
	BasicMemInfoAddSwapTotal(bldr, sh.SwapTotal)
	BasicMemInfoAddSwapFree(bldr, sh.SwapFree)
	bldr.Finish(MemInfoEnd(bldr))
	return bldr.Bytes[bldr.Head():]
}

// BenchMemInfo runs the MemInfo benches for Serialize/Deserialize.
func BenchMemInfo() []benchutil.Bench {
	var results []benchutil.Bench
	memInfo = make([][]byte, shared.Len)

	bench := newBench(shared.MemInfo.String())
	bench.Desc = shared.Marshal.String()
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(memInfoSerialize))
	results = append(results, bench)
	bench.Desc = shared.Unmarshal.String()
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(memInfoDeserialize))
	results = append(results, bench)
	memInfo = nil
	return results
}

func memInfoSerialize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			memInfo[j] = serializeMemInfo(shared.MemInfoData[j])
		}
	}
}

func memInfoDeserialize(b *testing.B) {
	var tmp *MemInfo
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			tmp = GetRootAsMemInfo(memInfo[j], 0)
		}
	}
	_ = tmp
}

func serializeMemInfo(sh shared.ShMemInfo) []byte {
	bldr.Reset()
	MemInfoStart(bldr)
	MemInfoAddMemTotal(bldr, sh.MemTotal)
	MemInfoAddMemFree(bldr, sh.MemFree)
	MemInfoAddMemAvailable(bldr, sh.MemAvailable)
	MemInfoAddBuffers(bldr, sh.Buffers)
	MemInfoAddCached(bldr, sh.Cached)
	MemInfoAddSwapCached(bldr, sh.SwapCached)
	MemInfoAddActive(bldr, sh.Active)
	MemInfoAddInactive(bldr, sh.Inactive)
	MemInfoAddActiveAnon(bldr, sh.ActiveAnon)
	MemInfoAddInactiveAnon(bldr, sh.InactiveAnon)
	MemInfoAddActiveFile(bldr, sh.ActiveFile)
	MemInfoAddInactiveFile(bldr, sh.InactiveFile)
	MemInfoAddUnevictable(bldr, sh.Unevictable)
	MemInfoAddMlocked(bldr, sh.Mlocked)
	MemInfoAddSwapTotal(bldr, sh.SwapTotal)
	MemInfoAddSwapFree(bldr, sh.SwapFree)
	MemInfoAddDirty(bldr, sh.Dirty)
	MemInfoAddWriteback(bldr, sh.Writeback)
	MemInfoAddAnonPages(bldr, sh.AnonPages)
	MemInfoAddMapped(bldr, sh.Mapped)
	MemInfoAddShmem(bldr, sh.Shmem)
	MemInfoAddSlab(bldr, sh.Slab)
	MemInfoAddSReclaimable(bldr, sh.SReclaimable)
	MemInfoAddSUnreclaim(bldr, sh.SUnreclaim)
	MemInfoAddKernelStack(bldr, sh.KernelStack)
	MemInfoAddNFSUnstable(bldr, sh.NFSUnstable)
	MemInfoAddBounce(bldr, sh.Bounce)
	MemInfoAddWritebackTmp(bldr, sh.WritebackTmp)
	MemInfoAddCommitLimit(bldr, sh.CommitLimit)
	MemInfoAddVmallocTotal(bldr, sh.VmallocTotal)
	MemInfoAddVmallocUsed(bldr, sh.VmallocUsed)
	MemInfoAddVmallocChunk(bldr, sh.VmallocChunk)
	MemInfoAddHardwareCorrupted(bldr, sh.HardwareCorrupted)
	MemInfoAddAnonHugePages(bldr, sh.AnonHugePages)
	MemInfoAddHugePagesTotal(bldr, sh.HugePagesTotal)
	MemInfoAddHugePagesFree(bldr, sh.HugePagesFree)
	MemInfoAddHugePagesRsvd(bldr, sh.HugePagesRsvd)
	MemInfoAddHugepagesize(bldr, sh.Hugepagesize)
	MemInfoAddDirectMap4k(bldr, sh.DirectMap4k)
	MemInfoAddDirectMap2M(bldr, sh.DirectMap2M)
	bldr.Finish(MemInfoEnd(bldr))
	return bldr.Bytes[bldr.Head():]
}

// BenchMessage runs the Message benches for Serialize/Deserialize.
func BenchMessage(l int) []benchutil.Bench {
	var results []benchutil.Bench
	message = make([][]byte, shared.Len)

	bench := newBench(fmt.Sprintf("%s: %d", shared.Message.String(), l))
	bench.Desc = shared.Marshal.String()
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(messageSerialize))
	results = append(results, bench)
	bench.Desc = shared.Unmarshal.String()
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(messageDeserialize))
	results = append(results, bench)
	message = nil
	return results
}

func messageSerialize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			message[j] = serializeMessage(shared.MessageData[j])
		}
	}
}

func messageDeserialize(b *testing.B) {
	var tmp *Message
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			tmp = GetRootAsMessage(message[j], 0)
		}
	}
	_ = tmp
}

func serializeMessage(sh shared.ShMessage) []byte {
	bldr.Reset()
	id := bldr.CreateByteVector(sh.ID)
	data := bldr.CreateByteVector(sh.Data)
	MessageStart(bldr)
	MessageAddID(bldr, id)
	MessageAddType(bldr, sh.Type)
	MessageAddKind(bldr, sh.Kind)
	MessageAddData(bldr, data)
	bldr.Finish(MessageEnd(bldr))
	return bldr.Bytes[bldr.Head():]
}

// BenchRedditAccount runs the RedditAccount benches for Serialize/Deserialize.
func BenchRedditAccount() []benchutil.Bench {
	var results []benchutil.Bench
	redditAccount = make([][]byte, shared.Len)

	bench := newBench(shared.RedditAccount.String())
	bench.Desc = shared.Marshal.String()
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(redditAccountSerialize))
	results = append(results, bench)
	bench.Desc = shared.Unmarshal.String()
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(redditAccountDeserialize))
	results = append(results, bench)
	redditAccount = nil
	return results
}

func redditAccountSerialize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			redditAccount[j] = serializeRedditAccount(shared.RedditAccountData[j])
		}
	}
}

func redditAccountDeserialize(b *testing.B) {
	var tmp *RedditAccount
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			tmp = GetRootAsRedditAccount(redditAccount[j], 0)
		}
	}
	_ = tmp
}

func serializeRedditAccount(sh shared.ShRedditAccount) []byte {
	bldr.Reset()
	rid := bldr.CreateString(sh.ID)
	rname := bldr.CreateString(sh.Name)
	kind := bldr.CreateString(sh.Kind)
	aid := bldr.CreateString(sh.Data.ID)
	hash := bldr.CreateString(sh.Data.ModHash)
	aname := bldr.CreateString(sh.Data.Name)

	AccountDataStart(bldr)
	AccountDataAddCommentKarma(bldr, sh.Data.CommentKarma)
	AccountDataAddHasMail(bldr, boolToByte(sh.Data.HasMail))
	AccountDataAddHasModMail(bldr, boolToByte(sh.Data.HasModMail))
	AccountDataAddID(bldr, aid)
	AccountDataAddInboxCount(bldr, sh.Data.InboxCount)
	AccountDataAddIsFriend(bldr, boolToByte(sh.Data.IsFriend))
	AccountDataAddIsGold(bldr, boolToByte(sh.Data.IsGold))
	AccountDataAddLinkKarma(bldr, sh.Data.LinkKarma)
	AccountDataAddModHash(bldr, hash)
	AccountDataAddName(bldr, aname)
	AccountDataAddOver18(bldr, boolToByte(sh.Data.Over18))
	acc := AccountDataEnd(bldr)

	RedditAccountStart(bldr)
	RedditAccountAddID(bldr, rid)
	RedditAccountAddName(bldr, rname)
	RedditAccountAddKind(bldr, kind)
	RedditAccountAddData(bldr, acc)
	bldr.Finish(RedditAccountEnd(bldr))
	return bldr.Bytes[bldr.Head():]
}

// BenchCPUInfo runs the CPUInfo benches for Serialize/Deserialize.
func BenchCPUInfo(l int) []benchutil.Bench {
	var results []benchutil.Bench
	cpuInfo = make([][]byte, 0, shared.Len)

	bench := newBench(fmt.Sprintf("%s: %d", shared.CPUInfo, l))
	bench.Desc = shared.Marshal.String()
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(cpuInfoSerialize))
	results = append(results, bench)
	bench.Desc = shared.Unmarshal.String()
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(cpuInfoDeserialize))
	results = append(results, bench)
	message = nil
	return results
}

func cpuInfoSerialize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			cpuInfo = append(cpuInfo, serializeCPUInfo(shared.CPUInfoData[j]))
		}
	}
}

func cpuInfoDeserialize(b *testing.B) {
	var tmp *CPUInfo
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			tmp = GetRootAsCPUInfo(cpuInfo[j], 0)
		}
	}
	_ = tmp
}

func serializeCPUInfo(sh shared.ShCPUInfo) []byte {
	bldr.Reset()
	cpus := make([]flatbuffers.UOffsetT, len(sh.CPUs))
	vendorIDs := make([]flatbuffers.UOffsetT, len(sh.CPUs))
	cpuFamilies := make([]flatbuffers.UOffsetT, len(sh.CPUs))
	models := make([]flatbuffers.UOffsetT, len(sh.CPUs))
	modelNames := make([]flatbuffers.UOffsetT, len(sh.CPUs))
	steppings := make([]flatbuffers.UOffsetT, len(sh.CPUs))
	microcodes := make([]flatbuffers.UOffsetT, len(sh.CPUs))
	cacheSizes := make([]flatbuffers.UOffsetT, len(sh.CPUs))
	fpus := make([]flatbuffers.UOffsetT, len(sh.CPUs))
	fpuExceptions := make([]flatbuffers.UOffsetT, len(sh.CPUs))
	cpuIDLevels := make([]flatbuffers.UOffsetT, len(sh.CPUs))
	wps := make([]flatbuffers.UOffsetT, len(sh.CPUs))
	flags := make([]flatbuffers.UOffsetT, len(sh.CPUs))
	clFlushSizes := make([]flatbuffers.UOffsetT, len(sh.CPUs))
	cacheAlignments := make([]flatbuffers.UOffsetT, len(sh.CPUs))
	addressSizes := make([]flatbuffers.UOffsetT, len(sh.CPUs))
	powerManagements := make([]flatbuffers.UOffsetT, len(sh.CPUs))
	// create the strings
	for i := 0; i < len(sh.CPUs); i++ {
		vendorIDs[i] = bldr.CreateString(sh.CPUs[i].VendorID)
		cpuFamilies[i] = bldr.CreateString(sh.CPUs[i].CPUFamily)
		models[i] = bldr.CreateString(sh.CPUs[i].Model)
		modelNames[i] = bldr.CreateString(sh.CPUs[i].ModelName)
		steppings[i] = bldr.CreateString(sh.CPUs[i].Stepping)
		microcodes[i] = bldr.CreateString(sh.CPUs[i].Microcode)
		cacheSizes[i] = bldr.CreateString(sh.CPUs[i].CacheSize)
		fpus[i] = bldr.CreateString(sh.CPUs[i].FPU)
		fpuExceptions[i] = bldr.CreateString(sh.CPUs[i].FPUException)
		cpuIDLevels[i] = bldr.CreateString(sh.CPUs[i].CPUIDLevel)
		wps[i] = bldr.CreateString(sh.CPUs[i].WP)
		flags[i] = bldr.CreateString(sh.CPUs[i].Flags)
		clFlushSizes[i] = bldr.CreateString(sh.CPUs[i].CLFlushSize)
		cacheAlignments[i] = bldr.CreateString(sh.CPUs[i].CacheAlignment)
		addressSizes[i] = bldr.CreateString(sh.CPUs[i].AddressSizes)
		powerManagements[i] = bldr.CreateString(sh.CPUs[i].PowerManagement)
	}
	// create the CPUs
	for i := 0; i < len(sh.CPUs); i++ {
		CPUStart(bldr)
		CPUAddProcessor(bldr, sh.CPUs[i].Processor)
		CPUAddVendorID(bldr, vendorIDs[i])
		CPUAddCPUFamily(bldr, cpuFamilies[i])
		CPUAddModel(bldr, models[i])
		CPUAddModelName(bldr, modelNames[i])
		CPUAddStepping(bldr, steppings[i])
		CPUAddMicrocode(bldr, microcodes[i])
		CPUAddCPUMHz(bldr, sh.CPUs[i].CPUMHz)
		CPUAddCacheSize(bldr, cacheSizes[i])
		CPUAddPhysicalID(bldr, sh.CPUs[i].PhysicalID)
		CPUAddSiblings(bldr, sh.CPUs[i].Siblings)
		CPUAddCoreID(bldr, sh.CPUs[i].CoreID)
		CPUAddCPUCores(bldr, sh.CPUs[i].CPUCores)
		CPUAddApicID(bldr, sh.CPUs[i].ApicID)
		CPUAddInitialApicID(bldr, sh.CPUs[i].InitialApicID)
		CPUAddFPU(bldr, fpus[i])
		CPUAddFPUException(bldr, fpuExceptions[i])
		CPUAddCPUIDLevel(bldr, cpuIDLevels[i])
		CPUAddWP(bldr, wps[i])
		CPUAddFlags(bldr, flags[i])
		CPUAddBogoMIPS(bldr, sh.CPUs[i].BogoMIPS)
		CPUAddCLFlushSize(bldr, clFlushSizes[i])
		CPUAddCacheAlignment(bldr, cacheAlignments[i])
		CPUAddAddressSizes(bldr, addressSizes[i])
		CPUAddPowerManagement(bldr, powerManagements[i])
		cpus[i] = CPUEnd(bldr)
	}
	// Process the flat.CPUs vector
	CPUInfoStartCPUsVector(bldr, len(cpus))
	for i := len(cpus) - 1; i >= 0; i-- {
		bldr.PrependUOffsetT(cpus[i])
	}
	cpusV := bldr.EndVector(len(cpus))
	CPUInfoStart(bldr)
	CPUInfoAddCPUs(bldr, cpusV)
	bldr.Finish(CPUInfoEnd(bldr))
	return bldr.Bytes[bldr.Head():]
}

func boolToByte(b bool) byte {
	if b {
		return 0x00
	}
	return 0x01
}
