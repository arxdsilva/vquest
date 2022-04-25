package config

type Config struct {
	GameName  string
	Display   Display
	TargetFPS int32
}

func NewDefault() Config {
	return Config{
		GameName:  "vQuest",
		TargetFPS: 60,
		Display: Display{
			Height: 720,
			Width:  1280,
		},
	}
}
