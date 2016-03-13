// automatically generated, do not modify

package fb

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type AccountData struct {
	_tab flatbuffers.Table
}

func (rcv *AccountData) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AccountData) CommentKarma() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *AccountData) HasMail() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *AccountData) HasModMail() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *AccountData) HasVerifiedEmail() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *AccountData) ID() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *AccountData) InboxCount() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *AccountData) IsFriend() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *AccountData) IsGold() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *AccountData) IsMod() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *AccountData) LinkKarma() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *AccountData) ModHash() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *AccountData) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *AccountData) Over18() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func AccountDataStart(builder *flatbuffers.Builder) { builder.StartObject(13) }
func AccountDataAddCommentKarma(builder *flatbuffers.Builder, CommentKarma int64) { builder.PrependInt64Slot(0, CommentKarma, 0) }
func AccountDataAddHasMail(builder *flatbuffers.Builder, HasMail byte) { builder.PrependByteSlot(1, HasMail, 0) }
func AccountDataAddHasModMail(builder *flatbuffers.Builder, HasModMail byte) { builder.PrependByteSlot(2, HasModMail, 0) }
func AccountDataAddHasVerifiedEmail(builder *flatbuffers.Builder, HasVerifiedEmail byte) { builder.PrependByteSlot(3, HasVerifiedEmail, 0) }
func AccountDataAddID(builder *flatbuffers.Builder, ID flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(ID), 0) }
func AccountDataAddInboxCount(builder *flatbuffers.Builder, InboxCount int64) { builder.PrependInt64Slot(5, InboxCount, 0) }
func AccountDataAddIsFriend(builder *flatbuffers.Builder, IsFriend byte) { builder.PrependByteSlot(6, IsFriend, 0) }
func AccountDataAddIsGold(builder *flatbuffers.Builder, IsGold byte) { builder.PrependByteSlot(7, IsGold, 0) }
func AccountDataAddIsMod(builder *flatbuffers.Builder, IsMod byte) { builder.PrependByteSlot(8, IsMod, 0) }
func AccountDataAddLinkKarma(builder *flatbuffers.Builder, LinkKarma int64) { builder.PrependInt64Slot(9, LinkKarma, 0) }
func AccountDataAddModHash(builder *flatbuffers.Builder, ModHash flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(10, flatbuffers.UOffsetT(ModHash), 0) }
func AccountDataAddName(builder *flatbuffers.Builder, Name flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(11, flatbuffers.UOffsetT(Name), 0) }
func AccountDataAddOver18(builder *flatbuffers.Builder, Over18 byte) { builder.PrependByteSlot(12, Over18, 0) }
func AccountDataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
