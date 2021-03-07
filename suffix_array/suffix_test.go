package suffix_array

import (
	"index/suffixarray"
	"math/rand"
	"testing"
	"time"
)

const (
	length   = 65536
	leftPos  = 10200
	rightPos = 10450
)

func initSeed() {
	rand.Seed(time.Now().UnixNano())
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generateTestString(size int) string {
	b := make([]byte, size)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func BenchmarkFindAllIndex(b *testing.B) {
	initSeed()
	testStr := generateTestString(length)
	sa := suffixarray.New([]byte(testStr))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		offset := sa.Lookup([]byte(testStr[leftPos: rightPos]), -1)
		if len(offset) < 1 || offset[0] != leftPos {
			b.Fatalf("mis-match: %v", offset)
		}
	}
}
