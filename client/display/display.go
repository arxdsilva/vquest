package display

import "github.com/gen2brain/raylib-go/raylib"


type Display struct {
	Height         int32
	Width          int32
	CustomSettings []Display
}

func New() Display {
	settings := []Display{
		Display{
			Width: 1920,
			Height: 1080,
		},
		Display{
			Width: 1280,
			Height: 720,
		},
		Display{
			Width: 720,
			Height: 480,
		},
	}
	width := rl.GetScreenWidth()
	height := rl.GetScreenHeight()
	var drawH, drawW int32
	for _, s := range settings {
		if height > s.Height {
			drawH = s.Height
			drawW = s.Width
			break
		}
	}
	return &Display{
		Height: drawH,
		Width: drawW,
		CustomSettings: settings
	}
}
