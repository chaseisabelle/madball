package main

import "math/rand"

func imbetween(x float64, a float64, b float64) bool {
	return x >= a && x <= b
}

func randFloat64n(min float64, max float64) float64 {
	return rand.Float64() * min + max
}
