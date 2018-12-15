package main

import (
	"fmt"

	"github.com/arxdsilva/vquest/client/display"
	"github.com/gen2brain/raylib-go/raylib"
)

func main() {
	d := display.New()
	rl.InitWindow(int32(d.Width), int32(d.Height), "vQuest")
	rl.SetTargetFPS(60)
	recColor := rl.Color{R: 255, G: 50, B: 0, A: 172}
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
	// var loginLetters, passLetters []string
	loginLetters := [10]string{}
	var loginLettersCount, passLettersCount int
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), loginRect) {
			loginMouse = true
		} else {
			loginMouse = false
		}
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), passRect) {
			passMouse = true
		} else {
			passMouse = false
		}
		if loginMouse || passMouse {
			key := rl.GetKeyPressed()
			if (key >= 32) && (key <= 125) {
				if loginMouse && (loginLettersCount < 10) {
					loginLetters[loginLettersCount] = fmt.Sprintf("%c", key)
					loginLettersCount++
					fmt.Println(loginLetters, loginLettersCount)
				}
				if passMouse && (passLettersCount < 10) {
					// loginLetters
					fmt.Println(passLettersCount)
				}
			}
		}
		rl.ClearBackground(rl.RayWhite)
		wdt := float32(d.Width) * 0.3
		rl.DrawRectangleRec(loginRect, recColor)
		rl.DrawRectangleRec(passRect, recColor)
		rl.DrawText("Login:", int32(wdt), int32(tH1), 20, rl.LightGray)
		rl.DrawText("Password:", int32(wdt), int32(tH2), 20, rl.LightGray)
		rl.EndDrawing()
	}
	rl.CloseWindow()
}
