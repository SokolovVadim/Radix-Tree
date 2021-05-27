package csa

import (
	"fmt"
	"reflect"
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

func TestCsa(t *testing.T) {
	input := "abbaabbaaababbb$"
	suffixArray := []int{15, 7, 8, 3, 9, 4, 0, 11, 14, 6, 2, 10, 13, 5, 1, 12}
	psiArray    := []uint32{EofMarker, 2, 4, 5, 11, 13, 14, 15, 0, 1, 3, 7, 8, 9, 10, 12}

	// bitString := "{1,4,7,9,16,19,21,23}"
	csa := newCsa(input)
	if !reflect.DeepEqual(suffixArray, csa.suffixOffsets) {
		t.Errorf("suffix array is wrong")
	}
	if !reflect.DeepEqual(psiArray, csa.psi) {
		t.Errorf("psi-array is wrong")
	}
	// csa.efCompress()
/*	if bitString != csa.ef.b.String() {
		t.Errorf("bit map is wrong")
	}*/
/*	if !reflect.DeepEqual(csa.ef.getMany(len(csa.psi) / alphabetLength), psiArray[:len(csa.psi) / alphabetLength]) {
		t.Errorf("converted psi-array is wrong")
	}*/
	fmt.Println(csa.psi)
	x := csa.getSaFromPsi(13, csa.psi)
	fmt.Println(x)
	// csa.printContents()
}