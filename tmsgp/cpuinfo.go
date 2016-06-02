package tmsgp

//go:generate msgp
type CPUInfo struct {
	CPUs []CPU `msg:cpus`
}

type CPU struct {
	Processor       int16   `msg:"processor"`
	VendorID        string  `msg:"vendor_id"`
	CPUFamily       string  `msg:"cpu_family"`
	Model           string  `msg:"model"`
	ModelName       string  `msg:"model_name"`
	Stepping        string  `msg:"stepping"`
	Microcode       string  `msg:"microcode"`
	CPUMHz          float32 `msg:"cpu_mhz"`
	CacheSize       string  `msg:"cache_size"`
	PhysicalID      int16   `msg:"physical_id"`
	Siblings        int16   `msg:"siblings"`
	CoreID          int16   `msg:"core_id"`
	CPUCores        int16   `msg:"cpu_cores"`
	ApicID          int16   `msg:"apicid"`
	InitialApicID   int16   `msg:"initial_apicid"`
	FPU             string  `msg:"fpu"`
	FPUException    string  `msg:"fpu_exception"`
	CPUIDLevel      string  `msg:"cpuid_level"`
	WP              string  `msg:"wp"`
	Flags           string  `msg:"flags"`
	BogoMIPS        float32 `msg:"bogomips"`
	CLFlushSize     string  `msg:"clflush_size"`
	CacheAlignment  string  `msg:"cache_alignment"`
	AddressSizes    string  `msg:"address_sizes"`
	PowerManagement string  `msg:"power_management"`
}
