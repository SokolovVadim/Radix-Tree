package bwt

import (
	"fmt"
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

func inverseBWT(bwtArr string) string {
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

func decode(x int, bwtArr []byte) int {
	hopsToEnd := 0
	for {
		if bwtArr[x] == '$'{
			break
		}
		hopsToEnd++
		x = int(bwtArr[x])
	}
	return len(bwtArr) - 1 - hopsToEnd
}

func regenerateSuffix(bwtArr []byte) []byte {
	regArr := make([]byte, len(bwtArr))
	for i := 0; i < len(bwtArr); i++ {
		regArr[i] = byte(decode(i, bwtArr))
	}
	return regArr
}

func BenchmarkBWT(b *testing.B) {
	var input = "SIX.MIXED.PIXIES.SIFT.SIXTY.PIXIE.DUST.BOXES$"
	sa := NewSuffixArrayX(input)
	println("Sa.index:")
	for i := 0; i < len(sa.index); i++ {
		fmt.Printf("%d ", sa.index[i])
	}

	bwtArr := naiveBWT(sa, input)
	println("\nBWT: ", string(bwtArr))
	for i := 0; i < len(bwtArr); i++ {
		fmt.Printf("%d ", bwtArr[i])
	}
	if string(bwtArr) != "STEXYDST.E.IXXIIXXSSMPPS.B..EE.$.USFXDIIOIIIT" {
		b.Errorf("NaiveBWT is wrong, %s", string(bwtArr))
	}
	ibwtArr := inverseBWT(string(bwtArr))
	println("\niBWT: ", string(ibwtArr))
	for i := 0; i < len(ibwtArr); i++ {
		fmt.Printf("%d ", ibwtArr[i])
	}
	if string(ibwtArr) != "SIX.MIXED.PIXIES.SIFT.SIXTY.PIXIE.DUST.BOXES" {
		b.Errorf("iBWT is wrong, %s", string(bwtArr))
	}

	regArr := regenerateSuffix(bwtArr)
	println("regenerated suffix: ", string(regArr))
}