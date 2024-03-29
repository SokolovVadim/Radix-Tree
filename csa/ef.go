package csa
// https://github.com/amallia/go-ef
import (
	"errors"
	"github.com/RoaringBitmap/roaring"
	"log"
	"math"
)

const (
	efInfo = `Universe: %d
Elements: %d
Lower_bits: %d
Higher_bits_length: %d
Mask: 0b%b
Lower_bits offset: %d
Bitvector length: %d
`
)

// EliasFano codec structure
type CompressedText struct {
	universe         uint64
	n                uint64
	lowerBits        uint64
	higherBitsLength uint64
	mask             uint64
	lowerBitsOffset  uint64
	bvLen            uint64
	b                *roaring.Bitmap
	curValue         uint64
	position         uint64
	highBitsPos      uint64
}

// New creates a new empty EliasFano object
func NewEF(universe uint64, n uint64) *CompressedText {
	var lowerBits uint64
	if lowerBits = 0; universe > n {
		lowerBits = msb(universe / n)
	}
	higherBitsLength := n + (universe >> lowerBits) + 2
	mask := (uint64(1) << lowerBits) - 1
	lowerBitsOffset := higherBitsLength
	bvLen := lowerBitsOffset + n*uint64(lowerBits)
	b := roaring.NewBitmap()
	return &CompressedText{universe, n, lowerBits, higherBitsLength, mask, lowerBitsOffset, bvLen, b, 0, 0, 0}
}

// Compress a monotone increasing array of positive integers. It sets the position at the beginning.
func (ef *CompressedText) Compress(elems []uint64) {
	// fmt.Println("compressible:", elems)
	last := uint64(0)

	for i, elem := range elems {
		if i > 0 && elem < last {
			log.Fatal("Sequence is not sorted")
		}
		if elem > ef.universe {
			log.Fatalf("Element %d is greater than universe", elem)
		}
		ef.b.Add(uint32(elem) + uint32(i) + 1)
	}
	// fmt.Println("decode after compressing: ", ef.b.String())
}

func (ef* CompressedText) getVal(k uint32) uint32 {
	value, _ := ef.b.Select(k)
	return value - k - 1
}

func (ef* CompressedText) getMany(length int) []uint32 {
	arr := make([]uint32, length)
	for i := 0; i < length; i++ {
		arr[i], _ = ef.b.Select(uint32(i))
		arr[i] -= uint32(i) + 1
	}
	return arr
}

// Next moves the internal iterator to the next position and returns a value or an error.
func (ef *CompressedText) Next() (uint64, error) {
	ef.position++
	if ef.position >= ef.Size() {
		return 0, errors.New("End reached")
	}
	ef.readCurrentValue()
	return ef.Value(), nil
}

// Position return the current position of the internal iterator.
func (ef *CompressedText) Position() uint64 {
	return ef.position
}

// Reset moves the internal iterator to the beginning.
func (ef *CompressedText) Reset() {
	ef.highBitsPos = 0
	ef.position = 0
	ef.readCurrentValue()
}

// Info prints info regarding the EliasFano codec.
func (ef *CompressedText) Info() {
	log.Printf(efInfo, ef.universe, ef.n, ef.lowerBits, ef.higherBitsLength, ef.mask, ef.lowerBitsOffset, ef.bvLen)
}

// Value returns the value of the current element.
func (ef *CompressedText) Value() uint64 {
	return ef.curValue
}

// Size returns the number of elements encoded.
func (ef *CompressedText) Size() uint64 {
	return ef.n
}

// Bitsize returns the size of the internal bitvector.
func (ef *CompressedText) Bitsize() uint64 {
	return uint64(ef.b.GetSizeInBytes())
}

func SetTo(b* roaring.Bitmap, i uint32, value bool) {
	if value {
		b.Add(i)
	} else {
		b.Remove(i)
	}
}

func setBits(b *roaring.Bitmap, offset uint64, bits uint64, length uint64) {
	for i := uint64(0); i < length; i++ {
		val := bits & (1 << (length - i - 1))
		SetTo(b, uint32(offset+i+1), val > 0)
	}
}

func (ef *CompressedText) readCurrentValue() uint32 {
	value, _ := ef.b.Select(uint32(ef.highBitsPos))
	ef.curValue = uint64(value)
	return value
}

func round(a float64) int64 {
	if a < 0 {
		return int64(a - 0.5)
	}
	return int64(a + 0.5)
}

func msb(x uint64) uint64 {
	return uint64(round(math.Log2(float64(x))))
}
