package radix_test

import (
	"github.com/SokolovVadim/Radix-Tree"
	"reflect"
	"testing"
)

const (
	Length = 65535
	LeftPos = 10200
	RightPos = 10450
)

func GetSubstring(b *testing.B, r *radix.Tree, subString string,
	              testStr string) (str string, val interface{}) {
	var out string
	var pos interface{}
	// there is only one possible match for a string
	fn := func(s string, v interface{}) bool {
		out, pos = s, v
		return false
	}

	r.WalkPrefix(subString, fn)
	// fmt.Println("out:", out, "pos =", pos)
	if !reflect.DeepEqual(out, testStr[LeftPos: Length]) {
		b.Fatalf("mis-match: %v %v", out, subString)
	}
	return out, pos
}

func BenchmarkGet(b *testing.B) {
	InitSeed()
	testStr := GenerateTestString(Length)
	// fmt.Println(testStr)
	r := radix.New()
	var substringArray = CreateSubstrings(testStr)
	FillRadixTree(Length, r, substringArray)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetSubstring(b, r, testStr[LeftPos: RightPos], testStr)
	}
}
