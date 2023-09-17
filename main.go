package main

import (
	"gotetris/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/layout"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Tetris")
	board := models.NewBoard()
	render(myWindow, board)

	myWindow.Resize(fyne.NewSize(400, 600))
	myWindow.ShowAndRun()
}

func render(myWindow fyne.Window, board *models.Board) {
	container := container.New(layout.NewHBoxLayout())
	container.Add(board.DrawBoard(myWindow))
	myWindow.SetContent(container)
}
