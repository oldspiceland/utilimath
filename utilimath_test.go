package utilimath

import (
	"testing"
)

func TestUInt64Pow10to2(t *testing.T) {
	base := uint64(10)
	exp := uint64(2)
	powered := UInt64Pow(base, exp)
	if powered != uint64(100) {
		t.Error("Expected 100, got ", powered)
	}
}

func TestUInt64Pow15to3(t *testing.T) {
	base := uint64(15)
	exp := uint64(3)
	powered := UInt64Pow(base, exp)
	if powered != uint64(3375) {
		t.Error("Expected 3375, got ", powered)
	}
}

func TestIntPow10to2(t *testing.T) {
	base := 10
	exp := 2
	powered := IntPow(base, exp)
	if powered != 100 {
		t.Error("Expected 100, got ", powered)
	}
}

func TestIntPow15to3(t *testing.T) {
	base := 15
	exp := 3
	powered := IntPow(base, exp)
	if powered != 3375 {
		t.Error("Expected 3375, got ", powered)
	}
}
