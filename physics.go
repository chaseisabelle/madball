package main

import (
	"github.com/levenlabs/golib/timeutil"
	"github.com/vova616/chipmunk"
	"github.com/vova616/chipmunk/vect"
)

type Physics struct {
	Space *chipmunk.Space
	Last float64
}

func InitPhysics() {
	space := chipmunk.NewSpace()

	space.Gravity = vect.Vect{
		X: 0,
		Y: -900,
	}

	body := chipmunk.NewBodyStatic()

	lines := []*chipmunk.Shape{
		chipmunk.NewSegment(vect.Vect{
			X: 111.0,
			Y: 280.0,
		}, vect.Vect{
			X: 407.0,
			Y: 246.0,
		}, 0),

		chipmunk.NewSegment(vect.Vect{
			X: 407.0,
			Y: 246.0,
		}, vect.Vect{
			X: 407.0,
			Y: 343.0,
		}, 0),
	}

	for _, line := range lines {
		line.SetElasticity(0.6)
		body.AddShape(line)
	}

	space.AddBody(body)

	physics = &Physics{
		Space: space,
	}
}

func (c *Physics) Step() {
	now := timeutil.TimestampNow().Float64()

	if c.Last == 0 {
		c.Last = now
	}

	c.Space.Step(floater(now - c.Last))
}
