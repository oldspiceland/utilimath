package utilimath

import ()

//IntPow is an Integer-only version of math.Pow
func IntPow(base, exp uint64) (powered uint64) {
	for i := uint64(1); i <= exp; i++ {
		base = base + base
	}
	powered = base
	return
}
