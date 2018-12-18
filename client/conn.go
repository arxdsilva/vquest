package main

import (
	"github.com/arxdsilva/vquest/client/display"
	ui "github.com/gen2brain/raylib-go/raygui"
	"github.com/gen2brain/raylib-go/raylib"
)

func drawConnectedScreen(d display.Display) {
	rect := rl.Rectangle{
		X:      float32(d.Width / 2),
		Y:      float32(d.Height / 2),
		Width:  100,
		Height: 20,
	}
	ui.Label(rect, resp+" you are connected")
	return
}
