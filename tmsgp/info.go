package tmsgp

//go:generate msgp
type MemInfo struct {
	MemTotal          int64 `msg:"mem_total"`
	MemFree           int64 `msg:"mem_free"`
	MemAvailable      int64 `msg:"mem_available"`
	Buffers           int64 `msg:"buffers"`
	Cached            int64 `msg:"cached"`
	SwapCached        int64 `msg:"swap_cached"`
	Active            int64 `msg:"active"`
	Inactive          int64 `msg:"inactive"`
	ActiveAnon        int64 `msg:"active_anon"`
	InactiveAnon      int64 `msg:"inactive_anon"`
	ActiveFile        int64 `msg:"active_file"`
	InactiveFile      int64 `msg:"inactive_file"`
	Unevictable       int64 `msg:"unevictable"`
	Mlocked           int64 `msg:"mlocked"`
	SwapTotal         int64 `msg:"swap_total"`
	SwapFree          int64 `msg:"swap_free"`
	Dirty             int64 `msg:"dirty"`
	Writeback         int64 `msg:"writeback"`
	AnonPages         int64 `msg:"anon_pages"`
	Mapped            int64 `msg:"mapped"`
	Shmem             int64 `msg:"sh_mem"`
	Slab              int64 `msg:"slab"`
	SReclaimable      int64 `msg:"s_reclaimable"`
	SUnreclaim        int64 `msg:"s_unreclaim"`
	KernelStack       int64 `msg:"kernel_stack"`
	NFSUnstable       int64 `msg:"nfs_unstable"`
	Bounce            int64 `msg:"bounce"`
	WritebackTmp      int64 `msg:"writeback_tmp"`
	CommitLimit       int64 `msg:"commit_limit"`
	VmallocTotal      int64 `msg:"vmalloc_total"`
	VmallocUsed       int64 `msg:"vmalloc_used"`
	VmallocChunk      int64 `msg:"vmalloc_chunked"`
	HardwareCorrupted int64 `msg:"hardware_corrupted"`
	AnonHugePages     int64 `msg:"anon_huge_pages"`
	HugePagesTotal    int64 `msg:"huge_pages_total"`
	HugePagesFree     int64 `msg:"huge_pages_free"`
	HugePagesRsvd     int64 `msg:"huge_pages_rsvd"`
	Hugepagesize      int64 `msg:"huge_pages_size"`
	DirectMap4k       int64 `msg:"direct_map_4k"`
	DirectMap2M       int64 `msg:"direct_map_2m"`
}

type BasicMemInfo struct {
	MemTotal     int64 `msg:"mem_total"`
	MemFree      int64 `msg:"mem_free"`
	MemAvailable int64 `msg:"mem_available"`
	Buffers      int64 `msg:"buffers"`
	Cached       int64 `msg:"cached"`
	SwapCached   int64 `msg:"swap_cached"`
	SwapTotal    int64 `msg:"swap_total"`
	SwapFree     int64 `msg:"swap_free"`
}
