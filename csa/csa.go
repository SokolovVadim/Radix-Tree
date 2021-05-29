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
	bv   []*CompressedText
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
	// csa.ef = NewEF(uint64(len(csa.psi)), uint64(len(csa.psi)))
	seqLen := make([]int, alphabetLength - 1)
	csa.bv = make([]*CompressedText, alphabetLength)
	j := 0
	for i := 1; i < csa.length - 1; i++ {
		if csa.psi[i] > csa.psi[i + 1] {
			seqLen[j] = i
			println(i)
			if j > 0 {
				fmt.Println("here")
				curLen := uint64(seqLen[j] - seqLen[j - 1])
				csa.bv[j] = NewEF(uint64(i), curLen)
				csa.bv[j].Compress(csa.psi[seqLen[j - 1] + 1:i])
			} else {
				curLen := uint64(i)
				csa.bv[j] = NewEF(uint64(csa.length), curLen)
				csa.bv[j].Compress(csa.psi[1:curLen])
			}
			j++
		}
	}
	// seqLen[alphabetLength - 1] = csa.length - seqLen[alphabetLength - 2]
	for i := range seqLen {
		fmt.Println(seqLen[i])
	}

/*	for i := range seqLen {
		csa.bv[i] = NewEF(uint64(seqLen[i]), uint64(seqLen[i]))
		csa.bv[i].Compress(csa.psi[1:csa.length / 2])
	}*/

	// csa.ef.Compress(csa.psi[1:csa.length / 2])
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
