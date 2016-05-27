package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	pcg "github.com/dgryski/go-pcgr"
	"github.com/mohae/benchutil"
	"github.com/mohae/serial-bowl/capnp"
	"github.com/mohae/serial-bowl/fb"
	"github.com/mohae/serial-bowl/ffjson"
	"github.com/mohae/serial-bowl/ffjsonbuf"
	"github.com/mohae/serial-bowl/gencode"
	"github.com/mohae/serial-bowl/gobs"
	"github.com/mohae/serial-bowl/jsn"
	"github.com/mohae/serial-bowl/pb"
	"github.com/mohae/serial-bowl/shared"
	"github.com/mohae/serial-bowl/vmsgpack"
)

// flags
var (
	output         string
	format         string
	section        bool
	sectionHeaders bool
	nameSections   bool
	systemInfo     bool
)

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
}

func main() {
	flag.Parse()
	done := make(chan struct{})
	// start the visual ticker
	go benchutil.Dot(done)

	// set the output
	var w io.Writer
	var err error
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
		bench.(*benchutil.MDBench).GroupAsSectionName = nameSections
	default:
		bench = benchutil.NewStringBench(w)
	}
	bench.SectionPerGroup(section)
	bench.SectionHeaders(sectionHeaders)
	bench.IncludeSystemInfo(systemInfo)

	// set the header info
	bench.SetGroupColumnHeader("Data Type")
	bench.SetSubGroupColumnHeader("Format")
	bench.SetNameColumnHeader("Package")
	bench.SetDescColumnHeader("Operation")
	// Run the benchmarks
	benchBasicMemInfo(bench)
	benchMemInfo(bench)
	benchMessage(bench)
	benchRedditAccount(bench)

	close(done)
	fmt.Println("\ngenerating output...\n")
	err = bench.Out()
	if err != nil {
		fmt.Printf("error generating output: %s\n", err)
	}
}

func benchBasicMemInfo(bench benchutil.Benchmarker) {
	// CapnProto2
	b := capnp.BenchBasicMemInfo()
	bench.Append(b...)
	// Flatbuffers
	b = fb.BenchBasicMemInfo()
	bench.Append(b...)
	// Gencode
	b = gencode.BenchBasicMemInfo()
	bench.Append(b...)
	// Gob
	b = gobs.BenchBasicMemInfo()
	bench.Append(b...)
	// JSON
	b = jsn.BenchBasicMemInfo()
	bench.Append(b...)
	// FFJSON
	b = ffjson.BenchBasicMemInfo()
	bench.Append(b...)
	// FFJSON Buf
	b = ffjsonbuf.BenchBasicMemInfo()
	bench.Append(b...)
	// vmhailenco/msgPack
	b = vmsgpack.BenchBasicMemInfo()
	bench.Append(b...)
	// PBv3
	b = pb.BenchBasicMemInfo()
	bench.Append(b...)
}

func benchMemInfo(bench benchutil.Benchmarker) {
	// CapnProto2
	b := capnp.BenchMemInfo()
	bench.Append(b...)
	// Flatbuffers
	b = fb.BenchMemInfo()
	bench.Append(b...)
	// Gencode
	b = gencode.BenchMemInfo()
	bench.Append(b...)
	// Gob
	b = gobs.BenchMemInfo()
	bench.Append(b...)
	// JSON
	b = jsn.BenchMemInfo()
	bench.Append(b...)
	// FFJSON
	b = ffjson.BenchMemInfo()
	bench.Append(b...)
	// FFJSON Buf
	b = ffjsonbuf.BenchMemInfo()
	bench.Append(b...)
	// vmhailenco/msgPack
	b = vmsgpack.BenchMemInfo()
	bench.Append(b...)
	// PBv3
	b = pb.BenchMemInfo()
	bench.Append(b...)
}

func benchMessage(bench benchutil.Benchmarker) {
	var b []benchutil.Bench

	// Message Data
	dataLen := []int{16, 64, 256, 1024, 2048, 4096}
	for _, v := range dataLen {
		var rnd pcg.Rand
		rnd.Seed(benchutil.Seed())
		shared.GenMessageData(v, shared.Len, rnd)
		// CapnProto2
		// TODO: 4096 Bytes of data causes the following error:
		// capnp: NewMessage called on arena with data
		// figure out how to resolve.
		if v > 2048 {
			goto skipCap
		}
		b = capnp.BenchMessage(v)
		bench.Append(b...)
	skipCap:
		// Flatbuffers
		b = fb.BenchMessage(v)
		bench.Append(b...)
		// Gencode
		b = gencode.BenchMessage(v)
		bench.Append(b...)
		// Gob
		b = gobs.BenchMessage(v)
		bench.Append(b...)
		// JSON
		b = jsn.BenchMessage(v)
		bench.Append(b...)
		// FFJSON
		b = ffjson.BenchMessage(v)
		bench.Append(b...)
		// FFJSONBuf
		b = ffjsonbuf.BenchMessage(v)
		bench.Append(b...)
		// vmhailenco/msgPack
		b = vmsgpack.BenchMessage(v)
		bench.Append(b...)
		// PBv3
		b = pb.BenchMessage(v)
		bench.Append(b...)
	}
}

func benchRedditAccount(bench benchutil.Benchmarker) {
	// CapnProto2
	b := capnp.BenchRedditAccount()
	bench.Append(b...)
	// Flatbuffers
	b = fb.BenchRedditAccount()
	bench.Append(b...)
	// Gencode
	b = gencode.BenchRedditAccount()
	bench.Append(b...)
	// Gob
	b = gobs.BenchRedditAccount()
	bench.Append(b...)
	// JSON
	b = jsn.BenchRedditAccount()
	bench.Append(b...)
	// FFJSON
	b = ffjson.BenchRedditAccount()
	bench.Append(b...)
	// FFJSONBuf
	b = ffjsonbuf.BenchRedditAccount()
	bench.Append(b...)
	// vmhailenco/msgPack
	b = vmsgpack.BenchRedditAccount()
	bench.Append(b...)
	// PB v3
	b = pb.BenchRedditAccount()
	bench.Append(b...)
}
