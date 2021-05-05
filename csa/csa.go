package csa

import (
	"fmt"
)
const err = -1

type Csa struct {
	psi  []int
	sa   []int
	text []rune
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

func findIndex(sa* suffixarrayx, idx int) int {
	for i := 0; i < sa.n; i++ {
		if idx == sa.index[i] {
			return i
		}
	}
	return err
}

func naivePsi(sa* suffixarrayx, text string) []int {
	textLen := len(text)
	psiArr := make([]int, textLen)
	// assume PSI[0] = '$'
	psiArr[0] = '$'
	for i := 1; i < textLen; i++ {
		psiArr[i] = findIndex(sa, sa.index[i] + 1)
	}
	return psiArr
}
