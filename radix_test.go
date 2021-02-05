package radix

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRadixInterface(t *testing.T) {
	r := New()
	r.Insert("foo", 1)
	r.Insert("bar", 2)
	r.Insert("foobar", 2)

	value, ok := r.Get("foo")
	if(ok) {
		fmt.Println(value)
	}
	// fmt.Printf("%s", r.Get("foo"))
}

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
		fmt.Println(substring)
	}
	return substringArray
}

func FillRadixTree(substringArray []string) {
	r := New()
	for i := 0; i < len(substringArray); i++ {
		r.Insert(substringArray[i], 1)
		// fmt.Println(substringArray[i], " inserted")
	}
}

func TestInsert(t *testing.T) {
	// r := New()
	InitSeed()
	test_str := GenerateTestString(16)
	fmt.Println(test_str)
	var substringArray []string = CreateSubstrings(test_str)
	FillRadixTree(substringArray)
}
