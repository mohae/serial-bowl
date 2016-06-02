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

type CPUInfo struct {
	CPUs []CPU
}

func (d *CPUInfo) Size() (s uint64) {

	{
		l := uint64(len(d.CPUs))

		{

			t := l
			for t >= 0x80 {
				t <<= 7
				s++
			}
			s++

		}
		for k := range d.CPUs {

			{
				s += d.CPUs[k].Size()
			}

		}
	}
	return
}
func (d *CPUInfo) Marshal(buf []byte) ([]byte, error) {
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
		l := uint64(len(d.CPUs))

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
		for k := range d.CPUs {

			{
				nbuf, err := d.CPUs[k].Marshal(buf[i+0:])
				if err != nil {
					return nil, err
				}
				i += uint64(len(nbuf))
			}

		}
	}
	return buf[:i+0], nil
}

func (d *CPUInfo) Unmarshal(buf []byte) (uint64, error) {
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
		if uint64(cap(d.CPUs)) >= l {
			d.CPUs = d.CPUs[:l]
		} else {
			d.CPUs = make([]CPU, l)
		}
		for k := range d.CPUs {

			{
				ni, err := d.CPUs[k].Unmarshal(buf[i+0:])
				if err != nil {
					return 0, err
				}
				i += ni
			}

		}
	}
	return i + 0, nil
}

type CPU struct {
	Processor       int16
	VendorID        string
	CPUFamily       string
	Model           string
	ModelName       string
	Stepping        string
	Microcode       string
	CPUMHz          float32
	CacheSize       string
	PhysicalID      int16
	Siblings        int16
	CoreID          int16
	CPUCores        int16
	ApicID          int16
	InitialApicID   int16
	FPU             string
	FPUException    string
	CPUIDLevel      string
	WP              string
	Flags           string
	BogoMIPS        float32
	CLFlushSize     string
	CacheAlignment  string
	AddressSizes    string
	PowerManagement string
}

func (d *CPU) Size() (s uint64) {

	{
		l := uint64(len(d.VendorID))

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
		l := uint64(len(d.CPUFamily))

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
		l := uint64(len(d.Model))

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
		l := uint64(len(d.ModelName))

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
		l := uint64(len(d.Stepping))

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
		l := uint64(len(d.Microcode))

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
		l := uint64(len(d.CacheSize))

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
		l := uint64(len(d.FPU))

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
		l := uint64(len(d.FPUException))

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
		l := uint64(len(d.CPUIDLevel))

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
		l := uint64(len(d.WP))

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
		l := uint64(len(d.Flags))

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
		l := uint64(len(d.CLFlushSize))

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
		l := uint64(len(d.CacheAlignment))

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
		l := uint64(len(d.AddressSizes))

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
		l := uint64(len(d.PowerManagement))

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
	s += 22
	return
}
func (d *CPU) Marshal(buf []byte) ([]byte, error) {
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

		buf[i+0+0] = byte(d.Processor >> 0)

		buf[i+1+0] = byte(d.Processor >> 8)

	}
	{
		l := uint64(len(d.VendorID))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+2] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+2] = byte(t)
			i++

		}
		copy(buf[i+2:], d.VendorID)
		i += l
	}
	{
		l := uint64(len(d.CPUFamily))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+2] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+2] = byte(t)
			i++

		}
		copy(buf[i+2:], d.CPUFamily)
		i += l
	}
	{
		l := uint64(len(d.Model))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+2] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+2] = byte(t)
			i++

		}
		copy(buf[i+2:], d.Model)
		i += l
	}
	{
		l := uint64(len(d.ModelName))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+2] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+2] = byte(t)
			i++

		}
		copy(buf[i+2:], d.ModelName)
		i += l
	}
	{
		l := uint64(len(d.Stepping))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+2] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+2] = byte(t)
			i++

		}
		copy(buf[i+2:], d.Stepping)
		i += l
	}
	{
		l := uint64(len(d.Microcode))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+2] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+2] = byte(t)
			i++

		}
		copy(buf[i+2:], d.Microcode)
		i += l
	}
	{

		v := *(*uint32)(unsafe.Pointer(&(d.CPUMHz)))

		buf[i+0+2] = byte(v >> 0)

		buf[i+1+2] = byte(v >> 8)

		buf[i+2+2] = byte(v >> 16)

		buf[i+3+2] = byte(v >> 24)

	}
	{
		l := uint64(len(d.CacheSize))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+6] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+6] = byte(t)
			i++

		}
		copy(buf[i+6:], d.CacheSize)
		i += l
	}
	{

		buf[i+0+6] = byte(d.PhysicalID >> 0)

		buf[i+1+6] = byte(d.PhysicalID >> 8)

	}
	{

		buf[i+0+8] = byte(d.Siblings >> 0)

		buf[i+1+8] = byte(d.Siblings >> 8)

	}
	{

		buf[i+0+10] = byte(d.CoreID >> 0)

		buf[i+1+10] = byte(d.CoreID >> 8)

	}
	{

		buf[i+0+12] = byte(d.CPUCores >> 0)

		buf[i+1+12] = byte(d.CPUCores >> 8)

	}
	{

		buf[i+0+14] = byte(d.ApicID >> 0)

		buf[i+1+14] = byte(d.ApicID >> 8)

	}
	{

		buf[i+0+16] = byte(d.InitialApicID >> 0)

		buf[i+1+16] = byte(d.InitialApicID >> 8)

	}
	{
		l := uint64(len(d.FPU))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+18] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+18] = byte(t)
			i++

		}
		copy(buf[i+18:], d.FPU)
		i += l
	}
	{
		l := uint64(len(d.FPUException))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+18] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+18] = byte(t)
			i++

		}
		copy(buf[i+18:], d.FPUException)
		i += l
	}
	{
		l := uint64(len(d.CPUIDLevel))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+18] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+18] = byte(t)
			i++

		}
		copy(buf[i+18:], d.CPUIDLevel)
		i += l
	}
	{
		l := uint64(len(d.WP))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+18] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+18] = byte(t)
			i++

		}
		copy(buf[i+18:], d.WP)
		i += l
	}
	{
		l := uint64(len(d.Flags))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+18] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+18] = byte(t)
			i++

		}
		copy(buf[i+18:], d.Flags)
		i += l
	}
	{

		v := *(*uint32)(unsafe.Pointer(&(d.BogoMIPS)))

		buf[i+0+18] = byte(v >> 0)

		buf[i+1+18] = byte(v >> 8)

		buf[i+2+18] = byte(v >> 16)

		buf[i+3+18] = byte(v >> 24)

	}
	{
		l := uint64(len(d.CLFlushSize))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+22] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+22] = byte(t)
			i++

		}
		copy(buf[i+22:], d.CLFlushSize)
		i += l
	}
	{
		l := uint64(len(d.CacheAlignment))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+22] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+22] = byte(t)
			i++

		}
		copy(buf[i+22:], d.CacheAlignment)
		i += l
	}
	{
		l := uint64(len(d.AddressSizes))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+22] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+22] = byte(t)
			i++

		}
		copy(buf[i+22:], d.AddressSizes)
		i += l
	}
	{
		l := uint64(len(d.PowerManagement))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+22] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+22] = byte(t)
			i++

		}
		copy(buf[i+22:], d.PowerManagement)
		i += l
	}
	return buf[:i+22], nil
}

func (d *CPU) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		d.Processor = 0 | (int16(buf[i+0+0]) << 0) | (int16(buf[i+1+0]) << 8)

	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+2] & 0x7F)
			for buf[i+2]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+2]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.VendorID = string(buf[i+2 : i+2+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+2] & 0x7F)
			for buf[i+2]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+2]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.CPUFamily = string(buf[i+2 : i+2+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+2] & 0x7F)
			for buf[i+2]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+2]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Model = string(buf[i+2 : i+2+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+2] & 0x7F)
			for buf[i+2]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+2]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.ModelName = string(buf[i+2 : i+2+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+2] & 0x7F)
			for buf[i+2]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+2]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Stepping = string(buf[i+2 : i+2+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+2] & 0x7F)
			for buf[i+2]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+2]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Microcode = string(buf[i+2 : i+2+l])
		i += l
	}
	{

		v := 0 | (uint32(buf[i+0+2]) << 0) | (uint32(buf[i+1+2]) << 8) | (uint32(buf[i+2+2]) << 16) | (uint32(buf[i+3+2]) << 24)
		d.CPUMHz = *(*float32)(unsafe.Pointer(&v))

	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+6] & 0x7F)
			for buf[i+6]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+6]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.CacheSize = string(buf[i+6 : i+6+l])
		i += l
	}
	{

		d.PhysicalID = 0 | (int16(buf[i+0+6]) << 0) | (int16(buf[i+1+6]) << 8)

	}
	{

		d.Siblings = 0 | (int16(buf[i+0+8]) << 0) | (int16(buf[i+1+8]) << 8)

	}
	{

		d.CoreID = 0 | (int16(buf[i+0+10]) << 0) | (int16(buf[i+1+10]) << 8)

	}
	{

		d.CPUCores = 0 | (int16(buf[i+0+12]) << 0) | (int16(buf[i+1+12]) << 8)

	}
	{

		d.ApicID = 0 | (int16(buf[i+0+14]) << 0) | (int16(buf[i+1+14]) << 8)

	}
	{

		d.InitialApicID = 0 | (int16(buf[i+0+16]) << 0) | (int16(buf[i+1+16]) << 8)

	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+18] & 0x7F)
			for buf[i+18]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+18]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.FPU = string(buf[i+18 : i+18+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+18] & 0x7F)
			for buf[i+18]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+18]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.FPUException = string(buf[i+18 : i+18+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+18] & 0x7F)
			for buf[i+18]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+18]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.CPUIDLevel = string(buf[i+18 : i+18+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+18] & 0x7F)
			for buf[i+18]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+18]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.WP = string(buf[i+18 : i+18+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+18] & 0x7F)
			for buf[i+18]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+18]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Flags = string(buf[i+18 : i+18+l])
		i += l
	}
	{

		v := 0 | (uint32(buf[i+0+18]) << 0) | (uint32(buf[i+1+18]) << 8) | (uint32(buf[i+2+18]) << 16) | (uint32(buf[i+3+18]) << 24)
		d.BogoMIPS = *(*float32)(unsafe.Pointer(&v))

	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+22] & 0x7F)
			for buf[i+22]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+22]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.CLFlushSize = string(buf[i+22 : i+22+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+22] & 0x7F)
			for buf[i+22]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+22]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.CacheAlignment = string(buf[i+22 : i+22+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+22] & 0x7F)
			for buf[i+22]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+22]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.AddressSizes = string(buf[i+22 : i+22+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+22] & 0x7F)
			for buf[i+22]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+22]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.PowerManagement = string(buf[i+22 : i+22+l])
		i += l
	}
	return i + 22, nil
}
