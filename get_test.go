package radix

import (
	"testing"
)

func GetSubstring(b *testing.B, r *Tree, substringArray []string) {
	for i := 0; i < len(substringArray); i++ {
		_, ok := r.Get(substringArray[i])
		if !ok {
			b.Fatalf("missing key: %v", substringArray[i])
		}
	}
}

func BenchmarkGet(b *testing.B) {
	InitSeed()
	testStr := GenerateTestString(LENGTH)
	r := New()
	var substringArray []string = CreateSubstrings(testStr)
	FillRadixTree(r, substringArray)
	GetSubstring(b, r, substringArray)
}
