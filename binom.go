package intmath

import (
	"math/bits"

	"golang.org/x/exp/constraints"
)

// Binomial(n, k) == ‹(n k), false›.
// Unless the result overflows uint64, in which case it returns ‹0, true›.
func Binomial[V constraints.Unsigned](n, k V) (uint64, bool) {
	if k > n {
		return 0, false
	}
	if k == 0 || k == n {
		return 1, false
	}
	k = min(k, n-k)
	n64 := uint64(n)
	k64 := uint64(k)
	c := uint64(1)

	for i := uint64(1); i <= k64; i, n64 = i+1, n64-1 {
		q, r := bits.Div64(0, c, i)
		overflow, a := bits.Mul64(q, n64)
		if overflow != 0 {
			return 0, true
		}
		q, r = bits.Mul64(r, n64)
		q, _ = bits.Div64(q, r, i)
		c, overflow = bits.Add64(a, q, 0)
		if overflow != 0 {
			return 0, true
		}
	}
	return c, false
}
