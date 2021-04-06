package bwt

import (
	"sort"
	"strings"
	"testing"
)

func naiveBWT(sa* suffixarrayx, text string) []byte {
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

func regenerateSuffix(bwtArr string) string {
	// make an empty table
	bwtArr += "^"
	var table = make([]string, len(bwtArr))
	for range table {
		for i := 0; i < len(bwtArr); i++ {
			table[i] = bwtArr[i:i+1] + table[i]
		}
		sort.Strings(table)
	}
	for _, row := range table {
		if strings.HasSuffix(row, "$") {
			return row[1 : len(bwtArr) - 1]
		}
	}
	return ""
}

func BenchmarkBWT(b *testing.B) {
	var input = "SIX.MIXED.PIXIES.SIFT.SIXTY.PIXIE.DUST.BOXES$"
	sa := NewSuffixArrayX(input)

	bwtArr := naiveBWT(sa, input)
	println(string(bwtArr))
	if string(bwtArr) != "STEXYDST.E.IXXIIXXSSMPPS.B..EE.$.USFXDIIOIIIT" {
		b.Errorf("NaiveBWT is wrong, %s", string(bwtArr))
	}
	ibwtArr := regenerateSuffix(string(bwtArr))
	println("ibwt: ", string(ibwtArr))
	if string(ibwtArr) != "SIX.MIXED.PIXIES.SIFT.SIXTY.PIXIE.DUST.BOXES" {
		b.Errorf("iBWT is wrong, %s", string(bwtArr))
	}
}