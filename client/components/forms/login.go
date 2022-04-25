package forms

import (
	"image/color"

	"github.com/arxdsilva/vquest/client/components/input"
	"github.com/arxdsilva/vquest/client/components/text"
	"github.com/arxdsilva/vquest/client/config"
	ui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const lineDistancingFactor = 1.15
const loginLineDistancingFactor = 1.4
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
	authenticateErr           error // display this later
	Authenticated             bool
}

func New(x, y, fontSize int32, color color.RGBA, cfg config.Display) Login {
	return Login{
		x:        x,
		y:        y,
		fontSize: fontSize,
		color:    color,
		userText: text.New(
			"username:", x, y, fontSize, color),
		userPasswordText: text.New(
			"password:",
			x,
			int32(float32(y)*lineDistancingFactor),
			fontSize, color),
		userNameInput: input.New(
			float32(x)*columnDistancingFactor,
			float32(y),
			float32(cfg.Width)*0.2,
			float32(fontSize),
			false,
			true,
			rl.LightGray,
		),
		userPasswordInput: input.New(
			float32(x)*columnDistancingFactor,
			float32(y)*lineDistancingFactor,
			float32(cfg.Width)*0.2,
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
	clicked := ui.Button(rl.Rectangle{
		X: float32(l.x),
		Y: float32(l.y) * loginLineDistancingFactor},
		"login")
	// fmt.Println("clicked: ", clicked)
	if clicked {
		l.send(l.userText.Text, l.userPasswordText.Text)
	}
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

func (l *Login) send(username, password string) {
	// todo

	// if err == nil
	l.Authenticated = true
}
