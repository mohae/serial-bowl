// package shared contains things shared across tests and the various
// serialization protocols.  This exists to prevent import cycles.
package shared

import (
	"math/rand"

	pcg "github.com/dgryski/go-pcgr"
	"github.com/mohae/benchutil"
)

const alphanum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// The length of the slice holding the data structures.
var Len = 1000

// Data for the benchmarks.  All benchmarks operate off of the same set of
// randomly generated data.
var BasicMemInfoData []ShBasicMemInfo
var MemInfoData []ShMemInfo
var MessageData []ShMessage
var RedditAccountData []ShRedditAccount
var CPUInfoData []ShCPUInfo

//go:generate stringer -type=Op
type Op int

const (
	Marshal Op = iota
	Unmarshal
	Serialize
	Deserialize
)

// Ops holds a slice of Op values.  This allows consistency in ordering of
// output from Bench.Results.
var Ops []Op

//go:generate stringer -type=Proto
type Proto int

const (
	UnknownProt Proto = iota
	Flatbuffers
	GenCode
	Gob
	JSON
	ProtobufV3
	CapnProto2
	MessagePack
)

// Protos holds a slice of Protocol types.
var Protos []Proto

//go:generate stringer -type=StructType
type StructType int

const (
	UnknownStruct StructType = iota
	BasicMemInfo
	MemInfo
	Message
	RedditAccount
	CPUInfo
)

// StructTypes holds a slice of StructTypes.
var StructTypes []StructType

// Max Length of various types (for formatting purposes)
// and other formatting info
var (
	maxOpLen         int
	maxProtoLen      int
	maxStructTypeLen int
	PadLen           int = 2
)

func init() {
	Ops = []Op{Marshal, Unmarshal, Serialize, Deserialize}
	for _, v := range Ops {
		if len(v.String()) > maxOpLen {
			maxOpLen = len(v.String())
		}
	}
	maxOpLen += PadLen
	Protos = []Proto{Flatbuffers, GenCode, Gob, JSON, ProtobufV3, CapnProto2, MessagePack}
	for _, v := range Protos {
		if len(v.String()) > maxProtoLen {
			maxProtoLen = len(v.String())
		}
	}
	maxProtoLen += PadLen
	StructTypes = []StructType{BasicMemInfo, MemInfo, Message, RedditAccount, CPUInfo}
	for _, v := range StructTypes {
		if len(v.String()) > maxStructTypeLen {
			maxStructTypeLen = len(v.String())
		}
	}
	maxStructTypeLen += PadLen
}

func GenData() {
	GenBasicMemInfoData(Len)
	GenMemInfoData(Len)
	GenRedditAccountData(Len)
	// Message data is generated at start of each message data length test
	// As each block of message data tests use a different message.Data length.

	// CPUInfo data is generated at start of each test as each block of tests
	// use a different number of CPUs.
}

// GenBasicMemInfoData generates the random data for the BasicMemInfo struct.
// The resulting slice of structs will have l elements.
func GenBasicMemInfoData(l int) {
	var rnd pcg.Rand
	rnd.Seed(benchutil.Seed())
	BasicMemInfoData = make([]ShBasicMemInfo, l)
	for i := 0; i < l; i++ {
		BasicMemInfoData[i] = ShBasicMemInfo{
			MemTotal:     rnd.Int63(),
			MemFree:      rnd.Int63(),
			MemAvailable: rnd.Int63(),
			Buffers:      rnd.Int63(),
			Cached:       rnd.Int63(),
			SwapCached:   rnd.Int63(),
			SwapTotal:    rnd.Int63(),
			SwapFree:     rnd.Int63(),
		}
	}
}

// GenMemInfoData generates the random data for the BasicMemInfo struct.  The
// resulting slice of structs will have l elements.
func GenMemInfoData(l int) {
	var rnd pcg.Rand
	rnd.Seed(benchutil.Seed())
	MemInfoData = make([]ShMemInfo, 0, l)
	for i := 0; i < l; i++ {
		tmp := ShMemInfo{
			MemTotal:          rnd.Int63(),
			MemFree:           rnd.Int63(),
			MemAvailable:      rnd.Int63(),
			Buffers:           rnd.Int63(),
			Cached:            rnd.Int63(),
			SwapCached:        rnd.Int63(),
			Active:            rnd.Int63(),
			Inactive:          rnd.Int63(),
			ActiveAnon:        rnd.Int63(),
			InactiveAnon:      rnd.Int63(),
			ActiveFile:        rnd.Int63(),
			InactiveFile:      rnd.Int63(),
			Unevictable:       rnd.Int63(),
			Mlocked:           rnd.Int63(),
			SwapTotal:         rnd.Int63(),
			SwapFree:          rnd.Int63(),
			Dirty:             rnd.Int63(),
			Writeback:         rnd.Int63(),
			AnonPages:         rnd.Int63(),
			Mapped:            rnd.Int63(),
			Shmem:             rnd.Int63(),
			Slab:              rnd.Int63(),
			SReclaimable:      rnd.Int63(),
			SUnreclaim:        rnd.Int63(),
			KernelStack:       rnd.Int63(),
			NFSUnstable:       rnd.Int63(),
			Bounce:            rnd.Int63(),
			WritebackTmp:      rnd.Int63(),
			CommitLimit:       rnd.Int63(),
			VmallocTotal:      rnd.Int63(),
			VmallocUsed:       rnd.Int63(),
			VmallocChunk:      rnd.Int63(),
			HardwareCorrupted: rnd.Int63(),
			AnonHugePages:     rnd.Int63(),
			HugePagesTotal:    rnd.Int63(),
			HugePagesFree:     rnd.Int63(),
			HugePagesRsvd:     rnd.Int63(),
			Hugepagesize:      rnd.Int63(),
			DirectMap4k:       rnd.Int63(),
			DirectMap2M:       rnd.Int63(),
		}
		MemInfoData = append(MemInfoData, tmp)
	}
}

// GenMessageData generates the random data for the Message struct whose data
// element being n bytes in length.  The resulting slice of structs will have
// l elements.
func GenMessageData(n, l int) {
	var rnd pcg.Rand
	rnd.Seed(benchutil.Seed())
	MessageData = make([]ShMessage, 0, l)
	for i := 0; i < l; i++ {
		id := benchutil.RandBytes(8)
		data := benchutil.RandBytes(uint32(n))
		tmp := ShMessage{
			ID:     id,
			DestID: rnd.Next(),
			Type:   int8(rnd.Bound(1<<7 - 1)),
			Kind:   int16(rnd.Bound(1<<15 - 1)),
			Data:   data,
		}
		MessageData = append(MessageData, tmp)
	}
}

func GenRedditAccountData(l int) {
	var rnd pcg.Rand
	rnd.Seed(benchutil.Seed())
	RedditAccountData = make([]ShRedditAccount, 0, l)
	for i := 0; i < l; i++ {
		tmp := ShRedditAccount{
			ID:   benchutil.RandString(20),
			Name: benchutil.RandString(uint32(rnd.Bound(30))),
			Kind: benchutil.RandString(5),
			Data: AccountData{
				CommentKarma: rand.Int63(),
				HasMail:      benchutil.RandBool(),
				HasModMail:   benchutil.RandBool(),
				ID:           benchutil.RandString(20),
				InboxCount:   rand.Int63(),
				IsFriend:     benchutil.RandBool(),
				IsGold:       benchutil.RandBool(),
				LinkKarma:    rand.Int63(),
				ModHash:      benchutil.RandString(88),
				Name:         benchutil.RandString(uint32(rnd.Bound(30))),
				Over18:       benchutil.RandBool(),
			},
		}
		RedditAccountData = append(RedditAccountData, tmp)
	}
}

// n = number of cpus the data structure should have
func GenCPUInfoData(n, l int) []ShCPUInfo {
	CPUInfoData = make([]ShCPUInfo, 0, l)
	for i := 0; i < l; i++ {
		tmp := ShCPUInfo{CPUs: make([]ShCPU, 0, n)}
		for j := 0; j < n; j++ {
			cpu := ShCPU{
				Processor:       int16(j),
				VendorID:        "GenuineIntel",
				Model:           "6",
				ModelName:       "CPUInfoData",
				Stepping:        "9",
				Microcode:       "0x19",
				CPUMHz:          3400.480,
				CacheSize:       "6144KB",
				PhysicalID:      int16(j),
				Siblings:        int16(n),
				CoreID:          int16(j),
				CPUCores:        int16(n),
				ApicID:          int16(j),
				InitialApicID:   int16(j),
				FPU:             "yes",
				FPUException:    "yes",
				CPUIDLevel:      "13",
				WP:              "yes",
				Flags:           "pu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 syscall nx rdtscp lm constant_tsc rep_good nopl xtopology nonstop_tsc pni pclmulqdq monitor ssse3 cx16 sse4_1 sse4_2 popcnt aes xsave avx rdrand hypervisor lahf_lm",
				BogoMIPS:        6800.96,
				CLFlushSize:     "64",
				CacheAlignment:  "64",
				AddressSizes:    "36 bits physical, 48 bits virtual",
				PowerManagement: "",
			}
			tmp.CPUs = append(tmp.CPUs, cpu)
		}
		CPUInfoData = append(CPUInfoData, tmp)
	}
	return CPUInfoData
}
