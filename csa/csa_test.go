package csa

import (
	"fmt"
	"io/ioutil"
	"log"
	"runtime"
	"testing"
)

func BenchmarkCSA(b *testing.B) {
	PrintMemUsage()
	content, err := ioutil.ReadFile("C:\\Users\\Vadim\\GolandProjects\\Radix-Tree\\utils\\data.txt")
	if err != nil {
		log.Fatal(err)
	}
	testStr := string(content)

	csa := newCsa(testStr)
	fmt.Println("csa len:", csa.length)
	// fmt.Println("sa:", csa.suffixOffsets)
	csa.efCompress()
	runtime.GC()
	PrintMemUsage()

	b.ResetTimer()
    //	for i := 0; i < b.N; i++ {
	//	csa.lookup(testStr[leftPos: rightPos])
	// }
	csa.lookup(testStr[leftPos: rightPos])
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

/*func TestCsa(t *testing.T) {
	// PrintMemUsage()
	// defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
	input := "adbbaabcbdaaccaaabcdabaabcabbacbadbdbabcdbcbc$"
	csa := newCsa(input)
	// csa.printContents()
	csa.efCompress()
	// runtime.GC()
	// PrintMemUsage()
	// fmt.Println("Lookup:")
	csa.lookup("dbdba")
	// PrintMemUsage()

	// Force GC to clear up, should see a memory drop
	// runtime.GC()
	// PrintMemUsage()
}*/