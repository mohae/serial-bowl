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

type Message struct {
	ID     []byte
	DestID uint32
	Type   int8
	Kind   int16
	Data   []byte
}

func (d *Message) Size() (s uint64) {

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
		l := uint64(len(d.Data))

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
	s += 7
	return
}
func (d *Message) Marshal(buf []byte) ([]byte, error) {
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
		copy(buf[i:], d.ID)
		i += l
	}
	{

		buf[i+0+0] = byte(d.DestID >> 0)

		buf[i+1+0] = byte(d.DestID >> 8)

		buf[i+2+0] = byte(d.DestID >> 16)

		buf[i+3+0] = byte(d.DestID >> 24)

	}
	{

		buf[i+0+4] = byte(d.Type >> 0)

	}
	{

		buf[i+0+5] = byte(d.Kind >> 0)

		buf[i+1+5] = byte(d.Kind >> 8)

	}
	{
		l := uint64(len(d.Data))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+7] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+7] = byte(t)
			i++

		}
		copy(buf[i:], d.Data)
		i += l
	}
	return buf[:i+7], nil
}

func (d *Message) Unmarshal(buf []byte) (uint64, error) {
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
		if uint64(cap(d.ID)) >= l {
			d.ID = d.ID[:l]
		} else {
			d.ID = make([]byte, l)
		}
		copy(d.ID, buf[i:])
		i += l
	}
	{

		d.DestID = 0 | (uint32(buf[i+0+0]) << 0) | (uint32(buf[i+1+0]) << 8) | (uint32(buf[i+2+0]) << 16) | (uint32(buf[i+3+0]) << 24)

	}
	{

		d.Type = 0 | (int8(buf[i+0+4]) << 0)

	}
	{

		d.Kind = 0 | (int16(buf[i+0+5]) << 0) | (int16(buf[i+1+5]) << 8)

	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+7] & 0x7F)
			for buf[i+7]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+7]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.Data)) >= l {
			d.Data = d.Data[:l]
		} else {
			d.Data = make([]byte, l)
		}
		copy(d.Data, buf[i:])
		i += l
	}
	return i + 7, nil
}
