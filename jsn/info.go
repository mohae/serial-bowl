package jsn

type MemInfo struct {
	MemTotal          int
	MemFree           int
	MemAvailable      int
	Buffers           int
	Cached            int
	SwapCached        int
	Active            int
	Inactive          int
	ActiveAnon        int
	InactiveAnon      int
	ActiveFile        int
	InactiveFile      int
	Unevictable       int
	Mlocked           int
	SwapTotal         int
	SwapFree          int
	Dirty             int
	Writeback         int
	AnonPages         int
	Mapped            int
	Shmem             int
	Slab              int
	SReclaimable      int
	SUnreclaim        int
	KernelStack       int
	NFSUnstable       int
	Bounce            int
	WritebackTmp      int
	CommitLimit       int
	VmallocTotal      int
	VmallocUsed       int
	VmallocChunk      int
	HardwareCorrupted int
	AnonHugePages     int
	HugePagesTotal    int
	HugePagesFree     int
	HugePagesRsvd     int
	Hugepagesize      int
	DirectMap4k       int
	DirectMap2M       int
}

type BasicMemInfo struct {
	MemTotal     int
	MemFree      int
	MemAvailable int
	Buffers      int
	Cached       int
	SwapCached   int
	SwapTotal    int
	SwapFree     int
}
