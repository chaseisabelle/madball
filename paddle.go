package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"image/color"
)

type Paddle struct {
	Drawer      imdraw.IMDraw
	Rectangle   pixel.Rect
	Velocity    pixel.Vec
	Color       color.Color
	Restitution float64
	Mass        float64
}

func InitPaddle() {
	width, height := screen.Window.Monitor().Size()
	min := pixvec(0, 0)
	max := pixvec(width / 5, height / 20)
	rectangle := pixel.R(min.X, min.Y, max.X, max.Y).Moved(min.Add(pixvec(0, height / 40)))
	velocity := pixvec(0, 0)
	drawer := imdraw.New(nil)

	paddle = &Paddle{
		Drawer:      *drawer,
		Rectangle:   rectangle,
		Velocity:    velocity,
		Color:       color.Black,
		Restitution: 0,
		Mass:        1,
	}
}

func (p *Paddle) Draw() {
	p.Drawer.Color = p.Color
	p.Drawer.Push(p.Rectangle.Min, p.Rectangle.Max)
	p.Drawer.Rectangle(0) //<< 0 thickness draws filled rect
	p.Drawer.Draw(screen.Window)
}

func (p *Paddle) Clear() {
	p.Drawer.Clear()
}

func (p *Paddle) Move() {
	// move to a new center based on mouse center
	origin := p.Rectangle.Center()
	center := pixvec(screen.Window.MousePosition().X, origin.Y)

	// make sure it stays inside the screen
	halfW := p.Rectangle.W() / 2
	halfH := p.Rectangle.H() / 2

	monW, monH := screen.Window.Monitor().Size()

	if center.X - halfW < 0 {
		center.X = halfW
	} else if center.X + halfW > monW {
		center.X = monW - halfW
	}

	if center.Y - halfH < 0 {
		center.Y = halfH
	} else if center.Y + halfH > monH {
		center.Y = monH - halfH
	}

	// update the velocity
	p.Velocity.X = center.X - origin.X
	p.Velocity.Y = center.Y - origin.Y

	// set the new center
	p.Rectangle = p.Rectangle.Moved(p.Velocity)
}


