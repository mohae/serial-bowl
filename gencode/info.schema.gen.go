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

type MemInfo struct {
	MemTotal          int64
	MemFree           int64
	MemAvailable      int64
	Buffers           int64
	Cached            int64
	SwapCached        int64
	Active            int64
	Inactive          int64
	ActiveAnon        int64
	InactiveAnon      int64
	ActiveFile        int64
	InactiveFile      int64
	Unevictable       int64
	Mlocked           int64
	SwapTotal         int64
	SwapFree          int64
	Dirty             int64
	Writeback         int64
	AnonPages         int64
	Mapped            int64
	Shmem             int64
	Slab              int64
	SReclaimable      int64
	SUnreclaim        int64
	KernelStack       int64
	NFSUnstable       int64
	Bounce            int64
	WritebackTmp      int64
	CommitLimit       int64
	VmallocTotal      int64
	VmallocUsed       int64
	VmallocChunk      int64
	HardwareCorrupted int64
	AnonHugePages     int64
	HugePagesTotal    int64
	HugePagesFree     int64
	HugePagesRsvd     int64
	Hugepagesize      int64
	DirectMap4k       int64
	DirectMap2M       int64
}

func (d *MemInfo) Size() (s uint64) {

	s += 320
	return
}
func (d *MemInfo) Marshal(buf []byte) ([]byte, error) {
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

		buf[0+0] = byte(d.MemTotal >> 0)

		buf[1+0] = byte(d.MemTotal >> 8)

		buf[2+0] = byte(d.MemTotal >> 16)

		buf[3+0] = byte(d.MemTotal >> 24)

		buf[4+0] = byte(d.MemTotal >> 32)

		buf[5+0] = byte(d.MemTotal >> 40)

		buf[6+0] = byte(d.MemTotal >> 48)

		buf[7+0] = byte(d.MemTotal >> 56)

	}
	{

		buf[0+8] = byte(d.MemFree >> 0)

		buf[1+8] = byte(d.MemFree >> 8)

		buf[2+8] = byte(d.MemFree >> 16)

		buf[3+8] = byte(d.MemFree >> 24)

		buf[4+8] = byte(d.MemFree >> 32)

		buf[5+8] = byte(d.MemFree >> 40)

		buf[6+8] = byte(d.MemFree >> 48)

		buf[7+8] = byte(d.MemFree >> 56)

	}
	{

		buf[0+16] = byte(d.MemAvailable >> 0)

		buf[1+16] = byte(d.MemAvailable >> 8)

		buf[2+16] = byte(d.MemAvailable >> 16)

		buf[3+16] = byte(d.MemAvailable >> 24)

		buf[4+16] = byte(d.MemAvailable >> 32)

		buf[5+16] = byte(d.MemAvailable >> 40)

		buf[6+16] = byte(d.MemAvailable >> 48)

		buf[7+16] = byte(d.MemAvailable >> 56)

	}
	{

		buf[0+24] = byte(d.Buffers >> 0)

		buf[1+24] = byte(d.Buffers >> 8)

		buf[2+24] = byte(d.Buffers >> 16)

		buf[3+24] = byte(d.Buffers >> 24)

		buf[4+24] = byte(d.Buffers >> 32)

		buf[5+24] = byte(d.Buffers >> 40)

		buf[6+24] = byte(d.Buffers >> 48)

		buf[7+24] = byte(d.Buffers >> 56)

	}
	{

		buf[0+32] = byte(d.Cached >> 0)

		buf[1+32] = byte(d.Cached >> 8)

		buf[2+32] = byte(d.Cached >> 16)

		buf[3+32] = byte(d.Cached >> 24)

		buf[4+32] = byte(d.Cached >> 32)

		buf[5+32] = byte(d.Cached >> 40)

		buf[6+32] = byte(d.Cached >> 48)

		buf[7+32] = byte(d.Cached >> 56)

	}
	{

		buf[0+40] = byte(d.SwapCached >> 0)

		buf[1+40] = byte(d.SwapCached >> 8)

		buf[2+40] = byte(d.SwapCached >> 16)

		buf[3+40] = byte(d.SwapCached >> 24)

		buf[4+40] = byte(d.SwapCached >> 32)

		buf[5+40] = byte(d.SwapCached >> 40)

		buf[6+40] = byte(d.SwapCached >> 48)

		buf[7+40] = byte(d.SwapCached >> 56)

	}
	{

		buf[0+48] = byte(d.Active >> 0)

		buf[1+48] = byte(d.Active >> 8)

		buf[2+48] = byte(d.Active >> 16)

		buf[3+48] = byte(d.Active >> 24)

		buf[4+48] = byte(d.Active >> 32)

		buf[5+48] = byte(d.Active >> 40)

		buf[6+48] = byte(d.Active >> 48)

		buf[7+48] = byte(d.Active >> 56)

	}
	{

		buf[0+56] = byte(d.Inactive >> 0)

		buf[1+56] = byte(d.Inactive >> 8)

		buf[2+56] = byte(d.Inactive >> 16)

		buf[3+56] = byte(d.Inactive >> 24)

		buf[4+56] = byte(d.Inactive >> 32)

		buf[5+56] = byte(d.Inactive >> 40)

		buf[6+56] = byte(d.Inactive >> 48)

		buf[7+56] = byte(d.Inactive >> 56)

	}
	{

		buf[0+64] = byte(d.ActiveAnon >> 0)

		buf[1+64] = byte(d.ActiveAnon >> 8)

		buf[2+64] = byte(d.ActiveAnon >> 16)

		buf[3+64] = byte(d.ActiveAnon >> 24)

		buf[4+64] = byte(d.ActiveAnon >> 32)

		buf[5+64] = byte(d.ActiveAnon >> 40)

		buf[6+64] = byte(d.ActiveAnon >> 48)

		buf[7+64] = byte(d.ActiveAnon >> 56)

	}
	{

		buf[0+72] = byte(d.InactiveAnon >> 0)

		buf[1+72] = byte(d.InactiveAnon >> 8)

		buf[2+72] = byte(d.InactiveAnon >> 16)

		buf[3+72] = byte(d.InactiveAnon >> 24)

		buf[4+72] = byte(d.InactiveAnon >> 32)

		buf[5+72] = byte(d.InactiveAnon >> 40)

		buf[6+72] = byte(d.InactiveAnon >> 48)

		buf[7+72] = byte(d.InactiveAnon >> 56)

	}
	{

		buf[0+80] = byte(d.ActiveFile >> 0)

		buf[1+80] = byte(d.ActiveFile >> 8)

		buf[2+80] = byte(d.ActiveFile >> 16)

		buf[3+80] = byte(d.ActiveFile >> 24)

		buf[4+80] = byte(d.ActiveFile >> 32)

		buf[5+80] = byte(d.ActiveFile >> 40)

		buf[6+80] = byte(d.ActiveFile >> 48)

		buf[7+80] = byte(d.ActiveFile >> 56)

	}
	{

		buf[0+88] = byte(d.InactiveFile >> 0)

		buf[1+88] = byte(d.InactiveFile >> 8)

		buf[2+88] = byte(d.InactiveFile >> 16)

		buf[3+88] = byte(d.InactiveFile >> 24)

		buf[4+88] = byte(d.InactiveFile >> 32)

		buf[5+88] = byte(d.InactiveFile >> 40)

		buf[6+88] = byte(d.InactiveFile >> 48)

		buf[7+88] = byte(d.InactiveFile >> 56)

	}
	{

		buf[0+96] = byte(d.Unevictable >> 0)

		buf[1+96] = byte(d.Unevictable >> 8)

		buf[2+96] = byte(d.Unevictable >> 16)

		buf[3+96] = byte(d.Unevictable >> 24)

		buf[4+96] = byte(d.Unevictable >> 32)

		buf[5+96] = byte(d.Unevictable >> 40)

		buf[6+96] = byte(d.Unevictable >> 48)

		buf[7+96] = byte(d.Unevictable >> 56)

	}
	{

		buf[0+104] = byte(d.Mlocked >> 0)

		buf[1+104] = byte(d.Mlocked >> 8)

		buf[2+104] = byte(d.Mlocked >> 16)

		buf[3+104] = byte(d.Mlocked >> 24)

		buf[4+104] = byte(d.Mlocked >> 32)

		buf[5+104] = byte(d.Mlocked >> 40)

		buf[6+104] = byte(d.Mlocked >> 48)

		buf[7+104] = byte(d.Mlocked >> 56)

	}
	{

		buf[0+112] = byte(d.SwapTotal >> 0)

		buf[1+112] = byte(d.SwapTotal >> 8)

		buf[2+112] = byte(d.SwapTotal >> 16)

		buf[3+112] = byte(d.SwapTotal >> 24)

		buf[4+112] = byte(d.SwapTotal >> 32)

		buf[5+112] = byte(d.SwapTotal >> 40)

		buf[6+112] = byte(d.SwapTotal >> 48)

		buf[7+112] = byte(d.SwapTotal >> 56)

	}
	{

		buf[0+120] = byte(d.SwapFree >> 0)

		buf[1+120] = byte(d.SwapFree >> 8)

		buf[2+120] = byte(d.SwapFree >> 16)

		buf[3+120] = byte(d.SwapFree >> 24)

		buf[4+120] = byte(d.SwapFree >> 32)

		buf[5+120] = byte(d.SwapFree >> 40)

		buf[6+120] = byte(d.SwapFree >> 48)

		buf[7+120] = byte(d.SwapFree >> 56)

	}
	{

		buf[0+128] = byte(d.Dirty >> 0)

		buf[1+128] = byte(d.Dirty >> 8)

		buf[2+128] = byte(d.Dirty >> 16)

		buf[3+128] = byte(d.Dirty >> 24)

		buf[4+128] = byte(d.Dirty >> 32)

		buf[5+128] = byte(d.Dirty >> 40)

		buf[6+128] = byte(d.Dirty >> 48)

		buf[7+128] = byte(d.Dirty >> 56)

	}
	{

		buf[0+136] = byte(d.Writeback >> 0)

		buf[1+136] = byte(d.Writeback >> 8)

		buf[2+136] = byte(d.Writeback >> 16)

		buf[3+136] = byte(d.Writeback >> 24)

		buf[4+136] = byte(d.Writeback >> 32)

		buf[5+136] = byte(d.Writeback >> 40)

		buf[6+136] = byte(d.Writeback >> 48)

		buf[7+136] = byte(d.Writeback >> 56)

	}
	{

		buf[0+144] = byte(d.AnonPages >> 0)

		buf[1+144] = byte(d.AnonPages >> 8)

		buf[2+144] = byte(d.AnonPages >> 16)

		buf[3+144] = byte(d.AnonPages >> 24)

		buf[4+144] = byte(d.AnonPages >> 32)

		buf[5+144] = byte(d.AnonPages >> 40)

		buf[6+144] = byte(d.AnonPages >> 48)

		buf[7+144] = byte(d.AnonPages >> 56)

	}
	{

		buf[0+152] = byte(d.Mapped >> 0)

		buf[1+152] = byte(d.Mapped >> 8)

		buf[2+152] = byte(d.Mapped >> 16)

		buf[3+152] = byte(d.Mapped >> 24)

		buf[4+152] = byte(d.Mapped >> 32)

		buf[5+152] = byte(d.Mapped >> 40)

		buf[6+152] = byte(d.Mapped >> 48)

		buf[7+152] = byte(d.Mapped >> 56)

	}
	{

		buf[0+160] = byte(d.Shmem >> 0)

		buf[1+160] = byte(d.Shmem >> 8)

		buf[2+160] = byte(d.Shmem >> 16)

		buf[3+160] = byte(d.Shmem >> 24)

		buf[4+160] = byte(d.Shmem >> 32)

		buf[5+160] = byte(d.Shmem >> 40)

		buf[6+160] = byte(d.Shmem >> 48)

		buf[7+160] = byte(d.Shmem >> 56)

	}
	{

		buf[0+168] = byte(d.Slab >> 0)

		buf[1+168] = byte(d.Slab >> 8)

		buf[2+168] = byte(d.Slab >> 16)

		buf[3+168] = byte(d.Slab >> 24)

		buf[4+168] = byte(d.Slab >> 32)

		buf[5+168] = byte(d.Slab >> 40)

		buf[6+168] = byte(d.Slab >> 48)

		buf[7+168] = byte(d.Slab >> 56)

	}
	{

		buf[0+176] = byte(d.SReclaimable >> 0)

		buf[1+176] = byte(d.SReclaimable >> 8)

		buf[2+176] = byte(d.SReclaimable >> 16)

		buf[3+176] = byte(d.SReclaimable >> 24)

		buf[4+176] = byte(d.SReclaimable >> 32)

		buf[5+176] = byte(d.SReclaimable >> 40)

		buf[6+176] = byte(d.SReclaimable >> 48)

		buf[7+176] = byte(d.SReclaimable >> 56)

	}
	{

		buf[0+184] = byte(d.SUnreclaim >> 0)

		buf[1+184] = byte(d.SUnreclaim >> 8)

		buf[2+184] = byte(d.SUnreclaim >> 16)

		buf[3+184] = byte(d.SUnreclaim >> 24)

		buf[4+184] = byte(d.SUnreclaim >> 32)

		buf[5+184] = byte(d.SUnreclaim >> 40)

		buf[6+184] = byte(d.SUnreclaim >> 48)

		buf[7+184] = byte(d.SUnreclaim >> 56)

	}
	{

		buf[0+192] = byte(d.KernelStack >> 0)

		buf[1+192] = byte(d.KernelStack >> 8)

		buf[2+192] = byte(d.KernelStack >> 16)

		buf[3+192] = byte(d.KernelStack >> 24)

		buf[4+192] = byte(d.KernelStack >> 32)

		buf[5+192] = byte(d.KernelStack >> 40)

		buf[6+192] = byte(d.KernelStack >> 48)

		buf[7+192] = byte(d.KernelStack >> 56)

	}
	{

		buf[0+200] = byte(d.NFSUnstable >> 0)

		buf[1+200] = byte(d.NFSUnstable >> 8)

		buf[2+200] = byte(d.NFSUnstable >> 16)

		buf[3+200] = byte(d.NFSUnstable >> 24)

		buf[4+200] = byte(d.NFSUnstable >> 32)

		buf[5+200] = byte(d.NFSUnstable >> 40)

		buf[6+200] = byte(d.NFSUnstable >> 48)

		buf[7+200] = byte(d.NFSUnstable >> 56)

	}
	{

		buf[0+208] = byte(d.Bounce >> 0)

		buf[1+208] = byte(d.Bounce >> 8)

		buf[2+208] = byte(d.Bounce >> 16)

		buf[3+208] = byte(d.Bounce >> 24)

		buf[4+208] = byte(d.Bounce >> 32)

		buf[5+208] = byte(d.Bounce >> 40)

		buf[6+208] = byte(d.Bounce >> 48)

		buf[7+208] = byte(d.Bounce >> 56)

	}
	{

		buf[0+216] = byte(d.WritebackTmp >> 0)

		buf[1+216] = byte(d.WritebackTmp >> 8)

		buf[2+216] = byte(d.WritebackTmp >> 16)

		buf[3+216] = byte(d.WritebackTmp >> 24)

		buf[4+216] = byte(d.WritebackTmp >> 32)

		buf[5+216] = byte(d.WritebackTmp >> 40)

		buf[6+216] = byte(d.WritebackTmp >> 48)

		buf[7+216] = byte(d.WritebackTmp >> 56)

	}
	{

		buf[0+224] = byte(d.CommitLimit >> 0)

		buf[1+224] = byte(d.CommitLimit >> 8)

		buf[2+224] = byte(d.CommitLimit >> 16)

		buf[3+224] = byte(d.CommitLimit >> 24)

		buf[4+224] = byte(d.CommitLimit >> 32)

		buf[5+224] = byte(d.CommitLimit >> 40)

		buf[6+224] = byte(d.CommitLimit >> 48)

		buf[7+224] = byte(d.CommitLimit >> 56)

	}
	{

		buf[0+232] = byte(d.VmallocTotal >> 0)

		buf[1+232] = byte(d.VmallocTotal >> 8)

		buf[2+232] = byte(d.VmallocTotal >> 16)

		buf[3+232] = byte(d.VmallocTotal >> 24)

		buf[4+232] = byte(d.VmallocTotal >> 32)

		buf[5+232] = byte(d.VmallocTotal >> 40)

		buf[6+232] = byte(d.VmallocTotal >> 48)

		buf[7+232] = byte(d.VmallocTotal >> 56)

	}
	{

		buf[0+240] = byte(d.VmallocUsed >> 0)

		buf[1+240] = byte(d.VmallocUsed >> 8)

		buf[2+240] = byte(d.VmallocUsed >> 16)

		buf[3+240] = byte(d.VmallocUsed >> 24)

		buf[4+240] = byte(d.VmallocUsed >> 32)

		buf[5+240] = byte(d.VmallocUsed >> 40)

		buf[6+240] = byte(d.VmallocUsed >> 48)

		buf[7+240] = byte(d.VmallocUsed >> 56)

	}
	{

		buf[0+248] = byte(d.VmallocChunk >> 0)

		buf[1+248] = byte(d.VmallocChunk >> 8)

		buf[2+248] = byte(d.VmallocChunk >> 16)

		buf[3+248] = byte(d.VmallocChunk >> 24)

		buf[4+248] = byte(d.VmallocChunk >> 32)

		buf[5+248] = byte(d.VmallocChunk >> 40)

		buf[6+248] = byte(d.VmallocChunk >> 48)

		buf[7+248] = byte(d.VmallocChunk >> 56)

	}
	{

		buf[0+256] = byte(d.HardwareCorrupted >> 0)

		buf[1+256] = byte(d.HardwareCorrupted >> 8)

		buf[2+256] = byte(d.HardwareCorrupted >> 16)

		buf[3+256] = byte(d.HardwareCorrupted >> 24)

		buf[4+256] = byte(d.HardwareCorrupted >> 32)

		buf[5+256] = byte(d.HardwareCorrupted >> 40)

		buf[6+256] = byte(d.HardwareCorrupted >> 48)

		buf[7+256] = byte(d.HardwareCorrupted >> 56)

	}
	{

		buf[0+264] = byte(d.AnonHugePages >> 0)

		buf[1+264] = byte(d.AnonHugePages >> 8)

		buf[2+264] = byte(d.AnonHugePages >> 16)

		buf[3+264] = byte(d.AnonHugePages >> 24)

		buf[4+264] = byte(d.AnonHugePages >> 32)

		buf[5+264] = byte(d.AnonHugePages >> 40)

		buf[6+264] = byte(d.AnonHugePages >> 48)

		buf[7+264] = byte(d.AnonHugePages >> 56)

	}
	{

		buf[0+272] = byte(d.HugePagesTotal >> 0)

		buf[1+272] = byte(d.HugePagesTotal >> 8)

		buf[2+272] = byte(d.HugePagesTotal >> 16)

		buf[3+272] = byte(d.HugePagesTotal >> 24)

		buf[4+272] = byte(d.HugePagesTotal >> 32)

		buf[5+272] = byte(d.HugePagesTotal >> 40)

		buf[6+272] = byte(d.HugePagesTotal >> 48)

		buf[7+272] = byte(d.HugePagesTotal >> 56)

	}
	{

		buf[0+280] = byte(d.HugePagesFree >> 0)

		buf[1+280] = byte(d.HugePagesFree >> 8)

		buf[2+280] = byte(d.HugePagesFree >> 16)

		buf[3+280] = byte(d.HugePagesFree >> 24)

		buf[4+280] = byte(d.HugePagesFree >> 32)

		buf[5+280] = byte(d.HugePagesFree >> 40)

		buf[6+280] = byte(d.HugePagesFree >> 48)

		buf[7+280] = byte(d.HugePagesFree >> 56)

	}
	{

		buf[0+288] = byte(d.HugePagesRsvd >> 0)

		buf[1+288] = byte(d.HugePagesRsvd >> 8)

		buf[2+288] = byte(d.HugePagesRsvd >> 16)

		buf[3+288] = byte(d.HugePagesRsvd >> 24)

		buf[4+288] = byte(d.HugePagesRsvd >> 32)

		buf[5+288] = byte(d.HugePagesRsvd >> 40)

		buf[6+288] = byte(d.HugePagesRsvd >> 48)

		buf[7+288] = byte(d.HugePagesRsvd >> 56)

	}
	{

		buf[0+296] = byte(d.Hugepagesize >> 0)

		buf[1+296] = byte(d.Hugepagesize >> 8)

		buf[2+296] = byte(d.Hugepagesize >> 16)

		buf[3+296] = byte(d.Hugepagesize >> 24)

		buf[4+296] = byte(d.Hugepagesize >> 32)

		buf[5+296] = byte(d.Hugepagesize >> 40)

		buf[6+296] = byte(d.Hugepagesize >> 48)

		buf[7+296] = byte(d.Hugepagesize >> 56)

	}
	{

		buf[0+304] = byte(d.DirectMap4k >> 0)

		buf[1+304] = byte(d.DirectMap4k >> 8)

		buf[2+304] = byte(d.DirectMap4k >> 16)

		buf[3+304] = byte(d.DirectMap4k >> 24)

		buf[4+304] = byte(d.DirectMap4k >> 32)

		buf[5+304] = byte(d.DirectMap4k >> 40)

		buf[6+304] = byte(d.DirectMap4k >> 48)

		buf[7+304] = byte(d.DirectMap4k >> 56)

	}
	{

		buf[0+312] = byte(d.DirectMap2M >> 0)

		buf[1+312] = byte(d.DirectMap2M >> 8)

		buf[2+312] = byte(d.DirectMap2M >> 16)

		buf[3+312] = byte(d.DirectMap2M >> 24)

		buf[4+312] = byte(d.DirectMap2M >> 32)

		buf[5+312] = byte(d.DirectMap2M >> 40)

		buf[6+312] = byte(d.DirectMap2M >> 48)

		buf[7+312] = byte(d.DirectMap2M >> 56)

	}
	return buf[:i+320], nil
}

func (d *MemInfo) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		d.MemTotal = 0 | (int64(buf[0+0]) << 0) | (int64(buf[1+0]) << 8) | (int64(buf[2+0]) << 16) | (int64(buf[3+0]) << 24) | (int64(buf[4+0]) << 32) | (int64(buf[5+0]) << 40) | (int64(buf[6+0]) << 48) | (int64(buf[7+0]) << 56)

	}
	{

		d.MemFree = 0 | (int64(buf[0+8]) << 0) | (int64(buf[1+8]) << 8) | (int64(buf[2+8]) << 16) | (int64(buf[3+8]) << 24) | (int64(buf[4+8]) << 32) | (int64(buf[5+8]) << 40) | (int64(buf[6+8]) << 48) | (int64(buf[7+8]) << 56)

	}
	{

		d.MemAvailable = 0 | (int64(buf[0+16]) << 0) | (int64(buf[1+16]) << 8) | (int64(buf[2+16]) << 16) | (int64(buf[3+16]) << 24) | (int64(buf[4+16]) << 32) | (int64(buf[5+16]) << 40) | (int64(buf[6+16]) << 48) | (int64(buf[7+16]) << 56)

	}
	{

		d.Buffers = 0 | (int64(buf[0+24]) << 0) | (int64(buf[1+24]) << 8) | (int64(buf[2+24]) << 16) | (int64(buf[3+24]) << 24) | (int64(buf[4+24]) << 32) | (int64(buf[5+24]) << 40) | (int64(buf[6+24]) << 48) | (int64(buf[7+24]) << 56)

	}
	{

		d.Cached = 0 | (int64(buf[0+32]) << 0) | (int64(buf[1+32]) << 8) | (int64(buf[2+32]) << 16) | (int64(buf[3+32]) << 24) | (int64(buf[4+32]) << 32) | (int64(buf[5+32]) << 40) | (int64(buf[6+32]) << 48) | (int64(buf[7+32]) << 56)

	}
	{

		d.SwapCached = 0 | (int64(buf[0+40]) << 0) | (int64(buf[1+40]) << 8) | (int64(buf[2+40]) << 16) | (int64(buf[3+40]) << 24) | (int64(buf[4+40]) << 32) | (int64(buf[5+40]) << 40) | (int64(buf[6+40]) << 48) | (int64(buf[7+40]) << 56)

	}
	{

		d.Active = 0 | (int64(buf[0+48]) << 0) | (int64(buf[1+48]) << 8) | (int64(buf[2+48]) << 16) | (int64(buf[3+48]) << 24) | (int64(buf[4+48]) << 32) | (int64(buf[5+48]) << 40) | (int64(buf[6+48]) << 48) | (int64(buf[7+48]) << 56)

	}
	{

		d.Inactive = 0 | (int64(buf[0+56]) << 0) | (int64(buf[1+56]) << 8) | (int64(buf[2+56]) << 16) | (int64(buf[3+56]) << 24) | (int64(buf[4+56]) << 32) | (int64(buf[5+56]) << 40) | (int64(buf[6+56]) << 48) | (int64(buf[7+56]) << 56)

	}
	{

		d.ActiveAnon = 0 | (int64(buf[0+64]) << 0) | (int64(buf[1+64]) << 8) | (int64(buf[2+64]) << 16) | (int64(buf[3+64]) << 24) | (int64(buf[4+64]) << 32) | (int64(buf[5+64]) << 40) | (int64(buf[6+64]) << 48) | (int64(buf[7+64]) << 56)

	}
	{

		d.InactiveAnon = 0 | (int64(buf[0+72]) << 0) | (int64(buf[1+72]) << 8) | (int64(buf[2+72]) << 16) | (int64(buf[3+72]) << 24) | (int64(buf[4+72]) << 32) | (int64(buf[5+72]) << 40) | (int64(buf[6+72]) << 48) | (int64(buf[7+72]) << 56)

	}
	{

		d.ActiveFile = 0 | (int64(buf[0+80]) << 0) | (int64(buf[1+80]) << 8) | (int64(buf[2+80]) << 16) | (int64(buf[3+80]) << 24) | (int64(buf[4+80]) << 32) | (int64(buf[5+80]) << 40) | (int64(buf[6+80]) << 48) | (int64(buf[7+80]) << 56)

	}
	{

		d.InactiveFile = 0 | (int64(buf[0+88]) << 0) | (int64(buf[1+88]) << 8) | (int64(buf[2+88]) << 16) | (int64(buf[3+88]) << 24) | (int64(buf[4+88]) << 32) | (int64(buf[5+88]) << 40) | (int64(buf[6+88]) << 48) | (int64(buf[7+88]) << 56)

	}
	{

		d.Unevictable = 0 | (int64(buf[0+96]) << 0) | (int64(buf[1+96]) << 8) | (int64(buf[2+96]) << 16) | (int64(buf[3+96]) << 24) | (int64(buf[4+96]) << 32) | (int64(buf[5+96]) << 40) | (int64(buf[6+96]) << 48) | (int64(buf[7+96]) << 56)

	}
	{

		d.Mlocked = 0 | (int64(buf[0+104]) << 0) | (int64(buf[1+104]) << 8) | (int64(buf[2+104]) << 16) | (int64(buf[3+104]) << 24) | (int64(buf[4+104]) << 32) | (int64(buf[5+104]) << 40) | (int64(buf[6+104]) << 48) | (int64(buf[7+104]) << 56)

	}
	{

		d.SwapTotal = 0 | (int64(buf[0+112]) << 0) | (int64(buf[1+112]) << 8) | (int64(buf[2+112]) << 16) | (int64(buf[3+112]) << 24) | (int64(buf[4+112]) << 32) | (int64(buf[5+112]) << 40) | (int64(buf[6+112]) << 48) | (int64(buf[7+112]) << 56)

	}
	{

		d.SwapFree = 0 | (int64(buf[0+120]) << 0) | (int64(buf[1+120]) << 8) | (int64(buf[2+120]) << 16) | (int64(buf[3+120]) << 24) | (int64(buf[4+120]) << 32) | (int64(buf[5+120]) << 40) | (int64(buf[6+120]) << 48) | (int64(buf[7+120]) << 56)

	}
	{

		d.Dirty = 0 | (int64(buf[0+128]) << 0) | (int64(buf[1+128]) << 8) | (int64(buf[2+128]) << 16) | (int64(buf[3+128]) << 24) | (int64(buf[4+128]) << 32) | (int64(buf[5+128]) << 40) | (int64(buf[6+128]) << 48) | (int64(buf[7+128]) << 56)

	}
	{

		d.Writeback = 0 | (int64(buf[0+136]) << 0) | (int64(buf[1+136]) << 8) | (int64(buf[2+136]) << 16) | (int64(buf[3+136]) << 24) | (int64(buf[4+136]) << 32) | (int64(buf[5+136]) << 40) | (int64(buf[6+136]) << 48) | (int64(buf[7+136]) << 56)

	}
	{

		d.AnonPages = 0 | (int64(buf[0+144]) << 0) | (int64(buf[1+144]) << 8) | (int64(buf[2+144]) << 16) | (int64(buf[3+144]) << 24) | (int64(buf[4+144]) << 32) | (int64(buf[5+144]) << 40) | (int64(buf[6+144]) << 48) | (int64(buf[7+144]) << 56)

	}
	{

		d.Mapped = 0 | (int64(buf[0+152]) << 0) | (int64(buf[1+152]) << 8) | (int64(buf[2+152]) << 16) | (int64(buf[3+152]) << 24) | (int64(buf[4+152]) << 32) | (int64(buf[5+152]) << 40) | (int64(buf[6+152]) << 48) | (int64(buf[7+152]) << 56)

	}
	{

		d.Shmem = 0 | (int64(buf[0+160]) << 0) | (int64(buf[1+160]) << 8) | (int64(buf[2+160]) << 16) | (int64(buf[3+160]) << 24) | (int64(buf[4+160]) << 32) | (int64(buf[5+160]) << 40) | (int64(buf[6+160]) << 48) | (int64(buf[7+160]) << 56)

	}
	{

		d.Slab = 0 | (int64(buf[0+168]) << 0) | (int64(buf[1+168]) << 8) | (int64(buf[2+168]) << 16) | (int64(buf[3+168]) << 24) | (int64(buf[4+168]) << 32) | (int64(buf[5+168]) << 40) | (int64(buf[6+168]) << 48) | (int64(buf[7+168]) << 56)

	}
	{

		d.SReclaimable = 0 | (int64(buf[0+176]) << 0) | (int64(buf[1+176]) << 8) | (int64(buf[2+176]) << 16) | (int64(buf[3+176]) << 24) | (int64(buf[4+176]) << 32) | (int64(buf[5+176]) << 40) | (int64(buf[6+176]) << 48) | (int64(buf[7+176]) << 56)

	}
	{

		d.SUnreclaim = 0 | (int64(buf[0+184]) << 0) | (int64(buf[1+184]) << 8) | (int64(buf[2+184]) << 16) | (int64(buf[3+184]) << 24) | (int64(buf[4+184]) << 32) | (int64(buf[5+184]) << 40) | (int64(buf[6+184]) << 48) | (int64(buf[7+184]) << 56)

	}
	{

		d.KernelStack = 0 | (int64(buf[0+192]) << 0) | (int64(buf[1+192]) << 8) | (int64(buf[2+192]) << 16) | (int64(buf[3+192]) << 24) | (int64(buf[4+192]) << 32) | (int64(buf[5+192]) << 40) | (int64(buf[6+192]) << 48) | (int64(buf[7+192]) << 56)

	}
	{

		d.NFSUnstable = 0 | (int64(buf[0+200]) << 0) | (int64(buf[1+200]) << 8) | (int64(buf[2+200]) << 16) | (int64(buf[3+200]) << 24) | (int64(buf[4+200]) << 32) | (int64(buf[5+200]) << 40) | (int64(buf[6+200]) << 48) | (int64(buf[7+200]) << 56)

	}
	{

		d.Bounce = 0 | (int64(buf[0+208]) << 0) | (int64(buf[1+208]) << 8) | (int64(buf[2+208]) << 16) | (int64(buf[3+208]) << 24) | (int64(buf[4+208]) << 32) | (int64(buf[5+208]) << 40) | (int64(buf[6+208]) << 48) | (int64(buf[7+208]) << 56)

	}
	{

		d.WritebackTmp = 0 | (int64(buf[0+216]) << 0) | (int64(buf[1+216]) << 8) | (int64(buf[2+216]) << 16) | (int64(buf[3+216]) << 24) | (int64(buf[4+216]) << 32) | (int64(buf[5+216]) << 40) | (int64(buf[6+216]) << 48) | (int64(buf[7+216]) << 56)

	}
	{

		d.CommitLimit = 0 | (int64(buf[0+224]) << 0) | (int64(buf[1+224]) << 8) | (int64(buf[2+224]) << 16) | (int64(buf[3+224]) << 24) | (int64(buf[4+224]) << 32) | (int64(buf[5+224]) << 40) | (int64(buf[6+224]) << 48) | (int64(buf[7+224]) << 56)

	}
	{

		d.VmallocTotal = 0 | (int64(buf[0+232]) << 0) | (int64(buf[1+232]) << 8) | (int64(buf[2+232]) << 16) | (int64(buf[3+232]) << 24) | (int64(buf[4+232]) << 32) | (int64(buf[5+232]) << 40) | (int64(buf[6+232]) << 48) | (int64(buf[7+232]) << 56)

	}
	{

		d.VmallocUsed = 0 | (int64(buf[0+240]) << 0) | (int64(buf[1+240]) << 8) | (int64(buf[2+240]) << 16) | (int64(buf[3+240]) << 24) | (int64(buf[4+240]) << 32) | (int64(buf[5+240]) << 40) | (int64(buf[6+240]) << 48) | (int64(buf[7+240]) << 56)

	}
	{

		d.VmallocChunk = 0 | (int64(buf[0+248]) << 0) | (int64(buf[1+248]) << 8) | (int64(buf[2+248]) << 16) | (int64(buf[3+248]) << 24) | (int64(buf[4+248]) << 32) | (int64(buf[5+248]) << 40) | (int64(buf[6+248]) << 48) | (int64(buf[7+248]) << 56)

	}
	{

		d.HardwareCorrupted = 0 | (int64(buf[0+256]) << 0) | (int64(buf[1+256]) << 8) | (int64(buf[2+256]) << 16) | (int64(buf[3+256]) << 24) | (int64(buf[4+256]) << 32) | (int64(buf[5+256]) << 40) | (int64(buf[6+256]) << 48) | (int64(buf[7+256]) << 56)

	}
	{

		d.AnonHugePages = 0 | (int64(buf[0+264]) << 0) | (int64(buf[1+264]) << 8) | (int64(buf[2+264]) << 16) | (int64(buf[3+264]) << 24) | (int64(buf[4+264]) << 32) | (int64(buf[5+264]) << 40) | (int64(buf[6+264]) << 48) | (int64(buf[7+264]) << 56)

	}
	{

		d.HugePagesTotal = 0 | (int64(buf[0+272]) << 0) | (int64(buf[1+272]) << 8) | (int64(buf[2+272]) << 16) | (int64(buf[3+272]) << 24) | (int64(buf[4+272]) << 32) | (int64(buf[5+272]) << 40) | (int64(buf[6+272]) << 48) | (int64(buf[7+272]) << 56)

	}
	{

		d.HugePagesFree = 0 | (int64(buf[0+280]) << 0) | (int64(buf[1+280]) << 8) | (int64(buf[2+280]) << 16) | (int64(buf[3+280]) << 24) | (int64(buf[4+280]) << 32) | (int64(buf[5+280]) << 40) | (int64(buf[6+280]) << 48) | (int64(buf[7+280]) << 56)

	}
	{

		d.HugePagesRsvd = 0 | (int64(buf[0+288]) << 0) | (int64(buf[1+288]) << 8) | (int64(buf[2+288]) << 16) | (int64(buf[3+288]) << 24) | (int64(buf[4+288]) << 32) | (int64(buf[5+288]) << 40) | (int64(buf[6+288]) << 48) | (int64(buf[7+288]) << 56)

	}
	{

		d.Hugepagesize = 0 | (int64(buf[0+296]) << 0) | (int64(buf[1+296]) << 8) | (int64(buf[2+296]) << 16) | (int64(buf[3+296]) << 24) | (int64(buf[4+296]) << 32) | (int64(buf[5+296]) << 40) | (int64(buf[6+296]) << 48) | (int64(buf[7+296]) << 56)

	}
	{

		d.DirectMap4k = 0 | (int64(buf[0+304]) << 0) | (int64(buf[1+304]) << 8) | (int64(buf[2+304]) << 16) | (int64(buf[3+304]) << 24) | (int64(buf[4+304]) << 32) | (int64(buf[5+304]) << 40) | (int64(buf[6+304]) << 48) | (int64(buf[7+304]) << 56)

	}
	{

		d.DirectMap2M = 0 | (int64(buf[0+312]) << 0) | (int64(buf[1+312]) << 8) | (int64(buf[2+312]) << 16) | (int64(buf[3+312]) << 24) | (int64(buf[4+312]) << 32) | (int64(buf[5+312]) << 40) | (int64(buf[6+312]) << 48) | (int64(buf[7+312]) << 56)

	}
	return i + 320, nil
}

type BasicMemInfo struct {
	MemTotal     int64
	MemFree      int64
	MemAvailable int64
	Buffers      int64
	Cached       int64
	SwapCached   int64
	SwapTotal    int64
	SwapFree     int64
}

func (d *BasicMemInfo) Size() (s uint64) {

	s += 64
	return
}
func (d *BasicMemInfo) Marshal(buf []byte) ([]byte, error) {
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

		buf[0+0] = byte(d.MemTotal >> 0)

		buf[1+0] = byte(d.MemTotal >> 8)

		buf[2+0] = byte(d.MemTotal >> 16)

		buf[3+0] = byte(d.MemTotal >> 24)

		buf[4+0] = byte(d.MemTotal >> 32)

		buf[5+0] = byte(d.MemTotal >> 40)

		buf[6+0] = byte(d.MemTotal >> 48)

		buf[7+0] = byte(d.MemTotal >> 56)

	}
	{

		buf[0+8] = byte(d.MemFree >> 0)

		buf[1+8] = byte(d.MemFree >> 8)

		buf[2+8] = byte(d.MemFree >> 16)

		buf[3+8] = byte(d.MemFree >> 24)

		buf[4+8] = byte(d.MemFree >> 32)

		buf[5+8] = byte(d.MemFree >> 40)

		buf[6+8] = byte(d.MemFree >> 48)

		buf[7+8] = byte(d.MemFree >> 56)

	}
	{

		buf[0+16] = byte(d.MemAvailable >> 0)

		buf[1+16] = byte(d.MemAvailable >> 8)

		buf[2+16] = byte(d.MemAvailable >> 16)

		buf[3+16] = byte(d.MemAvailable >> 24)

		buf[4+16] = byte(d.MemAvailable >> 32)

		buf[5+16] = byte(d.MemAvailable >> 40)

		buf[6+16] = byte(d.MemAvailable >> 48)

		buf[7+16] = byte(d.MemAvailable >> 56)

	}
	{

		buf[0+24] = byte(d.Buffers >> 0)

		buf[1+24] = byte(d.Buffers >> 8)

		buf[2+24] = byte(d.Buffers >> 16)

		buf[3+24] = byte(d.Buffers >> 24)

		buf[4+24] = byte(d.Buffers >> 32)

		buf[5+24] = byte(d.Buffers >> 40)

		buf[6+24] = byte(d.Buffers >> 48)

		buf[7+24] = byte(d.Buffers >> 56)

	}
	{

		buf[0+32] = byte(d.Cached >> 0)

		buf[1+32] = byte(d.Cached >> 8)

		buf[2+32] = byte(d.Cached >> 16)

		buf[3+32] = byte(d.Cached >> 24)

		buf[4+32] = byte(d.Cached >> 32)

		buf[5+32] = byte(d.Cached >> 40)

		buf[6+32] = byte(d.Cached >> 48)

		buf[7+32] = byte(d.Cached >> 56)

	}
	{

		buf[0+40] = byte(d.SwapCached >> 0)

		buf[1+40] = byte(d.SwapCached >> 8)

		buf[2+40] = byte(d.SwapCached >> 16)

		buf[3+40] = byte(d.SwapCached >> 24)

		buf[4+40] = byte(d.SwapCached >> 32)

		buf[5+40] = byte(d.SwapCached >> 40)

		buf[6+40] = byte(d.SwapCached >> 48)

		buf[7+40] = byte(d.SwapCached >> 56)

	}
	{

		buf[0+48] = byte(d.SwapTotal >> 0)

		buf[1+48] = byte(d.SwapTotal >> 8)

		buf[2+48] = byte(d.SwapTotal >> 16)

		buf[3+48] = byte(d.SwapTotal >> 24)

		buf[4+48] = byte(d.SwapTotal >> 32)

		buf[5+48] = byte(d.SwapTotal >> 40)

		buf[6+48] = byte(d.SwapTotal >> 48)

		buf[7+48] = byte(d.SwapTotal >> 56)

	}
	{

		buf[0+56] = byte(d.SwapFree >> 0)

		buf[1+56] = byte(d.SwapFree >> 8)

		buf[2+56] = byte(d.SwapFree >> 16)

		buf[3+56] = byte(d.SwapFree >> 24)

		buf[4+56] = byte(d.SwapFree >> 32)

		buf[5+56] = byte(d.SwapFree >> 40)

		buf[6+56] = byte(d.SwapFree >> 48)

		buf[7+56] = byte(d.SwapFree >> 56)

	}
	return buf[:i+64], nil
}

func (d *BasicMemInfo) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		d.MemTotal = 0 | (int64(buf[0+0]) << 0) | (int64(buf[1+0]) << 8) | (int64(buf[2+0]) << 16) | (int64(buf[3+0]) << 24) | (int64(buf[4+0]) << 32) | (int64(buf[5+0]) << 40) | (int64(buf[6+0]) << 48) | (int64(buf[7+0]) << 56)

	}
	{

		d.MemFree = 0 | (int64(buf[0+8]) << 0) | (int64(buf[1+8]) << 8) | (int64(buf[2+8]) << 16) | (int64(buf[3+8]) << 24) | (int64(buf[4+8]) << 32) | (int64(buf[5+8]) << 40) | (int64(buf[6+8]) << 48) | (int64(buf[7+8]) << 56)

	}
	{

		d.MemAvailable = 0 | (int64(buf[0+16]) << 0) | (int64(buf[1+16]) << 8) | (int64(buf[2+16]) << 16) | (int64(buf[3+16]) << 24) | (int64(buf[4+16]) << 32) | (int64(buf[5+16]) << 40) | (int64(buf[6+16]) << 48) | (int64(buf[7+16]) << 56)

	}
	{

		d.Buffers = 0 | (int64(buf[0+24]) << 0) | (int64(buf[1+24]) << 8) | (int64(buf[2+24]) << 16) | (int64(buf[3+24]) << 24) | (int64(buf[4+24]) << 32) | (int64(buf[5+24]) << 40) | (int64(buf[6+24]) << 48) | (int64(buf[7+24]) << 56)

	}
	{

		d.Cached = 0 | (int64(buf[0+32]) << 0) | (int64(buf[1+32]) << 8) | (int64(buf[2+32]) << 16) | (int64(buf[3+32]) << 24) | (int64(buf[4+32]) << 32) | (int64(buf[5+32]) << 40) | (int64(buf[6+32]) << 48) | (int64(buf[7+32]) << 56)

	}
	{

		d.SwapCached = 0 | (int64(buf[0+40]) << 0) | (int64(buf[1+40]) << 8) | (int64(buf[2+40]) << 16) | (int64(buf[3+40]) << 24) | (int64(buf[4+40]) << 32) | (int64(buf[5+40]) << 40) | (int64(buf[6+40]) << 48) | (int64(buf[7+40]) << 56)

	}
	{

		d.SwapTotal = 0 | (int64(buf[0+48]) << 0) | (int64(buf[1+48]) << 8) | (int64(buf[2+48]) << 16) | (int64(buf[3+48]) << 24) | (int64(buf[4+48]) << 32) | (int64(buf[5+48]) << 40) | (int64(buf[6+48]) << 48) | (int64(buf[7+48]) << 56)

	}
	{

		d.SwapFree = 0 | (int64(buf[0+56]) << 0) | (int64(buf[1+56]) << 8) | (int64(buf[2+56]) << 16) | (int64(buf[3+56]) << 24) | (int64(buf[4+56]) << 32) | (int64(buf[5+56]) << 40) | (int64(buf[6+56]) << 48) | (int64(buf[7+56]) << 56)

	}
	return i + 64, nil
}
