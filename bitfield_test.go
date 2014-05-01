package bitfield

import (
	"testing"
	"math/rand"
	"time"
	"fmt"
)

func TestSet(t *testing.T) {
	b := New(100)
	for i := uint32(0); i < 100; i++ {
		if b.Test(i) {
			t.Errorf("%d shouldn't be set", i)
		}
		b.Set(i)
		if !b.Test(i) {
			t.Errorf("%d should be set", i)
		}
	}
}

func TestClear(t *testing.T) {
	b := New(100)
	for i := uint32(0); i < 100; i++ {
		b.Set(i)
	}

	for i := uint32(0); i < 100; i += 3 {
		b.Clear(i)
	}

	for i := uint32(0); i < 100; i++ {
		if b.Test(i) != (i % 3 != 0) {
			t.Errorf("Clear is broken!")
		}
	}
}

func TestVsMap(t *testing.T) {
	N := int(1e6)
	b := New(N)
	m := make(map[uint32]struct{})
	indexes := make([]uint32, N)

	for i := range indexes {
		indexes[i] = uint32(rand.Float64() * float64(N))
	}

	t0 := time.Now()
	for _, idx := range indexes {
		m[idx] = struct{}{}
	}

	for _, idx := range indexes {
		if _, ok := m[idx]; !ok {
			t.Errorf("%d index not found", idx)
			t.FailNow()
		}
	}

	fmt.Println("Time using map:", (time.Now().Sub(t0)))

	t0 = time.Now()
	for _, idx := range indexes {
		b.Set(idx)
	}

	for _, idx := range indexes {
		if !b.Test(idx) {
			t.Errorf("%d index not found", idx)
			t.FailNow()
		}
	}

	fmt.Println("Time using bitfield:", (time.Now().Sub(t0)))
}
