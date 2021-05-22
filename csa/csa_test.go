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

// returns number of 1s from bv
// to the left of position x
func rank(bv []byte, x int) int {
	var sum int
	for i := 0; i < x; i++ {
		if bv[i] == 1 {
			sum++
		}
	}
	return sum
}

func BenchmarkCSA(b *testing.B) {
	println("hello!")
	// csa := newCsa()
	input := "abbaabbaaababbb$"
	sa := NewSuffixArrayX(input)
	csa := newCsaFromSa(sa)
	csa.psi = naivePsi(csa)
	createBitVector(csa)
	efCompress(csa)
	printContents(csa)

	// 
}