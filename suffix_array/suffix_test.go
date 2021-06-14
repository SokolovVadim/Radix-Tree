package suffix_array

import (
	"fmt"
	"index/suffixarray"
	"io/ioutil"
	"log"
	"runtime"
	"testing"
)

const (
	leftPos  = 1020
	rightPos = 1045
)

func BenchmarkFindAllIndex(b *testing.B) {
	PrintMemUsage()
	content, err := ioutil.ReadFile("C:\\Users\\Vadim\\GolandProjects\\Radix-Tree\\utils\\data.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string
	testStr := string(content)
	sa := suffixarray.New([]byte(testStr))
	runtime.GC()
	PrintMemUsage()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		offset := sa.Lookup([]byte(testStr[leftPos: rightPos]), -1)
		if len(offset) < 1 || offset[0] != leftPos {
			b.Fatalf("mis-match: %v", offset)
		}
	}
	PrintMemUsage()

	// Force GC to clear up, should see a memory drop
	runtime.GC()
	PrintMemUsage()
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	// fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("Alloc = %v kB", bToKb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v kB", bToKb(m.TotalAlloc))
	fmt.Printf("\tSys = %v kB", bToKb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func bToKb(b uint64) uint64 {
	return b / 1024
}

