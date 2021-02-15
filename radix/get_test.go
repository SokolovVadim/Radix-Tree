package radix_test

import (
	"github.com/SokolovVadim/Radix-Tree"
	"testing"
)

const (
	Length = 65536;
	LeftPos = 1024;
	RightPos = 1032;
)

func GetSubstring(b *testing.B, r *radix.Tree, testStr string) {
	_, ok := r.Get(testStr[LeftPos: RightPos])
	if !ok {
		b.Fatalf("missing key: %v", testStr[LeftPos:RightPos])
	}
}

func BenchmarkGet(b *testing.B) {
	InitSeed()
	testStr := GenerateTestString(Length)
	// println(testStr)
	r := radix.New()
	var substringArray = CreateSubstrings(testStr)
	FillRadixTree(Length, r, substringArray)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetSubstring(b, r, testStr)
	}
}
