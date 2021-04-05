package bwt

import (
	"testing"
)

func NaiveBWT(sa* suffixarrayx, text string) []byte {
	textLen := len(text)
	btwArr := make([]byte, textLen)
	for i := 0; i < textLen; i++ {
		j := sa.index[i] - 1
		if j < 0 {
			j += textLen
		}
		btwArr[i] = text[j]
	}
	return btwArr
}

func BenchmarkBWT(b *testing.B) {
	var input = "SIX.MIXED.PIXIES.SIFT.SIXTY.PIXIE.DUST.BOXES$"
	sa := NewSuffixArrayX(input)

	btwArr := NaiveBWT(sa, input)
	println(string(btwArr))
	if string(btwArr) != "STEXYDST.E.IXXIIXXSSMPPS.B..EE.$.USFXDIIOIIIT" {
		b.Errorf("NaiveBWT is wrong, %s", string(btwArr))
	}
}