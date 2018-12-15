package main

import (
	"fmt"
	"strings"

	"github.com/arxdsilva/vquest/client/display"
	"github.com/gen2brain/raylib-go/raylib"
)

func main() {
	d := display.New()
	rl.InitWindow(int32(d.Width), int32(d.Height), "vQuest")
	rl.SetTargetFPS(60)
	rectX := float32(d.Width) / 2
	rectWidth := float32(d.Width) * 0.1
	rectHeight := float32(d.Height) * 0.05
	tH1 := float32(d.Height) * 0.4
	tH2 := float32(d.Height) * 0.5
	loginRect := rl.Rectangle{
		X:      rectX,
		Y:      tH1,
		Width:  rectWidth,
		Height: rectHeight,
	}
	passRect := rl.Rectangle{
		X:      rectX,
		Y:      tH2,
		Width:  rectWidth,
		Height: rectHeight,
	}
	var loginMouse, passMouse bool
	var loginLetters, passLetters []string
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		loginMouse = mouseCollision(loginRect)
		passMouse = mouseCollision(passRect)
		if loginMouse || passMouse {
			key := rl.GetKeyPressed()
			if (key >= 32) && (key <= 125) {
				if loginMouse && (len(loginLetters) < 10) {
					loginLetters = append(loginLetters, fmt.Sprintf("%c", key))
				}
				if passMouse && (len(passLetters) < 10) {
					passLetters = append(passLetters, fmt.Sprintf("%c", key))
				}
			}
			if rl.IsKeyPressed(rl.KeyBackspace) {
				if (loginMouse) && (len(loginLetters) > 0) {
					loginLetters = loginLetters[:len(loginLetters)-1]
				}
				if (passMouse) && (len(passLetters) > 0) {
					passLetters = passLetters[:len(passLetters)-1]
				}
			}
		}
		rl.DrawRectangleRec(loginRect, rl.LightGray)
		rl.DrawRectangleRec(passRect, rl.LightGray)
		drawOutlineOnCollision(loginRect)
		drawOutlineOnCollision(passRect)
		lInt := loginRect.ToInt32()
		rl.DrawText(strings.Join(loginLetters, ""), lInt.X+5, lInt.Y+(lInt.Height/2), 20, rl.Maroon)
		pInt := passRect.ToInt32()
		rl.DrawText(strings.Repeat("-", len(passLetters)), pInt.X+5, pInt.Y+(pInt.Height/2), 20, rl.Maroon)
		rl.ClearBackground(rl.RayWhite)
		wdt := float32(d.Width) * 0.3
		rl.DrawText("Login:", int32(wdt), int32(tH1), 20, rl.LightGray)
		rl.DrawText("Password:", int32(wdt), int32(tH2), 20, rl.LightGray)
		rl.EndDrawing()
	}
	rl.CloseWindow()
}

func mouseCollision(rec rl.Rectangle) bool {
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), rec) {
		return true
	}
	return false
}

func drawOutlineOnCollision(r rl.Rectangle) {
	rInt := r.ToInt32()
	if mouseCollision(r) {
		rl.DrawRectangleLines(rInt.X, rInt.Y, rInt.Width, rInt.Height, rl.Red)
		return
	}
	rl.DrawRectangleLines(rInt.X, rInt.Y, rInt.Width, rInt.Height, rl.DarkGray)
}
