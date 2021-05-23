package csa

import (
	"testing"
)

func select_(csa Csa, k int, c int) int {
	if k <= 0 {
		return indexNotFound
	}
	idx := csa.suffixOffsets[c]
	if idx == 0 && c != 0 {
		return csa.len
	}
	if idx == 255 {
		return indexNotFound
	}
	if int(csa.suffixOffsets[idx]) + k - 1 < int(csa.suffixOffsets[idx + 1]) {
		return int(csa.psi[int(csa.suffixOffsets[idx]) + k - 1])
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
	/*input := "abbaabbaaababbb$"

	csa := newCsa(input)
	csa.psi = naivePsi(csa)
	createBitVector(csa)
	efCompress(csa)
	printContents(csa)*/
}

func testEq(first, second []int) bool {
	// If one is nil, the other must also be nil.
	if (first == nil) != (second == nil) {
		return false
	}
	if len(first) != len(second) {
		return false
	}
	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			return false
		}
	}
	return true
}

func TestCsa(t *testing.T) {
	input := "abbaabbaaababbb$"
	psiArray := []int{0, 2, 4, 5, 11, 13, 14, 15, 0, 1, 3, 7, 8, 9, 10, 12}
	csa := newCsa(input)
	if !testEq(psiArray, csa.psi) {
		t.Errorf("psi-array is wrong")
	}
	createBitVector(csa)
	efCompress(csa)
	csa.printContents()
}