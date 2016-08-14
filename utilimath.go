package utilimath

import (
	"fmt"
)

//IntPow is an Integer-only version of math.Pow
func IntPow(base, exp uint64) (powered uint64) {
	for i := 1; i <= exp; i++ {
		base = base * 2
	}
	powered = base
	return
}
