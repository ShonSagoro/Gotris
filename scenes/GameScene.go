package scenes

import (
	"fmt"
	"gotetris/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var t *models.Tetromino
var b *models.Board
var d *models.DataGame

type GameScene struct {
	window fyne.Window
}

func NewGameScene(window fyne.Window) *GameScene {
	gameScene := &GameScene{window: window}
	gameScene.Render()
	gameScene.StartGame()
	return gameScene
}

func (g *GameScene) Render() {
	d = models.NewDataGame()
	b = models.NewBoard(d)
	t = models.NewTetromino(b)
	DrawSceneGame(g, b, d)
}

func (g *GameScene) StartGame() {
	go t.FallShape(g.window)
	go b.CheckAndClear()
	t.SetKeys(g.window)
}

func DrawSceneGame(g *GameScene, board *models.Board, data *models.DataGame) {
	container := container.New(layout.NewHBoxLayout())
	container.Add(board.DrawBoard())
	container.Add(g.DrawDataGame(data))
	g.window.SetContent(container)
}

func (a *GameScene) DrawDataGame(data *models.DataGame) *fyne.Container {
	scoreLabel := widget.NewLabel("Score:")
	scoreContainer := container.NewHBox(scoreLabel, data.GetScoreLabel())

	levelLabel := widget.NewLabel(fmt.Sprintf("Level: %s", data.GetLevel()))

	linesLabel := widget.NewLabel(fmt.Sprintf("Lines: %d", data.GetLines()))

	goBackMenu := widget.NewButton("Abrir Menú", func() {
		dialog.ShowConfirm("Salir", "¿Desea salir de la aplicación?", func(response bool) {
			if response {
				a.BackToMenu()
			}
		}, a.window)
	})

	return container.NewVBox(
		scoreContainer,
		levelLabel,
		linesLabel,
		layout.NewSpacer(),
		layout.NewSpacer(),
		layout.NewSpacer(),
		goBackMenu,
		layout.NewSpacer(),
	)
}

func (a *GameScene) BackToMenu() {
	NewMainScene(a.window)
}
