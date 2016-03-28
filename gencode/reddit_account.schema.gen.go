package gencode

import (
	"io"
	"time"
	"unsafe"
)

var (
	_ = unsafe.Sizeof(0)
	_ = io.ReadFull
	_ = time.Now()
)

type RedditAccount struct {
	ID   string
	Name string
	Kind string
	Data AccountData
}

func (d *RedditAccount) Size() (s uint64) {

	{
		l := uint64(len(d.ID))

		{

			t := l
			for t >= 0x80 {
				t <<= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.Name))

		{

			t := l
			for t >= 0x80 {
				t <<= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.Kind))

		{

			t := l
			for t >= 0x80 {
				t <<= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		s += d.Data.Size()
	}
	return
}
func (d *RedditAccount) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{
		l := uint64(len(d.ID))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		copy(buf[i+0:], d.ID)
		i += l
	}
	{
		l := uint64(len(d.Name))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		copy(buf[i+0:], d.Name)
		i += l
	}
	{
		l := uint64(len(d.Kind))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		copy(buf[i+0:], d.Kind)
		i += l
	}
	{
		nbuf, err := d.Data.Marshal(buf[i+0:])
		if err != nil {
			return nil, err
		}
		i += uint64(len(nbuf))
	}
	return buf[:i+0], nil
}

func (d *RedditAccount) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.ID = string(buf[i+0 : i+0+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Name = string(buf[i+0 : i+0+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Kind = string(buf[i+0 : i+0+l])
		i += l
	}
	{
		ni, err := d.Data.Unmarshal(buf[i+0:])
		if err != nil {
			return 0, err
		}
		i += ni
	}
	return i + 0, nil
}

type AccountData struct {
	CommentKarma int64
	HasMail      bool
	HasModMail   bool
	ID           string
	InboxCount   int64
	IsFriend     bool
	IsGold       bool
	LinkKarma    int64
	ModHash      string
	Name         string
	Over18       bool
}

func (d *AccountData) Size() (s uint64) {

	{
		l := uint64(len(d.ID))

		{

			t := l
			for t >= 0x80 {
				t <<= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.ModHash))

		{

			t := l
			for t >= 0x80 {
				t <<= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.Name))

		{

			t := l
			for t >= 0x80 {
				t <<= 7
				s++
			}
			s++

		}
		s += l
	}
	s += 29
	return
}
func (d *AccountData) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{

		buf[i+0+0] = byte(d.CommentKarma >> 0)

		buf[i+1+0] = byte(d.CommentKarma >> 8)

		buf[i+2+0] = byte(d.CommentKarma >> 16)

		buf[i+3+0] = byte(d.CommentKarma >> 24)

		buf[i+4+0] = byte(d.CommentKarma >> 32)

		buf[i+5+0] = byte(d.CommentKarma >> 40)

		buf[i+6+0] = byte(d.CommentKarma >> 48)

		buf[i+7+0] = byte(d.CommentKarma >> 56)

	}
	{
		if d.HasMail {
			buf[i+8] = 1
		} else {
			buf[i+8] = 0
		}
	}
	{
		if d.HasModMail {
			buf[i+9] = 1
		} else {
			buf[i+9] = 0
		}
	}
	{
		l := uint64(len(d.ID))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+10] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+10] = byte(t)
			i++

		}
		copy(buf[i+10:], d.ID)
		i += l
	}
	{

		buf[i+0+10] = byte(d.InboxCount >> 0)

		buf[i+1+10] = byte(d.InboxCount >> 8)

		buf[i+2+10] = byte(d.InboxCount >> 16)

		buf[i+3+10] = byte(d.InboxCount >> 24)

		buf[i+4+10] = byte(d.InboxCount >> 32)

		buf[i+5+10] = byte(d.InboxCount >> 40)

		buf[i+6+10] = byte(d.InboxCount >> 48)

		buf[i+7+10] = byte(d.InboxCount >> 56)

	}
	{
		if d.IsFriend {
			buf[i+18] = 1
		} else {
			buf[i+18] = 0
		}
	}
	{
		if d.IsGold {
			buf[i+19] = 1
		} else {
			buf[i+19] = 0
		}
	}
	{

		buf[i+0+20] = byte(d.LinkKarma >> 0)

		buf[i+1+20] = byte(d.LinkKarma >> 8)

		buf[i+2+20] = byte(d.LinkKarma >> 16)

		buf[i+3+20] = byte(d.LinkKarma >> 24)

		buf[i+4+20] = byte(d.LinkKarma >> 32)

		buf[i+5+20] = byte(d.LinkKarma >> 40)

		buf[i+6+20] = byte(d.LinkKarma >> 48)

		buf[i+7+20] = byte(d.LinkKarma >> 56)

	}
	{
		l := uint64(len(d.ModHash))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+28] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+28] = byte(t)
			i++

		}
		copy(buf[i+28:], d.ModHash)
		i += l
	}
	{
		l := uint64(len(d.Name))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+28] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+28] = byte(t)
			i++

		}
		copy(buf[i+28:], d.Name)
		i += l
	}
	{
		if d.Over18 {
			buf[i+28] = 1
		} else {
			buf[i+28] = 0
		}
	}
	return buf[:i+29], nil
}

func (d *AccountData) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		d.CommentKarma = 0 | (int64(buf[i+0+0]) << 0) | (int64(buf[i+1+0]) << 8) | (int64(buf[i+2+0]) << 16) | (int64(buf[i+3+0]) << 24) | (int64(buf[i+4+0]) << 32) | (int64(buf[i+5+0]) << 40) | (int64(buf[i+6+0]) << 48) | (int64(buf[i+7+0]) << 56)

	}
	{
		d.HasMail = buf[i+8] == 1
	}
	{
		d.HasModMail = buf[i+9] == 1
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+10] & 0x7F)
			for buf[i+10]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+10]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.ID = string(buf[i+10 : i+10+l])
		i += l
	}
	{

		d.InboxCount = 0 | (int64(buf[i+0+10]) << 0) | (int64(buf[i+1+10]) << 8) | (int64(buf[i+2+10]) << 16) | (int64(buf[i+3+10]) << 24) | (int64(buf[i+4+10]) << 32) | (int64(buf[i+5+10]) << 40) | (int64(buf[i+6+10]) << 48) | (int64(buf[i+7+10]) << 56)

	}
	{
		d.IsFriend = buf[i+18] == 1
	}
	{
		d.IsGold = buf[i+19] == 1
	}
	{

		d.LinkKarma = 0 | (int64(buf[i+0+20]) << 0) | (int64(buf[i+1+20]) << 8) | (int64(buf[i+2+20]) << 16) | (int64(buf[i+3+20]) << 24) | (int64(buf[i+4+20]) << 32) | (int64(buf[i+5+20]) << 40) | (int64(buf[i+6+20]) << 48) | (int64(buf[i+7+20]) << 56)

	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+28] & 0x7F)
			for buf[i+28]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+28]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.ModHash = string(buf[i+28 : i+28+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+28] & 0x7F)
			for buf[i+28]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+28]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Name = string(buf[i+28 : i+28+l])
		i += l
	}
	{
		d.Over18 = buf[i+28] == 1
	}
	return i + 29, nil
}
