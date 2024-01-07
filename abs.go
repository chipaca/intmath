package intmath

import "golang.org/x/exp/constraints"

// Abs(n) computes the absolute value of n, without branching.
func Abs[V constraints.Signed](n V) V {
	mask := n >> 63
	return (n + mask) ^ mask
}
