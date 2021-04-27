package csa
import (
	"testing"
)

func select_(sa Csa, k int, c int) int {
	if k <= 0 {
		return err
	}
	idx := sa.sa[c]
	if idx == 0 && c != 0 {
		return sa.len
	}
	if idx == 255 {
		return err
	}
	if int(sa.sa[idx]) + k - 1 < int(sa.sa[idx + 1]) {
		return int(sa.psi[int(sa.sa[idx]) + k - 1])
	} else {
		return sa.len
	}
}

func rank(bv []byte, x int) {

}

func BenchmarkCSA(b *testing.B) {
	println("hello!")
	csa := newCsa()
	printContents(csa)
}