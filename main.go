package main

import (
	"gotetris/scenes"

	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()
	window := app.NewWindow("Tetris")
	window.CenterOnScreen()
	scenes.NewMainScene(window)
	window.ShowAndRun()
}
