package radix

import (
	"math/rand"
	"testing"
	"time"
)

const (
	LENGTH = 65536
)

func InitSeed() {
	rand.Seed(time.Now().UnixNano())
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateTestString(size int64) string {
	b := make([]byte, size)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func CreateSubstrings(str string) []string {
	runes := []rune(str)
	// Loop over string in order to acquire substrings
	// from the end to the beginning
	var substringArray []string
	for start := 1; start < len(runes); start++ {
		substring := string(runes[start:len(runes)])
		substringArray = append(substringArray, substring)
		// fmt.Println(substring)
	}
	return substringArray
}

func FillRadixTree(r *Tree, substringArray []string) {
	for i := 0; i < len(substringArray); i++ {
		r.Insert(substringArray[i], i)
		// fmt.Println(substringArray[i], " inserted")
	}
}

// go test -bench=. -benchmem
func BenchmarkInsert(b *testing.B) {
	InitSeed()
	test_str := GenerateTestString(LENGTH)
	// fmt.Println(test_str)
	r := New()
	var substringArray []string = CreateSubstrings(test_str)
	FillRadixTree(r, substringArray)
}
