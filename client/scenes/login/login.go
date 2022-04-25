package login

import (
	"github.com/arxdsilva/vquest/client/components/forms"
	"github.com/arxdsilva/vquest/client/components/text"
	"github.com/arxdsilva/vquest/client/config"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type LoginScene struct {
	display config.Display
	title   text.Component
	form    forms.Login
}

func New(cfg config.Config) *LoginScene {
	// start main window
	rl.InitWindow(
		int32(cfg.Display.Width), int32(cfg.Display.Height), "vQuest")
	rl.SetTargetFPS(cfg.TargetFPS)

	return &LoginScene{
		display: cfg.Display,
		title: text.New(
			cfg.GameName,
			int32(float32(cfg.Display.Width)*0.4*0.65),
			int32(float32(cfg.Display.Height)*0.4*0.2),
			180, rl.LightGray),
		form: forms.New(
			int32(float32(cfg.Display.Width)*0.4),
			int32(float32(cfg.Display.Height)*0.4),
			20, rl.LightGray, cfg.Display),
	}
}

func (l *LoginScene) Render() config.GameState {
	l.title.Draw()
	l.form.Draw()
	if l.form.Authenticated {
		return config.InGameState
	}
	return config.LoginState
}
