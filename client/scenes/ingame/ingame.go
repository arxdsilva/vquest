package ingame

import (
	"github.com/arxdsilva/vquest/client/config"
	ui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type InGameScene struct {
	display config.Display
}

func New(cfg config.Config) *InGameScene {
	return &InGameScene{display: cfg.Display}
}

func (l *InGameScene) Render() config.GameState {
	rect := rl.Rectangle{
		X:      float32(l.display.Width / 2),
		Y:      float32(l.display.Height / 2),
		Width:  100,
		Height: 20,
	}
	ui.Label(rect, "you are connected")
	return config.InGameState
}
