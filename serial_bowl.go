package main

import (
	"fmt"

	"github.com/mohae/serial-bowl/fb"
	"github.com/mohae/serial-bowl/jsn"
	"github.com/mohae/serial-bowl/pb"
	"github.com/mohae/serial-bowl/shared"
)

func main() {
	shared.GenData()

	// BasicMemInfo
	// Flatbuffers
	b := fb.BenchBasicMemInfo()
	txt := b.Output()
	for _, v := range txt {
		fmt.Println(v)
	}
	// JSON
	b = jsn.BenchBasicMemInfo()
	txt = b.Output()
	for _, v := range txt {
		fmt.Println(v)
	}
	// PBv3
	b = pb.BenchBasicMemInfo()
	txt = b.Output()
	for _, v := range txt {
		fmt.Println(v)
	}

	// MemInfo
	fmt.Println("")
	// Flatbuffers
	b = fb.BenchMemInfo()
	txt = b.Output()
	for _, v := range txt {
		fmt.Println(v)
	}
	// JSON
	b = jsn.BenchMemInfo()
	txt = b.Output()
	for _, v := range txt {
		fmt.Println(v)
	}
	// PBv3
	b = pb.BenchMemInfo()
	txt = b.Output()
	for _, v := range txt {
		fmt.Println(v)
	}

	// Message Data
	dataLen := []int{16, 256, 1024, 4096}
	for _, v := range dataLen {
		fmt.Println("")
		shared.GenMessageData(v, shared.Len)
		fmt.Printf("Data: %d bytes\n", v)
		// Flatbuffers
		b = fb.BenchMessage()
		txt = b.Output()
		for _, v := range txt {
			fmt.Println(v)
		}
		// JSON Marshal
		b = jsn.BenchMessage()
		txt = b.Output()
		for _, v := range txt {
			fmt.Println(v)
		}
		// PBv3
		b = pb.BenchMessage()
		txt = b.Output()
		for _, v := range txt {
			fmt.Println(v)
		}
	}

	// Reddit Account
	fmt.Println("")
	// Flatbuffers
	b = fb.BenchRedditAccount()
	txt = b.Output()
	for _, v := range txt {
		fmt.Println(v)
	} // JSON Marshal
	b = jsn.BenchRedditAccount()
	txt = b.Output()
	for _, v := range txt {
		fmt.Println(v)
	} // PB v3
	b = pb.BenchRedditAccount()
	txt = b.Output()
	for _, v := range txt {
		fmt.Println(v)
	}
}
