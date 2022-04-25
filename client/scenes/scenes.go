package scenes

import (
	"github.com/arxdsilva/vquest/client/config"
	"github.com/arxdsilva/vquest/client/scenes/ingame"
	"github.com/arxdsilva/vquest/client/scenes/login"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Manager interface {
	Render()
}

type SceneManager struct {
	cfg    config.Config
	state  config.GameState
	login  *login.LoginScene
	ingame *ingame.InGameScene
}

func NewManager(cfg config.Config) *SceneManager {
	return &SceneManager{
		cfg:    cfg,
		state:  config.LoginState,
		login:  login.New(cfg),
		ingame: ingame.New(cfg),
	}
}

func (m *SceneManager) Render() {
	switch m.state {
	case config.LoginState:
		m.state = m.login.Render()
	case config.InGameState:
		m.state = m.ingame.Render()
	}
}

func Start(m Manager) {
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		m.Render()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
