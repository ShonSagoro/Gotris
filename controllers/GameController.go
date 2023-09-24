package controllers

import (
	"fmt"
	"gotetris/models"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type App struct {
	app    fyne.App
	window fyne.Window
}

func NewApp(title string) *App {
	appTetris := app.New()
	appWindow := appTetris.NewWindow(title)
	return &App{
		app:    appTetris,
		window: appWindow,
	}
}

func InitApp() {
	app := NewApp("Tetris")
	app.DrawMenu()
	app.window.CenterOnScreen()
	app.window.ShowAndRun()
}

func (a *App) DrawMenu() {
	titleImg := canvas.NewImageFromFile("resources/logo.png")
	titleImg.Resize(fyne.NewSize(300, 75))
	titleImg.Move(fyne.NewPos(50, 10))
	titleContainer := container.NewWithoutLayout(titleImg)

	start := widget.NewButton("Start", a.StartGame)

	credits := widget.NewButton("Credits", a.Credits)

	exit := widget.NewButton("Exit", a.ExitGame)

	container_center := container.NewVBox(
		titleContainer,
		layout.NewSpacer(),
		start,
		credits,
		exit,
		layout.NewSpacer(),
	)

	a.window.SetContent(container_center)
	a.window.Resize(fyne.NewSize(400, 500))
	a.window.SetFixedSize(true)
}

func (a *App) StartGame() {
	board := models.NewBoard()
	data := models.NewDataGame()
	DrawGame(a, board, data)
}

func (a *App) ExitGame() {
	a.window.Close()
}

func (a *App) Credits() {
	shonPhoto := canvas.NewImageFromFile("resources/shon_credits.png")
	shonPhoto.Resize(fyne.NewSize(300, 250))
	shonPhoto.Move(fyne.NewPos(0, 10))
	shonPhotoContainer := container.NewWithoutLayout(shonPhoto)
	goBackMenu := widget.NewButton("Menu", a.DrawMenu)

	shonCard := widget.NewCard("Jonathan G. Shon sagoro", "Programmer", shonPhotoContainer)

	githubURL := "https://github.com/ShonSagoro"

	parsedURL, err := url.Parse(githubURL)
	if err != nil {
		panic(err)
	}

	github_widget := widget.NewHyperlink("Checa mi Github", parsedURL)

	container_center := container.NewHBox(
		container.NewCenter(github_widget),
		container.NewVBox(
			shonCard,
			layout.NewSpacer(),
			goBackMenu,
		),
	)

	a.window.SetContent(container_center)
	a.window.Resize(fyne.NewSize(300, 400))
}

func DrawGame(app *App, board *models.Board, data *models.DataGame) {
	container := container.New(layout.NewHBoxLayout())
	container.Add(board.DrawBoard(app.window))
	container.Add(app.DrawDataGame(data))
	app.window.SetContent(container)
}

func (a *App) DrawDataGame(data *models.DataGame) *fyne.Container {
	scoreLabel := widget.NewLabel(fmt.Sprintf("Score: %d", data.GetScore()))

	levelLabel := widget.NewLabel(fmt.Sprintf("Level: %s", data.GetLevel()))

	linesLabel := widget.NewLabel(fmt.Sprintf("Lines: %d", data.GetLines()))

	goBackMenu := widget.NewButton("Menu", a.DrawMenu)

	return container.NewVBox(
		scoreLabel,
		levelLabel,
		linesLabel,
		layout.NewSpacer(),
		layout.NewSpacer(),
		layout.NewSpacer(),
		goBackMenu,
		layout.NewSpacer(),
	)
}
