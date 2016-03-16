package capnp

// AUTO GENERATED - DO NOT EDIT

import (
	capnp "zombiezen.com/go/capnproto2"
)

type BasicMemInfo struct{ capnp.Struct }

func NewBasicMemInfo(s *capnp.Segment) (BasicMemInfo, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 64, PointerCount: 0})
	if err != nil {
		return BasicMemInfo{}, err
	}
	return BasicMemInfo{st}, nil
}

func NewRootBasicMemInfo(s *capnp.Segment) (BasicMemInfo, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 64, PointerCount: 0})
	if err != nil {
		return BasicMemInfo{}, err
	}
	return BasicMemInfo{st}, nil
}

func ReadRootBasicMemInfo(msg *capnp.Message) (BasicMemInfo, error) {
	root, err := msg.Root()
	if err != nil {
		return BasicMemInfo{}, err
	}
	st := capnp.ToStruct(root)
	return BasicMemInfo{st}, nil
}

func (s BasicMemInfo) MemTotal() int64 {
	return int64(s.Struct.Uint64(0))
}

func (s BasicMemInfo) SetMemTotal(v int64) {

	s.Struct.SetUint64(0, uint64(v))
}

func (s BasicMemInfo) MemFree() int64 {
	return int64(s.Struct.Uint64(8))
}

func (s BasicMemInfo) SetMemFree(v int64) {

	s.Struct.SetUint64(8, uint64(v))
}

func (s BasicMemInfo) MemAvailable() int64 {
	return int64(s.Struct.Uint64(16))
}

func (s BasicMemInfo) SetMemAvailable(v int64) {

	s.Struct.SetUint64(16, uint64(v))
}

func (s BasicMemInfo) Buffers() int64 {
	return int64(s.Struct.Uint64(24))
}

func (s BasicMemInfo) SetBuffers(v int64) {

	s.Struct.SetUint64(24, uint64(v))
}

func (s BasicMemInfo) Cached() int64 {
	return int64(s.Struct.Uint64(32))
}

func (s BasicMemInfo) SetCached(v int64) {

	s.Struct.SetUint64(32, uint64(v))
}

func (s BasicMemInfo) SwapCached() int64 {
	return int64(s.Struct.Uint64(40))
}

func (s BasicMemInfo) SetSwapCached(v int64) {

	s.Struct.SetUint64(40, uint64(v))
}

func (s BasicMemInfo) SwapTotal() int64 {
	return int64(s.Struct.Uint64(48))
}

func (s BasicMemInfo) SetSwapTotal(v int64) {

	s.Struct.SetUint64(48, uint64(v))
}

func (s BasicMemInfo) SwapFree() int64 {
	return int64(s.Struct.Uint64(56))
}

func (s BasicMemInfo) SetSwapFree(v int64) {

	s.Struct.SetUint64(56, uint64(v))
}

// BasicMemInfo_List is a list of BasicMemInfo.
type BasicMemInfo_List struct{ capnp.List }

// NewBasicMemInfo creates a new list of BasicMemInfo.
func NewBasicMemInfo_List(s *capnp.Segment, sz int32) (BasicMemInfo_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 64, PointerCount: 0}, sz)
	if err != nil {
		return BasicMemInfo_List{}, err
	}
	return BasicMemInfo_List{l}, nil
}

func (s BasicMemInfo_List) At(i int) BasicMemInfo           { return BasicMemInfo{s.List.Struct(i)} }
func (s BasicMemInfo_List) Set(i int, v BasicMemInfo) error { return s.List.SetStruct(i, v.Struct) }

// BasicMemInfo_Promise is a wrapper for a BasicMemInfo promised by a client call.
type BasicMemInfo_Promise struct{ *capnp.Pipeline }

func (p BasicMemInfo_Promise) Struct() (BasicMemInfo, error) {
	s, err := p.Pipeline.Struct()
	return BasicMemInfo{s}, err
}

type MemInfo struct{ capnp.Struct }

func NewMemInfo(s *capnp.Segment) (MemInfo, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 320, PointerCount: 0})
	if err != nil {
		return MemInfo{}, err
	}
	return MemInfo{st}, nil
}

func NewRootMemInfo(s *capnp.Segment) (MemInfo, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 320, PointerCount: 0})
	if err != nil {
		return MemInfo{}, err
	}
	return MemInfo{st}, nil
}

func ReadRootMemInfo(msg *capnp.Message) (MemInfo, error) {
	root, err := msg.Root()
	if err != nil {
		return MemInfo{}, err
	}
	st := capnp.ToStruct(root)
	return MemInfo{st}, nil
}

func (s MemInfo) MemTotal() int64 {
	return int64(s.Struct.Uint64(0))
}

func (s MemInfo) SetMemTotal(v int64) {

	s.Struct.SetUint64(0, uint64(v))
}

func (s MemInfo) MemFree() int64 {
	return int64(s.Struct.Uint64(8))
}

func (s MemInfo) SetMemFree(v int64) {

	s.Struct.SetUint64(8, uint64(v))
}

func (s MemInfo) MemAvailable() int64 {
	return int64(s.Struct.Uint64(16))
}

func (s MemInfo) SetMemAvailable(v int64) {

	s.Struct.SetUint64(16, uint64(v))
}

func (s MemInfo) Buffers() int64 {
	return int64(s.Struct.Uint64(24))
}

func (s MemInfo) SetBuffers(v int64) {

	s.Struct.SetUint64(24, uint64(v))
}

func (s MemInfo) Cached() int64 {
	return int64(s.Struct.Uint64(32))
}

func (s MemInfo) SetCached(v int64) {

	s.Struct.SetUint64(32, uint64(v))
}

func (s MemInfo) SwapCached() int64 {
	return int64(s.Struct.Uint64(40))
}

func (s MemInfo) SetSwapCached(v int64) {

	s.Struct.SetUint64(40, uint64(v))
}

func (s MemInfo) Active() int64 {
	return int64(s.Struct.Uint64(48))
}

func (s MemInfo) SetActive(v int64) {

	s.Struct.SetUint64(48, uint64(v))
}

func (s MemInfo) Inactive() int64 {
	return int64(s.Struct.Uint64(56))
}

func (s MemInfo) SetInactive(v int64) {

	s.Struct.SetUint64(56, uint64(v))
}

func (s MemInfo) ActiveAnon() int64 {
	return int64(s.Struct.Uint64(64))
}

func (s MemInfo) SetActiveAnon(v int64) {

	s.Struct.SetUint64(64, uint64(v))
}

func (s MemInfo) InactiveAnon() int64 {
	return int64(s.Struct.Uint64(72))
}

func (s MemInfo) SetInactiveAnon(v int64) {

	s.Struct.SetUint64(72, uint64(v))
}

func (s MemInfo) ActiveFile() int64 {
	return int64(s.Struct.Uint64(80))
}

func (s MemInfo) SetActiveFile(v int64) {

	s.Struct.SetUint64(80, uint64(v))
}

func (s MemInfo) InactiveFile() int64 {
	return int64(s.Struct.Uint64(88))
}

func (s MemInfo) SetInactiveFile(v int64) {

	s.Struct.SetUint64(88, uint64(v))
}

func (s MemInfo) Unevictable() int64 {
	return int64(s.Struct.Uint64(96))
}

func (s MemInfo) SetUnevictable(v int64) {

	s.Struct.SetUint64(96, uint64(v))
}

func (s MemInfo) Mlocked() int64 {
	return int64(s.Struct.Uint64(104))
}

func (s MemInfo) SetMlocked(v int64) {

	s.Struct.SetUint64(104, uint64(v))
}

func (s MemInfo) SwapTotal() int64 {
	return int64(s.Struct.Uint64(112))
}

func (s MemInfo) SetSwapTotal(v int64) {

	s.Struct.SetUint64(112, uint64(v))
}

func (s MemInfo) SwapFree() int64 {
	return int64(s.Struct.Uint64(120))
}

func (s MemInfo) SetSwapFree(v int64) {

	s.Struct.SetUint64(120, uint64(v))
}

func (s MemInfo) Dirty() int64 {
	return int64(s.Struct.Uint64(128))
}

func (s MemInfo) SetDirty(v int64) {

	s.Struct.SetUint64(128, uint64(v))
}

func (s MemInfo) Writeback() int64 {
	return int64(s.Struct.Uint64(136))
}

func (s MemInfo) SetWriteback(v int64) {

	s.Struct.SetUint64(136, uint64(v))
}

func (s MemInfo) AnonPages() int64 {
	return int64(s.Struct.Uint64(144))
}

func (s MemInfo) SetAnonPages(v int64) {

	s.Struct.SetUint64(144, uint64(v))
}

func (s MemInfo) Mapped() int64 {
	return int64(s.Struct.Uint64(152))
}

func (s MemInfo) SetMapped(v int64) {

	s.Struct.SetUint64(152, uint64(v))
}

func (s MemInfo) Shmem() int64 {
	return int64(s.Struct.Uint64(160))
}

func (s MemInfo) SetShmem(v int64) {

	s.Struct.SetUint64(160, uint64(v))
}

func (s MemInfo) Slab() int64 {
	return int64(s.Struct.Uint64(168))
}

func (s MemInfo) SetSlab(v int64) {

	s.Struct.SetUint64(168, uint64(v))
}

func (s MemInfo) SReclaimable() int64 {
	return int64(s.Struct.Uint64(176))
}

func (s MemInfo) SetSReclaimable(v int64) {

	s.Struct.SetUint64(176, uint64(v))
}

func (s MemInfo) SUnreclaim() int64 {
	return int64(s.Struct.Uint64(184))
}

func (s MemInfo) SetSUnreclaim(v int64) {

	s.Struct.SetUint64(184, uint64(v))
}

func (s MemInfo) KernelStack() int64 {
	return int64(s.Struct.Uint64(192))
}

func (s MemInfo) SetKernelStack(v int64) {

	s.Struct.SetUint64(192, uint64(v))
}

func (s MemInfo) NFSUnstable() int64 {
	return int64(s.Struct.Uint64(200))
}

func (s MemInfo) SetNFSUnstable(v int64) {

	s.Struct.SetUint64(200, uint64(v))
}

func (s MemInfo) Bounce() int64 {
	return int64(s.Struct.Uint64(208))
}

func (s MemInfo) SetBounce(v int64) {

	s.Struct.SetUint64(208, uint64(v))
}

func (s MemInfo) WritebackTmp() int64 {
	return int64(s.Struct.Uint64(216))
}

func (s MemInfo) SetWritebackTmp(v int64) {

	s.Struct.SetUint64(216, uint64(v))
}

func (s MemInfo) CommitLimit() int64 {
	return int64(s.Struct.Uint64(224))
}

func (s MemInfo) SetCommitLimit(v int64) {

	s.Struct.SetUint64(224, uint64(v))
}

func (s MemInfo) VmallocTotal() int64 {
	return int64(s.Struct.Uint64(232))
}

func (s MemInfo) SetVmallocTotal(v int64) {

	s.Struct.SetUint64(232, uint64(v))
}

func (s MemInfo) VmallocUsed() int64 {
	return int64(s.Struct.Uint64(240))
}

func (s MemInfo) SetVmallocUsed(v int64) {

	s.Struct.SetUint64(240, uint64(v))
}

func (s MemInfo) VmallocChunk() int64 {
	return int64(s.Struct.Uint64(248))
}

func (s MemInfo) SetVmallocChunk(v int64) {

	s.Struct.SetUint64(248, uint64(v))
}

func (s MemInfo) HardwareCorrupted() int64 {
	return int64(s.Struct.Uint64(256))
}

func (s MemInfo) SetHardwareCorrupted(v int64) {

	s.Struct.SetUint64(256, uint64(v))
}

func (s MemInfo) AnonHugePages() int64 {
	return int64(s.Struct.Uint64(264))
}

func (s MemInfo) SetAnonHugePages(v int64) {

	s.Struct.SetUint64(264, uint64(v))
}

func (s MemInfo) HugePagesTotal() int64 {
	return int64(s.Struct.Uint64(272))
}

func (s MemInfo) SetHugePagesTotal(v int64) {

	s.Struct.SetUint64(272, uint64(v))
}

func (s MemInfo) HugePagesFree() int64 {
	return int64(s.Struct.Uint64(280))
}

func (s MemInfo) SetHugePagesFree(v int64) {

	s.Struct.SetUint64(280, uint64(v))
}

func (s MemInfo) HugePagesRsvd() int64 {
	return int64(s.Struct.Uint64(288))
}

func (s MemInfo) SetHugePagesRsvd(v int64) {

	s.Struct.SetUint64(288, uint64(v))
}

func (s MemInfo) Hugepagesize() int64 {
	return int64(s.Struct.Uint64(296))
}

func (s MemInfo) SetHugepagesize(v int64) {

	s.Struct.SetUint64(296, uint64(v))
}

func (s MemInfo) DirectMap4k() int64 {
	return int64(s.Struct.Uint64(304))
}

func (s MemInfo) SetDirectMap4k(v int64) {

	s.Struct.SetUint64(304, uint64(v))
}

func (s MemInfo) DirectMap2M() int64 {
	return int64(s.Struct.Uint64(312))
}

func (s MemInfo) SetDirectMap2M(v int64) {

	s.Struct.SetUint64(312, uint64(v))
}

// MemInfo_List is a list of MemInfo.
type MemInfo_List struct{ capnp.List }

// NewMemInfo creates a new list of MemInfo.
func NewMemInfo_List(s *capnp.Segment, sz int32) (MemInfo_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 320, PointerCount: 0}, sz)
	if err != nil {
		return MemInfo_List{}, err
	}
	return MemInfo_List{l}, nil
}

func (s MemInfo_List) At(i int) MemInfo           { return MemInfo{s.List.Struct(i)} }
func (s MemInfo_List) Set(i int, v MemInfo) error { return s.List.SetStruct(i, v.Struct) }

// MemInfo_Promise is a wrapper for a MemInfo promised by a client call.
type MemInfo_Promise struct{ *capnp.Pipeline }

func (p MemInfo_Promise) Struct() (MemInfo, error) {
	s, err := p.Pipeline.Struct()
	return MemInfo{s}, err
}

type Message struct{ capnp.Struct }

func NewMessage(s *capnp.Segment) (Message, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	if err != nil {
		return Message{}, err
	}
	return Message{st}, nil
}

func NewRootMessage(s *capnp.Segment) (Message, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	if err != nil {
		return Message{}, err
	}
	return Message{st}, nil
}

func ReadRootMessage(msg *capnp.Message) (Message, error) {
	root, err := msg.Root()
	if err != nil {
		return Message{}, err
	}
	st := capnp.ToStruct(root)
	return Message{st}, nil
}

func (s Message) ID() ([]byte, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return nil, err
	}

	return []byte(capnp.ToData(p)), nil

}

func (s Message) SetID(v []byte) error {

	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, d)
}

func (s Message) DestID() uint32 {
	return s.Struct.Uint32(0)
}

func (s Message) SetDestID(v uint32) {

	s.Struct.SetUint32(0, v)
}

func (s Message) Type() int8 {
	return int8(s.Struct.Uint8(4))
}

func (s Message) SetType(v int8) {

	s.Struct.SetUint8(4, uint8(v))
}

func (s Message) Kind() int16 {
	return int16(s.Struct.Uint16(6))
}

func (s Message) SetKind(v int16) {

	s.Struct.SetUint16(6, uint16(v))
}

func (s Message) Data() ([]byte, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return nil, err
	}

	return []byte(capnp.ToData(p)), nil

}

func (s Message) SetData(v []byte) error {

	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(1, d)
}

// Message_List is a list of Message.
type Message_List struct{ capnp.List }

// NewMessage creates a new list of Message.
func NewMessage_List(s *capnp.Segment, sz int32) (Message_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	if err != nil {
		return Message_List{}, err
	}
	return Message_List{l}, nil
}

func (s Message_List) At(i int) Message           { return Message{s.List.Struct(i)} }
func (s Message_List) Set(i int, v Message) error { return s.List.SetStruct(i, v.Struct) }

// Message_Promise is a wrapper for a Message promised by a client call.
type Message_Promise struct{ *capnp.Pipeline }

func (p Message_Promise) Struct() (Message, error) {
	s, err := p.Pipeline.Struct()
	return Message{s}, err
}

type RedditAccount struct{ capnp.Struct }

func NewRedditAccount(s *capnp.Segment) (RedditAccount, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4})
	if err != nil {
		return RedditAccount{}, err
	}
	return RedditAccount{st}, nil
}

func NewRootRedditAccount(s *capnp.Segment) (RedditAccount, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4})
	if err != nil {
		return RedditAccount{}, err
	}
	return RedditAccount{st}, nil
}

func ReadRootRedditAccount(msg *capnp.Message) (RedditAccount, error) {
	root, err := msg.Root()
	if err != nil {
		return RedditAccount{}, err
	}
	st := capnp.ToStruct(root)
	return RedditAccount{st}, nil
}

func (s RedditAccount) ID() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s RedditAccount) IDBytes() ([]byte, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return nil, err
	}
	return capnp.ToData(p), nil
}

func (s RedditAccount) SetID(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

func (s RedditAccount) Name() (string, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s RedditAccount) NameBytes() ([]byte, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return nil, err
	}
	return capnp.ToData(p), nil
}

func (s RedditAccount) SetName(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(1, t)
}

func (s RedditAccount) Kind() (string, error) {
	p, err := s.Struct.Pointer(2)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s RedditAccount) KindBytes() ([]byte, error) {
	p, err := s.Struct.Pointer(2)
	if err != nil {
		return nil, err
	}
	return capnp.ToData(p), nil
}

func (s RedditAccount) SetKind(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(2, t)
}

func (s RedditAccount) Data() (AccountData, error) {
	p, err := s.Struct.Pointer(3)
	if err != nil {
		return AccountData{}, err
	}

	ss := capnp.ToStruct(p)

	return AccountData{Struct: ss}, nil
}

func (s RedditAccount) SetData(v AccountData) error {

	return s.Struct.SetPointer(3, v.Struct)
}

// NewData sets the data field to a newly
// allocated AccountData struct, preferring placement in s's segment.
func (s RedditAccount) NewData() (AccountData, error) {

	ss, err := NewAccountData(s.Struct.Segment())
	if err != nil {
		return AccountData{}, err
	}
	err = s.Struct.SetPointer(3, ss)
	return ss, err
}

// RedditAccount_List is a list of RedditAccount.
type RedditAccount_List struct{ capnp.List }

// NewRedditAccount creates a new list of RedditAccount.
func NewRedditAccount_List(s *capnp.Segment, sz int32) (RedditAccount_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4}, sz)
	if err != nil {
		return RedditAccount_List{}, err
	}
	return RedditAccount_List{l}, nil
}

func (s RedditAccount_List) At(i int) RedditAccount           { return RedditAccount{s.List.Struct(i)} }
func (s RedditAccount_List) Set(i int, v RedditAccount) error { return s.List.SetStruct(i, v.Struct) }

// RedditAccount_Promise is a wrapper for a RedditAccount promised by a client call.
type RedditAccount_Promise struct{ *capnp.Pipeline }

func (p RedditAccount_Promise) Struct() (RedditAccount, error) {
	s, err := p.Pipeline.Struct()
	return RedditAccount{s}, err
}

func (p RedditAccount_Promise) Data() AccountData_Promise {
	return AccountData_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
}

type AccountData struct{ capnp.Struct }

func NewAccountData(s *capnp.Segment) (AccountData, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 32, PointerCount: 3})
	if err != nil {
		return AccountData{}, err
	}
	return AccountData{st}, nil
}

func NewRootAccountData(s *capnp.Segment) (AccountData, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 32, PointerCount: 3})
	if err != nil {
		return AccountData{}, err
	}
	return AccountData{st}, nil
}

func ReadRootAccountData(msg *capnp.Message) (AccountData, error) {
	root, err := msg.Root()
	if err != nil {
		return AccountData{}, err
	}
	st := capnp.ToStruct(root)
	return AccountData{st}, nil
}

func (s AccountData) CommentKarma() int64 {
	return int64(s.Struct.Uint64(0))
}

func (s AccountData) SetCommentKarma(v int64) {

	s.Struct.SetUint64(0, uint64(v))
}

func (s AccountData) HasMail() bool {
	return s.Struct.Bit(64)
}

func (s AccountData) SetHasMail(v bool) {

	s.Struct.SetBit(64, v)
}

func (s AccountData) HasModMail() bool {
	return s.Struct.Bit(65)
}

func (s AccountData) SetHasModMail(v bool) {

	s.Struct.SetBit(65, v)
}

func (s AccountData) ID() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s AccountData) IDBytes() ([]byte, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return nil, err
	}
	return capnp.ToData(p), nil
}

func (s AccountData) SetID(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

func (s AccountData) InboxCount() int64 {
	return int64(s.Struct.Uint64(16))
}

func (s AccountData) SetInboxCount(v int64) {

	s.Struct.SetUint64(16, uint64(v))
}

func (s AccountData) IsFriend() bool {
	return s.Struct.Bit(66)
}

func (s AccountData) SetIsFriend(v bool) {

	s.Struct.SetBit(66, v)
}

func (s AccountData) IsGold() bool {
	return s.Struct.Bit(67)
}

func (s AccountData) SetIsGold(v bool) {

	s.Struct.SetBit(67, v)
}

func (s AccountData) LinkKarma() int64 {
	return int64(s.Struct.Uint64(24))
}

func (s AccountData) SetLinkKarma(v int64) {

	s.Struct.SetUint64(24, uint64(v))
}

func (s AccountData) ModHash() (string, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s AccountData) ModHashBytes() ([]byte, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return nil, err
	}
	return capnp.ToData(p), nil
}

func (s AccountData) SetModHash(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(1, t)
}

func (s AccountData) Name() (string, error) {
	p, err := s.Struct.Pointer(2)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s AccountData) NameBytes() ([]byte, error) {
	p, err := s.Struct.Pointer(2)
	if err != nil {
		return nil, err
	}
	return capnp.ToData(p), nil
}

func (s AccountData) SetName(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(2, t)
}

func (s AccountData) Over18() bool {
	return s.Struct.Bit(68)
}

func (s AccountData) SetOver18(v bool) {

	s.Struct.SetBit(68, v)
}

// AccountData_List is a list of AccountData.
type AccountData_List struct{ capnp.List }

// NewAccountData creates a new list of AccountData.
func NewAccountData_List(s *capnp.Segment, sz int32) (AccountData_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 32, PointerCount: 3}, sz)
	if err != nil {
		return AccountData_List{}, err
	}
	return AccountData_List{l}, nil
}

func (s AccountData_List) At(i int) AccountData           { return AccountData{s.List.Struct(i)} }
func (s AccountData_List) Set(i int, v AccountData) error { return s.List.SetStruct(i, v.Struct) }

// AccountData_Promise is a wrapper for a AccountData promised by a client call.
type AccountData_Promise struct{ *capnp.Pipeline }

func (p AccountData_Promise) Struct() (AccountData, error) {
	s, err := p.Pipeline.Struct()
	return AccountData{s}, err
}
