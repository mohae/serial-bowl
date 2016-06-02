// automatically generated, do not modify

package fb

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type CPUInfo struct {
	_tab flatbuffers.Table
}

func GetRootAsCPUInfo(buf []byte, offset flatbuffers.UOffsetT) *CPUInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &CPUInfo{}
	x.Init(buf, n + offset)
	return x
}

func (rcv *CPUInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *CPUInfo) CPUs(obj *CPU, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
	if obj == nil {
		obj = new(CPU)
	}
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *CPUInfo) CPUsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func CPUInfoStart(builder *flatbuffers.Builder) { builder.StartObject(1) }
func CPUInfoAddCPUs(builder *flatbuffers.Builder, CPUs flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(CPUs), 0) }
func CPUInfoStartCPUsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT { return builder.StartVector(4, numElems, 4)
}
func CPUInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
