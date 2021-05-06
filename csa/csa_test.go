package csa

import (
	"testing"
)

func select_(csa Csa, k int, c int) int {
	if k <= 0 {
		return err
	}
	idx := csa.sa[c]
	if idx == 0 && c != 0 {
		return csa.len
	}
	if idx == 255 {
		return err
	}
	if int(csa.sa[idx]) + k - 1 < int(csa.sa[idx + 1]) {
		return int(csa.psi[int(csa.sa[idx]) + k - 1])
	} else {
		return csa.len
	}
}

func rank(bv []byte, x int) {

}

func BenchmarkCSA(b *testing.B) {
	println("hello!")
	// csa := newCsa()
	input := "abbaabbaaababbb$"
	sa := NewSuffixArrayX(input)
	csa := newCsaFromSa(sa)
	csa.psi = naivePsi(csa)
	createBitVector(csa)
	printContents(csa)
}