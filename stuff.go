package main

import (
	"github.com/vova616/chipmunk"
	"github.com/vova616/chipmunk/vect"
)

type Stuff struct {
	Space *chipmunk.Space
}

func InitStuff() {
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

	stuff = &Stuff{
		Space: space,
	}
}
