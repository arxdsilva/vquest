package main

import (
	"github.com/arxdsilva/vquest/client/display"
	"github.com/gen2brain/raylib-go/raylib"
)

var clientState int

const loginScreen = 0
const connected = 1

func main() {
	d := display.New()
	rl.InitWindow(int32(d.Width), int32(d.Height), "vQuest")
	rl.SetTargetFPS(60)
	loadLoginObjects(d)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		if clientState == 0 {
			drawLoginScreen(d)
		} else {
			drawConnectedScreen(d)
		}
		rl.EndDrawing()
	}
	rl.CloseWindow()
}

func mouseCollision(rec rl.Rectangle) bool {
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), rec) {
		return true
	}
	return false
}
