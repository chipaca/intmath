package intmath_test

import (
	"math"
	"math/bits"
	"testing"
	"testing/quick"

	"chipaca.com/intmath"
)

func TestSqrtBasic(t *testing.T) {
	for i := uint64(0); i < 1000; i++ {
		sqrt := intmath.Sqrt(i * i)
		if sqrt != i {
			t.Errorf("Sqrt(%d) should be %d, got %d", i*i, i, sqrt)
		}
	}
}

func TestSqrtQuick(t *testing.T) {
	f := func(n uint64) bool {
		overflow, pow := bits.Mul64(n, n)
		if overflow > 0 {
			return true
		}
		sqrt := intmath.Sqrt(pow)
		return sqrt == n
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSqrtQuick2(t *testing.T) {
	f := func(n uint64) bool {
		fsqrt := uint64(math.Floor(math.Sqrt(float64(n))))
		isqrt := intmath.Sqrt(n)
		return fsqrt == isqrt
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func BenchmarkSqrt(b *testing.B) {
	for name, f := range map[string]func(uint64) uint64{
		"Sqrt":  intmath.Sqrt[uint64],
		"SqrtI": SqrtI,
	} {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				f(uint64(i))
			}
		})
	}
}

// actually slower (on machines with hardware float sqrt):
func SqrtI(n uint64) uint64 {
	if n == 0 {
		return 0
	}
	c := 31 - (63-bits.LeadingZeros64(n))/2
	u := asqrt(n<<(2*c)) >> c
	if u*u-1 >= n {
		u--
	}
	return u
}

func asqrt(n uint64) uint64 {
	var u = 1 + (n >> 62)
	u = (u << 1) + (n>>59)/u
	u = (u << 3) + (n>>53)/u
	u = (u << 7) + (n>>41)/u
	return (u << 15) + (n>>17)/u
}
