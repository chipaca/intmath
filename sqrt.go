package intmath

import (
	"math"

	"golang.org/x/exp/constraints"
)

// Sqrt(u) returns √u
//
// NOTE the current implementation uses floating point as that is
// implemented in hardware in most places (see SqrtI and associated
// benchmarks in the tests). This might change if somebody points me
// to a faster version.
func Sqrt[V constraints.Unsigned](n V) V {
	return V(math.Sqrt(float64(n)))
}
