package scenes

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type MainScene struct {
	window fyne.Window
}

func NewMainScene(window fyne.Window) *MainScene {
	MainScene := &MainScene{
		window: window,
	}
	MainScene.InitApp()
	return MainScene
}

func (m *MainScene) InitApp() {
	m.DrawSceneMenu()
}

func (m *MainScene) DrawSceneMenu() {
	titleImg := canvas.NewImageFromFile("resources/logo.png")
	titleImg.Resize(fyne.NewSize(300, 75))
	titleImg.Move(fyne.NewPos(50, 10))
	titleContainer := container.NewWithoutLayout(titleImg)

	start := widget.NewButton("Start", m.StartGame)

	credits := widget.NewButton("Credits", m.DrawCredits)

	exit := widget.NewButton("Exit", m.ExitGame)

	container_center := container.NewVBox(
		titleContainer,
		layout.NewSpacer(),
		start,
		credits,
		exit,
		layout.NewSpacer(),
	)

	m.window.SetContent(container_center)
	m.window.Resize(fyne.NewSize(400, 500))
	m.window.SetFixedSize(true)
}

func (m *MainScene) ExitGame() {
	m.window.Close()
}

func (m *MainScene) DrawCredits() {
	shonPhoto := canvas.NewImageFromFile("resources/shon_credits.png")
	shonPhoto.Resize(fyne.NewSize(300, 250))
	shonPhoto.Move(fyne.NewPos(0, 10))
	shonPhotoContainer := container.NewWithoutLayout(shonPhoto)
	goBackMenu := widget.NewButton("Menu", m.DrawSceneMenu)

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

	m.window.Resize(fyne.NewSize(300, 400))
	m.window.SetContent(container_center)
}

func (m *MainScene) StartGame() {
	NewGameScene(m.window)
}
