package csa

import (
	"fmt"
)

const (
	indexNotFound = -1
	alphabetLength = 3
	EofMarker = 1000
)

type Csa struct {
	text string
	suffixOffsets   []int
	psi  []uint64
	ef   *CompressedText
	bv   []*CompressedText
	seqOffset []int
	seqLen []int
	seqChar[]byte
	length int
}

func newCsa(text string) *Csa {
	sa := NewSuffixArrayX(text)
	csa := new(Csa)
	csa.text = text
	csa.suffixOffsets = sa.index
	csa.length = sa.n
	csa.psi = csa.naivePsi()
	return csa
}

func (csa *Csa)printContents() {
	println("========== Printing contents ==========")
	println("Text:")
	fmt.Println(csa.text)
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
	if csa.ef != nil {
		println("\nBitmap:")
		fmt.Println(csa.ef.b.String())
	}
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

func (csa* Csa)naivePsi() []uint64 {
	psiArr := make([]uint64, csa.length)
	// assume PSI[0] = '$'
	// PSI[0] = index, where SA[index] = 0
	psiArr[0] = EofMarker
	for i := 1; i < csa.length; i++ {
		psiArr[i] = uint64(findIndex(csa.suffixOffsets, csa.suffixOffsets[i] + 1))
	}
	csa.psi = psiArr
	return psiArr
}

func (csa* Csa)efCompressOne() {
	array := []uint64{0, 2, 3, 8, 11, 13, 14, 16, 29}
	size := uint64(len(array))
	max := array[size-1]
	ef := NewEF(max, size)
	ef.Compress(array)
	fmt.Println(ef.getMany(int(size)))
}

func (csa* Csa)efCompress() {
	// create an array for storing lengths of sequences
	csa.seqLen = make([]int, alphabetLength)
	csa.seqOffset = make([]int, alphabetLength)
	csa.bv = make([]*CompressedText, alphabetLength)
	csa.seqChar = make([]byte, alphabetLength)
	j := 0
	for i := 1; i < csa.length - 1; i++ {
		// check the start of each new ascending sequence
		// to be compressed
		if csa.psi[i + 1] < csa.psi[i] {
			// the first entry
			if j == 0 {
				csa.seqOffset[j] = 1
				csa.seqOffset[j + 1] = i + 1
				curLen := uint64(i + 1)
				csa.seqLen[j] = i
				csa.seqChar[j] = csa.text[csa.suffixOffsets[i]]
				csa.bv[j] = NewEF(uint64(csa.length), curLen)
				csa.bv[j].Compress(csa.psi[1:curLen])
				// fmt.Println("initial bitmap:", csa.bv[j].b.String())
			} else {
				if j < alphabetLength - 1 {
					csa.seqOffset[j + 1] = i + 1
					csa.seqChar[j + 1] = csa.text[csa.psi[i]]
				}
				csa.seqChar[j] = csa.text[csa.suffixOffsets[i]]
				curLen := uint64(i - csa.seqOffset[j])
				csa.seqLen[j] = i + 1 - csa.seqOffset[j]
				csa.bv[j] = NewEF(uint64(csa.length), curLen)
				csa.bv[j].Compress(csa.psi[csa.seqOffset[j]:i + 1])
			}
			j++
		}
	}
	j = alphabetLength - 1
	csa.seqLen[j] = csa.length - csa.seqOffset[j]
	csa.bv[j] = NewEF(uint64(csa.length), uint64(csa.seqLen[j]))
	csa.bv[j].Compress(csa.psi[csa.seqOffset[j]:])
	fmt.Println("SeqLen:", csa.seqLen)
	fmt.Println("SeqOffset:", csa.seqOffset)
	fmt.Println("SeqChar:", string(csa.seqChar))
	psi := csa.bv[0].getMany(csa.seqLen[0])
	fmt.Println("decoded:", psi)
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

func (csa* Csa)getPsiFromBv(x uint32) uint32 {
	idx := 0
	// iterate over bitmaps and find appropriate
	for i := 0; i < alphabetLength - 1; i++ {
		if int(x) >= csa.seqOffset[i] && int(x) < csa.seqOffset[i + 1] {
			idx = i
			break
		}
	}
	if int(x) >= csa.seqOffset[alphabetLength - 1] && int(x) < csa.length {
		idx = alphabetLength - 1
	}
	return csa.bv[idx].getVal(x - uint32(csa.seqOffset[idx]))
}

func (csa* Csa)getSaFromBitmap(x int, c byte) int {
	hopsToEnd := 0
	initialX := x
	i := 0
	for {
		psiX := csa.getPsiFromBv(uint32(x))// bv.getVal(uint32(x))
		// fmt.Println("psiX:", psiX)
		if psiX == 0 || i > 20 {
			break
		}
		hopsToEnd++
		i++
		x = int(psiX)
	}
	if initialX == 0 {
		return csa.length - hopsToEnd - 1
	}
	return csa.length - hopsToEnd - 2
}

func (csa* Csa)Lookup(str string) {
	l := 0
	r := csa.length - 1
	m := len(str)
	index := 0
	for {
		if l <= r {
			mid := l + (r - l) / 2

			index = csa.getSaFromBitmap(mid, str[0])
			if str == csa.text[index:index + m] {
				fmt.Println("pattern found at index", index)
			}
			if str < csa.text[index:index + m] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			break
		}
	}
}
