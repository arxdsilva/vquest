package text

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Component struct {
	x, y     int32
	text     string
	fontSize int32
	color    color.RGBA
}

func New(text string, x, y, fontSize int32, color color.RGBA) Component {
	return Component{
		text:     text,
		x:        x,
		y:        y,
		fontSize: fontSize,
		color:    color,
	}
}

func (c *Component) Draw() {
	rl.DrawText(c.text, c.x, c.y, c.fontSize, c.color)
}

func (c *Component) OnHover() {

}
