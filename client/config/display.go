package config

type Display struct {
	Height int
	Width  int
}

func NewDisplay() Display {
	return Display{
		Width:  1280,
		Height: 720,
	}
}
