package intmath_test

import (
	"math"
	"testing"
	"testing/quick"

	"chipaca.com/intmath"
)

func TestPowBasic(t *testing.T) {
	for n, expected := range map[uint64]uint64{
		0: 1,
		1: 2,
		2: 4,
	} {
		pow := intmath.Pow(2, n)
		if pow != expected {
			t.Errorf("2**%d should be %d, got %d", n, expected, pow)
		}
	}
}

func TestPowQuick(t *testing.T) {
	f := func(m, n uint64) bool {
		p := math.Pow(float64(m), float64(n))
		if p > math.MaxUint64 {
			return true
		}
		return uint64(p) == intmath.Pow(m, n)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestPowXQuick(t *testing.T) {
	f := func(m, n uint64) bool {
		pow := math.Pow(float64(m), float64(n))

		intpow, overflow := intmath.PowX(m, n)
		if pow > math.MaxUint64 {
			return overflow
		}

		return uint64(pow) == intpow
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func floatPow(m, n uint64) uint64 {
	return uint64(math.Pow(float64(m), float64(n)))
}

func BenchmarkPow(b *testing.B) {
	for name, f := range map[string]func(m, n uint64) uint64{
		"intmath.Pow": intmath.Pow[uint64],
		"floatPow":    floatPow,
	} {
		b.Run(name, func(b *testing.B) {
			n := uint64(b.N)
			for i := uint64(0); i < n; i++ {
				for j := uint64(0); j < n; j++ {
					f(i, j)
				}
			}
		})
	}
}

func floatPowX(m, n uint64) (uint64, bool) {
	p := math.Pow(float64(m), float64(n))
	if p > math.MaxUint64 {
		return 0, true
	}
	return uint64(p), false
}

func BenchmarkPowX(b *testing.B) {
	for name, f := range map[string]func(m, n uint64) (uint64, bool){
		"intmath.PowX": intmath.PowX[uint64],
		"floatPowX":    floatPowX,
	} {
		b.Run(name, func(b *testing.B) {
			n := uint64(b.N)
			for i := uint64(0); i < n; i++ {
				for j := uint64(0); j < n; j++ {
					f(i, j)
				}
			}
		})
	}
}
