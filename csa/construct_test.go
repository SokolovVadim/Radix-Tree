package csa

import (
	"fmt"
	"io/ioutil"
	"log"
	"runtime"
	"testing"
)

func BenchmarkCSAConstruct(b *testing.B) {
	PrintMemUsage()
	content, err := ioutil.ReadFile("C:\\Users\\Vadim\\GolandProjects\\Radix-Tree\\utils\\data.txt")
	if err != nil {
		log.Fatal(err)
	}
	testStr := string(content)

	csa := newCsa(testStr)
	fmt.Println("csa len:", csa.length)
	csa.efCompress()

	runtime.GC()
	PrintMemUsage()
}
