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
	blinkFontSize  = 18
)

type Component struct {
	active       bool
	password     bool
	color        color.RGBA
	rectangle    rl.Rectangle
	text         text.Component
	userInput    []string
	frameCounter int
}

func New(x, y, width, height float32, password, active bool, color color.RGBA) Component {
	return Component{
		active:   active,
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
	// draw red lines on edge
	rInt := c.rectangle.ToInt32()
	rl.DrawRectangleLines(
		rInt.X, rInt.Y,
		rInt.Width, rInt.Height,
		rl.Red)

	// get input from user (add)
	key := rl.GetKeyPressed()
	if (key >= 32) && (key <= 125) {
		c.userInput = append(c.userInput, fmt.Sprintf("%c", key))
	}
	// get input from user (remove - backspace)
	// len to avoid panic
	if rl.IsKeyPressed(rl.KeyBackspace) && len(c.userInput) > 0 {
		c.userInput = c.userInput[:len(c.userInput)-1]
	}

	// display input as needed
	displayTxt := strings.Join(c.userInput, "")
	if c.password {
		displayTxt = strings.Repeat("*", len(displayTxt))
	}
	c.text.Text = displayTxt

	// blink "|" to give txt wait feedback sensation to user
	if ((c.frameCounter / 20) % 3) == 0 {
		rl.DrawText("|",
			c.rectangle.ToInt32().X+8+rl.MeasureText(displayTxt, fontSize),
			c.rectangle.ToInt32().Y, blinkFontSize, rl.Black)
	}

	c.frameCounter++
}

func (c *Component) Draw() {
	rl.DrawRectangleRec(c.rectangle, c.color)
	c.text.Draw()
	if c.active {
		// if c.onHover() || c.Active {
		c.highlight()
		return
	}
}

func (c *Component) ToggleActive() {
	c.active = !c.active
}
