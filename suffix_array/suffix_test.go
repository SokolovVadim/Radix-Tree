package suffix_array

import (
	"index/suffixarray"
	"io/ioutil"
	"log"
	"testing"
)

const (
	leftPos  = 10200
	rightPos = 10450
)

func BenchmarkFindAllIndex(b *testing.B) {
	content, err := ioutil.ReadFile("C:\\Users\\Vadim\\GolandProjects\\Radix-Tree\\utils\\data.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string
	testStr := string(content)
	sa := suffixarray.New([]byte(testStr))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		offset := sa.Lookup([]byte(testStr[leftPos: rightPos]), -1)
		if len(offset) < 1 || offset[0] != leftPos {
			b.Fatalf("mis-match: %v", offset)
		}
	}
}
