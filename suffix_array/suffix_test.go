package suffix_array

import (
	"index/suffixarray"
	"math/rand"
	"regexp"
	"strings"
	"testing"
	"time"
)

const (
	Length   = 65536
	LeftPos  = 10200
	RightPos = 10450
)

func InitSeed() {
	rand.Seed(time.Now().UnixNano())
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateTestString(size int) string {
	b := make([]byte, size)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func CreateSubstrings(str string) ([]string, string) {
	runes := []rune(str)
	// Loop over string in order to acquire substrings
	// from the end to the beginning
	var substringArray []string
	for start := 1; start < len(runes); start++ {
		substring := string(runes[start:len(runes)])
		substringArray = append(substringArray, substring)
		// fmt.Println(substring)
	}
	// use \x00 to start each string
	joinedStrings := "\x00" + strings.Join(substringArray, "\x00")
	return substringArray, joinedStrings
}

func GetSubstr(sa* suffixarray.Index, joinedStrings string,
	           subString string) (string, int, int) {
	// fmt.Println("substr:", subString)
	match, err := regexp.Compile("\x00" + subString + "[^\x00]*")
	if err != nil {
		panic(err)
	}
	ms := sa.FindAllIndex(match, -1)

	/*for _, m := range ms {
		start, end := m[0], m[1]
		fmt.Printf("match = %q\n", joinedStrings[start + 1: end])
	}*/

	// there is only one possible match for a string
	if len(ms) > 0 {
		start, end := ms[0][0], ms[0][1]
		return joinedStrings[start + 1: end], start, end
	} else {
		return "", 0, 0
	}
}

func BenchmarkFindAllIndex(b *testing.B) {
	InitSeed()
	testStr := GenerateTestString(Length)
	// fmt.Println(testStr)
	var _, joinedStrings = CreateSubstrings(testStr)
	sa := suffixarray.New([]byte(joinedStrings))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetSubstr(sa, joinedStrings, testStr[LeftPos: RightPos])
	}
}
