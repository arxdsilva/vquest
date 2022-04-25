package input

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Component struct {
	active    bool
	password  bool
	color     color.RGBA
	rectangle rl.Rectangle
}

func New(x, y, width, height float32, password bool, color color.RGBA) Component {
	return Component{
		password: password,
		color:    color,
		rectangle: rl.Rectangle{
			X:      x,
			Y:      y,
			Width:  width,
			Height: height,
		}}
}

func (c *Component) onHover() bool {
	return rl.CheckCollisionPointRec(
		rl.GetMousePosition(), c.rectangle)
}

func (c *Component) highlight() {
	c.active = true
	rInt := c.rectangle.ToInt32()
	rl.DrawRectangleLines(
		rInt.X, rInt.Y,
		rInt.Width, rInt.Height,
		rl.Red)
}

func (c *Component) Draw() {
	rl.DrawRectangleRec(c.rectangle, c.color)
	if c.onHover() {
		c.highlight()
		return
	}
	c.active = false
}
