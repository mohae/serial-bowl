package main

import (
	"github.com/mohae/serial-bowl/jsn"
	"github.com/mohae/serial-bowl/pb"
)

func PreparePBBasicMemInfoData(data []jsn.BasicMemInfo) []pb.BasicMemInfo {
	tmp := make([]pb.BasicMemInfo, len(data))
	for i := 0; i < len(data); i++ {
		tmp[i] = pb.BasicMemInfo{
			MemTotal:     int64(data[i].MemTotal),
			MemFree:      int64(data[i].MemFree),
			MemAvailable: int64(data[i].MemAvailable),
			Buffers:      int64(data[i].Buffers),
			Cached:       int64(data[i].Cached),
			SwapCached:   int64(data[i].SwapCached),
			SwapTotal:    int64(data[i].SwapTotal),
			SwapFree:     int64(data[i].SwapFree),
		}
	}
	return tmp
}

func PreparePBMemInfoData(data []jsn.MemInfo) []pb.MemInfo {
	tmp := make([]pb.MemInfo, len(data))
	for i := 0; i < len(data); i++ {
		tmp[i] = pb.MemInfo{
			MemTotal:          int64(data[i].MemTotal),
			MemFree:           int64(data[i].MemFree),
			MemAvailable:      int64(data[i].MemAvailable),
			Buffers:           int64(data[i].Buffers),
			Cached:            int64(data[i].Cached),
			SwapCached:        int64(data[i].SwapCached),
			Active:            int64(data[i].Active),
			Inactive:          int64(data[i].Inactive),
			ActiveAnon:        int64(data[i].ActiveAnon),
			InactiveAnon:      int64(data[i].InactiveAnon),
			ActiveFile:        int64(data[i].ActiveFile),
			InactiveFile:      int64(data[i].InactiveFile),
			Unevictable:       int64(data[i].Unevictable),
			Mlocked:           int64(data[i].Mlocked),
			SwapTotal:         int64(data[i].SwapTotal),
			SwapFree:          int64(data[i].SwapFree),
			Dirty:             int64(data[i].Dirty),
			Writeback:         int64(data[i].Writeback),
			AnonPages:         int64(data[i].AnonPages),
			Mapped:            int64(data[i].Mapped),
			Shmem:             int64(data[i].Shmem),
			Slab:              int64(data[i].Slab),
			SReclaimable:      int64(data[i].SReclaimable),
			SUnreclaim:        int64(data[i].SUnreclaim),
			KernelStack:       int64(data[i].KernelStack),
			NFSUnstable:       int64(data[i].NFSUnstable),
			Bounce:            int64(data[i].Bounce),
			WritebackTmp:      int64(data[i].WritebackTmp),
			CommitLimit:       int64(data[i].CommitLimit),
			VmallocTotal:      int64(data[i].VmallocTotal),
			VmallocUsed:       int64(data[i].VmallocUsed),
			VmallocChunk:      int64(data[i].VmallocChunk),
			HardwareCorrupted: int64(data[i].HardwareCorrupted),
			AnonHugePages:     int64(data[i].AnonHugePages),
			HugePagesTotal:    int64(data[i].HugePagesTotal),
			HugePagesFree:     int64(data[i].HugePagesFree),
			HugePagesRsvd:     int64(data[i].HugePagesRsvd),
			Hugepagesize:      int64(data[i].Hugepagesize),
			DirectMap4K:       int64(data[i].DirectMap4k),
			DirectMap2M:       int64(data[i].DirectMap2M),
		}
	}
	return tmp
}

// PreparePBMessageData generates the protobuf version of the data
func PreparePBMessageData(data []jsn.Message) []pb.Message {
	tmp := make([]pb.Message, len(data))
	for i := 0; i < len(data); i++ {
		tmp[i] = pb.Message{
			ID:     data[i].ID,
			DestID: data[i].DestID,
			Type:   int32(data[i].Type),
			Kind:   int32(data[i].Kind),
			Data:   data[i].Data,
		}
	}
	return tmp
}
