package csa

import (
	"fmt"
	"testing"
)

func BenchmarkCSA(b *testing.B) {

}

func TestCsa(t *testing.T) {
	input := "abbaabcbaaccaaabcabaabcabbacbabbabcbcbc$"

	csa := newCsa(input)
	csa.printContents()
	csa.efCompress()
	fmt.Println("Lookup:")
	csa.Lookup("abca")
}