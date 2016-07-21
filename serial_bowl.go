package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/mohae/benchutil"
	"github.com/mohae/serial-bowl/capnp"
	"github.com/mohae/serial-bowl/fb"
	"github.com/mohae/serial-bowl/ffjson"
	"github.com/mohae/serial-bowl/ffjsonbuf"
	"github.com/mohae/serial-bowl/gencode"
	"github.com/mohae/serial-bowl/gobs"
	"github.com/mohae/serial-bowl/gogo/gofast"
	"github.com/mohae/serial-bowl/gogo/gogofast"
	"github.com/mohae/serial-bowl/gogo/gogofaster"
	"github.com/mohae/serial-bowl/gogo/gogoslick"
	"github.com/mohae/serial-bowl/jsn"
	"github.com/mohae/serial-bowl/pb"
	"github.com/mohae/serial-bowl/shared"
	"github.com/mohae/serial-bowl/tmsgp"
	"github.com/mohae/serial-bowl/vmsgpack"
)

// flags
var (
	output         string
	format         string
	benchFormats   stringArray
	section        bool
	sectionHeaders bool
	nameSections   bool
	systemInfo     bool
	datasetSize    int // number of elements in test slice
)

type stringArray []string

func (f *stringArray) String() string {
	var s string
	for i, v := range *f {
		if i == 0 {
			s = v
			continue
		}
		s += ", " + v

	}
	return s
}

func (f *stringArray) Set(v string) error {
	*f = append(*f, v)
	return nil
}

func init() {
	flag.StringVar(&output, "output", "stdout", "output destination")
	flag.StringVar(&output, "o", "stdout", "output destination (short)")
	flag.StringVar(&format, "format", "txt", "format of output")
	flag.StringVar(&format, "f", "txt", "format of output")
	flag.BoolVar(&nameSections, "namesections", false, "use group as section name: some restrictions apply")
	flag.BoolVar(&nameSections, "n", false, "use group as section name: some restrictions apply")
	flag.BoolVar(&section, "sections", false, "don't separate groups of tests into sections")
	flag.BoolVar(&section, "s", false, "don't separate groups of tests into sections")
	flag.BoolVar(&sectionHeaders, "sectionheader", false, "if there are sections, add a section header row")
	flag.BoolVar(&sectionHeaders, "h", false, "if there are sections, add a section header row")
	flag.BoolVar(&systemInfo, "sysinfo", false, "add the system information to the output")
	flag.BoolVar(&systemInfo, "i", false, "add the system information to the output")
	flag.IntVar(&datasetSize, "datasetsize", 1000, "the size of the test dataset for each benchmark")
	flag.IntVar(&datasetSize, "d", 1000, "the size of the test dataset for each benchmark")
	flag.Var(&benchFormats, "benchformat", "format(s) to benchmark")
	flag.Var(&benchFormats, "b", "format(s) to benchmark")
}

func main() {
	flag.Parse()
	runFormats, err := processBenchFormats()
	if err != nil {
		fmt.Println(err)
		return
	}
	done := make(chan struct{})
	// start the visual ticker
	go benchutil.Dot(done)
	shared.Len = datasetSize
	// set the output
	var w io.Writer
	switch output {
	case "stdout":
		w = os.Stdout
	default:
		w, err = os.OpenFile(output, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer w.(*os.File).Close()
	}

	// generate all of the test data
	shared.GenData()

	// get the benchmark for the desired format
	// process the output
	var bench benchutil.Benchmarker
	switch format {
	case "csv":
		bench = benchutil.NewCSVBench(w)
	case "md":
		bench = benchutil.NewMDBench(w)
	default:
		bench = benchutil.NewStringBench(w)
	}
	bench.SectionPerGroup(section)
	bench.SectionHeaders(sectionHeaders)
	bench.IncludeSystemInfo(systemInfo)
	bench.NameSections(nameSections)

	// set the header info
	bench.SetGroupColumnHeader("Data Type")
	bench.SetSubGroupColumnHeader("Format")
	bench.SetNameColumnHeader("Package")
	bench.SetDescColumnHeader("Operation")

	benchBasicMemInfo(runFormats, bench)
	benchCPUInfo(runFormats, bench)
	benchMemInfo(runFormats, bench)
	benchMessage(runFormats, bench)
	benchRedditAccount(runFormats, bench)

	close(done)
	fmt.Println("\ngenerating output...\n")
	err = bench.Out()
	if err != nil {
		fmt.Printf("error generating output: %s\n", err)
	}
}

func benchBasicMemInfo(formats []shared.Format, bench benchutil.Benchmarker) {
	// Only run the benchmarks for the requested formats.
	var b []benchutil.Bench
	for _, v := range formats {
		switch v {
		case shared.CapnProto2: // Captain Proto 2
			b = capnp.BenchBasicMemInfo()
			bench.Append(b...)
		case shared.Flatbuffers: // Flatbuffers
			b = fb.BenchBasicMemInfo()
			bench.Append(b...)
		case shared.GenCode: // Gencode
			b = gencode.BenchBasicMemInfo()
			bench.Append(b...)
		case shared.Gob: // Gob
			b = gobs.BenchBasicMemInfo()
			bench.Append(b...)
		case shared.JSON: // JSON
			b = jsn.BenchBasicMemInfo()
			bench.Append(b...)
			// FFJSON
			b = ffjson.BenchBasicMemInfo()
			bench.Append(b...)
			// FFJSON Buf
			b = ffjsonbuf.BenchBasicMemInfo()
			bench.Append(b...)
		case shared.MessagePack: // Message Pack
			// vmhailenco/msgPack
			b = vmsgpack.BenchBasicMemInfo()
			bench.Append(b...)
			// tinylib/msgp
			b = tmsgp.BenchBasicMemInfo()
			bench.Append(b...)
		case shared.ProtobufV3: // Protocol Buffers V3
			// PBv3
			b = pb.BenchBasicMemInfo()
			bench.Append(b...)
			// gofast
			b = gofast.BenchBasicMemInfo()
			bench.Append(b...)
			// gogofast
			b = gogofast.BenchBasicMemInfo()
			bench.Append(b...)
			// gogofaster
			b = gogofaster.BenchBasicMemInfo()
			bench.Append(b...)
			// gogoslick
			b = gogoslick.BenchBasicMemInfo()
			bench.Append(b...)
		}
	}
}

func benchMemInfo(formats []shared.Format, bench benchutil.Benchmarker) {
	var b []benchutil.Bench
	// Only run the benchmarks for the requested formats.
	for _, v := range formats {
		switch v {
		case shared.CapnProto2: // Captain Proto 2
			b = capnp.BenchMemInfo()
			bench.Append(b...)
		case shared.Flatbuffers: // Flatbuffers
			b = fb.BenchMemInfo()
			bench.Append(b...)
		case shared.GenCode: // Gencode
			b = gencode.BenchMemInfo()
			bench.Append(b...)
		case shared.Gob: // Gob
			b = gobs.BenchMemInfo()
			bench.Append(b...)
		case shared.JSON: // JSON
			b = jsn.BenchMemInfo()
			bench.Append(b...)
			// FFJSON
			b = ffjson.BenchMemInfo()
			bench.Append(b...)
			// FFJSON Buf
			b = ffjsonbuf.BenchMemInfo()
			bench.Append(b...)
		case shared.MessagePack: // Message Pack
			// vmhailenco/msgPack
			b = vmsgpack.BenchMemInfo()
			bench.Append(b...)
			// tinylib/msgp
			b = tmsgp.BenchMemInfo()
			bench.Append(b...)
		case shared.ProtobufV3: // Protocol Buffers V3
			// PBv3
			b = pb.BenchMemInfo()
			bench.Append(b...)
			// gofast
			b = gofast.BenchMemInfo()
			bench.Append(b...)
			// gogofast
			b = gogofast.BenchMemInfo()
			bench.Append(b...)
			// gogofaster
			b = gogofaster.BenchMemInfo()
			bench.Append(b...)
			// gogoslick
			b = gogoslick.BenchMemInfo()
			bench.Append(b...)
		}
	}
}

func benchMessage(formats []shared.Format, bench benchutil.Benchmarker) {
	var b []benchutil.Bench

	// Message Data
	dataLen := []int{16, 64, 256, 1024, 2048, 4096}
	for _, v := range dataLen {
		shared.GenMessageData(v, shared.Len)
		// Only run the benchmarks for the requested formats.
		for _, f := range formats {
			switch f {
			case shared.CapnProto2: // Captain Proto 2
				// TODO: 4096 Bytes of data causes the following error:
				// capnp: NewMessage called on arena with data
				// figure out how to resolve.
				if v > 2048 {
					continue
				}
				b = capnp.BenchMessage(v)
				bench.Append(b...)
			case shared.Flatbuffers: // Flatbuffers
				b = fb.BenchMessage(v)
				bench.Append(b...)
			case shared.GenCode: // Gencode
				b = gencode.BenchMessage(v)
				bench.Append(b...)
			case shared.Gob: // Gob
				// Gob
				b = gobs.BenchMessage(v)
				bench.Append(b...)
			case shared.JSON: // JSON
				// JSON
				b = jsn.BenchMessage(v)
				bench.Append(b...)
				// FFJSON
				b = ffjson.BenchMessage(v)
				bench.Append(b...)
				// FFJSONBuf
				b = ffjsonbuf.BenchMessage(v)
				bench.Append(b...)
			case shared.MessagePack: // Message Pack
				// vmhailenco/msgPack
				b = vmsgpack.BenchMessage(v)
				bench.Append(b...)
				// tinylib/msgp
				b = tmsgp.BenchMessage(v)
				bench.Append(b...)
			case shared.ProtobufV3: // Protocol Buffer V3
				// PBv3
				b = pb.BenchMessage(v)
				bench.Append(b...)
				// gofast
				b = gofast.BenchMessage(v)
				bench.Append(b...)
				// gogofast
				b = gogofast.BenchMessage(v)
				bench.Append(b...)
				// gogofaster
				b = gogofaster.BenchMessage(v)
				bench.Append(b...)
				// gogoslick
				b = gogoslick.BenchMessage(v)
				bench.Append(b...)
			}
		}
	}
}

func benchRedditAccount(formats []shared.Format, bench benchutil.Benchmarker) {
	var b []benchutil.Bench
	// Only run the benchmarks for the requested formats.
	for _, v := range formats {
		switch v {
		case shared.CapnProto2: // Captain Proto 2
			b = capnp.BenchRedditAccount()
			bench.Append(b...)
		case shared.Flatbuffers: // Flatbuffers
			b = fb.BenchRedditAccount()
			bench.Append(b...)
		case shared.GenCode: // Gencode
			b = gencode.BenchRedditAccount()
			bench.Append(b...)
		case shared.Gob: // Gob
			b = gobs.BenchRedditAccount()
			bench.Append(b...)
		case shared.JSON: // JSON
			// JSON
			b = jsn.BenchRedditAccount()
			bench.Append(b...)
			// FFJSON
			b = ffjson.BenchRedditAccount()
			bench.Append(b...)
			// FFJSONBuf
			b = ffjsonbuf.BenchRedditAccount()
			bench.Append(b...)
		case shared.MessagePack: // Message Pack
			// vmhailenco/msgPack
			b = vmsgpack.BenchRedditAccount()
			bench.Append(b...)
			// tinylib/msgp
			b = tmsgp.BenchRedditAccount()
			bench.Append(b...)
		case shared.ProtobufV3: // Protocol Buffer V3
			// PB v3
			b = pb.BenchRedditAccount()
			bench.Append(b...)
			// gofast
			b = gofast.BenchRedditAccount()
			bench.Append(b...)
			// gogofast
			b = gogofast.BenchRedditAccount()
			bench.Append(b...)
			// gogofaster
			b = gogofaster.BenchRedditAccount()
			bench.Append(b...)
			// gogoslick
			b = gogoslick.BenchRedditAccount()
			bench.Append(b...)
		}
	}
}

func benchCPUInfo(formats []shared.Format, bench benchutil.Benchmarker) {
	var b []benchutil.Bench
	// num cpus: it's n + 1
	cpus := []int{1, 4, 8, 16}
	for _, n := range cpus {
		shared.GenCPUInfoData(n, shared.Len)
		// Only run the benchmarks for the requested formats.
		for _, v := range formats {
			switch v {
			case shared.CapnProto2: // Captain Proto 2
				b = capnp.BenchCPUInfo(n)
				bench.Append(b...)
			case shared.Flatbuffers: // Flatbuffers
				b = fb.BenchCPUInfo(n)
				bench.Append(b...)
			case shared.GenCode: // Gencode
				b = gencode.BenchCPUInfo(n)
				bench.Append(b...)
			case shared.Gob: // Gob
				b = gobs.BenchCPUInfo(n)
				bench.Append(b...)
			case shared.JSON: // JSON
				b = jsn.BenchCPUInfo(n)
				bench.Append(b...)
				// FFJSON
				b = ffjson.BenchCPUInfo(n)
				bench.Append(b...)
				// FFJSONBuf
				b = ffjsonbuf.BenchCPUInfo(n)
				bench.Append(b...)
			case shared.MessagePack: // Message Pack
				// vmhailenco/msgPack
				b = vmsgpack.BenchCPUInfo(n)
				bench.Append(b...)
				// tinylib/msgp
				b = tmsgp.BenchCPUInfo(n)
				bench.Append(b...)
			case shared.ProtobufV3: // Protocol Buffers V3
				// PB v3
				b = pb.BenchCPUInfo(n)
				bench.Append(b...)
				// gofast
				b = gofast.BenchCPUInfo(n)
				bench.Append(b...)
				// gogofast
				b = gogofast.BenchCPUInfo(n)
				bench.Append(b...)
				// gogofaster
				b = gogofaster.BenchCPUInfo(n)
				bench.Append(b...)
				// gogoslick
				b = gogoslick.BenchCPUInfo(n)
				bench.Append(b...)
			}
		}
	}
}

func processBenchFormats() ([]shared.Format, error) {
	var formats []shared.Format
	// If nothing was specified, everything should be run.
	if len(benchFormats) == 0 {
		return shared.Formats, nil
	}
	// otherwsie go through the input and try to match.  Error on anything we can't
	// figure out.
	for _, v := range benchFormats {
		switch v {
		case "fb", "flat", "flatbuffer", "flatbuffers":
			formats = append(formats, shared.Flatbuffers)
		case "gencode":
			formats = append(formats, shared.GenCode)
		case "gob":
			formats = append(formats, shared.Gob)
		case "json", "jsn":
			formats = append(formats, shared.JSON)
		case "capn", "capnp", "cpn", "capnproto", "captainproto", "capnproto2":
			formats = append(formats, shared.CapnProto2)
		case "protobuf", "pb", "protobufv3":
			formats = append(formats, shared.ProtobufV3)
		default:
			return nil, fmt.Errorf("%s: unknown serialization format", v)
		}
	}
	return formats, nil
}
