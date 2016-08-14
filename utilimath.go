package utilimath

import ()

//IntPow is an Integer-only version of math.Pow
func UInt64Pow(base, exp uint64) (powered uint64) {
	powered = base
	for i := uint64(1); i < exp; i++ {
		powered = powered * base
	}
	return
}
