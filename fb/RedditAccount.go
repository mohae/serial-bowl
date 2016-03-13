// automatically generated, do not modify

package fb

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type RedditAccount struct {
	_tab flatbuffers.Table
}

func GetRootAsRedditAccount(buf []byte, offset flatbuffers.UOffsetT) *RedditAccount {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &RedditAccount{}
	x.Init(buf, n + offset)
	return x
}

func (rcv *RedditAccount) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *RedditAccount) ID() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *RedditAccount) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *RedditAccount) Kind() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *RedditAccount) Data(obj *AccountData) *AccountData {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(AccountData)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func RedditAccountStart(builder *flatbuffers.Builder) { builder.StartObject(4) }
func RedditAccountAddID(builder *flatbuffers.Builder, ID flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(ID), 0) }
func RedditAccountAddName(builder *flatbuffers.Builder, Name flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(Name), 0) }
func RedditAccountAddKind(builder *flatbuffers.Builder, Kind flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(Kind), 0) }
func RedditAccountAddData(builder *flatbuffers.Builder, Data flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(Data), 0) }
func RedditAccountEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
