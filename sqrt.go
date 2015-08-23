package main

import (
// "math"
)

func Sqrt(x float64) (z float64) {
	z = 1.0
	lastZ := z
	for i := 0; i < 10; i++ {
		lastZ = z
		z = z - (z*z-x)/(2*z)
		if z == lastZ {
			break
		}
	}
	return
}
