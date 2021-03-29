package radix_test

import (
	"github.com/SokolovVadim/Radix-Tree"
	"io/ioutil"
	"log"
	"testing"
)

func createSubstrings(str string) []string {
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

func fillRadixTree(size int, r *radix.Tree, substringArray []string) {
	for i := 0; i < size - 1; i++ {
		r.Insert(substringArray[i], i)
		// fmt.Println(substringArray[i], " inserted")
	}
}

// go test -bench=. -benchmem -benchtime=100x
func BenchmarkInsert(b *testing.B) {
	content, err := ioutil.ReadFile("C:\\Users\\Vadim\\GolandProjects\\Radix-Tree\\utils\\data.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string
	testStr := string(content)
	r := radix.New()
	var substringArray []string = createSubstrings(testStr)

	b.ResetTimer()
	fillRadixTree(b.N, r, substringArray)
}