package gogofast

import (
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/mohae/benchutil"
	"github.com/mohae/serial-bowl/shared"
)

// marshal results holder.
var (
	redditAccount  [][]byte
	memInfo        [][]byte
	basicMemInfo   [][]byte
	message        [][]byte
	cpuInfo        [][]byte
	basicMemInfos  []GoGoFastBasicMemInfo
	memInfos       []GoGoFastMemInfo
	messages       []GoGoFastMessage
	redditAccounts []GoGoFastRedditAccount
	cpuInfos       []GoGoFastCPUInfo
)

func newBench(s string) benchutil.Bench {
	bench := benchutil.NewBench("gogo/protobuf/protoc-gen-gogofast")
	bench.Iterations = shared.Len
	bench.Group = s
	bench.SubGroup = shared.ProtobufV3.String()
	return bench
}

// BenchBasicMemInfo runs the BasicMemInfo benches for Marshal/Unmarshal.
func BenchBasicMemInfo() []benchutil.Bench {
	var results []benchutil.Bench
	basicMemInfo = make([][]byte, shared.Len)
	basicMemInfos = PrepareBasicMemInfoData(shared.BasicMemInfoData)

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
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			basicMemInfo[j], _ = proto.Marshal(&basicMemInfos[j])
		}
	}
}

func basicMemInfoUnmarshal(b *testing.B) {
	var tmp GoGoFastBasicMemInfo
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			proto.Unmarshal(basicMemInfo[j], &tmp)
		}
	}
	_ = tmp
}

// BenchMemInfo runs the MemInfo benches for Marshal/Unmarshal.
func BenchMemInfo() []benchutil.Bench {
	var results []benchutil.Bench
	memInfo = make([][]byte, shared.Len)
	memInfos = PrepareMemInfoData(shared.MemInfoData)

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
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			memInfo[j], _ = proto.Marshal(&memInfos[j])
		}
	}
}

func memInfoUnmarshal(b *testing.B) {
	var tmp GoGoFastMemInfo
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			proto.Unmarshal(memInfo[j], &tmp)
		}
	}
	_ = tmp
}

// BenchMessage runs the Message benches for Marshal/Unmarshal.
func BenchMessage(l int) []benchutil.Bench {
	var results []benchutil.Bench
	message = make([][]byte, shared.Len)
	messages = PrepareMessageData(shared.MessageData)

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
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			message[j], _ = proto.Marshal(&messages[j])
		}
	}
}

func messageUnmarshal(b *testing.B) {
	var tmp GoGoFastMessage
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			proto.Unmarshal(message[j], &tmp)
		}
	}
	_ = tmp
}

// BenchmarkRedditAccount runs the RedditAccount benches for Marshal/Unmarshal.
func BenchRedditAccount() []benchutil.Bench {
	var results []benchutil.Bench
	redditAccount = make([][]byte, shared.Len)
	redditAccounts = PrepareRedditAccountData(shared.RedditAccountData)

	bench := newBench(shared.RedditAccount.String())
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
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			redditAccount[j], _ = proto.Marshal(&redditAccounts[j])
		}
	}
}

func redditAccountUnmarshal(b *testing.B) {
	var tmp GoGoFastRedditAccount
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			proto.Unmarshal(redditAccount[j], &tmp)
		}
	}
	_ = tmp
}

// BenchCPUInfo runs the CPUInfo benches for Marshal/Unmarshal.
func BenchCPUInfo(l int) []benchutil.Bench {
	var results []benchutil.Bench
	cpuInfo = make([][]byte, shared.Len)
	cpuInfos = PrepareCPUInfoData(shared.CPUInfoData)

	bench := newBench(fmt.Sprintf("%s: %d", shared.CPUInfo.String(), l))
	bench.Desc = shared.Marshal.String()
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(cpuInfoMarshal))
	results = append(results, bench)
	bench.Desc = shared.Unmarshal.String()
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(cpuInfoUnmarshal))
	results = append(results, bench)
	cpuInfo = nil
	return results
}

func cpuInfoMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			cpuInfo[j], _ = proto.Marshal(&cpuInfos[j])
		}
	}
}

func cpuInfoUnmarshal(b *testing.B) {
	var tmp GoGoFastCPUInfo
	for i := 0; i < b.N; i++ {
		for j := 0; j < shared.Len; j++ {
			proto.Unmarshal(cpuInfo[j], &tmp)
		}
	}
	_ = tmp
}

// PrepareBasicMemInfoData generates the protobuf version of the data.
func PrepareBasicMemInfoData(data []shared.ShBasicMemInfo) []GoGoFastBasicMemInfo {
	tmp := make([]GoGoFastBasicMemInfo, 0, len(data))
	for i := 0; i < len(data); i++ {
		tmp = append(tmp, GoGoFastBasicMemInfo{
			MemTotal:     data[i].MemTotal,
			MemFree:      data[i].MemFree,
			MemAvailable: data[i].MemAvailable,
			Buffers:      data[i].Buffers,
			Cached:       data[i].Cached,
			SwapCached:   data[i].SwapCached,
			SwapTotal:    data[i].SwapTotal,
			SwapFree:     data[i].SwapFree,
		})
	}
	return tmp
}

// PrepareMemInfoData generates the protobuf version of the data.
func PrepareMemInfoData(data []shared.ShMemInfo) []GoGoFastMemInfo {
	tmp := make([]GoGoFastMemInfo, 0, len(data))
	for i := 0; i < len(data); i++ {
		tmp = append(tmp, GoGoFastMemInfo{
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
		})
	}
	return tmp
}

// PrepareMessageData generates the protobuf version of the data.
func PrepareMessageData(data []shared.ShMessage) []GoGoFastMessage {
	tmp := make([]GoGoFastMessage, 0, len(data))
	for i := 0; i < len(data); i++ {
		tmp = append(tmp, GoGoFastMessage{
			ID:     data[i].ID,
			DestID: data[i].DestID,
			Type:   int32(data[i].Type),
			Kind:   int32(data[i].Kind),
			Data:   data[i].Data,
		})
	}
	return tmp
}

// PrepareRedditAccountData generates the protobuf version of the data.
func PrepareRedditAccountData(data []shared.ShRedditAccount) []GoGoFastRedditAccount {
	tmp := make([]GoGoFastRedditAccount, len(data))
	for i := 0; i < len(data); i++ {
		tmp = append(tmp, GoGoFastRedditAccount{
			ID:   data[i].ID,
			Name: data[i].Name,
			Kind: data[i].Kind,
			Data: &GoGoFastAccountData{
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
		})
	}
	return tmp
}

// PrepareCPUInfoData generates the ProtobufV3 version of the data.
func PrepareCPUInfoData(data []shared.ShCPUInfo) []GoGoFastCPUInfo {
	tmp := make([]GoGoFastCPUInfo, 0, len(data))
	for i := 0; i < len(data); i++ {
		inf := GoGoFastCPUInfo{CPUs: make([]*GoGoFastCPU, 0, len(data[i].CPUs))}
		for j := 0; j < len(data[i].CPUs); j++ {
			cpu := &GoGoFastCPU{
				Processor:       int32(data[i].CPUs[j].Processor),
				VendorID:        data[i].CPUs[j].VendorID,
				CPUFamily:       data[i].CPUs[j].CPUFamily,
				Model:           data[i].CPUs[j].Model,
				ModelName:       data[i].CPUs[j].ModelName,
				Stepping:        data[i].CPUs[j].Stepping,
				Microcode:       data[i].CPUs[j].Microcode,
				CPUMHz:          data[i].CPUs[j].CPUMHz,
				CacheSize:       data[i].CPUs[j].CacheSize,
				PhysicalID:      int32(data[i].CPUs[j].PhysicalID),
				Siblings:        int32(data[i].CPUs[j].Siblings),
				CoreID:          int32(data[i].CPUs[j].CoreID),
				CPUCores:        int32(data[i].CPUs[j].CPUCores),
				ApicID:          int32(data[i].CPUs[j].ApicID),
				InitialApicID:   int32(data[i].CPUs[j].InitialApicID),
				FPU:             data[i].CPUs[j].FPU,
				FPUException:    data[i].CPUs[j].FPUException,
				CPUIDLevel:      data[i].CPUs[j].CPUIDLevel,
				WP:              data[i].CPUs[j].WP,
				Flags:           data[i].CPUs[j].Flags,
				BogoMIPS:        data[i].CPUs[j].BogoMIPS,
				CLFlushSize:     data[i].CPUs[j].CLFlushSize,
				CacheAlignment:  data[i].CPUs[j].CacheAlignment,
				AddressSizes:    data[i].CPUs[j].AddressSizes,
				PowerManagement: data[i].CPUs[j].PowerManagement,
			}
			inf.CPUs = append(inf.CPUs, cpu)
		}
		tmp = append(tmp, inf)
	}
	return tmp
}
