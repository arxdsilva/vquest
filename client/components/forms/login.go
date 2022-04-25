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
	x, y                      int32
	fontSize                  int32
	color                     color.RGBA
	userText                  text.Component
	userPasswordText          text.Component
	userNameInput             input.Component
	userPasswordInput         input.Component
	isUserNameInputActive     bool
	isUserPasswordInputActive bool
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
			true,
			rl.LightGray,
		),
		userPasswordInput: input.New(
			float32(x)*columnDistancingFactor,
			float32(y)*lineDistancingFactor,
			float32(cfg.Width)*0.1,
			float32(fontSize),
			true,
			false,
			rl.LightGray,
		),
		isUserNameInputActive:     true,
		isUserPasswordInputActive: false,
	}
}

func (l *Login) Draw() {
	l.manageInputHighlight()
	l.userText.Draw()
	l.userPasswordText.Draw()
	l.userNameInput.Draw()
	l.userPasswordInput.Draw()
}

func (l *Login) manageInputHighlight() {
	if rl.IsKeyReleased(rl.KeyTab) {
		l.toggleInputs()
		return
	}
	if l.userNameInput.Clicked() || l.userPasswordInput.Clicked() {
		l.toggleInputs()
	}
}

func (l *Login) toggleInputs() {
	l.userNameInput.ToggleActive()
	l.userPasswordInput.ToggleActive()
}
