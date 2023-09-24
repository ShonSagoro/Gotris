package controllers

import (
	"fmt"
	"gotetris/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func InitGame() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Tetris")
	board := models.NewBoard()
	data := models.NewDataGame()
	NewGame(myWindow, board, data)
	// myWindow.Resize(fyne.NewSize(400, 600))
	myWindow.CenterOnScreen()
	myWindow.ShowAndRun()
}

func NewGame(myWindow fyne.Window, board *models.Board, data *models.DataGame) {
	container := container.New(layout.NewHBoxLayout())
	container.Add(board.DrawBoard(myWindow))
	container.Add(DrawDataGame(data))
	myWindow.SetContent(container)
}

func DrawDataGame(data *models.DataGame) *fyne.Container {
	ScoreLabel := widget.NewLabel(fmt.Sprintf("Score: %d", data.GetScore()))
	ScoreLabel.Move(fyne.NewPos(0, 0))

	LevelLabel := widget.NewLabel(fmt.Sprintf("Level: %s", data.GetLevel()))
	LevelLabel.Move(fyne.NewPos(0, 50))

	LinesLabel := widget.NewLabel(fmt.Sprintf("Lines: %d", data.GetLines()))
	LinesLabel.Move(fyne.NewPos(0, 100))

	return container.NewWithoutLayout(ScoreLabel, LevelLabel, LinesLabel)
}
