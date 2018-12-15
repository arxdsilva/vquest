package display

type Display struct {
	Height         int
	Width          int
	CustomSettings []Display
}

func New() Display {
	settings := []Display{
		Display{
			Width:  1920,
			Height: 1080,
		},
		Display{
			Width:  1280,
			Height: 720,
		},
		Display{
			Width:  720,
			Height: 480,
		},
	}
	// width := rl.GetScreenWidth()
	// height := rl.GetScreenHeight()
	// fmt.Println(height, width, "<<<<<<<<<<<<<<")
	// var drawH, drawW int
	// for _, s := range settings {
	// 	if height > s.Height {
	// 		drawH = s.Height
	// 		drawW = s.Width
	// 		break
	// 	}
	// }
	return Display{
		Height:         720,
		Width:          1280,
		CustomSettings: settings,
	}
}
