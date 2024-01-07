package intmath_test

import (
	"math"
	"testing"
	"testing/quick"

	"chipaca.com/intmath"
	"golang.org/x/exp/constraints"
)

func TestAbsBasic(t *testing.T) {
	for i := int64(0); i < math.MaxInt64/10; i = i*10 + 9 {
		abs := intmath.Abs(i)
		if abs != i {
			t.Errorf("Abs(%d) should be %d, got %d", i, i, abs)
		}
		abs = intmath.Abs(-i)
		if abs != i {
			t.Errorf("Abs(-%d) should be %d, got %d", i, i, abs)
		}
	}
}

func testAbsQuickF[V constraints.Signed](n V) bool {
	abs := intmath.Abs(n)
	if n < 0 {
		return abs == -n
	}
	return abs == n
}

func TestAbsQuick(t *testing.T) {
	if err := quick.Check(testAbsQuickF[int8], nil); err != nil {
		t.Error(err)
	}
	if err := quick.Check(testAbsQuickF[int16], nil); err != nil {
		t.Error(err)
	}
	if err := quick.Check(testAbsQuickF[int32], nil); err != nil {
		t.Error(err)
	}
	if err := quick.Check(testAbsQuickF[int64], nil); err != nil {
		t.Error(err)
	}
	if err := quick.Check(testAbsQuickF[int], nil); err != nil {
		t.Error(err)
	}
}
