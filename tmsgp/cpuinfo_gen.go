package tmsgp

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *CPU) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "processor":
			z.Processor, err = dc.ReadInt16()
			if err != nil {
				return
			}
		case "vendor_id":
			z.VendorID, err = dc.ReadString()
			if err != nil {
				return
			}
		case "cpu_family":
			z.CPUFamily, err = dc.ReadString()
			if err != nil {
				return
			}
		case "model":
			z.Model, err = dc.ReadString()
			if err != nil {
				return
			}
		case "model_name":
			z.ModelName, err = dc.ReadString()
			if err != nil {
				return
			}
		case "stepping":
			z.Stepping, err = dc.ReadString()
			if err != nil {
				return
			}
		case "microcode":
			z.Microcode, err = dc.ReadString()
			if err != nil {
				return
			}
		case "cpu_mhz":
			z.CPUMHz, err = dc.ReadFloat32()
			if err != nil {
				return
			}
		case "cache_size":
			z.CacheSize, err = dc.ReadString()
			if err != nil {
				return
			}
		case "physical_id":
			z.PhysicalID, err = dc.ReadInt16()
			if err != nil {
				return
			}
		case "siblings":
			z.Siblings, err = dc.ReadInt16()
			if err != nil {
				return
			}
		case "core_id":
			z.CoreID, err = dc.ReadInt16()
			if err != nil {
				return
			}
		case "cpu_cores":
			z.CPUCores, err = dc.ReadInt16()
			if err != nil {
				return
			}
		case "apicid":
			z.ApicID, err = dc.ReadInt16()
			if err != nil {
				return
			}
		case "initial_apicid":
			z.InitialApicID, err = dc.ReadInt16()
			if err != nil {
				return
			}
		case "fpu":
			z.FPU, err = dc.ReadString()
			if err != nil {
				return
			}
		case "fpu_exception":
			z.FPUException, err = dc.ReadString()
			if err != nil {
				return
			}
		case "cpuid_level":
			z.CPUIDLevel, err = dc.ReadString()
			if err != nil {
				return
			}
		case "wp":
			z.WP, err = dc.ReadString()
			if err != nil {
				return
			}
		case "flags":
			z.Flags, err = dc.ReadString()
			if err != nil {
				return
			}
		case "bogomips":
			z.BogoMIPS, err = dc.ReadFloat32()
			if err != nil {
				return
			}
		case "clflush_size":
			z.CLFlushSize, err = dc.ReadString()
			if err != nil {
				return
			}
		case "cache_alignment":
			z.CacheAlignment, err = dc.ReadString()
			if err != nil {
				return
			}
		case "address_sizes":
			z.AddressSizes, err = dc.ReadString()
			if err != nil {
				return
			}
		case "power_management":
			z.PowerManagement, err = dc.ReadString()
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
func (z *CPU) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 25
	// write "processor"
	err = en.Append(0xde, 0x0, 0x19, 0xa9, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteInt16(z.Processor)
	if err != nil {
		return
	}
	// write "vendor_id"
	err = en.Append(0xa9, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteString(z.VendorID)
	if err != nil {
		return
	}
	// write "cpu_family"
	err = en.Append(0xaa, 0x63, 0x70, 0x75, 0x5f, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteString(z.CPUFamily)
	if err != nil {
		return
	}
	// write "model"
	err = en.Append(0xa5, 0x6d, 0x6f, 0x64, 0x65, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Model)
	if err != nil {
		return
	}
	// write "model_name"
	err = en.Append(0xaa, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ModelName)
	if err != nil {
		return
	}
	// write "stepping"
	err = en.Append(0xa8, 0x73, 0x74, 0x65, 0x70, 0x70, 0x69, 0x6e, 0x67)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Stepping)
	if err != nil {
		return
	}
	// write "microcode"
	err = en.Append(0xa9, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x63, 0x6f, 0x64, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Microcode)
	if err != nil {
		return
	}
	// write "cpu_mhz"
	err = en.Append(0xa7, 0x63, 0x70, 0x75, 0x5f, 0x6d, 0x68, 0x7a)
	if err != nil {
		return err
	}
	err = en.WriteFloat32(z.CPUMHz)
	if err != nil {
		return
	}
	// write "cache_size"
	err = en.Append(0xaa, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.CacheSize)
	if err != nil {
		return
	}
	// write "physical_id"
	err = en.Append(0xab, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt16(z.PhysicalID)
	if err != nil {
		return
	}
	// write "siblings"
	err = en.Append(0xa8, 0x73, 0x69, 0x62, 0x6c, 0x69, 0x6e, 0x67, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteInt16(z.Siblings)
	if err != nil {
		return
	}
	// write "core_id"
	err = en.Append(0xa7, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt16(z.CoreID)
	if err != nil {
		return
	}
	// write "cpu_cores"
	err = en.Append(0xa9, 0x63, 0x70, 0x75, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteInt16(z.CPUCores)
	if err != nil {
		return
	}
	// write "apicid"
	err = en.Append(0xa6, 0x61, 0x70, 0x69, 0x63, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt16(z.ApicID)
	if err != nil {
		return
	}
	// write "initial_apicid"
	err = en.Append(0xae, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x61, 0x70, 0x69, 0x63, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt16(z.InitialApicID)
	if err != nil {
		return
	}
	// write "fpu"
	err = en.Append(0xa3, 0x66, 0x70, 0x75)
	if err != nil {
		return err
	}
	err = en.WriteString(z.FPU)
	if err != nil {
		return
	}
	// write "fpu_exception"
	err = en.Append(0xad, 0x66, 0x70, 0x75, 0x5f, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteString(z.FPUException)
	if err != nil {
		return
	}
	// write "cpuid_level"
	err = en.Append(0xab, 0x63, 0x70, 0x75, 0x69, 0x64, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.CPUIDLevel)
	if err != nil {
		return
	}
	// write "wp"
	err = en.Append(0xa2, 0x77, 0x70)
	if err != nil {
		return err
	}
	err = en.WriteString(z.WP)
	if err != nil {
		return
	}
	// write "flags"
	err = en.Append(0xa5, 0x66, 0x6c, 0x61, 0x67, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Flags)
	if err != nil {
		return
	}
	// write "bogomips"
	err = en.Append(0xa8, 0x62, 0x6f, 0x67, 0x6f, 0x6d, 0x69, 0x70, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteFloat32(z.BogoMIPS)
	if err != nil {
		return
	}
	// write "clflush_size"
	err = en.Append(0xac, 0x63, 0x6c, 0x66, 0x6c, 0x75, 0x73, 0x68, 0x5f, 0x73, 0x69, 0x7a, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.CLFlushSize)
	if err != nil {
		return
	}
	// write "cache_alignment"
	err = en.Append(0xaf, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteString(z.CacheAlignment)
	if err != nil {
		return
	}
	// write "address_sizes"
	err = en.Append(0xad, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteString(z.AddressSizes)
	if err != nil {
		return
	}
	// write "power_management"
	err = en.Append(0xb0, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteString(z.PowerManagement)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *CPU) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 25
	// string "processor"
	o = append(o, 0xde, 0x0, 0x19, 0xa9, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72)
	o = msgp.AppendInt16(o, z.Processor)
	// string "vendor_id"
	o = append(o, 0xa9, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x69, 0x64)
	o = msgp.AppendString(o, z.VendorID)
	// string "cpu_family"
	o = append(o, 0xaa, 0x63, 0x70, 0x75, 0x5f, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79)
	o = msgp.AppendString(o, z.CPUFamily)
	// string "model"
	o = append(o, 0xa5, 0x6d, 0x6f, 0x64, 0x65, 0x6c)
	o = msgp.AppendString(o, z.Model)
	// string "model_name"
	o = append(o, 0xaa, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.ModelName)
	// string "stepping"
	o = append(o, 0xa8, 0x73, 0x74, 0x65, 0x70, 0x70, 0x69, 0x6e, 0x67)
	o = msgp.AppendString(o, z.Stepping)
	// string "microcode"
	o = append(o, 0xa9, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x63, 0x6f, 0x64, 0x65)
	o = msgp.AppendString(o, z.Microcode)
	// string "cpu_mhz"
	o = append(o, 0xa7, 0x63, 0x70, 0x75, 0x5f, 0x6d, 0x68, 0x7a)
	o = msgp.AppendFloat32(o, z.CPUMHz)
	// string "cache_size"
	o = append(o, 0xaa, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65)
	o = msgp.AppendString(o, z.CacheSize)
	// string "physical_id"
	o = append(o, 0xab, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x69, 0x64)
	o = msgp.AppendInt16(o, z.PhysicalID)
	// string "siblings"
	o = append(o, 0xa8, 0x73, 0x69, 0x62, 0x6c, 0x69, 0x6e, 0x67, 0x73)
	o = msgp.AppendInt16(o, z.Siblings)
	// string "core_id"
	o = append(o, 0xa7, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64)
	o = msgp.AppendInt16(o, z.CoreID)
	// string "cpu_cores"
	o = append(o, 0xa9, 0x63, 0x70, 0x75, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x73)
	o = msgp.AppendInt16(o, z.CPUCores)
	// string "apicid"
	o = append(o, 0xa6, 0x61, 0x70, 0x69, 0x63, 0x69, 0x64)
	o = msgp.AppendInt16(o, z.ApicID)
	// string "initial_apicid"
	o = append(o, 0xae, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x61, 0x70, 0x69, 0x63, 0x69, 0x64)
	o = msgp.AppendInt16(o, z.InitialApicID)
	// string "fpu"
	o = append(o, 0xa3, 0x66, 0x70, 0x75)
	o = msgp.AppendString(o, z.FPU)
	// string "fpu_exception"
	o = append(o, 0xad, 0x66, 0x70, 0x75, 0x5f, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e)
	o = msgp.AppendString(o, z.FPUException)
	// string "cpuid_level"
	o = append(o, 0xab, 0x63, 0x70, 0x75, 0x69, 0x64, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c)
	o = msgp.AppendString(o, z.CPUIDLevel)
	// string "wp"
	o = append(o, 0xa2, 0x77, 0x70)
	o = msgp.AppendString(o, z.WP)
	// string "flags"
	o = append(o, 0xa5, 0x66, 0x6c, 0x61, 0x67, 0x73)
	o = msgp.AppendString(o, z.Flags)
	// string "bogomips"
	o = append(o, 0xa8, 0x62, 0x6f, 0x67, 0x6f, 0x6d, 0x69, 0x70, 0x73)
	o = msgp.AppendFloat32(o, z.BogoMIPS)
	// string "clflush_size"
	o = append(o, 0xac, 0x63, 0x6c, 0x66, 0x6c, 0x75, 0x73, 0x68, 0x5f, 0x73, 0x69, 0x7a, 0x65)
	o = msgp.AppendString(o, z.CLFlushSize)
	// string "cache_alignment"
	o = append(o, 0xaf, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74)
	o = msgp.AppendString(o, z.CacheAlignment)
	// string "address_sizes"
	o = append(o, 0xad, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x73)
	o = msgp.AppendString(o, z.AddressSizes)
	// string "power_management"
	o = append(o, 0xb0, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74)
	o = msgp.AppendString(o, z.PowerManagement)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *CPU) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "processor":
			z.Processor, bts, err = msgp.ReadInt16Bytes(bts)
			if err != nil {
				return
			}
		case "vendor_id":
			z.VendorID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "cpu_family":
			z.CPUFamily, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "model":
			z.Model, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "model_name":
			z.ModelName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "stepping":
			z.Stepping, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "microcode":
			z.Microcode, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "cpu_mhz":
			z.CPUMHz, bts, err = msgp.ReadFloat32Bytes(bts)
			if err != nil {
				return
			}
		case "cache_size":
			z.CacheSize, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "physical_id":
			z.PhysicalID, bts, err = msgp.ReadInt16Bytes(bts)
			if err != nil {
				return
			}
		case "siblings":
			z.Siblings, bts, err = msgp.ReadInt16Bytes(bts)
			if err != nil {
				return
			}
		case "core_id":
			z.CoreID, bts, err = msgp.ReadInt16Bytes(bts)
			if err != nil {
				return
			}
		case "cpu_cores":
			z.CPUCores, bts, err = msgp.ReadInt16Bytes(bts)
			if err != nil {
				return
			}
		case "apicid":
			z.ApicID, bts, err = msgp.ReadInt16Bytes(bts)
			if err != nil {
				return
			}
		case "initial_apicid":
			z.InitialApicID, bts, err = msgp.ReadInt16Bytes(bts)
			if err != nil {
				return
			}
		case "fpu":
			z.FPU, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "fpu_exception":
			z.FPUException, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "cpuid_level":
			z.CPUIDLevel, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "wp":
			z.WP, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "flags":
			z.Flags, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "bogomips":
			z.BogoMIPS, bts, err = msgp.ReadFloat32Bytes(bts)
			if err != nil {
				return
			}
		case "clflush_size":
			z.CLFlushSize, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "cache_alignment":
			z.CacheAlignment, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "address_sizes":
			z.AddressSizes, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "power_management":
			z.PowerManagement, bts, err = msgp.ReadStringBytes(bts)
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

func (z *CPU) Msgsize() (s int) {
	s = 3 + 10 + msgp.Int16Size + 10 + msgp.StringPrefixSize + len(z.VendorID) + 11 + msgp.StringPrefixSize + len(z.CPUFamily) + 6 + msgp.StringPrefixSize + len(z.Model) + 11 + msgp.StringPrefixSize + len(z.ModelName) + 9 + msgp.StringPrefixSize + len(z.Stepping) + 10 + msgp.StringPrefixSize + len(z.Microcode) + 8 + msgp.Float32Size + 11 + msgp.StringPrefixSize + len(z.CacheSize) + 12 + msgp.Int16Size + 9 + msgp.Int16Size + 8 + msgp.Int16Size + 10 + msgp.Int16Size + 7 + msgp.Int16Size + 15 + msgp.Int16Size + 4 + msgp.StringPrefixSize + len(z.FPU) + 14 + msgp.StringPrefixSize + len(z.FPUException) + 12 + msgp.StringPrefixSize + len(z.CPUIDLevel) + 3 + msgp.StringPrefixSize + len(z.WP) + 6 + msgp.StringPrefixSize + len(z.Flags) + 9 + msgp.Float32Size + 13 + msgp.StringPrefixSize + len(z.CLFlushSize) + 16 + msgp.StringPrefixSize + len(z.CacheAlignment) + 14 + msgp.StringPrefixSize + len(z.AddressSizes) + 17 + msgp.StringPrefixSize + len(z.PowerManagement)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *CPUInfo) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "CPUs":
			var xsz uint32
			xsz, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.CPUs) >= int(xsz) {
				z.CPUs = z.CPUs[:xsz]
			} else {
				z.CPUs = make([]CPU, xsz)
			}
			for xvk := range z.CPUs {
				err = z.CPUs[xvk].DecodeMsg(dc)
				if err != nil {
					return
				}
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
func (z *CPUInfo) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "CPUs"
	err = en.Append(0x81, 0xa4, 0x43, 0x50, 0x55, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.CPUs)))
	if err != nil {
		return
	}
	for xvk := range z.CPUs {
		err = z.CPUs[xvk].EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *CPUInfo) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "CPUs"
	o = append(o, 0x81, 0xa4, 0x43, 0x50, 0x55, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.CPUs)))
	for xvk := range z.CPUs {
		o, err = z.CPUs[xvk].MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *CPUInfo) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "CPUs":
			var xsz uint32
			xsz, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.CPUs) >= int(xsz) {
				z.CPUs = z.CPUs[:xsz]
			} else {
				z.CPUs = make([]CPU, xsz)
			}
			for xvk := range z.CPUs {
				bts, err = z.CPUs[xvk].UnmarshalMsg(bts)
				if err != nil {
					return
				}
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

func (z *CPUInfo) Msgsize() (s int) {
	s = 1 + 5 + msgp.ArrayHeaderSize
	for xvk := range z.CPUs {
		s += z.CPUs[xvk].Msgsize()
	}
	return
}
