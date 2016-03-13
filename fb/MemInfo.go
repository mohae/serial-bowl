// automatically generated, do not modify

package fb

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type MemInfo struct {
	_tab flatbuffers.Table
}

func GetRootAsMemInfo(buf []byte, offset flatbuffers.UOffsetT) *MemInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MemInfo{}
	x.Init(buf, n + offset)
	return x
}

func (rcv *MemInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MemInfo) MemTotal() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) MemFree() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) MemAvailable() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) Buffers() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) Cached() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) SwapCached() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) Active() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) Inactive() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) ActiveAnon() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) InactiveAnon() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) ActiveFile() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) InactiveFile() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) Unevictable() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) Mlocked() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) SwapTotal() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(32))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) SwapFree() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(34))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) Dirty() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(36))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) Writeback() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(38))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) AnonPages() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(40))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) Mapped() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(42))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) Shmem() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(44))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) Slab() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(46))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) SReclaimable() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(48))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) SUnreclaim() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(50))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) KernelStack() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(52))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) NFSUnstable() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(54))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) Bounce() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(56))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) WritebackTmp() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(58))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) CommitLimit() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(60))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) VmallocTotal() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(62))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) VmallocUsed() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(64))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) VmallocChunk() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(66))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) HardwareCorrupted() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(68))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) AnonHugePages() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(70))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) HugePagesTotal() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(72))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) HugePagesFree() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(74))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) HugePagesRsvd() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(76))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) Hugepagesize() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(78))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) DirectMap4k() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(80))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemInfo) DirectMap2M() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(82))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func MemInfoStart(builder *flatbuffers.Builder) { builder.StartObject(40) }
func MemInfoAddMemTotal(builder *flatbuffers.Builder, MemTotal int64) { builder.PrependInt64Slot(0, MemTotal, 0) }
func MemInfoAddMemFree(builder *flatbuffers.Builder, MemFree int64) { builder.PrependInt64Slot(1, MemFree, 0) }
func MemInfoAddMemAvailable(builder *flatbuffers.Builder, MemAvailable int64) { builder.PrependInt64Slot(2, MemAvailable, 0) }
func MemInfoAddBuffers(builder *flatbuffers.Builder, Buffers int64) { builder.PrependInt64Slot(3, Buffers, 0) }
func MemInfoAddCached(builder *flatbuffers.Builder, Cached int64) { builder.PrependInt64Slot(4, Cached, 0) }
func MemInfoAddSwapCached(builder *flatbuffers.Builder, SwapCached int64) { builder.PrependInt64Slot(5, SwapCached, 0) }
func MemInfoAddActive(builder *flatbuffers.Builder, Active int64) { builder.PrependInt64Slot(6, Active, 0) }
func MemInfoAddInactive(builder *flatbuffers.Builder, Inactive int64) { builder.PrependInt64Slot(7, Inactive, 0) }
func MemInfoAddActiveAnon(builder *flatbuffers.Builder, ActiveAnon int64) { builder.PrependInt64Slot(8, ActiveAnon, 0) }
func MemInfoAddInactiveAnon(builder *flatbuffers.Builder, InactiveAnon int64) { builder.PrependInt64Slot(9, InactiveAnon, 0) }
func MemInfoAddActiveFile(builder *flatbuffers.Builder, ActiveFile int64) { builder.PrependInt64Slot(10, ActiveFile, 0) }
func MemInfoAddInactiveFile(builder *flatbuffers.Builder, InactiveFile int64) { builder.PrependInt64Slot(11, InactiveFile, 0) }
func MemInfoAddUnevictable(builder *flatbuffers.Builder, Unevictable int64) { builder.PrependInt64Slot(12, Unevictable, 0) }
func MemInfoAddMlocked(builder *flatbuffers.Builder, Mlocked int64) { builder.PrependInt64Slot(13, Mlocked, 0) }
func MemInfoAddSwapTotal(builder *flatbuffers.Builder, SwapTotal int64) { builder.PrependInt64Slot(14, SwapTotal, 0) }
func MemInfoAddSwapFree(builder *flatbuffers.Builder, SwapFree int64) { builder.PrependInt64Slot(15, SwapFree, 0) }
func MemInfoAddDirty(builder *flatbuffers.Builder, Dirty int64) { builder.PrependInt64Slot(16, Dirty, 0) }
func MemInfoAddWriteback(builder *flatbuffers.Builder, Writeback int64) { builder.PrependInt64Slot(17, Writeback, 0) }
func MemInfoAddAnonPages(builder *flatbuffers.Builder, AnonPages int64) { builder.PrependInt64Slot(18, AnonPages, 0) }
func MemInfoAddMapped(builder *flatbuffers.Builder, Mapped int64) { builder.PrependInt64Slot(19, Mapped, 0) }
func MemInfoAddShmem(builder *flatbuffers.Builder, Shmem int64) { builder.PrependInt64Slot(20, Shmem, 0) }
func MemInfoAddSlab(builder *flatbuffers.Builder, Slab int64) { builder.PrependInt64Slot(21, Slab, 0) }
func MemInfoAddSReclaimable(builder *flatbuffers.Builder, SReclaimable int64) { builder.PrependInt64Slot(22, SReclaimable, 0) }
func MemInfoAddSUnreclaim(builder *flatbuffers.Builder, SUnreclaim int64) { builder.PrependInt64Slot(23, SUnreclaim, 0) }
func MemInfoAddKernelStack(builder *flatbuffers.Builder, KernelStack int64) { builder.PrependInt64Slot(24, KernelStack, 0) }
func MemInfoAddNFSUnstable(builder *flatbuffers.Builder, NFSUnstable int64) { builder.PrependInt64Slot(25, NFSUnstable, 0) }
func MemInfoAddBounce(builder *flatbuffers.Builder, Bounce int64) { builder.PrependInt64Slot(26, Bounce, 0) }
func MemInfoAddWritebackTmp(builder *flatbuffers.Builder, WritebackTmp int64) { builder.PrependInt64Slot(27, WritebackTmp, 0) }
func MemInfoAddCommitLimit(builder *flatbuffers.Builder, CommitLimit int64) { builder.PrependInt64Slot(28, CommitLimit, 0) }
func MemInfoAddVmallocTotal(builder *flatbuffers.Builder, VmallocTotal int64) { builder.PrependInt64Slot(29, VmallocTotal, 0) }
func MemInfoAddVmallocUsed(builder *flatbuffers.Builder, VmallocUsed int64) { builder.PrependInt64Slot(30, VmallocUsed, 0) }
func MemInfoAddVmallocChunk(builder *flatbuffers.Builder, VmallocChunk int64) { builder.PrependInt64Slot(31, VmallocChunk, 0) }
func MemInfoAddHardwareCorrupted(builder *flatbuffers.Builder, HardwareCorrupted int64) { builder.PrependInt64Slot(32, HardwareCorrupted, 0) }
func MemInfoAddAnonHugePages(builder *flatbuffers.Builder, AnonHugePages int64) { builder.PrependInt64Slot(33, AnonHugePages, 0) }
func MemInfoAddHugePagesTotal(builder *flatbuffers.Builder, HugePagesTotal int64) { builder.PrependInt64Slot(34, HugePagesTotal, 0) }
func MemInfoAddHugePagesFree(builder *flatbuffers.Builder, HugePagesFree int64) { builder.PrependInt64Slot(35, HugePagesFree, 0) }
func MemInfoAddHugePagesRsvd(builder *flatbuffers.Builder, HugePagesRsvd int64) { builder.PrependInt64Slot(36, HugePagesRsvd, 0) }
func MemInfoAddHugepagesize(builder *flatbuffers.Builder, Hugepagesize int64) { builder.PrependInt64Slot(37, Hugepagesize, 0) }
func MemInfoAddDirectMap4k(builder *flatbuffers.Builder, DirectMap4k int64) { builder.PrependInt64Slot(38, DirectMap4k, 0) }
func MemInfoAddDirectMap2M(builder *flatbuffers.Builder, DirectMap2M int64) { builder.PrependInt64Slot(39, DirectMap2M, 0) }
func MemInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
