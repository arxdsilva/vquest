package forms

import (
	"image/color"

	"github.com/arxdsilva/vquest/client/components/text"
)

type Login struct {
	x, y     int32
	fontSize int32
	color    color.RGBA
	userText text.Component
}

func New(x, y, fontSize int32, color color.RGBA) Login {
	return Login{
		x:        x,
		y:        y,
		fontSize: fontSize,
		color:    color,
		userText: text.New("User:", x, y, fontSize, color),
	}
}

func (l *Login) Draw() {
	l.userText.Draw()
}
