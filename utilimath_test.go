package utilimath

import (
	"testing"
)

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
