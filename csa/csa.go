package csa

import (
	"fmt"
	"github.com/RoaringBitmap/roaring"
)
const err = -1

type Csa struct {
	text []rune
	sa   []int
	psi  []int
	rb   *roaring.Bitmap
	bv   []uint32
	ef   *EliasFano
	len int
}

func newCsa() *Csa {
	sa := new(Csa)
	return sa
}

func newCsaFromSa(sa* suffixarrayx) *Csa {
	csa := new(Csa)
	csa.sa = sa.index
	csa.len = sa.n
	csa.text = sa.text
	csa.rb = roaring.NewBitmap()
	return csa
}

func printContents(this *Csa) () {
	println("========== Printing contents ==========")
	println("Offset:")
	for i := 0; i < this.len; i++ {
		fmt.Printf("%v ", i)
	}
	println("\nText:")
	for _, i := range this.text {
		fmt.Printf("%c ", i)
	}
	println("\nSuffix array:")
	for _, i := range this.sa {
		fmt.Printf("%v ", i)
	}
	println("\nPsi array:")
	for _, i := range this.psi {
		fmt.Printf("%v ", i)
	}
	println("\nBit vector:")
	for _, i := range this.bv {
		fmt.Printf("%v ", i)
	}
	println("\nBitmap:")
	fmt.Println(this.rb.String())
	println("\n=========== End of printing ===========")
}

func naiveBWT(sa* suffixarrayx, text string) []byte {
	textLen := len(text)
	bwtArr := make([]byte, textLen)
	for i := 0; i < textLen; i++ {
		j := sa.index[i] - 1
		if j < 0 {
			j += textLen
		}
		bwtArr[i] = text[j]
	}
	return bwtArr
}

func findIndex(saIndex []int, idx int) int {
	for i := 0; i < len(saIndex); i++ {
		if idx == saIndex[i] {
			return i
		}
	}
	return err
}

func naivePsi(csa* Csa) []int {
	psiArr := make([]int, csa.len)
	// assume PSI[0] = '$'
	psiArr[0] = 0
	for i := 1; i < csa.len; i++ {
		psiArr[i] = findIndex(csa.sa, csa.sa[i] + 1)
	}
	csa.psi = psiArr
	return psiArr
}

func createBitVector(csa* Csa) {
	// length = N + #items ~= length * 2
	length := csa.len * 2
	bv := make([]uint32, length)
	for i := 0; i < length; i++ {
		bv[i] = 0
	}
	for i := 0; i < csa.len; i++ {
		println("i = ", i, ", psi[i] = ", csa.psi[i], ", i + psi[i] = ", i + csa.psi[i])
		bv[i + csa.psi[i]] = 1
	}
	csa.bv = bv
	csa.rb.AddMany(bv)
}

func efCompress(csa* Csa) {
	// create an Elias-Fano sequence with maximum element from psi
	csa.ef = NewEF(uint64(csa.psi[csa.len-1]), uint64(csa.len))
	// csa.ef.Compress(csa.psi)
}
