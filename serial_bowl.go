package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/mohae/serial-bowl/fb"
	"github.com/mohae/serial-bowl/jsn"
	"github.com/mohae/serial-bowl/pb"
	"github.com/mohae/serial-bowl/shared"
)

// flags
var (
	output string
	format string
)

func init() {
	flag.StringVar(&output, "output", "stdout", "output destination")
	flag.StringVar(&output, "o", "stdout", "output destination (short)")
	flag.StringVar(&format, "format", "txt", "format of output")
	flag.StringVar(&format, "f", "txt", "format of output")
}

func main() {
	flag.Parse()
	done := make(chan struct{})
	go dot(done)
	shared.GenData()
	var results []shared.Bench
	// BasicMemInfo
	// Flatbuffers
	b := fb.BenchBasicMemInfo()
	results = append(results, b)
	// JSON
	b = jsn.BenchBasicMemInfo()
	results = append(results, b)
	// PBv3
	b = pb.BenchBasicMemInfo()
	results = append(results, b)
	// MemInfo
	// Flatbuffers
	b = fb.BenchMemInfo()
	results = append(results, b)
	// JSON
	b = jsn.BenchMemInfo()
	results = append(results, b)
	// PBv3
	b = pb.BenchMemInfo()
	results = append(results, b)

	// Message Data
	dataLen := []int{16, 256, 1024, 4096}
	for _, v := range dataLen {
		shared.GenMessageData(v, shared.Len)
		// Flatbuffers
		b = fb.BenchMessage()
		results = append(results, b)
		// JSON Marshal
		b = jsn.BenchMessage()
		results = append(results, b)
		// PBv3
		b = pb.BenchMessage()
		results = append(results, b)
	}

	// Reddit Account
	// Flatbuffers
	b = fb.BenchRedditAccount()
	results = append(results, b)
	// JSON Marshal
	b = jsn.BenchRedditAccount()
	results = append(results, b)
	// PB v3
	b = pb.BenchRedditAccount()
	results = append(results, b)
	close(done)
	fmt.Println("")
	fmt.Println("generating output...")
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
	// process the output
	switch format {
	case "txt":
		shared.TXTOut(w, results)
	case "csv":
		err := shared.CSVOut(w, results)
		if err != nil {
			fmt.Printf("error creating CSV output: %s\n", err)
		}
	case "md":
		err := shared.MDOut(w, results)
		if err != nil {
			fmt.Printf("error creating MarkDown output: %s\n", err)
		}
	default:
		fmt.Printf("unknown output format: %q; defaulting to \"txt\"\n", format)
		shared.TXTOut(w, results)
	}
}

func dot(done chan struct{}) {
	t := time.NewTicker(time.Second)
	defer t.Stop()
	for {
		select {
		case <-done:
			return
		case <-t.C:
			fmt.Print(".")
		}
	}
}
