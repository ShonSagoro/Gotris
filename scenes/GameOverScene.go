package scenes

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type GameOverScene struct {
	window fyne.Window
}

func NewGameOverScene(window fyne.Window) *GameOverScene {
	gameOverScene := &GameOverScene{window: window}
	gameOverScene.Render()
	return gameOverScene
}

func (o *GameOverScene) Render() {
	o.DrawGameOver()
}

func (o *GameOverScene) DrawGameOver() {
	gameOverLabel := widget.NewLabelWithStyle("GAME OVER", fyne.TextAlignCenter, fyne.TextStyle{})
	retryButton := widget.NewButton("Retry", o.Retry)
	goBackMenuButton := widget.NewButton("Go Back Menu", o.Menu)

	container_center := container.NewCenter(
		container.NewHBox(
			container.NewVBox(
				gameOverLabel,
				container.NewHBox(
					container.NewCenter(retryButton),
					container.NewCenter(goBackMenuButton),
				),
			),
		),
	)

	o.window.Resize(fyne.NewSize(300, 400))
	o.window.SetContent(container_center)
}

func (o *GameOverScene) Retry() {
	NewGameScene(o.window)
}
func (o *GameOverScene) Menu() {
	NewMainScene(o.window)
}
