// automatically generated, do not modify

package fb

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type BasicMemInfo struct {
	_tab flatbuffers.Table
}

func GetRootAsBasicMemInfo(buf []byte, offset flatbuffers.UOffsetT) *BasicMemInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &BasicMemInfo{}
	x.Init(buf, n + offset)
	return x
}

func (rcv *BasicMemInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *BasicMemInfo) MemTotal() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *BasicMemInfo) MemFree() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *BasicMemInfo) MemAvailable() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *BasicMemInfo) Buffers() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *BasicMemInfo) Cached() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *BasicMemInfo) SwapCached() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *BasicMemInfo) SwapTotal() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *BasicMemInfo) SwapFree() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func BasicMemInfoStart(builder *flatbuffers.Builder) { builder.StartObject(8) }
func BasicMemInfoAddMemTotal(builder *flatbuffers.Builder, MemTotal int64) { builder.PrependInt64Slot(0, MemTotal, 0) }
func BasicMemInfoAddMemFree(builder *flatbuffers.Builder, MemFree int64) { builder.PrependInt64Slot(1, MemFree, 0) }
func BasicMemInfoAddMemAvailable(builder *flatbuffers.Builder, MemAvailable int64) { builder.PrependInt64Slot(2, MemAvailable, 0) }
func BasicMemInfoAddBuffers(builder *flatbuffers.Builder, Buffers int64) { builder.PrependInt64Slot(3, Buffers, 0) }
func BasicMemInfoAddCached(builder *flatbuffers.Builder, Cached int64) { builder.PrependInt64Slot(4, Cached, 0) }
func BasicMemInfoAddSwapCached(builder *flatbuffers.Builder, SwapCached int64) { builder.PrependInt64Slot(5, SwapCached, 0) }
func BasicMemInfoAddSwapTotal(builder *flatbuffers.Builder, SwapTotal int64) { builder.PrependInt64Slot(6, SwapTotal, 0) }
func BasicMemInfoAddSwapFree(builder *flatbuffers.Builder, SwapFree int64) { builder.PrependInt64Slot(7, SwapFree, 0) }
func BasicMemInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
