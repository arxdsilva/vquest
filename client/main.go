package main

import (
	"github.com/arxdsilva/vquest/client/config"
	// "github.com/arxdsilva/vquest/client/display"
	// "github.com/arxdsilva/vquest/client/msg"
	"github.com/arxdsilva/vquest/client/scenes"
	// rl "github.com/gen2brain/raylib-go/raylib"
)

var clientState int

const loginScreen = 0
const connected = 1

var resp string

func main() {
	cfg := config.NewDefault()
	scenes.Start(cfg)
}

// func main2() {
// 	d := display.New()
// 	rl.InitWindow(int32(d.Width), int32(d.Height), "vQuest")
// 	rl.SetTargetFPS(60)
// 	loadLoginObjects(d)
// 	c := &msg.Client{}
// 	for !rl.WindowShouldClose() {
// 		rl.BeginDrawing()
// 		rl.ClearBackground(rl.RayWhite)
// 		if clientState == 0 {
// 			drawLoginScreen(d)
// 		} else {
// 			if !c.Active {
// 				c.Connect()
// 			}
// 			c.Communicate(userName)
// 			resp = c.Response
// 			drawConnectedScreen(d)
// 		}
// 		rl.EndDrawing()
// 	}
// 	rl.CloseWindow()
// }

// func mouseCollision(rec rl.Rectangle) bool {
// 	if rl.CheckCollisionPointRec(rl.GetMousePosition(), rec) {
// 		return true
// 	}
// 	return false
// }
