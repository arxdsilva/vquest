package forms

import (
	"image/color"

	"github.com/arxdsilva/vquest/client/components/input"
	"github.com/arxdsilva/vquest/client/components/text"
	"github.com/arxdsilva/vquest/client/config"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const lineDistancingFactor = 1.15
const columnDistancingFactor = 1.4

type Login struct {
	x, y              int32
	fontSize          int32
	color             color.RGBA
	userText          text.Component
	userPasswordText  text.Component
	userNameInput     input.Component
	userPasswordInput input.Component
}

func New(x, y, fontSize int32, color color.RGBA, cfg config.Display) Login {
	return Login{
		x:        x,
		y:        y,
		fontSize: fontSize,
		color:    color,
		userText: text.New(
			"User:", x, y, fontSize, color),
		userPasswordText: text.New(
			"Password:",
			x,
			int32(float32(y)*lineDistancingFactor),
			fontSize, color),
		userNameInput: input.New(
			float32(x)*columnDistancingFactor,
			float32(y),
			float32(cfg.Width)*0.1,
			float32(fontSize),
			false,
			rl.LightGray,
		),
		userPasswordInput: input.New(
			float32(x)*columnDistancingFactor,
			float32(y)*lineDistancingFactor,
			float32(cfg.Width)*0.1,
			float32(fontSize),
			true,
			rl.LightGray,
		),
	}
}

func (l *Login) Draw() {
	l.userText.Draw()
	l.userPasswordText.Draw()
	l.userNameInput.Draw()
	l.userPasswordInput.Draw()
}
