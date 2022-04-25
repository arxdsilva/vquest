package main

import (
	"github.com/arxdsilva/vquest/client/config"
	"github.com/arxdsilva/vquest/client/scenes"
)

func main() {
	cfg := config.NewDefault()
	manager := scenes.NewManager(cfg)
	scenes.Start(manager)
}
