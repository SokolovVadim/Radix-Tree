package radix_test

import (
	"github.com/SokolovVadim/Radix-Tree"
	"testing"
)

const (
	length   = 65536
	leftPos  = 10200
	rightPos = 10450
)

func getSubstring(b *testing.B, r *radix.Tree, subString string,
	              testStr string) (str string, val interface{}) {
	var out interface{}
	var pos interface{}
	// there is only one possible match for a string
	fn := func(s string, v interface{}) bool {
		out, pos = s, v
		return false
	}

	r.WalkPrefix(subString, fn)
	if result, ok := out.(string); !ok || result != testStr[leftPos:length] {
		b.Fatalf("mis-match: %v %v", out, subString)
	}
	return out.(string), pos
}

func BenchmarkGet(b *testing.B) {
	initSeed()
	testStr := generateTestString(length)
	r := radix.New()
	var substringArray = createSubstrings(testStr)
	fillRadixTree(length, r, substringArray)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getSubstring(b, r, testStr[leftPos: rightPos], testStr)
	}
}
