package intmath

import (
	"math/bits"

	"golang.org/x/exp/constraints"
)

// Pow(m, n) == mⁿ
//
// from TAOCPv2, §4.6.3 (p462 in the 3rd edition)
func Pow[V constraints.Unsigned](m, n V) V {
	pow := V(1)

	for n > 0 {
		if n&1 == 1 {
			pow = m * pow
		}
		m = m * m
		n = n >> 1
	}
	return pow
}

// PowX(m, n) == ‹mⁿ, false› as long as it fits in a uint64. As soon
// as it overflows it returns ‹0, true›.
//
// from TAOCPv2, §4.6.3 (p462 in the 3rd edition)
func PowX[V constraints.Unsigned](m, n V) (pow uint64, overflowed bool) {
	// I'm still not convinced about returning uint64 instead of V here
	// but the overflow detection would be a bit more expensive and that ends up deciding it for now

	pow = uint64(1)
	var overflow uint64
	mu := uint64(m)
	for n > 0 {
		if n&1 == 1 {
			overflow, pow = bits.Mul64(mu, pow)
			if overflow > 0 {
				return 0, true
			}
		}
		overflow, mu = bits.Mul64(mu, mu)
		if overflow > 0 {
			return 0, true
		}
		n = n >> 1
	}
	return pow, false
}
