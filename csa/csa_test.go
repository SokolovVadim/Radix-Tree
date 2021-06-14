package csa
import (
	"testing"
)

func select_(bv []byte, k int) {

}

func rank(bv []byte, x int) {

}

func BenchmarkCSA(b *testing.B) {
	println("hello!")
	csa := newCsa()
	printContents(csa)
}