package csa

import "fmt"

type Csa struct {
	psi []byte
	sa  []byte
	len int
}

func newCsa() *Csa {
	sa := new(Csa)
	return sa
}

func printContents(this *Csa) () {
	println("Suffix array:")
	for i := 0; i < this.len; i++ {
		fmt.Printf("%d", this.sa[i])
	}
}