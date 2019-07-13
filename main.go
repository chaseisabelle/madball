package main

import (
	"github.com/faiface/pixel/pixelgl"
)

var screen *Screen
var physics *Physics
var paddle *Paddle
var ball *Ball

func main() {
	pixelgl.Run(run)
}

func run() {
	InitScreen()
	InitPhysics()
	InitBall()
	InitPaddle()

	for !screen.Window.Closed() {
		if !screen.Window.Focused() {
			continue
		}

		screen.Clear()
		paddle.Clear()
		ball.Clear()

		physics.Step()

		paddle.Move()
		ball.Move()

		paddle.Draw()
		ball.Draw()
		screen.Draw()
	}
}
