// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: cpuinfo.go
// DO NOT EDIT!

package ffjsonbuf

import (
	"bytes"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *CPU) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *CPU) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"processor":`)
	fflib.FormatBits2(buf, uint64(mj.Processor), 10, mj.Processor < 0)
	buf.WriteString(`,"vendor_id":`)
	fflib.WriteJsonString(buf, string(mj.VendorID))
	buf.WriteString(`,"cpu_family":`)
	fflib.WriteJsonString(buf, string(mj.CPUFamily))
	buf.WriteString(`,"model":`)
	fflib.WriteJsonString(buf, string(mj.Model))
	buf.WriteString(`,"model_name":`)
	fflib.WriteJsonString(buf, string(mj.ModelName))
	buf.WriteString(`,"stepping":`)
	fflib.WriteJsonString(buf, string(mj.Stepping))
	buf.WriteString(`,"microcode":`)
	fflib.WriteJsonString(buf, string(mj.Microcode))
	buf.WriteString(`,"cpu_mhz":`)
	fflib.AppendFloat(buf, float64(mj.CPUMHz), 'g', -1, 32)
	buf.WriteString(`,"cache_size":`)
	fflib.WriteJsonString(buf, string(mj.CacheSize))
	buf.WriteString(`,"physical_id":`)
	fflib.FormatBits2(buf, uint64(mj.PhysicalID), 10, mj.PhysicalID < 0)
	buf.WriteString(`,"siblings":`)
	fflib.FormatBits2(buf, uint64(mj.Siblings), 10, mj.Siblings < 0)
	buf.WriteString(`,"core_id":`)
	fflib.FormatBits2(buf, uint64(mj.CoreID), 10, mj.CoreID < 0)
	buf.WriteString(`,"cpu_cores":`)
	fflib.FormatBits2(buf, uint64(mj.CPUCores), 10, mj.CPUCores < 0)
	buf.WriteString(`,"apicid":`)
	fflib.FormatBits2(buf, uint64(mj.ApicID), 10, mj.ApicID < 0)
	buf.WriteString(`,"initial_apicid":`)
	fflib.FormatBits2(buf, uint64(mj.InitialApicID), 10, mj.InitialApicID < 0)
	buf.WriteString(`,"fpu":`)
	fflib.WriteJsonString(buf, string(mj.FPU))
	buf.WriteString(`,"fpu_exception":`)
	fflib.WriteJsonString(buf, string(mj.FPUException))
	buf.WriteString(`,"cpuid_level":`)
	fflib.WriteJsonString(buf, string(mj.CPUIDLevel))
	buf.WriteString(`,"wp":`)
	fflib.WriteJsonString(buf, string(mj.WP))
	buf.WriteString(`,"flags":`)
	fflib.WriteJsonString(buf, string(mj.Flags))
	buf.WriteString(`,"bogomips":`)
	fflib.AppendFloat(buf, float64(mj.BogoMIPS), 'g', -1, 32)
	buf.WriteString(`,"clflush_size":`)
	fflib.WriteJsonString(buf, string(mj.CLFlushSize))
	buf.WriteString(`,"cache_alignment":`)
	fflib.WriteJsonString(buf, string(mj.CacheAlignment))
	buf.WriteString(`,"address_sizes":`)
	fflib.WriteJsonString(buf, string(mj.AddressSizes))
	buf.WriteString(`,"power_management":`)
	fflib.WriteJsonString(buf, string(mj.PowerManagement))
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_CPUbase = iota
	ffj_t_CPUno_such_key

	ffj_t_CPU_Processor

	ffj_t_CPU_VendorID

	ffj_t_CPU_CPUFamily

	ffj_t_CPU_Model

	ffj_t_CPU_ModelName

	ffj_t_CPU_Stepping

	ffj_t_CPU_Microcode

	ffj_t_CPU_CPUMHz

	ffj_t_CPU_CacheSize

	ffj_t_CPU_PhysicalID

	ffj_t_CPU_Siblings

	ffj_t_CPU_CoreID

	ffj_t_CPU_CPUCores

	ffj_t_CPU_ApicID

	ffj_t_CPU_InitialApicID

	ffj_t_CPU_FPU

	ffj_t_CPU_FPUException

	ffj_t_CPU_CPUIDLevel

	ffj_t_CPU_WP

	ffj_t_CPU_Flags

	ffj_t_CPU_BogoMIPS

	ffj_t_CPU_CLFlushSize

	ffj_t_CPU_CacheAlignment

	ffj_t_CPU_AddressSizes

	ffj_t_CPU_PowerManagement
)

var ffj_key_CPU_Processor = []byte("processor")

var ffj_key_CPU_VendorID = []byte("vendor_id")

var ffj_key_CPU_CPUFamily = []byte("cpu_family")

var ffj_key_CPU_Model = []byte("model")

var ffj_key_CPU_ModelName = []byte("model_name")

var ffj_key_CPU_Stepping = []byte("stepping")

var ffj_key_CPU_Microcode = []byte("microcode")

var ffj_key_CPU_CPUMHz = []byte("cpu_mhz")

var ffj_key_CPU_CacheSize = []byte("cache_size")

var ffj_key_CPU_PhysicalID = []byte("physical_id")

var ffj_key_CPU_Siblings = []byte("siblings")

var ffj_key_CPU_CoreID = []byte("core_id")

var ffj_key_CPU_CPUCores = []byte("cpu_cores")

var ffj_key_CPU_ApicID = []byte("apicid")

var ffj_key_CPU_InitialApicID = []byte("initial_apicid")

var ffj_key_CPU_FPU = []byte("fpu")

var ffj_key_CPU_FPUException = []byte("fpu_exception")

var ffj_key_CPU_CPUIDLevel = []byte("cpuid_level")

var ffj_key_CPU_WP = []byte("wp")

var ffj_key_CPU_Flags = []byte("flags")

var ffj_key_CPU_BogoMIPS = []byte("bogomips")

var ffj_key_CPU_CLFlushSize = []byte("clflush_size")

var ffj_key_CPU_CacheAlignment = []byte("cache_alignment")

var ffj_key_CPU_AddressSizes = []byte("address_sizes")

var ffj_key_CPU_PowerManagement = []byte("power_management")

func (uj *CPU) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *CPU) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_CPUbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffj_t_CPUno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffj_key_CPU_ApicID, kn) {
						currentKey = ffj_t_CPU_ApicID
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_CPU_AddressSizes, kn) {
						currentKey = ffj_t_CPU_AddressSizes
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'b':

					if bytes.Equal(ffj_key_CPU_BogoMIPS, kn) {
						currentKey = ffj_t_CPU_BogoMIPS
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'c':

					if bytes.Equal(ffj_key_CPU_CPUFamily, kn) {
						currentKey = ffj_t_CPU_CPUFamily
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_CPU_CPUMHz, kn) {
						currentKey = ffj_t_CPU_CPUMHz
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_CPU_CacheSize, kn) {
						currentKey = ffj_t_CPU_CacheSize
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_CPU_CoreID, kn) {
						currentKey = ffj_t_CPU_CoreID
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_CPU_CPUCores, kn) {
						currentKey = ffj_t_CPU_CPUCores
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_CPU_CPUIDLevel, kn) {
						currentKey = ffj_t_CPU_CPUIDLevel
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_CPU_CLFlushSize, kn) {
						currentKey = ffj_t_CPU_CLFlushSize
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_CPU_CacheAlignment, kn) {
						currentKey = ffj_t_CPU_CacheAlignment
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'f':

					if bytes.Equal(ffj_key_CPU_FPU, kn) {
						currentKey = ffj_t_CPU_FPU
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_CPU_FPUException, kn) {
						currentKey = ffj_t_CPU_FPUException
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_CPU_Flags, kn) {
						currentKey = ffj_t_CPU_Flags
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffj_key_CPU_InitialApicID, kn) {
						currentKey = ffj_t_CPU_InitialApicID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'm':

					if bytes.Equal(ffj_key_CPU_Model, kn) {
						currentKey = ffj_t_CPU_Model
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_CPU_ModelName, kn) {
						currentKey = ffj_t_CPU_ModelName
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_CPU_Microcode, kn) {
						currentKey = ffj_t_CPU_Microcode
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'p':

					if bytes.Equal(ffj_key_CPU_Processor, kn) {
						currentKey = ffj_t_CPU_Processor
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_CPU_PhysicalID, kn) {
						currentKey = ffj_t_CPU_PhysicalID
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_CPU_PowerManagement, kn) {
						currentKey = ffj_t_CPU_PowerManagement
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 's':

					if bytes.Equal(ffj_key_CPU_Stepping, kn) {
						currentKey = ffj_t_CPU_Stepping
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_CPU_Siblings, kn) {
						currentKey = ffj_t_CPU_Siblings
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'v':

					if bytes.Equal(ffj_key_CPU_VendorID, kn) {
						currentKey = ffj_t_CPU_VendorID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'w':

					if bytes.Equal(ffj_key_CPU_WP, kn) {
						currentKey = ffj_t_CPU_WP
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.AsciiEqualFold(ffj_key_CPU_PowerManagement, kn) {
					currentKey = ffj_t_CPU_PowerManagement
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_CPU_AddressSizes, kn) {
					currentKey = ffj_t_CPU_AddressSizes
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffj_key_CPU_CacheAlignment, kn) {
					currentKey = ffj_t_CPU_CacheAlignment
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_CPU_CLFlushSize, kn) {
					currentKey = ffj_t_CPU_CLFlushSize
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_CPU_BogoMIPS, kn) {
					currentKey = ffj_t_CPU_BogoMIPS
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_CPU_Flags, kn) {
					currentKey = ffj_t_CPU_Flags
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_CPU_WP, kn) {
					currentKey = ffj_t_CPU_WP
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffj_key_CPU_CPUIDLevel, kn) {
					currentKey = ffj_t_CPU_CPUIDLevel
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffj_key_CPU_FPUException, kn) {
					currentKey = ffj_t_CPU_FPUException
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_CPU_FPU, kn) {
					currentKey = ffj_t_CPU_FPU
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffj_key_CPU_InitialApicID, kn) {
					currentKey = ffj_t_CPU_InitialApicID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_CPU_ApicID, kn) {
					currentKey = ffj_t_CPU_ApicID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_CPU_CPUCores, kn) {
					currentKey = ffj_t_CPU_CPUCores
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffj_key_CPU_CoreID, kn) {
					currentKey = ffj_t_CPU_CoreID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_CPU_Siblings, kn) {
					currentKey = ffj_t_CPU_Siblings
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_CPU_PhysicalID, kn) {
					currentKey = ffj_t_CPU_PhysicalID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_CPU_CacheSize, kn) {
					currentKey = ffj_t_CPU_CacheSize
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffj_key_CPU_CPUMHz, kn) {
					currentKey = ffj_t_CPU_CPUMHz
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_CPU_Microcode, kn) {
					currentKey = ffj_t_CPU_Microcode
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_CPU_Stepping, kn) {
					currentKey = ffj_t_CPU_Stepping
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffj_key_CPU_ModelName, kn) {
					currentKey = ffj_t_CPU_ModelName
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_CPU_Model, kn) {
					currentKey = ffj_t_CPU_Model
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffj_key_CPU_CPUFamily, kn) {
					currentKey = ffj_t_CPU_CPUFamily
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffj_key_CPU_VendorID, kn) {
					currentKey = ffj_t_CPU_VendorID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_CPU_Processor, kn) {
					currentKey = ffj_t_CPU_Processor
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_CPUno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffj_t_CPU_Processor:
					goto handle_Processor

				case ffj_t_CPU_VendorID:
					goto handle_VendorID

				case ffj_t_CPU_CPUFamily:
					goto handle_CPUFamily

				case ffj_t_CPU_Model:
					goto handle_Model

				case ffj_t_CPU_ModelName:
					goto handle_ModelName

				case ffj_t_CPU_Stepping:
					goto handle_Stepping

				case ffj_t_CPU_Microcode:
					goto handle_Microcode

				case ffj_t_CPU_CPUMHz:
					goto handle_CPUMHz

				case ffj_t_CPU_CacheSize:
					goto handle_CacheSize

				case ffj_t_CPU_PhysicalID:
					goto handle_PhysicalID

				case ffj_t_CPU_Siblings:
					goto handle_Siblings

				case ffj_t_CPU_CoreID:
					goto handle_CoreID

				case ffj_t_CPU_CPUCores:
					goto handle_CPUCores

				case ffj_t_CPU_ApicID:
					goto handle_ApicID

				case ffj_t_CPU_InitialApicID:
					goto handle_InitialApicID

				case ffj_t_CPU_FPU:
					goto handle_FPU

				case ffj_t_CPU_FPUException:
					goto handle_FPUException

				case ffj_t_CPU_CPUIDLevel:
					goto handle_CPUIDLevel

				case ffj_t_CPU_WP:
					goto handle_WP

				case ffj_t_CPU_Flags:
					goto handle_Flags

				case ffj_t_CPU_BogoMIPS:
					goto handle_BogoMIPS

				case ffj_t_CPU_CLFlushSize:
					goto handle_CLFlushSize

				case ffj_t_CPU_CacheAlignment:
					goto handle_CacheAlignment

				case ffj_t_CPU_AddressSizes:
					goto handle_AddressSizes

				case ffj_t_CPU_PowerManagement:
					goto handle_PowerManagement

				case ffj_t_CPUno_such_key:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_Processor:

	/* handler: uj.Processor type=int16 kind=int16 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int16", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 16)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.Processor = int16(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_VendorID:

	/* handler: uj.VendorID type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.VendorID = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_CPUFamily:

	/* handler: uj.CPUFamily type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.CPUFamily = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Model:

	/* handler: uj.Model type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Model = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ModelName:

	/* handler: uj.ModelName type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.ModelName = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Stepping:

	/* handler: uj.Stepping type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Stepping = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Microcode:

	/* handler: uj.Microcode type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Microcode = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_CPUMHz:

	/* handler: uj.CPUMHz type=float32 kind=float32 quoted=false*/

	{
		if tok != fflib.FFTok_double && tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for float32", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseFloat(fs.Output.Bytes(), 32)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.CPUMHz = float32(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_CacheSize:

	/* handler: uj.CacheSize type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.CacheSize = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_PhysicalID:

	/* handler: uj.PhysicalID type=int16 kind=int16 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int16", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 16)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.PhysicalID = int16(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Siblings:

	/* handler: uj.Siblings type=int16 kind=int16 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int16", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 16)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.Siblings = int16(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_CoreID:

	/* handler: uj.CoreID type=int16 kind=int16 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int16", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 16)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.CoreID = int16(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_CPUCores:

	/* handler: uj.CPUCores type=int16 kind=int16 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int16", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 16)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.CPUCores = int16(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ApicID:

	/* handler: uj.ApicID type=int16 kind=int16 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int16", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 16)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.ApicID = int16(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_InitialApicID:

	/* handler: uj.InitialApicID type=int16 kind=int16 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int16", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 16)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.InitialApicID = int16(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_FPU:

	/* handler: uj.FPU type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.FPU = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_FPUException:

	/* handler: uj.FPUException type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.FPUException = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_CPUIDLevel:

	/* handler: uj.CPUIDLevel type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.CPUIDLevel = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_WP:

	/* handler: uj.WP type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.WP = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Flags:

	/* handler: uj.Flags type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Flags = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_BogoMIPS:

	/* handler: uj.BogoMIPS type=float32 kind=float32 quoted=false*/

	{
		if tok != fflib.FFTok_double && tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for float32", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseFloat(fs.Output.Bytes(), 32)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.BogoMIPS = float32(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_CLFlushSize:

	/* handler: uj.CLFlushSize type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.CLFlushSize = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_CacheAlignment:

	/* handler: uj.CacheAlignment type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.CacheAlignment = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_AddressSizes:

	/* handler: uj.AddressSizes type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.AddressSizes = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_PowerManagement:

	/* handler: uj.PowerManagement type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.PowerManagement = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:
	return nil
}

func (mj *CPUInfo) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *CPUInfo) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"cpu":`)
	if mj.CPUs != nil {
		buf.WriteString(`[`)
		for i, v := range mj.CPUs {
			if i != 0 {
				buf.WriteString(`,`)
			}

			{

				err = v.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_CPUInfobase = iota
	ffj_t_CPUInfono_such_key

	ffj_t_CPUInfo_CPUs
)

var ffj_key_CPUInfo_CPUs = []byte("cpu")

func (uj *CPUInfo) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *CPUInfo) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_CPUInfobase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffj_t_CPUInfono_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'c':

					if bytes.Equal(ffj_key_CPUInfo_CPUs, kn) {
						currentKey = ffj_t_CPUInfo_CPUs
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffj_key_CPUInfo_CPUs, kn) {
					currentKey = ffj_t_CPUInfo_CPUs
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_CPUInfono_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffj_t_CPUInfo_CPUs:
					goto handle_CPUs

				case ffj_t_CPUInfono_such_key:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_CPUs:

	/* handler: uj.CPUs type=[]ffjsonbuf.CPU kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.CPUs = nil
		} else {

			uj.CPUs = make([]CPU, 0)

			wantVal := true

			for {

				var tmp_uj__CPUs CPU

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmp_uj__CPUs type=ffjsonbuf.CPU kind=struct quoted=false*/

				{
					if tok == fflib.FFTok_null {

						state = fflib.FFParse_after_value
						goto mainparse
					}

					err = tmp_uj__CPUs.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
					if err != nil {
						return err
					}
					state = fflib.FFParse_after_value
				}

				uj.CPUs = append(uj.CPUs, tmp_uj__CPUs)
				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:
	return nil
}
