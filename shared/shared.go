// package shared contains things shared across tests and the various
// serialization protocols.  This exists to prevent import cycles.
package shared

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/mohae/serial-bowl/jsn"
)

const alphanum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const maxInt64 = 1<<63 - 1

// The length of the slice holding the data structures.
var Len = 1000

// Data for the benchmarks.  All benchmarks operate off of the same set of
// randomly generated data.  The package jsn structs are used because they
// are just Go structs.
var BasicMemInfoData []jsn.BasicMemInfo
var MemInfoData []jsn.MemInfo
var MessageData []jsn.Message
var RedditAccountData []jsn.RedditAccount

type Result struct {
	Ops      int64
	NsOp     int64
	MBSec    float64
	BytesOp  int64
	AllocsOp int64
}

func (r *Result) SetFromBenchmarkResult(br testing.BenchmarkResult) {
	r.Ops = int64(br.N) * int64(Len)
	r.NsOp = br.T.Nanoseconds() / r.Ops
	r.MBSec = (float64(br.Bytes) * float64(r.Ops) / 1e6) / br.T.Seconds()
	r.BytesOp = int64(br.MemBytes) / r.Ops
	r.AllocsOp = int64(br.MemAllocs) / r.Ops
}

func (r Result) OpsString() string {
	return fmt.Sprintf("%d ops", r.Ops)
}

func (r Result) NsOpString() string {
	return fmt.Sprintf("%d ns/Op", r.NsOp)
}

func (r Result) MBSecString() string {
	return fmt.Sprintf("%9.2f MB/s", r.MBSec)
}

func (r Result) BytesOpString() string {
	return fmt.Sprintf("%d bytes/Op", r.BytesOp)
}

func (r Result) AllocsOpString() string {
	return fmt.Sprintf("%d allocs/Op", r.AllocsOp)
}

func (r Result) String() string {
	return fmt.Sprintf("%s%s%s%s", column(15, r.OpsString()), column(15, r.NsOpString()), column(18, r.BytesOpString()), column(16, r.AllocsOpString()))
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenData() {
	GenBasicMemInfoData(Len)
	GenMemInfoData(Len)
	GenRedditAccountData(Len)
	// Message data is generated at start of each message data length test
	// As each block of message data tests use a different message.Data length.
}

// GenBasicMemInfoData generates the random data for the BasicMemInfo struct.
// The resulting slice of structs will have l elements.
func GenBasicMemInfoData(l int) {
	BasicMemInfoData = make([]jsn.BasicMemInfo, l)
	for i := 0; i < l; i++ {
		BasicMemInfoData[i] = jsn.BasicMemInfo{
			MemTotal:     rand.Intn(maxInt64),
			MemFree:      rand.Intn(maxInt64),
			MemAvailable: rand.Intn(maxInt64),
			Buffers:      rand.Intn(maxInt64),
			Cached:       rand.Intn(maxInt64),
			SwapCached:   rand.Intn(maxInt64),
			SwapTotal:    rand.Intn(maxInt64),
			SwapFree:     rand.Intn(maxInt64),
		}
	}
}

// GenMemInfoData generates the random data for the BasicMemInfo struct.  The
// resulting slice of structs will have l elements.
func GenMemInfoData(l int) {
	MemInfoData = make([]jsn.MemInfo, l)
	for i := 0; i < l; i++ {
		MemInfoData[i] = jsn.MemInfo{
			MemTotal:          rand.Intn(maxInt64),
			MemFree:           rand.Intn(maxInt64),
			MemAvailable:      rand.Intn(maxInt64),
			Buffers:           rand.Intn(maxInt64),
			Cached:            rand.Intn(maxInt64),
			SwapCached:        rand.Intn(maxInt64),
			Active:            rand.Intn(maxInt64),
			Inactive:          rand.Intn(maxInt64),
			ActiveAnon:        rand.Intn(maxInt64),
			InactiveAnon:      rand.Intn(maxInt64),
			ActiveFile:        rand.Intn(maxInt64),
			InactiveFile:      rand.Intn(maxInt64),
			Unevictable:       rand.Intn(maxInt64),
			Mlocked:           rand.Intn(maxInt64),
			SwapTotal:         rand.Intn(maxInt64),
			SwapFree:          rand.Intn(maxInt64),
			Dirty:             rand.Intn(maxInt64),
			Writeback:         rand.Intn(maxInt64),
			AnonPages:         rand.Intn(maxInt64),
			Mapped:            rand.Intn(maxInt64),
			Shmem:             rand.Intn(maxInt64),
			Slab:              rand.Intn(maxInt64),
			SReclaimable:      rand.Intn(maxInt64),
			SUnreclaim:        rand.Intn(maxInt64),
			KernelStack:       rand.Intn(maxInt64),
			NFSUnstable:       rand.Intn(maxInt64),
			Bounce:            rand.Intn(maxInt64),
			WritebackTmp:      rand.Intn(maxInt64),
			CommitLimit:       rand.Intn(maxInt64),
			VmallocTotal:      rand.Intn(maxInt64),
			VmallocUsed:       rand.Intn(maxInt64),
			VmallocChunk:      rand.Intn(maxInt64),
			HardwareCorrupted: rand.Intn(maxInt64),
			AnonHugePages:     rand.Intn(maxInt64),
			HugePagesTotal:    rand.Intn(maxInt64),
			HugePagesFree:     rand.Intn(maxInt64),
			HugePagesRsvd:     rand.Intn(maxInt64),
			Hugepagesize:      rand.Intn(maxInt64),
			DirectMap4k:       rand.Intn(maxInt64),
			DirectMap2M:       rand.Intn(maxInt64),
		}
	}
}

// GenMessageData generates the random data for the Message struct whose data
// element being n bytes in length.  The resulting slice of structs will have
// l elements.
func GenMessageData(n, l int) {
	MessageData = make([]jsn.Message, l)
	for i := 0; i < l; i++ {
		id := RandBytes(8)
		data := RandBytes(n)
		MessageData[i] = jsn.Message{
			ID:     id,
			DestID: rand.Uint32(),
			Type:   int8(rand.Intn(1<<7 - 1)),
			Kind:   int16(rand.Intn(1<<15 - 1)),
			Data:   data,
		}
	}
}

func GenRedditAccountData(l int) {
	RedditAccountData = make([]jsn.RedditAccount, l)
	for i := 0; i < l; i++ {
		RedditAccountData[i] = jsn.RedditAccount{
			ID:   RandString(20),
			Name: RandString(rand.Intn(30)),
			Kind: RandString(5),
			Data: jsn.AccountData{
				CommentKarma: rand.Intn(maxInt64),
				HasMail:      RandBool(),
				HasModMail:   RandBool(),
				ID:           RandString(20),
				InboxCount:   rand.Intn(maxInt64),
				IsFriend:     RandBool(),
				IsGold:       RandBool(),
				LinkKarma:    rand.Intn(maxInt64),
				ModHash:      RandString(88),
				Name:         RandString(rand.Intn(30)),
				Over18:       RandBool(),
			},
		}
	}
}

// RandString returns a randomly generated string of length l.
func RandString(l int) string {
	return string(RandBytes(l))
}

// RandBytes returns a randomly generated []byte of length l.  The values of
// these bytes are restricted to the ASCII alphanum range; that doesn't matter
// for the purposes of these benchmarks.
func RandBytes(l int) []byte {
	b := make([]byte, l)
	for i := 0; i < l; i++ {
		b[i] = alphanum[rand.Intn(len(alphanum))]
	}
	return b
}

// RandBool returns a pseudo-random bool value.
func RandBool() bool {
	if rand.Int31()%2 == 0 {
		return false
	}
	return true
}

// column returns a right justified string of width w.
func column(w int, s string) string {
	pad := w - len(s)
	padding := make([]byte, pad)
	for i := 0; i < pad; i++ {
		padding[i] = 0x20
	}
	return fmt.Sprintf("%s%s", string(padding), s)
}
