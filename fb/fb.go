package fb

import (
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/mohae/serial-bowl/jsn"
)

var bldr = flatbuffers.NewBuilder(0)

func BoolToByte(b bool) byte {
	if b {
		return 0x00
	}
	return 0x01
}

func SerializeRedditAccount(js jsn.RedditAccount) []byte {
	bldr.Reset()
	rid := bldr.CreateString(js.ID)
	rname := bldr.CreateString(js.Name)
	kind := bldr.CreateString(js.Kind)
	aid := bldr.CreateString(js.Data.ID)
	hash := bldr.CreateString(js.Data.ModHash)
	aname := bldr.CreateString(js.Data.Name)

	AccountDataStart(bldr)
	AccountDataAddCommentKarma(bldr, int64(js.Data.CommentKarma))
	AccountDataAddHasMail(bldr, BoolToByte(js.Data.HasMail))
	AccountDataAddHasModMail(bldr, BoolToByte(js.Data.HasModMail))
	AccountDataAddID(bldr, aid)
	AccountDataAddInboxCount(bldr, int64(js.Data.InboxCount))
	AccountDataAddIsFriend(bldr, BoolToByte(js.Data.IsFriend))
	AccountDataAddIsGold(bldr, BoolToByte(js.Data.IsGold))
	AccountDataAddLinkKarma(bldr, int64(js.Data.LinkKarma))
	AccountDataAddModHash(bldr, hash)
	AccountDataAddName(bldr, aname)
	AccountDataAddOver18(bldr, BoolToByte(js.Data.Over18))
	acc := AccountDataEnd(bldr)

	RedditAccountStart(bldr)
	RedditAccountAddID(bldr, rid)
	RedditAccountAddName(bldr, rname)
	RedditAccountAddKind(bldr, kind)
	RedditAccountAddData(bldr, acc)
	bldr.Finish(RedditAccountEnd(bldr))
	return bldr.Bytes[bldr.Head():]
}

func SerializeMemInfo(js jsn.MemInfo) []byte {
	bldr.Reset()
	MemInfoStart(bldr)
	MemInfoAddMemTotal(bldr, int64(js.MemTotal))
	MemInfoAddMemFree(bldr, int64(js.MemFree))
	MemInfoAddMemAvailable(bldr, int64(js.MemAvailable))
	MemInfoAddBuffers(bldr, int64(js.Buffers))
	MemInfoAddCached(bldr, int64(js.Cached))
	MemInfoAddSwapCached(bldr, int64(js.SwapCached))
	MemInfoAddActive(bldr, int64(js.Active))
	MemInfoAddInactive(bldr, int64(js.Inactive))
	MemInfoAddActiveAnon(bldr, int64(js.ActiveAnon))
	MemInfoAddInactiveAnon(bldr, int64(js.InactiveAnon))
	MemInfoAddActiveFile(bldr, int64(js.ActiveFile))
	MemInfoAddInactiveFile(bldr, int64(js.InactiveFile))
	MemInfoAddUnevictable(bldr, int64(js.Unevictable))
	MemInfoAddMlocked(bldr, int64(js.Mlocked))
	MemInfoAddSwapTotal(bldr, int64(js.SwapTotal))
	MemInfoAddSwapFree(bldr, int64(js.SwapFree))
	MemInfoAddDirty(bldr, int64(js.Dirty))
	MemInfoAddWriteback(bldr, int64(js.Writeback))
	MemInfoAddAnonPages(bldr, int64(js.AnonPages))
	MemInfoAddMapped(bldr, int64(js.Mapped))
	MemInfoAddShmem(bldr, int64(js.Shmem))
	MemInfoAddSlab(bldr, int64(js.Slab))
	MemInfoAddSReclaimable(bldr, int64(js.SReclaimable))
	MemInfoAddSUnreclaim(bldr, int64(js.SUnreclaim))
	MemInfoAddKernelStack(bldr, int64(js.KernelStack))
	MemInfoAddNFSUnstable(bldr, int64(js.NFSUnstable))
	MemInfoAddBounce(bldr, int64(js.Bounce))
	MemInfoAddWritebackTmp(bldr, int64(js.WritebackTmp))
	MemInfoAddCommitLimit(bldr, int64(js.CommitLimit))
	MemInfoAddVmallocTotal(bldr, int64(js.VmallocTotal))
	MemInfoAddVmallocUsed(bldr, int64(js.VmallocUsed))
	MemInfoAddVmallocChunk(bldr, int64(js.VmallocChunk))
	MemInfoAddHardwareCorrupted(bldr, int64(js.HardwareCorrupted))
	MemInfoAddAnonHugePages(bldr, int64(js.AnonHugePages))
	MemInfoAddHugePagesTotal(bldr, int64(js.HugePagesTotal))
	MemInfoAddHugePagesFree(bldr, int64(js.HugePagesFree))
	MemInfoAddHugePagesRsvd(bldr, int64(js.HugePagesRsvd))
	MemInfoAddHugepagesize(bldr, int64(js.Hugepagesize))
	MemInfoAddDirectMap4k(bldr, int64(js.DirectMap4k))
	MemInfoAddDirectMap2M(bldr, int64(js.DirectMap2M))
	bldr.Finish(MemInfoEnd(bldr))
	return bldr.Bytes[bldr.Head():]
}

func SerializeBasicMemInfo(js jsn.BasicMemInfo) []byte {
	bldr.Reset()
	BasicMemInfoStart(bldr)
	BasicMemInfoAddMemTotal(bldr, int64(js.MemTotal))
	BasicMemInfoAddMemFree(bldr, int64(js.MemFree))
	BasicMemInfoAddMemAvailable(bldr, int64(js.MemAvailable))
	BasicMemInfoAddBuffers(bldr, int64(js.Buffers))
	BasicMemInfoAddCached(bldr, int64(js.Cached))
	BasicMemInfoAddSwapCached(bldr, int64(js.SwapCached))
	BasicMemInfoAddSwapTotal(bldr, int64(js.SwapTotal))
	BasicMemInfoAddSwapFree(bldr, int64(js.SwapFree))
	bldr.Finish(MemInfoEnd(bldr))
	return bldr.Bytes[bldr.Head():]
}

func SerializeMessage(js jsn.Message) []byte {
	bldr.Reset()
	id := bldr.CreateByteVector(js.ID)
	data := bldr.CreateByteVector(js.Data)
	MessageStart(bldr)
	MessageAddID(bldr, id)
	MessageAddType(bldr, js.Type)
	MessageAddKind(bldr, js.Kind)
	MessageAddData(bldr, data)
	bldr.Finish(MessageEnd(bldr))
	return bldr.Bytes[bldr.Head():]
}
