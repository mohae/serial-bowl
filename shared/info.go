package shared

type ShMemInfo struct {
	MemTotal          int64 `json:"mem_total"`
	MemFree           int64 `json:"mem_free"`
	MemAvailable      int64 `json:"mem_available"`
	Buffers           int64 `json:"buffers"`
	Cached            int64 `json:"cached"`
	SwapCached        int64 `json:"swap_cached"`
	Active            int64 `json:"active"`
	Inactive          int64 `json:"inactive"`
	ActiveAnon        int64 `json:"active_anon"`
	InactiveAnon      int64 `json:"inactive_anon"`
	ActiveFile        int64 `json:"active_file"`
	InactiveFile      int64 `json:"inactive_file"`
	Unevictable       int64 `json:"unevictable"`
	Mlocked           int64 `json:"mlocked"`
	SwapTotal         int64 `json:"swap_total"`
	SwapFree          int64 `json:"swap_free"`
	Dirty             int64 `json:"dirty"`
	Writeback         int64 `json:"writeback"`
	AnonPages         int64 `json:"anon_pages"`
	Mapped            int64 `json:"mapped"`
	Shmem             int64 `json:"sh_mem"`
	Slab              int64 `json:"slab"`
	SReclaimable      int64 `json:"s_reclaimable"`
	SUnreclaim        int64 `json:"s_unreclaim"`
	KernelStack       int64 `json:"kernel_stack"`
	NFSUnstable       int64 `json:"nfs_unstable"`
	Bounce            int64 `json:"bounce"`
	WritebackTmp      int64 `json:"writeback_tmp"`
	CommitLimit       int64 `json:"commit_limit"`
	VmallocTotal      int64 `json:"vmalloc_total"`
	VmallocUsed       int64 `json:"vmalloc_used"`
	VmallocChunk      int64 `json:"vmalloc_chunked"`
	HardwareCorrupted int64 `json:"hardware_corrupted"`
	AnonHugePages     int64 `json:"anon_huge_pages"`
	HugePagesTotal    int64 `json:"huge_pages_total"`
	HugePagesFree     int64 `json:"huge_pages_free"`
	HugePagesRsvd     int64 `json:"huge_pages_rsvd"`
	Hugepagesize      int64 `json:"huge_pages_size"`
	DirectMap4k       int64 `json:"direct_map_4k"`
	DirectMap2M       int64 `json:"direct_map_2m"`
}

type ShBasicMemInfo struct {
	MemTotal     int64 `json:"mem_total"`
	MemFree      int64 `json:"mem_free"`
	MemAvailable int64 `json:"mem_available"`
	Buffers      int64 `json:"buffers"`
	Cached       int64 `json:"cached"`
	SwapCached   int64 `json:"swap_cached"`
	SwapTotal    int64 `json:"swap_total"`
	SwapFree     int64 `json:"swap_free"`
}
