package intmath_test

import (
	"math"
	"math/big"
	"math/bits"
	"testing"
	"testing/quick"

	"chipaca.com/intmath"
	"golang.org/x/exp/constraints"
)

func TestBinomialBasic(t *testing.T) {
	// build pascal's triangle
	var sks [][]uint64
	kl := []uint64{1}
	sks = append(sks, kl)
	// 67 is the highest full row we can get to without overflowing
	for n := 1; n < 68; n++ {
		kn := make([]uint64, n+1)
		kn[0] = 1
		for j := 1; j < n; j++ {
			kn[j] = kl[j-1] + kl[j]
		}
		kn[n] = 1
		sks = append(sks, kn)
		kl = kn
	}

	for n, ks := range sks {
		for k, expected := range ks {
			binom, overflow := intmath.Binomial(uint64(n), uint64(k))
			if overflow || binom != expected {
				t.Errorf("Binomial(%d, %d) should be %d, got %d", n, k, expected, binom)
			}
		}
	}
}

func testBinomialQuick[V constraints.Unsigned](t *testing.T) {
	t.Parallel()
	if blen := bits.Len64(uint64(^V(0))); testing.Short() && blen >= 32 {
		t.Skipf("with %d bits this test is not short", blen)
	}
	if err := quick.Check(func(n, k V) bool {
		if uint64(n) > math.MaxInt64 || uint64(k) > math.MaxInt64 {
			// big binomial won't work
			return true
		}

		mine, overflow := intmath.Binomial(n, k)
		z := new(big.Int).Binomial(int64(n), int64(k))
		if !z.IsUint64() {
			if mine != 0 || !overflow {
				t.Logf("(%d, %d): big Binomial says !IsUint64, intmath gives %d [%v]", n, k, mine, overflow)
			}
			return mine == 0 && overflow
		}
		p := z.Uint64()
		if overflow {
			t.Logf("(%d, %d): big Binomial says IsUint64 (%d bits: %d), bits gives %d [%v]", n, k, z.BitLen(), p, mine, overflow)
		} else if p != mine {
			t.Logf("(%d, %d): big Binomial gives %d (%d bits), bits gives %d", n, k, p, z.BitLen(), mine)
		}
		return p == mine && !overflow
	}, nil); err != nil {
		t.Error(err)
	}
}

func TestBinomialQuickU8(t *testing.T)  { testBinomialQuick[uint8](t) }
func TestBinomialQuickU16(t *testing.T) { testBinomialQuick[uint16](t) }
func TestBinomailQuickU32(t *testing.T) { testBinomialQuick[uint32](t) }
func TestBinomailQuickU64(t *testing.T) { testBinomialQuick[uint64](t) }
func TestBinomailQuickU(t *testing.T)   { testBinomialQuick[uint](t) }

func bigBinomial[V constraints.Unsigned](n, k V) (uint64, bool) {
	z := new(big.Int).Binomial(int64(n), int64(k))
	if !z.IsUint64() {
		return 0, true
	}
	p := z.Uint64()
	return p, false
}

func benchmarkBinomial[V constraints.Unsigned](b *testing.B, f func(n, k V) (uint64, bool)) {
	for i := 0; i < b.N; i++ {
		for n := V(0); n < 100; n++ {
			for k := V(0); k < n; k++ {
				f(n, k)
			}
		}
	}
}

func BenchmarkMyBinomialU16(b *testing.B)  { benchmarkBinomial[uint16](b, intmath.Binomial) }
func BenchmarkBigBinomialU16(b *testing.B) { benchmarkBinomial[uint16](b, bigBinomial) }

func BenchmarkMyBinomialU64(b *testing.B)  { benchmarkBinomial[uint64](b, intmath.Binomial) }
func BenchmarkBigBinomialU64(b *testing.B) { benchmarkBinomial[uint64](b, bigBinomial) }
