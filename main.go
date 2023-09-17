package main

import (
	"gotetris/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

const (
	Width  = 10
	Height = 20
	Cell   = 30
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Tetris")
	board := models.NewBoard(Width, Height, Cell)
	tetromino := models.NewTetromino(board)
	render(myWindow, board, tetromino)
	myWindow.ShowAndRun()
}

func render(myWindow fyne.Window, board *models.Board, tetromino models.Tetromino) {
	container := container.New(layout.NewGridLayout(Width))
	container.Add(board.DrawBoard())
	container.Add(tetromino.DrawTetromino())
	myWindow.SetContent(container)
}
