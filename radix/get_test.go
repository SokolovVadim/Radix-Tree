package radix_test

import (
	"testing"
	"github.com/SokolovVadim/Radix-Tree"
)

func GetSubstring(b *testing.B, r *radix.Tree, substringArray []string) {
	for i := 0; i < len(substringArray); i++ {
		_, ok := r.Get(substringArray[i])
		if !ok {
			b.Fatalf("missing key: %v", substringArray[i])
		}
	}
}

func BenchmarkGet(b *testing.B) {
	InitSeed()
	testStr := GenerateTestString(Lenght)
	r := radix.New()
	var substringArray []string = CreateSubstrings(testStr)
	FillRadixTree(r, substringArray)
	GetSubstring(b, r, substringArray)
}
