package csa

import (
	"fmt"
	"testing"
)

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
	suffixArray := []int{15, 7, 8, 3, 9, 4, 0, 11, 14, 6, 2, 10, 13, 5, 1, 12}
	psiArray    := []int{0, 2, 4, 5, 11, 13, 14, 15, 0, 1, 3, 7, 8, 9, 10, 12}
	bitVector   := []int{1, 0, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0, 0}
	csa := newCsa(input)
	if !testEq(psiArray, csa.psi) {
		t.Errorf("psi-array is wrong")
	}
	if !testEq(suffixArray, csa.suffixOffsets) {
		t.Errorf("suffix array is wrong")
	}
	csa.createBitVector()
	if !testEq(bitVector, csa.bv) {
		t.Errorf("bit vector is wrong")
	}
	csa.efCompress()
	fmt.Println(csa.ef.getMany(len(csa.psi) / 2))
	csa.printContents()
}