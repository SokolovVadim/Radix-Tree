package csa

import (
	"fmt"
)

const (
	indexNotFound = -1
	alphabetLength = 2
	EofMarker = 1000
)

type Csa struct {
	suffixOffsets   []int
	psi  []uint32
	ef   *CompressedText
	length int
}

func newCsa(text string) *Csa {
	sa := NewSuffixArrayX(text)
	csa := new(Csa)
	csa.suffixOffsets = sa.index
	csa.length = sa.n
	csa.psi = csa.naivePsi()
	return csa
}

func (csa *Csa)printContents() {
	println("========== Printing contents ==========")
	println("Offset:")
	for i := 0; i < csa.length; i++ {
		fmt.Printf("%v ", i)
	}
	println("\nSuffix array:")
	for _, i := range csa.suffixOffsets {
		fmt.Printf("%v ", i)
	}
	println("\nPsi array:")
	for _, i := range csa.psi {
		fmt.Printf("%v ", i)
	}
	println("\nBitmap:")
	fmt.Println(csa.ef.b.String())
	println("\n=========== End of printing ===========")
}

func findIndex(saIndex []int, idx int) int {
	for i := 0; i < len(saIndex); i++ {
		if idx == saIndex[i] {
			return i
		}
	}
	return indexNotFound
}

func (csa* Csa)naivePsi() []uint32 {
	psiArr := make([]uint32, csa.length)
	// assume PSI[0] = '$'
	// PSI[0] = index, where SA[index] = 0
	psiArr[0] = EofMarker
	for i := 1; i < csa.length; i++ {
		psiArr[i] = uint32(findIndex(csa.suffixOffsets, csa.suffixOffsets[i] + 1))
	}
	csa.psi = psiArr
	return psiArr
}

func (csa* Csa)efCompress() {
	// create an Elias-Fano sequence with maximum element from psi
	csa.ef = NewEF(uint64(len(csa.psi)), uint64(len(csa.psi)))
	csa.ef.Compress(csa.psi[(len(csa.psi) / alphabetLength):])
}

func (csa* Csa)getSaFromPsi(x int, psi []uint32) int{
	// General relationship:
	// SA[x] = SA[PSI[X]] - 1
	// while(psi[x] != EOF_MARKER)
	hopsToEnd := 0
	for {
		if psi[x] == psi[0] {
			break
		}
		hopsToEnd++
		x = int(psi[x])
	}
	return csa.length - hopsToEnd - 1
}

func (csa* Csa)Lookup(str string) {

}
