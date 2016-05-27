package tmsgp

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *BasicMemInfo) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "mem_total":
			z.MemTotal, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "mem_free":
			z.MemFree, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "mem_available":
			z.MemAvailable, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "buffers":
			z.Buffers, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "cached":
			z.Cached, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "swap_cached":
			z.SwapCached, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "swap_total":
			z.SwapTotal, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "swap_free":
			z.SwapFree, err = dc.ReadInt64()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *BasicMemInfo) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 8
	// write "mem_total"
	err = en.Append(0x88, 0xa9, 0x6d, 0x65, 0x6d, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.MemTotal)
	if err != nil {
		return
	}
	// write "mem_free"
	err = en.Append(0xa8, 0x6d, 0x65, 0x6d, 0x5f, 0x66, 0x72, 0x65, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.MemFree)
	if err != nil {
		return
	}
	// write "mem_available"
	err = en.Append(0xad, 0x6d, 0x65, 0x6d, 0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.MemAvailable)
	if err != nil {
		return
	}
	// write "buffers"
	err = en.Append(0xa7, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Buffers)
	if err != nil {
		return
	}
	// write "cached"
	err = en.Append(0xa6, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Cached)
	if err != nil {
		return
	}
	// write "swap_cached"
	err = en.Append(0xab, 0x73, 0x77, 0x61, 0x70, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.SwapCached)
	if err != nil {
		return
	}
	// write "swap_total"
	err = en.Append(0xaa, 0x73, 0x77, 0x61, 0x70, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.SwapTotal)
	if err != nil {
		return
	}
	// write "swap_free"
	err = en.Append(0xa9, 0x73, 0x77, 0x61, 0x70, 0x5f, 0x66, 0x72, 0x65, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.SwapFree)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *BasicMemInfo) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 8
	// string "mem_total"
	o = append(o, 0x88, 0xa9, 0x6d, 0x65, 0x6d, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c)
	o = msgp.AppendInt64(o, z.MemTotal)
	// string "mem_free"
	o = append(o, 0xa8, 0x6d, 0x65, 0x6d, 0x5f, 0x66, 0x72, 0x65, 0x65)
	o = msgp.AppendInt64(o, z.MemFree)
	// string "mem_available"
	o = append(o, 0xad, 0x6d, 0x65, 0x6d, 0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65)
	o = msgp.AppendInt64(o, z.MemAvailable)
	// string "buffers"
	o = append(o, 0xa7, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x73)
	o = msgp.AppendInt64(o, z.Buffers)
	// string "cached"
	o = append(o, 0xa6, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64)
	o = msgp.AppendInt64(o, z.Cached)
	// string "swap_cached"
	o = append(o, 0xab, 0x73, 0x77, 0x61, 0x70, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64)
	o = msgp.AppendInt64(o, z.SwapCached)
	// string "swap_total"
	o = append(o, 0xaa, 0x73, 0x77, 0x61, 0x70, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c)
	o = msgp.AppendInt64(o, z.SwapTotal)
	// string "swap_free"
	o = append(o, 0xa9, 0x73, 0x77, 0x61, 0x70, 0x5f, 0x66, 0x72, 0x65, 0x65)
	o = msgp.AppendInt64(o, z.SwapFree)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *BasicMemInfo) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "mem_total":
			z.MemTotal, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "mem_free":
			z.MemFree, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "mem_available":
			z.MemAvailable, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "buffers":
			z.Buffers, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "cached":
			z.Cached, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "swap_cached":
			z.SwapCached, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "swap_total":
			z.SwapTotal, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "swap_free":
			z.SwapFree, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

func (z *BasicMemInfo) Msgsize() (s int) {
	s = 1 + 10 + msgp.Int64Size + 9 + msgp.Int64Size + 14 + msgp.Int64Size + 8 + msgp.Int64Size + 7 + msgp.Int64Size + 12 + msgp.Int64Size + 11 + msgp.Int64Size + 10 + msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *MemInfo) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "mem_total":
			z.MemTotal, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "mem_free":
			z.MemFree, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "mem_available":
			z.MemAvailable, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "buffers":
			z.Buffers, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "cached":
			z.Cached, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "swap_cached":
			z.SwapCached, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "active":
			z.Active, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "inactive":
			z.Inactive, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "active_anon":
			z.ActiveAnon, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "inactive_anon":
			z.InactiveAnon, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "active_file":
			z.ActiveFile, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "inactive_file":
			z.InactiveFile, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "unevictable":
			z.Unevictable, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "mlocked":
			z.Mlocked, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "swap_total":
			z.SwapTotal, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "swap_free":
			z.SwapFree, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "dirty":
			z.Dirty, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "writeback":
			z.Writeback, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "anon_pages":
			z.AnonPages, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "mapped":
			z.Mapped, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "sh_mem":
			z.Shmem, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "slab":
			z.Slab, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "s_reclaimable":
			z.SReclaimable, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "s_unreclaim":
			z.SUnreclaim, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "kernel_stack":
			z.KernelStack, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "nfs_unstable":
			z.NFSUnstable, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "bounce":
			z.Bounce, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "writeback_tmp":
			z.WritebackTmp, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "commit_limit":
			z.CommitLimit, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "vmalloc_total":
			z.VmallocTotal, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "vmalloc_used":
			z.VmallocUsed, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "vmalloc_chunked":
			z.VmallocChunk, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "hardware_corrupted":
			z.HardwareCorrupted, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "anon_huge_pages":
			z.AnonHugePages, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "huge_pages_total":
			z.HugePagesTotal, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "huge_pages_free":
			z.HugePagesFree, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "huge_pages_rsvd":
			z.HugePagesRsvd, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "huge_pages_size":
			z.Hugepagesize, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "direct_map_4k":
			z.DirectMap4k, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "direct_map_2m":
			z.DirectMap2M, err = dc.ReadInt64()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *MemInfo) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 40
	// write "mem_total"
	err = en.Append(0xde, 0x0, 0x28, 0xa9, 0x6d, 0x65, 0x6d, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.MemTotal)
	if err != nil {
		return
	}
	// write "mem_free"
	err = en.Append(0xa8, 0x6d, 0x65, 0x6d, 0x5f, 0x66, 0x72, 0x65, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.MemFree)
	if err != nil {
		return
	}
	// write "mem_available"
	err = en.Append(0xad, 0x6d, 0x65, 0x6d, 0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.MemAvailable)
	if err != nil {
		return
	}
	// write "buffers"
	err = en.Append(0xa7, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Buffers)
	if err != nil {
		return
	}
	// write "cached"
	err = en.Append(0xa6, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Cached)
	if err != nil {
		return
	}
	// write "swap_cached"
	err = en.Append(0xab, 0x73, 0x77, 0x61, 0x70, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.SwapCached)
	if err != nil {
		return
	}
	// write "active"
	err = en.Append(0xa6, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Active)
	if err != nil {
		return
	}
	// write "inactive"
	err = en.Append(0xa8, 0x69, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Inactive)
	if err != nil {
		return
	}
	// write "active_anon"
	err = en.Append(0xab, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x6e, 0x6f, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.ActiveAnon)
	if err != nil {
		return
	}
	// write "inactive_anon"
	err = en.Append(0xad, 0x69, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x6e, 0x6f, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.InactiveAnon)
	if err != nil {
		return
	}
	// write "active_file"
	err = en.Append(0xab, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.ActiveFile)
	if err != nil {
		return
	}
	// write "inactive_file"
	err = en.Append(0xad, 0x69, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.InactiveFile)
	if err != nil {
		return
	}
	// write "unevictable"
	err = en.Append(0xab, 0x75, 0x6e, 0x65, 0x76, 0x69, 0x63, 0x74, 0x61, 0x62, 0x6c, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Unevictable)
	if err != nil {
		return
	}
	// write "mlocked"
	err = en.Append(0xa7, 0x6d, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Mlocked)
	if err != nil {
		return
	}
	// write "swap_total"
	err = en.Append(0xaa, 0x73, 0x77, 0x61, 0x70, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.SwapTotal)
	if err != nil {
		return
	}
	// write "swap_free"
	err = en.Append(0xa9, 0x73, 0x77, 0x61, 0x70, 0x5f, 0x66, 0x72, 0x65, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.SwapFree)
	if err != nil {
		return
	}
	// write "dirty"
	err = en.Append(0xa5, 0x64, 0x69, 0x72, 0x74, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Dirty)
	if err != nil {
		return
	}
	// write "writeback"
	err = en.Append(0xa9, 0x77, 0x72, 0x69, 0x74, 0x65, 0x62, 0x61, 0x63, 0x6b)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Writeback)
	if err != nil {
		return
	}
	// write "anon_pages"
	err = en.Append(0xaa, 0x61, 0x6e, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.AnonPages)
	if err != nil {
		return
	}
	// write "mapped"
	err = en.Append(0xa6, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Mapped)
	if err != nil {
		return
	}
	// write "sh_mem"
	err = en.Append(0xa6, 0x73, 0x68, 0x5f, 0x6d, 0x65, 0x6d)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Shmem)
	if err != nil {
		return
	}
	// write "slab"
	err = en.Append(0xa4, 0x73, 0x6c, 0x61, 0x62)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Slab)
	if err != nil {
		return
	}
	// write "s_reclaimable"
	err = en.Append(0xad, 0x73, 0x5f, 0x72, 0x65, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x61, 0x62, 0x6c, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.SReclaimable)
	if err != nil {
		return
	}
	// write "s_unreclaim"
	err = en.Append(0xab, 0x73, 0x5f, 0x75, 0x6e, 0x72, 0x65, 0x63, 0x6c, 0x61, 0x69, 0x6d)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.SUnreclaim)
	if err != nil {
		return
	}
	// write "kernel_stack"
	err = en.Append(0xac, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x63, 0x6b)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.KernelStack)
	if err != nil {
		return
	}
	// write "nfs_unstable"
	err = en.Append(0xac, 0x6e, 0x66, 0x73, 0x5f, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x62, 0x6c, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.NFSUnstable)
	if err != nil {
		return
	}
	// write "bounce"
	err = en.Append(0xa6, 0x62, 0x6f, 0x75, 0x6e, 0x63, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Bounce)
	if err != nil {
		return
	}
	// write "writeback_tmp"
	err = en.Append(0xad, 0x77, 0x72, 0x69, 0x74, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x74, 0x6d, 0x70)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.WritebackTmp)
	if err != nil {
		return
	}
	// write "commit_limit"
	err = en.Append(0xac, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.CommitLimit)
	if err != nil {
		return
	}
	// write "vmalloc_total"
	err = en.Append(0xad, 0x76, 0x6d, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.VmallocTotal)
	if err != nil {
		return
	}
	// write "vmalloc_used"
	err = en.Append(0xac, 0x76, 0x6d, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x5f, 0x75, 0x73, 0x65, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.VmallocUsed)
	if err != nil {
		return
	}
	// write "vmalloc_chunked"
	err = en.Append(0xaf, 0x76, 0x6d, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x65, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.VmallocChunk)
	if err != nil {
		return
	}
	// write "hardware_corrupted"
	err = en.Append(0xb2, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x72, 0x72, 0x75, 0x70, 0x74, 0x65, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.HardwareCorrupted)
	if err != nil {
		return
	}
	// write "anon_huge_pages"
	err = en.Append(0xaf, 0x61, 0x6e, 0x6f, 0x6e, 0x5f, 0x68, 0x75, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.AnonHugePages)
	if err != nil {
		return
	}
	// write "huge_pages_total"
	err = en.Append(0xb0, 0x68, 0x75, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.HugePagesTotal)
	if err != nil {
		return
	}
	// write "huge_pages_free"
	err = en.Append(0xaf, 0x68, 0x75, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x66, 0x72, 0x65, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.HugePagesFree)
	if err != nil {
		return
	}
	// write "huge_pages_rsvd"
	err = en.Append(0xaf, 0x68, 0x75, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x72, 0x73, 0x76, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.HugePagesRsvd)
	if err != nil {
		return
	}
	// write "huge_pages_size"
	err = en.Append(0xaf, 0x68, 0x75, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x73, 0x69, 0x7a, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Hugepagesize)
	if err != nil {
		return
	}
	// write "direct_map_4k"
	err = en.Append(0xad, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x70, 0x5f, 0x34, 0x6b)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.DirectMap4k)
	if err != nil {
		return
	}
	// write "direct_map_2m"
	err = en.Append(0xad, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x70, 0x5f, 0x32, 0x6d)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.DirectMap2M)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *MemInfo) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 40
	// string "mem_total"
	o = append(o, 0xde, 0x0, 0x28, 0xa9, 0x6d, 0x65, 0x6d, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c)
	o = msgp.AppendInt64(o, z.MemTotal)
	// string "mem_free"
	o = append(o, 0xa8, 0x6d, 0x65, 0x6d, 0x5f, 0x66, 0x72, 0x65, 0x65)
	o = msgp.AppendInt64(o, z.MemFree)
	// string "mem_available"
	o = append(o, 0xad, 0x6d, 0x65, 0x6d, 0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65)
	o = msgp.AppendInt64(o, z.MemAvailable)
	// string "buffers"
	o = append(o, 0xa7, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x73)
	o = msgp.AppendInt64(o, z.Buffers)
	// string "cached"
	o = append(o, 0xa6, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64)
	o = msgp.AppendInt64(o, z.Cached)
	// string "swap_cached"
	o = append(o, 0xab, 0x73, 0x77, 0x61, 0x70, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64)
	o = msgp.AppendInt64(o, z.SwapCached)
	// string "active"
	o = append(o, 0xa6, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65)
	o = msgp.AppendInt64(o, z.Active)
	// string "inactive"
	o = append(o, 0xa8, 0x69, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65)
	o = msgp.AppendInt64(o, z.Inactive)
	// string "active_anon"
	o = append(o, 0xab, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x6e, 0x6f, 0x6e)
	o = msgp.AppendInt64(o, z.ActiveAnon)
	// string "inactive_anon"
	o = append(o, 0xad, 0x69, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x6e, 0x6f, 0x6e)
	o = msgp.AppendInt64(o, z.InactiveAnon)
	// string "active_file"
	o = append(o, 0xab, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x65)
	o = msgp.AppendInt64(o, z.ActiveFile)
	// string "inactive_file"
	o = append(o, 0xad, 0x69, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x65)
	o = msgp.AppendInt64(o, z.InactiveFile)
	// string "unevictable"
	o = append(o, 0xab, 0x75, 0x6e, 0x65, 0x76, 0x69, 0x63, 0x74, 0x61, 0x62, 0x6c, 0x65)
	o = msgp.AppendInt64(o, z.Unevictable)
	// string "mlocked"
	o = append(o, 0xa7, 0x6d, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64)
	o = msgp.AppendInt64(o, z.Mlocked)
	// string "swap_total"
	o = append(o, 0xaa, 0x73, 0x77, 0x61, 0x70, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c)
	o = msgp.AppendInt64(o, z.SwapTotal)
	// string "swap_free"
	o = append(o, 0xa9, 0x73, 0x77, 0x61, 0x70, 0x5f, 0x66, 0x72, 0x65, 0x65)
	o = msgp.AppendInt64(o, z.SwapFree)
	// string "dirty"
	o = append(o, 0xa5, 0x64, 0x69, 0x72, 0x74, 0x79)
	o = msgp.AppendInt64(o, z.Dirty)
	// string "writeback"
	o = append(o, 0xa9, 0x77, 0x72, 0x69, 0x74, 0x65, 0x62, 0x61, 0x63, 0x6b)
	o = msgp.AppendInt64(o, z.Writeback)
	// string "anon_pages"
	o = append(o, 0xaa, 0x61, 0x6e, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73)
	o = msgp.AppendInt64(o, z.AnonPages)
	// string "mapped"
	o = append(o, 0xa6, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x64)
	o = msgp.AppendInt64(o, z.Mapped)
	// string "sh_mem"
	o = append(o, 0xa6, 0x73, 0x68, 0x5f, 0x6d, 0x65, 0x6d)
	o = msgp.AppendInt64(o, z.Shmem)
	// string "slab"
	o = append(o, 0xa4, 0x73, 0x6c, 0x61, 0x62)
	o = msgp.AppendInt64(o, z.Slab)
	// string "s_reclaimable"
	o = append(o, 0xad, 0x73, 0x5f, 0x72, 0x65, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x61, 0x62, 0x6c, 0x65)
	o = msgp.AppendInt64(o, z.SReclaimable)
	// string "s_unreclaim"
	o = append(o, 0xab, 0x73, 0x5f, 0x75, 0x6e, 0x72, 0x65, 0x63, 0x6c, 0x61, 0x69, 0x6d)
	o = msgp.AppendInt64(o, z.SUnreclaim)
	// string "kernel_stack"
	o = append(o, 0xac, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x63, 0x6b)
	o = msgp.AppendInt64(o, z.KernelStack)
	// string "nfs_unstable"
	o = append(o, 0xac, 0x6e, 0x66, 0x73, 0x5f, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x62, 0x6c, 0x65)
	o = msgp.AppendInt64(o, z.NFSUnstable)
	// string "bounce"
	o = append(o, 0xa6, 0x62, 0x6f, 0x75, 0x6e, 0x63, 0x65)
	o = msgp.AppendInt64(o, z.Bounce)
	// string "writeback_tmp"
	o = append(o, 0xad, 0x77, 0x72, 0x69, 0x74, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x74, 0x6d, 0x70)
	o = msgp.AppendInt64(o, z.WritebackTmp)
	// string "commit_limit"
	o = append(o, 0xac, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74)
	o = msgp.AppendInt64(o, z.CommitLimit)
	// string "vmalloc_total"
	o = append(o, 0xad, 0x76, 0x6d, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c)
	o = msgp.AppendInt64(o, z.VmallocTotal)
	// string "vmalloc_used"
	o = append(o, 0xac, 0x76, 0x6d, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x5f, 0x75, 0x73, 0x65, 0x64)
	o = msgp.AppendInt64(o, z.VmallocUsed)
	// string "vmalloc_chunked"
	o = append(o, 0xaf, 0x76, 0x6d, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x65, 0x64)
	o = msgp.AppendInt64(o, z.VmallocChunk)
	// string "hardware_corrupted"
	o = append(o, 0xb2, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x72, 0x72, 0x75, 0x70, 0x74, 0x65, 0x64)
	o = msgp.AppendInt64(o, z.HardwareCorrupted)
	// string "anon_huge_pages"
	o = append(o, 0xaf, 0x61, 0x6e, 0x6f, 0x6e, 0x5f, 0x68, 0x75, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73)
	o = msgp.AppendInt64(o, z.AnonHugePages)
	// string "huge_pages_total"
	o = append(o, 0xb0, 0x68, 0x75, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c)
	o = msgp.AppendInt64(o, z.HugePagesTotal)
	// string "huge_pages_free"
	o = append(o, 0xaf, 0x68, 0x75, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x66, 0x72, 0x65, 0x65)
	o = msgp.AppendInt64(o, z.HugePagesFree)
	// string "huge_pages_rsvd"
	o = append(o, 0xaf, 0x68, 0x75, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x72, 0x73, 0x76, 0x64)
	o = msgp.AppendInt64(o, z.HugePagesRsvd)
	// string "huge_pages_size"
	o = append(o, 0xaf, 0x68, 0x75, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x73, 0x69, 0x7a, 0x65)
	o = msgp.AppendInt64(o, z.Hugepagesize)
	// string "direct_map_4k"
	o = append(o, 0xad, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x70, 0x5f, 0x34, 0x6b)
	o = msgp.AppendInt64(o, z.DirectMap4k)
	// string "direct_map_2m"
	o = append(o, 0xad, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x70, 0x5f, 0x32, 0x6d)
	o = msgp.AppendInt64(o, z.DirectMap2M)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MemInfo) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "mem_total":
			z.MemTotal, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "mem_free":
			z.MemFree, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "mem_available":
			z.MemAvailable, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "buffers":
			z.Buffers, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "cached":
			z.Cached, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "swap_cached":
			z.SwapCached, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "active":
			z.Active, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "inactive":
			z.Inactive, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "active_anon":
			z.ActiveAnon, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "inactive_anon":
			z.InactiveAnon, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "active_file":
			z.ActiveFile, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "inactive_file":
			z.InactiveFile, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "unevictable":
			z.Unevictable, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "mlocked":
			z.Mlocked, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "swap_total":
			z.SwapTotal, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "swap_free":
			z.SwapFree, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "dirty":
			z.Dirty, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "writeback":
			z.Writeback, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "anon_pages":
			z.AnonPages, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "mapped":
			z.Mapped, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "sh_mem":
			z.Shmem, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "slab":
			z.Slab, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "s_reclaimable":
			z.SReclaimable, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "s_unreclaim":
			z.SUnreclaim, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "kernel_stack":
			z.KernelStack, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "nfs_unstable":
			z.NFSUnstable, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "bounce":
			z.Bounce, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "writeback_tmp":
			z.WritebackTmp, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "commit_limit":
			z.CommitLimit, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "vmalloc_total":
			z.VmallocTotal, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "vmalloc_used":
			z.VmallocUsed, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "vmalloc_chunked":
			z.VmallocChunk, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "hardware_corrupted":
			z.HardwareCorrupted, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "anon_huge_pages":
			z.AnonHugePages, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "huge_pages_total":
			z.HugePagesTotal, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "huge_pages_free":
			z.HugePagesFree, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "huge_pages_rsvd":
			z.HugePagesRsvd, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "huge_pages_size":
			z.Hugepagesize, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "direct_map_4k":
			z.DirectMap4k, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "direct_map_2m":
			z.DirectMap2M, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

func (z *MemInfo) Msgsize() (s int) {
	s = 3 + 10 + msgp.Int64Size + 9 + msgp.Int64Size + 14 + msgp.Int64Size + 8 + msgp.Int64Size + 7 + msgp.Int64Size + 12 + msgp.Int64Size + 7 + msgp.Int64Size + 9 + msgp.Int64Size + 12 + msgp.Int64Size + 14 + msgp.Int64Size + 12 + msgp.Int64Size + 14 + msgp.Int64Size + 12 + msgp.Int64Size + 8 + msgp.Int64Size + 11 + msgp.Int64Size + 10 + msgp.Int64Size + 6 + msgp.Int64Size + 10 + msgp.Int64Size + 11 + msgp.Int64Size + 7 + msgp.Int64Size + 7 + msgp.Int64Size + 5 + msgp.Int64Size + 14 + msgp.Int64Size + 12 + msgp.Int64Size + 13 + msgp.Int64Size + 13 + msgp.Int64Size + 7 + msgp.Int64Size + 14 + msgp.Int64Size + 13 + msgp.Int64Size + 14 + msgp.Int64Size + 13 + msgp.Int64Size + 16 + msgp.Int64Size + 19 + msgp.Int64Size + 16 + msgp.Int64Size + 17 + msgp.Int64Size + 16 + msgp.Int64Size + 16 + msgp.Int64Size + 16 + msgp.Int64Size + 14 + msgp.Int64Size + 14 + msgp.Int64Size
	return
}
