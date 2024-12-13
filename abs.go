package intmath

import "golang.org/x/exp/constraints"

// Abs(n) computes the absolute value of n, without branching.
//
// I haven't been able to track down the original source of this algorithm.
// Read more on Sean Eron Anderson's “Bit Twiddling Hacks”,
// https://graphics.stanford.edu/~seander/bithacks.html#IntegerAbs
func Abs[V constraints.Signed](n V) V {
	mask := n >> 63
	return (n + mask) ^ mask
}
