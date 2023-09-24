package controllers

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type Window struct {
	app    fyne.App
	window fyne.Window
}

func NewWindows(title string) *Window {
	appTetris := app.New()
	appWindow := appTetris.NewWindow(title)
	return &Window{
		app:    appTetris,
		window: appWindow,
	}
}
