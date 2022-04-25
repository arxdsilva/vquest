package input

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Component struct {
	rectangle rl.Rectangle
}

func New(x, y, width, height float32) Component {
	return Component{
		rectangle: rl.Rectangle{
			X:      x,
			Y:      y,
			Width:  width,
			Height: height,
		}}
}

func (r *Component) OnHover(color color.RGBA) {
	rInt := r.rectangle.ToInt32()
	rl.DrawRectangleLines(
		rInt.X, rInt.Y,
		rInt.Width, rInt.Height,
		color)
}

func (r *Component) Draw(color color.RGBA) {
	rl.DrawRectangleRec(r.rectangle, color)
	r.OnHover(rl.Red)
}
