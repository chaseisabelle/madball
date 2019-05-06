package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
	"image/color"
	"math"
)

type Ball struct {
	Drawer      imdraw.IMDraw
	Circle      pixel.Circle
	Velocity    pixel.Vec
	Color       color.Color
	Restitution float64
	Mass        float64
	minVel      pixel.Vec
	maxVel      pixel.Vec
}

func InitBall() {
	monW, monH := screen.Window.Monitor().Size()

	drawer := imdraw.New(nil)
	circle := pixel.C(pixel.V(monW / 2, monH / 2), monW / 50)
	minVel := pixel.V(1, 5)
	maxVel := pixel.V(10, 10)
	velocity := pixel.V(randFloat64n(minVel.X, maxVel.X), -randFloat64n(minVel.Y, maxVel.Y))

	ball = &Ball{
		Drawer:      *drawer,
		Circle:      circle,
		Velocity:    velocity,
		Color:       colornames.Blue,
		Restitution: 0,
		Mass:        1,
		minVel:minVel,
		maxVel:maxVel,
	}
}

func (b *Ball) Draw() {
	b.Drawer.Color = b.Color
	b.Drawer.Push(b.Circle.Center)
	b.Drawer.Circle(b.Circle.Radius, 0) //<< 0 thickness for filled circle
	b.Drawer.Draw(screen.Window)
}

func (b *Ball) Clear() {
	b.Drawer.Clear()
}

func (b *Ball) Move() {
	// get the new center
	center := b.Circle.Center.Add(b.Velocity)
	radius := b.Circle.Radius

	// make sure we stay inside the screen
	monW, monH := screen.Window.Monitor().Size()

	if center.X - radius < 0 {
		center.X = radius
		b.Velocity.X = math.Abs(b.Velocity.X)
	} else if center.X + radius > monW {
		center.X = monW - radius
		b.Velocity.X = -math.Abs(b.Velocity.X)
	}

	if center.Y - radius < 0 {
		center.Y = radius
		b.Velocity.Y = math.Abs(b.Velocity.Y)
	} else if center.Y + radius > monH {
		center.Y = monH - radius
		b.Velocity.Y = -math.Abs(b.Velocity.Y)
	}

	// did we hit the paddle?
	intersect := b.Circle.IntersectRect(paddle.Rectangle)

	if intersect.X != 0 || intersect.Y != 0 {
		center = center.Add(intersect)
		b.Velocity = intersect

		if imbetween(b.Velocity.X, 0, b.minVel.X) {
			b.Velocity.X = b.minVel.X
		} else if b.Velocity.X > b.maxVel.X {
			b.Velocity.X = b.maxVel.X
		} else if imbetween(b.Velocity.X, -b.minVel.X, 0) {
			b.Velocity.X = -b.minVel.X
		} else if b.Velocity.X < -b.maxVel.X {
			b.Velocity.X = -b.maxVel.X
		}

		if imbetween(b.Velocity.Y, 0, b.minVel.Y) {
			b.Velocity.Y = b.minVel.Y
		} else if b.Velocity.Y > b.maxVel.Y {
			b.Velocity.Y = b.maxVel.Y
		} else if imbetween(b.Velocity.Y, -b.minVel.Y, 0) {
			b.Velocity.Y = -b.minVel.Y
		} else if b.Velocity.Y < -b.maxVel.Y {
			b.Velocity.Y = -b.maxVel.Y
		}

		paddleCenter := paddle.Rectangle.Center()

		if b.Velocity.X < 0 && b.Circle.Center.X > paddleCenter.X || b.Velocity.X > 0 && b.Circle.Center.X < paddleCenter.X {
			b.Velocity.X = -b.Velocity.X
		}

		if b.Velocity.Y < 0 && b.Circle.Center.Y > paddleCenter.Y || b.Velocity.Y > 0 && b.Circle.Center.Y < paddleCenter.Y {
			b.Velocity.Y = -b.Velocity.Y
		}
	}

	// set the new center
	b.Circle.Center = center
}

/*
	// did we hit the paddle?
	intersect := b.Circle.IntersectRect(paddle.Rectangle)

	// collision detection
	if intersect.X != 0 || intersect.Y != 0 {
		// collision resolution
		relVel := paddle.Velocity.Sub(b.Velocity)

		// intersect is the "normal" ?
		relVelAlongNorm := relVel.Dot(intersect.Normal()) //intersect.X * relVel.X + intersect.Y * relVel.Y

		if relVelAlongNorm <= 0 {
			e := math.Min(paddle.Restitution, b.Restitution) //<< elasticity

			j := -(1 + e) * relVelAlongNorm
			j /= 1 / paddle.Mass + 1 / b.Mass

			impulse := pixel.V(j * intersect.X, j * intersect.Y)

			b.Velocity.X += 1 / b.Mass * impulse.X
			b.Velocity.Y += 1 / b.Mass * impulse.Y
		}
	}

 */

/*
void ResolveCollision( Object A, Object B )
{
	// Calculate relative velocity
	Vec2 rv = B.velocity - A.velocity

	// Calculate relative velocity in terms of the normal direction
	float velAlongNormal = DotProduct( rv, normal )

	// Do not resolve if velocities are separating
	if(velAlongNormal > 0)
	return;

	// Calculate restitution
	float e = min( A.restitution, B.restitution)

	// Calculate impulse scalar
	float j = -(1 + e) * velAlongNormal
	j /= 1 / A.mass + 1 / B.mass

	// Apply impulse
	Vec2 impulse = j * normal
	A.velocity -= 1 / A.mass * impulse
	B.velocity += 1 / B.mass * impulse
}

func DotProduct(a pixel.Vec, b pixel.Vec) float64 {
	return a.X * b.X + a.Y * b.Y
}

*/