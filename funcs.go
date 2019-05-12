package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/vova616/chipmunk/vect"
	"math/rand"
)

func poop(bla ...interface{}) {
	for _, v := range bla {
		fmt.Printf("%+v\n", v)
	}
}

func imbetween(x float64, a float64, b float64) bool {
	return x >= a && x <= b
}

func randumb(min float64, max float64) float64 {
	return rand.Float64() * min + max
}

func floater(n float64) vect.Float {
	return vect.Float(n)
}

func phyvec(x float64, y float64) vect.Vect {
	return vect.Vect{
		X: floater(x),
		Y: floater(y),
	}
}

func pixvec(x float64, y float64) pixel.Vec {
	return pixel.V(x, y)
}
