package shared

type ShMemInfo struct {
	MemTotal          int `json:"mem_total"`
	MemFree           int `json:"mem_free"`
	MemAvailable      int `json:"mem_available"`
	Buffers           int `json:"buffers"`
	Cached            int `json:"cached"`
	SwapCached        int `json:"swap_cached"`
	Active            int `json:"active"`
	Inactive          int `json:"inactive"`
	ActiveAnon        int `json:"active_anon"`
	InactiveAnon      int `json:"inactive_anon"`
	ActiveFile        int `json:"active_file"`
	InactiveFile      int `json:"inactive_file"`
	Unevictable       int `json:"unevictable"`
	Mlocked           int `json:"mlocked"`
	SwapTotal         int `json:"swap_total"`
	SwapFree          int `json:"swap_free"`
	Dirty             int `json:"dirty"`
	Writeback         int `json:"writeback"`
	AnonPages         int `json:"anon_pages"`
	Mapped            int `json:"mapped"`
	Shmem             int `json:"sh_mem"`
	Slab              int `json:"slab"`
	SReclaimable      int `json:"s_reclaimable"`
	SUnreclaim        int `json:"s_unreclaim"`
	KernelStack       int `json:"kernel_stack"`
	NFSUnstable       int `json:"nfs_unstable"`
	Bounce            int `json:"bounce"`
	WritebackTmp      int `json:"writeback_tmp"`
	CommitLimit       int `json:"commit_limit"`
	VmallocTotal      int `json:"vmalloc_total"`
	VmallocUsed       int `json:"vmalloc_used"`
	VmallocChunk      int `json:"vmalloc_chunked"`
	HardwareCorrupted int `json:"hardware_corrupted"`
	AnonHugePages     int `json:"anon_huge_pages"`
	HugePagesTotal    int `json:"huge_pages_total"`
	HugePagesFree     int `json:"huge_pages_free"`
	HugePagesRsvd     int `json:"huge_pages_rsvd"`
	Hugepagesize      int `json:"huge_pages_size"`
	DirectMap4k       int `json:"direct_map_4k"`
	DirectMap2M       int `json:"direct_map_2m"`
}

type ShBasicMemInfo struct {
	MemTotal     int `json:""`
	MemFree      int `json:""`
	MemAvailable int `json:""`
	Buffers      int `json:""`
	Cached       int `json:""`
	SwapCached   int `json:""`
	SwapTotal    int `json:""`
	SwapFree     int `json:""`
}
