package bwt

import (
	"index/suffixarray"
	"testing"
)

type rotation struct {
	index int
	suffix string
}

func bwTransform(sa* suffixarray.Index, text string) []uint8 {
	textLen := len(text)
	var bwtArr []uint8
	for i := 0; i < textLen; i++ {
		println(sa.Bytes()[i])
		j := int(sa.Bytes()[i]) - 1
		if j < 0 {
			j = j + textLen
		}
		bwtArr = append(bwtArr, text[j])
	}
	return bwtArr
}

func BenchmarkBWT(b *testing.B) {
	var input = "abracadabra"
	sa := suffixarray.New([]byte(input))

	btwArr := bwTransform(sa, input)
	for i := 0; i < len(btwArr); i++ {
		println(btwArr)
	}
}