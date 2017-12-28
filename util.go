package ints

import "math"

func saturatingAdd(i, n uint64) uint64 {
	if j := i + n; j >= i {
		return j
	}
	return math.MaxUint64
}
