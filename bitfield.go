package bitfield

import (
	"math"
)

type BitField []byte

func New(n int) BitField {
	n = int(math.Ceil(float64(n) / 8.0))
	return BitField(make([]byte, n))
}

func (b BitField) Set(i uint32) {
	if (int(i) > len(b) * 8) {
		return
	}

	idx, offset := (i / 8), (i % 8)

	b[idx] |= (1 << uint(offset))
}

func (b BitField) Clear(i uint32) {
	if (int(i) > len(b) * 8) {
		return
	}

	idx, offset := (i / 8), (i % 8)
	b[idx] &= ^(1 << uint(offset))
}

func (b BitField) Test(i uint32) bool {
	if (int(i) > len(b) * 8) {
		return false
	}

	idx, offset := (i / 8), (i % 8)
	return (b[idx] & (1 << uint(offset))) != 0
}
