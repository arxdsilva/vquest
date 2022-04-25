package scenes

import (
	"github.com/arxdsilva/vquest/client/config"
	"github.com/arxdsilva/vquest/client/scenes/login"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Scene interface {
	Draw()
	// GoTo(Scene)
}

func Start(cfg config.Config) {
	// load possible scenes
	loginScene := login.New(cfg)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		// draw active scene
		loginScene.Draw()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
