package intmath

import (
	"math"
	"math/bits"

	"golang.org/x/exp/constraints"
)

var log2table = [...]uint64{
	63, 0, 58, 1, 59, 47, 53, 2,
	60, 39, 48, 27, 54, 33, 42, 3,
	61, 51, 37, 40, 49, 18, 28, 20,
	55, 30, 34, 11, 43, 14, 22, 4,
	62, 57, 46, 52, 38, 26, 32, 41,
	50, 36, 17, 19, 29, 10, 13, 21,
	56, 45, 25, 31, 35, 16, 9, 12,
	44, 24, 15, 8, 23, 7, 6, 5}

// FloorLog10(u) returns ⌊log₂u⌋
//
// If u is 0, you'll get a nonsense result.
//
// This algorithm is from “Find the log base 2 of an N-bit integer in O(lg(N))
// operations with multiply and lookup” from Sean Eron Anderson's “Bit Twiddling
// Hacks”
func FloorLog2[V constraints.Unsigned](u V) V {
	u |= u >> 1
	u |= u >> 2
	u |= u >> 4
	u |= u >> 8
	u |= u >> 16
	u |= u >> 32

	return V(log2table[(uint64(u-u>>1)*0x07EDD5E59A4E28C2)>>58])
}

// CeilLog2(u) returns ⌈log₂u⌉.
//
// If u is 0, you'll get a nonsense result.
//
// This one cheats, as ⌈log₂u⌉ is just `math/bits.Len64(u-1)`.
func CeilLog2[V constraints.Unsigned](u V) V {
	// when type switches on parametric types are done, use that
	// to call the potentially-cheaper functions from bits when
	// possible. At least on my machine it doesn't make a measurable
	// difference though (~¼ns per op)
	return V(bits.Len64(uint64(u - 1)))
}

var log10table = [65]struct {
	log uint64
	off uint64
}{
	{log: 1},
	{log: 1, off: 1}, {log: 1}, {log: 1},
	{log: 2, off: 10}, {log: 2}, {log: 2},
	{log: 3, off: 100}, {log: 3}, {log: 3},
	{log: 4, off: 1000}, {log: 4}, {log: 4}, {log: 4},
	{log: 5, off: 10000}, {log: 5}, {log: 5},
	{log: 6, off: 100000}, {log: 6}, {log: 6},
	{log: 7, off: 1000000}, {log: 7}, {log: 7}, {log: 7},
	{log: 8, off: 10000000}, {log: 8}, {log: 8},
	{log: 9, off: 100000000}, {log: 9}, {log: 9},
	{log: 10, off: 1000000000}, {log: 10}, {log: 10}, {log: 10},
	{log: 11, off: 10000000000}, {log: 11}, {log: 11},
	{log: 12, off: 100000000000}, {log: 12}, {log: 12},
	{log: 13, off: 1000000000000}, {log: 13}, {log: 13}, {log: 13},
	{log: 14, off: 10000000000000}, {log: 14}, {log: 14},
	{log: 15, off: 100000000000000}, {log: 15}, {log: 15},
	{log: 16, off: 1000000000000000}, {log: 16}, {log: 16}, {log: 16},
	{log: 17, off: 10000000000000000}, {log: 17}, {log: 17},
	{log: 18, off: 100000000000000000}, {log: 18}, {log: 18},
	{log: 19, off: 1000000000000000000}, {log: 19}, {log: 19}, {log: 19},
	{log: 20, off: 10000000000000000000},
}

// Len(n) returns the length of the base-10 string of n.
//
// Len(u) is essentially ⌊log₁₀u⌋+1.
func Len[V constraints.Integer](n V) V {
	var sgn uint64
	var v uint64
	if n < 0 {
		sgn = 1
		x := int64(n)
		if x == math.MinInt64 {
			// -x won't work for this one number
			return 20
		}
		v = uint64(-x)
	} else {
		v = uint64(n)
	}
	return V(floorLog10PlusOne(v) + sgn)
}

// FloorLog10(u) returns ⌊log₁₀u⌋
//
// If u is 0, you'll get a nonsense result.
//
// This starts from Sean Eron Anderson's “Bit Twiddling Hacks” entry on “Find
// integer log base 10 of an integer”, but combines that with the fact that
// ⌈log₂u⌉ is just `math/bits.Len64(u-1)` and that in most places that's a
// hardware instruction. There are some pesky offsets that need to be tracked in
// a table.
func FloorLog10[V constraints.Unsigned](u V) V {
	return V(floorLog10PlusOne(uint64(u)) - 1)
}

// a helper for both FloorLog10 and Len
func floorLog10PlusOne(v uint64) uint64 {
	x := log10table[bits.Len64(v)]
	var d uint64
	if x.off > v {
		d = 1
	}
	return x.log - d
}

// CeilLog10(u) returns ⌈log₁₀u⌉.
//
// If u is 0, you'll get a nonsense result.
//
// This is almost exactly the same as `FloorLog10`. The difference is a single
// `>=` instead of a `>` when deciding whether an offset is needed.
func CeilLog10[V constraints.Unsigned](u V) V {
	v := uint64(u)
	x := log10table[bits.Len64(v)]
	var d uint64
	if x.off >= v {
		d = 1
	}
	return V(x.log - d)
}
