package input

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/arxdsilva/vquest/client/components/text"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	textMarginLeft = 1.01
	fontSize       = 20
)

type Component struct {
	password  bool
	color     color.RGBA
	rectangle rl.Rectangle
	text      text.Component
	userInput []string
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
		},
		text: text.New(
			"", int32(x*textMarginLeft), int32(y),
			fontSize, rl.Black),
	}
}

func (c *Component) onHover() bool {
	return rl.CheckCollisionPointRec(
		rl.GetMousePosition(), c.rectangle)
}

func (c *Component) highlight() {
	rInt := c.rectangle.ToInt32()
	rl.DrawRectangleLines(
		rInt.X, rInt.Y,
		rInt.Width, rInt.Height,
		rl.Red)

	// get input from user (add)
	key := rl.GetKeyPressed()
	if (key >= 32) && (key <= 125) {
		c.userInput = append(c.userInput, fmt.Sprintf("%c", key))
		c.text.Text = strings.Join(c.userInput, "")
		return
	}
	// get input from user (remove - backspace)
	// len to avoid panic
	if rl.IsKeyPressed(rl.KeyBackspace) && len(c.userInput) > 0 {
		c.userInput = c.userInput[:len(c.userInput)-1]
		c.text.Text = strings.Join(c.userInput, "")
		return
	}
}

func (c *Component) Draw() {
	rl.DrawRectangleRec(c.rectangle, c.color)
	c.text.Draw()
	if c.onHover() {
		c.highlight()
		return
	}
}
