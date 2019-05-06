package main

func between(x float64, a float64, b float64) bool {
	return x > a && x < b
}

func inbetween(x float64, a float64, b float64) bool {
	return x >= a && x <= b
}
