package main

import (
	"fmt"
	"strings"

	"github.com/arxdsilva/vquest/client/display"
	"github.com/gen2brain/raylib-go/raylib"
)

var selectedField int
var userMouse, passMouse bool
var userLetters, passLetters []string
var userRect, passRect, loginRect rl.Rectangle

func loadLoginObjects(d display.Display) {
	rectX := float32(d.Width) / 2
	rectWidth := float32(d.Width) * 0.1
	rectHeight := float32(d.Height) * 0.05
	tH1 := float32(d.Height) * 0.4
	tH2 := float32(d.Height) * 0.5
	userRect = rl.Rectangle{
		X:      rectX,
		Y:      tH1,
		Width:  rectWidth,
		Height: rectHeight,
	}
	passRect = rl.Rectangle{
		X:      rectX,
		Y:      tH2,
		Width:  rectWidth,
		Height: rectHeight,
	}
	loginRect = rl.Rectangle{
		X:      rectX - rectWidth/2,
		Y:      float32(d.Height) * 0.7,
		Width:  rectWidth,
		Height: rectHeight,
	}
}

func drawLoginScreen(d display.Display) {
	userMouse = mouseCollision(userRect)
	passMouse = mouseCollision(passRect)
	if userMouse {
		selectedField = 0
	}
	if passMouse {
		selectedField = 1
	}
	key := rl.GetKeyPressed()
	if (rl.IsKeyPressed(rl.KeyTab)) && (selectedField == 0) {
		selectedField = 1
	} else if (rl.IsKeyPressed(rl.KeyTab)) && (selectedField == 1) {
		selectedField = 0
	}
	if (key >= 32) && (key <= 125) {
		if (selectedField == 0) && (len(userLetters) < 10) {
			userLetters = append(userLetters, fmt.Sprintf("%c", key))
		}
		if (selectedField == 1) && (len(passLetters) < 10) {
			passLetters = append(passLetters, fmt.Sprintf("%c", key))
		}
	}
	if rl.IsKeyPressed(rl.KeyBackspace) {
		if (selectedField == 0) && (len(userLetters) > 0) {
			userLetters = userLetters[:len(userLetters)-1]
		}
		if (selectedField == 1) && (len(passLetters) > 0) {
			passLetters = passLetters[:len(passLetters)-1]
		}
	}
	rl.DrawRectangleRec(userRect, rl.LightGray)
	rl.DrawRectangleRec(passRect, rl.LightGray)
	rl.DrawRectangleRec(loginRect, rl.DarkGray)
	drawOutlineOnCollision(userRect, passRect)
	lInt := userRect.ToInt32()
	rl.DrawText(strings.Join(userLetters, ""), lInt.X+5, lInt.Y+(lInt.Height/2), 20, rl.Maroon)
	pInt := passRect.ToInt32()
	rl.DrawText(strings.Repeat("-", len(passLetters)), pInt.X+5, pInt.Y+(pInt.Height/2), 20, rl.Maroon)
	rl.ClearBackground(rl.RayWhite)
	wdt := float32(d.Width) * 0.4
	hthUser := float32(d.Height) * 0.4
	hthPass := float32(d.Height) * 0.5
	drawLoginClick(loginRect)
	rl.DrawText("User:", int32(wdt), int32(hthUser), 20, rl.LightGray)
	rl.DrawText("Password:", int32(wdt), int32(hthPass), 20, rl.LightGray)
}

func drawOutlineOnCollision(r, l rl.Rectangle) {
	rInt := r.ToInt32()
	lInt := l.ToInt32()
	if selectedField == 0 {
		rl.DrawRectangleLines(rInt.X, rInt.Y, rInt.Width, rInt.Height, rl.Red)
		rl.DrawRectangleLines(lInt.X, lInt.Y, lInt.Width, lInt.Height, rl.DarkGray)
		return
	}
	rl.DrawRectangleLines(rInt.X, rInt.Y, rInt.Width, rInt.Height, rl.DarkGray)
	rl.DrawRectangleLines(lInt.X, lInt.Y, lInt.Width, lInt.Height, rl.Red)
}

func drawLoginClick(r rl.Rectangle) {
	lRInt := r.ToInt32()
	loginX := lRInt.X + int32(r.Width/3.3)
	loginY := lRInt.Y + int32(r.Height/4)
	if (mouseCollision(r)) && (rl.IsMouseButtonDown(rl.MouseLeftButton)) {
		rl.DrawText("Login", loginX, loginY, 20, rl.White)
		return
	}
	rl.DrawText("Login", loginX, loginY, 20, rl.LightGray)
}

// if (((framesCounter/20)%2) == 0) {
// 	DrawText("_", textBox.x + 8 + MeasureText(name, 40), textBox.y + 12, 40, MAROON);
// }
