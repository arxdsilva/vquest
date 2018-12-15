package main

import (
	"github.com/arxdsilva/vquest/client/display"
	"github.com/gen2brain/raylib-go/raylib"
)

func main() {
	d := display.New()
	rl.InitWindow(int32(d.Width), int32(d.Height), "vQuest")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		tH1 := float32(d.Height) * 0.4
		tH2 := float32(d.Height) * 0.5
		wdt := float32(d.Width) * 0.3
		rl.DrawText("Login:", int32(wdt), int32(tH1), 20, rl.LightGray)
		rl.DrawText("Password:", int32(wdt), int32(tH2), 20, rl.LightGray)
		rl.EndDrawing()
	}
	rl.CloseWindow()
}
