package fb

import (
	"testing"

	"github.com/google/flatbuffers/go"
	"github.com/mohae/serial-bowl/shared"
)

var basicMemInfo [][]byte
var memInfo [][]byte
var message [][]byte
var redditAccount [][]byte

var bldr = flatbuffers.NewBuilder(0)

// BenchBasicMemInfo runs the BasicMemInfo benches for Serialize/Deserialize.
func BenchBasicMemInfo() (m, u testing.BenchmarkResult) {
	basicMemInfo = make([][]byte, shared.Len)
	m = testing.Benchmark(basicMemInfoSerialize)
	u = testing.Benchmark(basicMemInfoDeserialize)
	basicMemInfo = nil
	return m, u
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

func serializeBasicMemInfo(sh shared.BasicMemInfo) []byte {
	bldr.Reset()
	BasicMemInfoStart(bldr)
	BasicMemInfoAddMemTotal(bldr, int64(sh.MemTotal))
	BasicMemInfoAddMemFree(bldr, int64(sh.MemFree))
	BasicMemInfoAddMemAvailable(bldr, int64(sh.MemAvailable))
	BasicMemInfoAddBuffers(bldr, int64(sh.Buffers))
	BasicMemInfoAddCached(bldr, int64(sh.Cached))
	BasicMemInfoAddSwapCached(bldr, int64(sh.SwapCached))
	BasicMemInfoAddSwapTotal(bldr, int64(sh.SwapTotal))
	BasicMemInfoAddSwapFree(bldr, int64(sh.SwapFree))
	bldr.Finish(MemInfoEnd(bldr))
	return bldr.Bytes[bldr.Head():]
}

// BenchMemInfo runs the MemInfo benches for Serialize/Deserialize.
func BenchMemInfo() (m, u testing.BenchmarkResult) {
	memInfo = make([][]byte, shared.Len)
	m = testing.Benchmark(memInfoSerialize)
	u = testing.Benchmark(memInfoDeserialize)
	memInfo = nil
	return m, u
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

func serializeMemInfo(sh shared.MemInfo) []byte {
	bldr.Reset()
	MemInfoStart(bldr)
	MemInfoAddMemTotal(bldr, int64(sh.MemTotal))
	MemInfoAddMemFree(bldr, int64(sh.MemFree))
	MemInfoAddMemAvailable(bldr, int64(sh.MemAvailable))
	MemInfoAddBuffers(bldr, int64(sh.Buffers))
	MemInfoAddCached(bldr, int64(sh.Cached))
	MemInfoAddSwapCached(bldr, int64(sh.SwapCached))
	MemInfoAddActive(bldr, int64(sh.Active))
	MemInfoAddInactive(bldr, int64(sh.Inactive))
	MemInfoAddActiveAnon(bldr, int64(sh.ActiveAnon))
	MemInfoAddInactiveAnon(bldr, int64(sh.InactiveAnon))
	MemInfoAddActiveFile(bldr, int64(sh.ActiveFile))
	MemInfoAddInactiveFile(bldr, int64(sh.InactiveFile))
	MemInfoAddUnevictable(bldr, int64(sh.Unevictable))
	MemInfoAddMlocked(bldr, int64(sh.Mlocked))
	MemInfoAddSwapTotal(bldr, int64(sh.SwapTotal))
	MemInfoAddSwapFree(bldr, int64(sh.SwapFree))
	MemInfoAddDirty(bldr, int64(sh.Dirty))
	MemInfoAddWriteback(bldr, int64(sh.Writeback))
	MemInfoAddAnonPages(bldr, int64(sh.AnonPages))
	MemInfoAddMapped(bldr, int64(sh.Mapped))
	MemInfoAddShmem(bldr, int64(sh.Shmem))
	MemInfoAddSlab(bldr, int64(sh.Slab))
	MemInfoAddSReclaimable(bldr, int64(sh.SReclaimable))
	MemInfoAddSUnreclaim(bldr, int64(sh.SUnreclaim))
	MemInfoAddKernelStack(bldr, int64(sh.KernelStack))
	MemInfoAddNFSUnstable(bldr, int64(sh.NFSUnstable))
	MemInfoAddBounce(bldr, int64(sh.Bounce))
	MemInfoAddWritebackTmp(bldr, int64(sh.WritebackTmp))
	MemInfoAddCommitLimit(bldr, int64(sh.CommitLimit))
	MemInfoAddVmallocTotal(bldr, int64(sh.VmallocTotal))
	MemInfoAddVmallocUsed(bldr, int64(sh.VmallocUsed))
	MemInfoAddVmallocChunk(bldr, int64(sh.VmallocChunk))
	MemInfoAddHardwareCorrupted(bldr, int64(sh.HardwareCorrupted))
	MemInfoAddAnonHugePages(bldr, int64(sh.AnonHugePages))
	MemInfoAddHugePagesTotal(bldr, int64(sh.HugePagesTotal))
	MemInfoAddHugePagesFree(bldr, int64(sh.HugePagesFree))
	MemInfoAddHugePagesRsvd(bldr, int64(sh.HugePagesRsvd))
	MemInfoAddHugepagesize(bldr, int64(sh.Hugepagesize))
	MemInfoAddDirectMap4k(bldr, int64(sh.DirectMap4k))
	MemInfoAddDirectMap2M(bldr, int64(sh.DirectMap2M))
	bldr.Finish(MemInfoEnd(bldr))
	return bldr.Bytes[bldr.Head():]
}

// BenchMessage runs the Message benches for Serialize/Deserialize.
func BenchMessage() (m, u testing.BenchmarkResult) {
	message = make([][]byte, shared.Len)
	m = testing.Benchmark(messageSerialize)
	u = testing.Benchmark(messageDeserialize)
	message = nil
	return m, u
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

func serializeMessage(sh shared.Message) []byte {
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
func BenchRedditAccount() (m, u testing.BenchmarkResult) {
	redditAccount = make([][]byte, shared.Len)
	m = testing.Benchmark(redditAccountSerialize)
	u = testing.Benchmark(redditAccountDeserialize)
	redditAccount = nil
	return m, u
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

func serializeRedditAccount(sh shared.RedditAccount) []byte {
	bldr.Reset()
	rid := bldr.CreateString(sh.ID)
	rname := bldr.CreateString(sh.Name)
	kind := bldr.CreateString(sh.Kind)
	aid := bldr.CreateString(sh.Data.ID)
	hash := bldr.CreateString(sh.Data.ModHash)
	aname := bldr.CreateString(sh.Data.Name)

	AccountDataStart(bldr)
	AccountDataAddCommentKarma(bldr, int64(sh.Data.CommentKarma))
	AccountDataAddHasMail(bldr, boolToByte(sh.Data.HasMail))
	AccountDataAddHasModMail(bldr, boolToByte(sh.Data.HasModMail))
	AccountDataAddID(bldr, aid)
	AccountDataAddInboxCount(bldr, int64(sh.Data.InboxCount))
	AccountDataAddIsFriend(bldr, boolToByte(sh.Data.IsFriend))
	AccountDataAddIsGold(bldr, boolToByte(sh.Data.IsGold))
	AccountDataAddLinkKarma(bldr, int64(sh.Data.LinkKarma))
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

func boolToByte(b bool) byte {
	if b {
		return 0x00
	}
	return 0x01
}
