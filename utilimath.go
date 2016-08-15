package utilimath

import (
	"math/big"
)

var (
	Version   string
	BuildTime string
)

//IntPow is an Integer-only version of math.Pow
func UInt64Pow(base, exp uint64) (powered uint64) {
	powered = base
	for i := uint64(1); i < exp; i++ {
		powered = powered * base
	}
	return
}

func IntPow(base, exp int) (powered int) {
	powered = base
	for i := 1; i < exp; i++ {
		powered = powered * base
	}
	return
}

func BigIntPow(base, exp big.Int) (powered big.Int) {
	powered = base
	for i := 1; i < exp; i++ {
		powered = powered * base
	}
	return
}
