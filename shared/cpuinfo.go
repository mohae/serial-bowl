package shared

type ShCPUInfo struct {
	CPUs []ShCPU `json:"cpu"`
}

type ShCPU struct {
	Processor       int16   `json:"processor"`
	VendorID        string  `json:"vendor_id"`
	CPUFamily       string  `json:"cpu_family"`
	Model           string  `json:"model"`
	ModelName       string  `json:"model_name"`
	Stepping        string  `json:"stepping"`
	Microcode       string  `json:"microcode"`
	CPUMHz          float32 `json:"cpu_mhz"`
	CacheSize       string  `json:"cache_size"`
	PhysicalID      int16   `json:"physical_id"`
	Siblings        int16   `json:"siblings"`
	CoreID          int16   `json:"core_id"`
	CPUCores        int16   `json:"cpu_cores"`
	ApicID          int16   `json:"apicid"`
	InitialApicID   int16   `json:"initial_apicid"`
	FPU             string  `json:"fpu"`
	FPUException    string  `json:"fpu_exception"`
	CPUIDLevel      string  `json:"cpuid_level"`
	WP              string  `json:"wp"`
	Flags           string  `json:"flags"`
	BogoMIPS        float32 `json:"bogomips"`
	CLFlushSize     string  `json:"clflush_size"`
	CacheAlignment  string  `json:"cache_alignment"`
	AddressSizes    string  `json:"address_sizes"`
	PowerManagement string  `json:"power_management"`
}
