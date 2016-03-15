// package shared contains things shared across tests and the various
// serialization protocols.  This exists to prevent import cycles.
package shared

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"testing"
	"time"

	"github.com/mohae/csv2md"
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
	JSON
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
var maxOpLen int
var maxProtoLen int
var maxStructTypeLen int

func init() {
	Ops = []Op{Marshal, Unmarshal, Serialize, Deserialize}
	for _, v := range Ops {
		if len(v.String()) > maxOpLen {
			maxOpLen = len(v.String())
		}
	}
	maxOpLen += 5
	Protos = []Proto{Flatbuffers, JSON, ProtobufV3, CapnProto2}
	for _, v := range Protos {
		if len(v.String()) > maxProtoLen {
			maxProtoLen = len(v.String())
		}
	}
	maxProtoLen += 5
	StructTypes = []StructType{BasicMemInfo, MemInfo, Message, RedditAccount}
	for _, v := range StructTypes {
		if len(v.String()) > maxStructTypeLen {
			maxStructTypeLen = len(v.String())
		}
	}
	maxStructTypeLen += 5
}

// Bench holds information about a serialization protocol's benchmark.
type Bench struct {
	Proto        Proto         // the type of protocol buffer
	StructString string        // the struct being benched
	Results      map[Op]Result // A map of Result keyed by Op.
}

func (b Bench) TXTOutput() []string {
	var out []string
	for _, v := range Ops {
		r, ok := b.Results[v]
		if !ok {
			continue
		}
		out = append(out, b.formatOutput(v, r))
	}
	return out
}

func (b Bench) formatOutput(op Op, r Result) string {
	return fmt.Sprintf("%s%s%s%s", column(maxProtoLen, b.Proto.String()), column(maxOpLen, op.String()), column(maxStructTypeLen, b.StructString), r.String())
}

func (b Bench) CSVOutput() [][]string {
	var out [][]string
	for _, v := range Ops {
		r, ok := b.Results[v]
		if !ok {
			continue
		}
		tmp := []string{b.Proto.String(), v.String(), b.StructString}
		tmp = append(tmp, r.CSV()...)
		out = append(out, tmp)
	}
	return out
}

// Holds result information
type Result struct {
	Ops      int64 // the number of operations performed
	NsOp     int64 // The amount of time, in Nanoseconds, per Op.
	BytesOp  int64 // The number of bytes allocated per Op.
	AllocsOp int64 // The number of Allocations per Op.
}

// ResultFromBenchmarkResult creates a Result{} from a testing.BenchmarkResult.
func ResultFromBenchmarkResult(br testing.BenchmarkResult) Result {
	var r Result
	r.Ops = int64(br.N) * int64(Len)
	r.NsOp = br.T.Nanoseconds() / r.Ops
	r.BytesOp = int64(br.MemBytes) / r.Ops
	r.AllocsOp = int64(br.MemAllocs) / r.Ops
	return r
}

func (r Result) OpsString() string {
	return fmt.Sprintf("%d ops", r.Ops)
}

func (r Result) NsOpString() string {
	return fmt.Sprintf("%d ns/Op", r.NsOp)
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

func (r Result) CSV() []string {
	return []string{fmt.Sprintf("%d", r.Ops), fmt.Sprintf("%d", r.NsOp), fmt.Sprintf("%d", r.BytesOp), fmt.Sprintf("%d", r.AllocsOp)}
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
func GenMemInfoData(l int) {
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
func GenMessageData(n, l int) {
	MessageData = make([]ShMessage, l)
	for i := 0; i < l; i++ {
		id := RandBytes(8)
		data := RandBytes(n)
		MessageData[i] = ShMessage{
			ID:     id,
			DestID: rand.Uint32(),
			Type:   int8(rand.Intn(1<<7 - 1)),
			Kind:   int16(rand.Intn(1<<15 - 1)),
			Data:   data,
		}
	}
}

func GenRedditAccountData(l int) {
	RedditAccountData = make([]ShRedditAccount, l)
	for i := 0; i < l; i++ {
		RedditAccountData[i] = ShRedditAccount{
			ID:   RandString(20),
			Name: RandString(rand.Intn(30)),
			Kind: RandString(5),
			Data: AccountData{
				CommentKarma: rand.Int63(),
				HasMail:      RandBool(),
				HasModMail:   RandBool(),
				ID:           RandString(20),
				InboxCount:   rand.Int63(),
				IsFriend:     RandBool(),
				IsGold:       RandBool(),
				LinkKarma:    rand.Int63(),
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

func TXTOut(w io.Writer, benchResults []Bench) {
	for _, v := range benchResults {
		lines := v.TXTOutput()
		for _, line := range lines {
			fmt.Fprintln(w, line)
		}
	}
}
func CSVOut(w io.Writer, benchResults []Bench) error {
	wr := csv.NewWriter(w)
	defer wr.Flush()
	// first write out the header
	err := wr.Write([]string{"Protocol", "Operation", "Data Type", "Operations", "Ns/Op", "Bytes/Op", "Allocs/Op"})
	if err != nil {
		return err
	}
	for _, bench := range benchResults {
		lines := bench.CSVOutput()
		for _, line := range lines {
			err := wr.Write(line)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func MDOut(w io.Writer, benchResults []Bench) error {
	var buff bytes.Buffer
	// first generate the csv
	err := CSVOut(&buff, benchResults)
	if err != nil {
		return fmt.Errorf("error while creating intermediate CSV: %s", err)
	}
	// then transmogrify to MD
	t := csv2md.NewTransmogrifier(&buff, w)
	t.SetFieldAlignment([]string{"l", "l", "l", "r", "r", "r", "r"})
	return t.MDTable()
}
