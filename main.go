package main

import (
	"fmt"
	"github.com/faiface/pixel/pixelgl"
)

var screen *Screen
var paddle *Paddle
var ball *Ball

func main() {
	pixelgl.Run(run)
}

func run() {
	InitScreen()
	InitBall()
	InitPaddle()

	for !screen.Window.Closed() {
		if !screen.Window.Focused() {
			continue
		}

		screen.Clear()
		paddle.Clear()
		ball.Clear()

		paddle.Move()
		ball.Move()

		paddle.Draw()
		ball.Draw()
		screen.Draw()
	}
}

func dbg(bla ...interface{}) {
	for _, v := range bla {
		fmt.Printf("%+v\n", v)
	}
}
