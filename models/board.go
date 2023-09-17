package models

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

var Gray = color.RGBA{R: 120, G: 124, B: 126, A: 255}

type Board struct {
	width  int
	height int
	cell   int
}

func NewBoard(w int, h int, c int) *Board {
	return (&Board{
		width:  w,
		height: h,
		cell:   c,
	})
}

func (b *Board) DrawBoard() *fyne.Container {
	board := b.TetrisRows()
	return board
}

func (b *Board) TetrisRows() *fyne.Container {
	board := container.New(layout.NewGridLayout(b.width))

	for y := 0; y < b.height; y++ {
		for x := 0; x < b.width; x++ {
			square := TetrisRow(b.cell)
			board.Add(square)
			fmt.Printf("[%d,%d] = position: %v", x, y, square.Position())
		}
	}
	return board

}

func TetrisRow(size int) *canvas.Rectangle {
	square := canvas.NewRectangle(Gray)
	square.SetMinSize(fyne.NewSquareSize(float32(size)))
	return square
}
