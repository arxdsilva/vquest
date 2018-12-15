package main

import "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "vQuest")
	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Login:", 190, 200, 20, rl.LightGray)
		rl.DrawText("Password:", 190, 180, 20, rl.LightGray)
		rl.EndDrawing()
		rl.ToggleFullscreen(true)
	}
	rl.CloseWindow()
}
