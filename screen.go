package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
)

type Screen struct {
	Window *pixelgl.Window
	Color color.Color
}

func InitScreen() {
	monitor := pixelgl.PrimaryMonitor()
	width, height := monitor.Size()
	window, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Bounds: pixel.R(0, 0, width, height),
		VSync:  true,
		Undecorated: true,
		Resizable: false,
	})

	if err != nil {
		panic(err)
	}

	window.SetMonitor(monitor) //<< fullscreen
	window.SetCursorVisible(false)

	screen = &Screen{
		Window: window,
		Color: color.White,
	}
}

func (s *Screen) Clear() {
	s.Window.Clear(s.Color)
}

func (s *Screen) Draw() {
	s.Window.Update()
}
