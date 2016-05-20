// package shared contains things shared across tests and the various
// serialization protocols.  This exists to prevent import cycles.
package shared

import (
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
	FFJSON
	FFJSONBuffer
	ProtobufV3
	CapnProto2
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
	Protos = []Proto{Flatbuffers, JSON, ProtobufV3, CapnProto2}
	for _, v := range Protos {
		if len(v.String()) > maxProtoLen {
			maxProtoLen = len(v.String())
		}
	}
	maxProtoLen += PadLen
	StructTypes = []StructType{BasicMemInfo, MemInfo, Message, RedditAccount}
	for _, v := range StructTypes {
		if len(v.String()) > maxStructTypeLen {
			maxStructTypeLen = len(v.String())
		}
	}
	maxStructTypeLen += PadLen
}

func GenData() {
	var rnd pcg.Rand
	rnd.Seed(benchutil.SeedVal())
	GenBasicMemInfoData(Len, rnd)
	GenMemInfoData(Len, rnd)
	GenRedditAccountData(Len, rnd)
	// Message data is generated at start of each message data length test
	// As each block of message data tests use a different message.Data length.
}

// GenBasicMemInfoData generates the random data for the BasicMemInfo struct.
// The resulting slice of structs will have l elements.
func GenBasicMemInfoData(l int, rand pcg.Rand) {
	BasicMemInfoData = make([]ShBasicMemInfo, l)
	for i := 0; i < l; i++ {
		BasicMemInfoData[i] = ShBasicMemInfo{
			MemTotal:     rand.Int63(),
			MemFree:      rand.Int63(),
			MemAvailable: rand.Int63(),
			Buffers:      rand.Int63(),
			Cached:       rand.Int63(),
			SwapCached:   rand.Int63(),
			SwapTotal:    rand.Int63(),
			SwapFree:     rand.Int63(),
		}
	}
}

// GenMemInfoData generates the random data for the BasicMemInfo struct.  The
// resulting slice of structs will have l elements.
func GenMemInfoData(l int, rand pcg.Rand) {
	MemInfoData = make([]ShMemInfo, l)
	for i := 0; i < l; i++ {
		MemInfoData[i] = ShMemInfo{
			MemTotal:          rand.Int63(),
			MemFree:           rand.Int63(),
			MemAvailable:      rand.Int63(),
			Buffers:           rand.Int63(),
			Cached:            rand.Int63(),
			SwapCached:        rand.Int63(),
			Active:            rand.Int63(),
			Inactive:          rand.Int63(),
			ActiveAnon:        rand.Int63(),
			InactiveAnon:      rand.Int63(),
			ActiveFile:        rand.Int63(),
			InactiveFile:      rand.Int63(),
			Unevictable:       rand.Int63(),
			Mlocked:           rand.Int63(),
			SwapTotal:         rand.Int63(),
			SwapFree:          rand.Int63(),
			Dirty:             rand.Int63(),
			Writeback:         rand.Int63(),
			AnonPages:         rand.Int63(),
			Mapped:            rand.Int63(),
			Shmem:             rand.Int63(),
			Slab:              rand.Int63(),
			SReclaimable:      rand.Int63(),
			SUnreclaim:        rand.Int63(),
			KernelStack:       rand.Int63(),
			NFSUnstable:       rand.Int63(),
			Bounce:            rand.Int63(),
			WritebackTmp:      rand.Int63(),
			CommitLimit:       rand.Int63(),
			VmallocTotal:      rand.Int63(),
			VmallocUsed:       rand.Int63(),
			VmallocChunk:      rand.Int63(),
			HardwareCorrupted: rand.Int63(),
			AnonHugePages:     rand.Int63(),
			HugePagesTotal:    rand.Int63(),
			HugePagesFree:     rand.Int63(),
			HugePagesRsvd:     rand.Int63(),
			Hugepagesize:      rand.Int63(),
			DirectMap4k:       rand.Int63(),
			DirectMap2M:       rand.Int63(),
		}
	}
}

// GenMessageData generates the random data for the Message struct whose data
// element being n bytes in length.  The resulting slice of structs will have
// l elements.
func GenMessageData(n, l int, rand pcg.Rand) {
	MessageData = make([]ShMessage, l)
	for i := 0; i < l; i++ {
		id := benchutil.RandBytes(8)
		data := benchutil.RandBytes(n)
		MessageData[i] = ShMessage{
			ID:     id,
			DestID: rand.Next(),
			Type:   int8(rand.Bound(1<<7 - 1)),
			Kind:   int16(rand.Bound(1<<15 - 1)),
			Data:   data,
		}
	}
}

func GenRedditAccountData(l int, rand pcg.Rand) {
	RedditAccountData = make([]ShRedditAccount, l)
	for i := 0; i < l; i++ {
		RedditAccountData[i] = ShRedditAccount{
			ID:   benchutil.RandString(20),
			Name: benchutil.RandString(int(rand.Bound(30))),
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
				Name:         benchutil.RandString(int(rand.Bound(30))),
				Over18:       benchutil.RandBool(),
			},
		}
	}
}
