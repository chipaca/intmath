package intmath_test

import (
	"fmt"
	"math"
	"strconv"
	"testing"
	"testing/quick"

	"golang.org/x/exp/constraints"

	"chipaca.com/intmath"
)

func TestLog10Basic(t *testing.T) {
	for i, n := uint64(0), uint64(1); n < math.MaxUint64/10; i, n = i+1, n*10 {
		log := intmath.FloorLog10(n)
		if log != i {
			t.Errorf("FloorLog10(%d) should be %d, got %d", n, i, log)
		}
		log = intmath.CeilLog10(n)
		if log != i {
			t.Errorf("CeilLog10(%d) should be %d, got %d", n, i, log)
		}
	}
}

func TestLog2Basic(t *testing.T) {
	for i, n := uint64(0), uint64(1); n < math.MaxUint64/10; i, n = i+1, n*2 {
		log := intmath.FloorLog2(n)
		if log != i {
			t.Errorf("FloorLog2(%d) should be %d, got %d", n, i, log)
		}
		log = intmath.CeilLog2(n)
		if log != i {
			t.Errorf("CeilLog2(%d) should be %d, got %d", n, i, log)
		}
	}
}

func TestFloorLog10Quick(t *testing.T) {
	f := func(n uint64) uint64 {
		return uint64(math.Floor(math.Log10(float64(n))))
	}
	if err := quick.CheckEqual(f, intmath.FloorLog10[uint64], nil); err != nil {
		t.Error(err)
	}
}

func TestCeilLog10Quick(t *testing.T) {
	f := func(n uint64) uint64 {
		return uint64(math.Ceil(math.Log10(float64(n))))
	}
	if err := quick.CheckEqual(f, intmath.CeilLog10[uint64], nil); err != nil {
		t.Error(err)
	}
}

func TestFloorLog2Quick(t *testing.T) {
	f := func(n uint64) uint64 {
		return uint64(math.Floor(math.Log2(float64(n))))
	}
	if err := quick.CheckEqual(f, intmath.FloorLog2[uint64], nil); err != nil {
		t.Error(err)
	}
}

func TestCeilLog2Quick(t *testing.T) {
	f := func(n uint64) uint64 {
		return uint64(math.Ceil(math.Log2(float64(n))))
	}
	if err := quick.CheckEqual(f, intmath.CeilLog2[uint64], nil); err != nil {
		t.Error(err)
	}
}

func testLen(t *testing.T, n uint64, expected uint64) {
	l := intmath.Len(n)
	if l != expected {
		t.Errorf("Len(%d) should be %d, got %d", n, expected, l)
	}
}

func TestLenBasic(t *testing.T) {
	for i, n := uint64(1), uint64(1); n < math.MaxUint64/10; i, n = i+1, n*10 {
		testLen(t, n, i)
	}
	testLen(t, 0x838178adfc68a64f, 19)
}

func lenFmt[V constraints.Integer](n V) V {
	return V(len(fmt.Sprint(n)))
}

func lenStrconv[V constraints.Integer](n V) V {
	if V(0)-1 < 0 {
		return V(len(strconv.FormatInt(int64(n), 10)))
	}
	return V(len(strconv.FormatUint(uint64(n), 10)))
}

func lenMath[V constraints.Integer](n V) V {
	extra := V(1)
	if n < 0 {
		extra = 2
		n = -n
	}
	return V(math.Floor(math.Log10(float64(n)))) + extra
}

// this one just checks that lenMath and lenStrconv give reasonable results
func TesOtherLenQuick(t *testing.T) {
	if err := quick.CheckEqual(lenStrconv[int32], intmath.Len[int32], nil); err != nil {
		t.Error(err)
	}
	if err := quick.CheckEqual(lenMath[int32], intmath.Len[int32], nil); err != nil {
		t.Error(err)
	}
	if err := quick.CheckEqual(lenFmt[int32], intmath.Len[int32], nil); err != nil {
		t.Error(err)
	}
	if err := quick.CheckEqual(lenStrconv[uint32], intmath.Len[uint32], nil); err != nil {
		t.Error(err)
	}
	if err := quick.CheckEqual(lenMath[uint32], intmath.Len[uint32], nil); err != nil {
		t.Error(err)
	}
	if err := quick.CheckEqual(lenFmt[uint32], intmath.Len[uint32], nil); err != nil {
		t.Error(err)
	}
}

func TestLenQuick(t *testing.T) {
	if err := quick.CheckEqual(lenStrconv[int8], intmath.Len[int8], nil); err != nil {
		t.Error(err)
	}
	if err := quick.CheckEqual(lenStrconv[int16], intmath.Len[int16], nil); err != nil {
		t.Error(err)
	}
	if err := quick.CheckEqual(lenStrconv[int32], intmath.Len[int32], nil); err != nil {
		t.Error(err)
	}
	if err := quick.CheckEqual(lenStrconv[int64], intmath.Len[int64], nil); err != nil {
		t.Error(err)
	}
	if err := quick.CheckEqual(lenStrconv[uint8], intmath.Len[uint8], nil); err != nil {
		t.Error(err)
	}
	if err := quick.CheckEqual(lenStrconv[uint16], intmath.Len[uint16], nil); err != nil {
		t.Error(err)
	}
	if err := quick.CheckEqual(lenStrconv[uint32], intmath.Len[uint32], nil); err != nil {
		t.Error(err)
	}
	if err := quick.CheckEqual(lenStrconv[uint64], intmath.Len[uint64], nil); err != nil {
		t.Error(err)
	}
}

func benchmarkLen(b *testing.B, f func(uint32) uint32) {
	for i := 0; i < b.N; i++ {
		f(uint32(i))
	}
}

func BenchmarkLenFmt(b *testing.B)     { benchmarkLen(b, lenFmt[uint32]) }
func BenchmarkLenStrconv(b *testing.B) { benchmarkLen(b, lenStrconv[uint32]) }
func BenchmarkLenMath(b *testing.B)    { benchmarkLen(b, lenMath[uint32]) }
func BenchmarkLenInt(b *testing.B)     { benchmarkLen(b, intmath.Len[uint32]) }

func BenchmarkLog(b *testing.B) {
	for name, f := range map[string]func(uint64) uint64{
		"Len":        intmath.Len[uint64],
		"CeilLog2":   intmath.CeilLog2[uint64],
		"CeilLog10":  intmath.CeilLog10[uint64],
		"FloorLog10": intmath.FloorLog10[uint64],
	} {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				f(uint64(i))
			}
		})
	}
}
