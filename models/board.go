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

const (
	Rows      = 6
	Columns   = 10
	Blocksize = 30
)

type Board struct {
	rows      int
	colums    int
	blocksize int
	blocks    [Rows][Columns]*canvas.Rectangle
}

func NewBoard() *Board {
	return (&Board{
		rows:      Rows,
		colums:    Columns,
		blocksize: Blocksize,
	})
}

func (b *Board) DrawBoard(myWindow fyne.Window) *fyne.Container {
	board := b.TetrisRows()
	current_tetromino := NewTetromino(b)
	current_tetromino.DrawTetromino()
	// go func() {
	// 	for range time.Tick(time.Second) {
	// 		if !current_tetromino.CanMoveDown() {
	// 			current_tetromino = NewTetromino(b)
	// 			current_tetromino.DrawTetromino()
	// 		} else {
	// 			current_tetromino.MoveDown()
	// 			myWindow.Canvas().Refresh(board)
	// 		}
	// 	}
	// }()

	myWindow.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {
		if keyEvent.Name == fyne.KeyUp {
			fmt.Println("rotating :D")
		} else if keyEvent.Name == fyne.KeyLeft {
			current_tetromino.MoveLeft()
		} else if keyEvent.Name == fyne.KeyRight {
			current_tetromino.MoveRight()
		} else if keyEvent.Name == fyne.KeyDown {
			current_tetromino.MoveDown()
		}
	})
	return board
}

func (b *Board) TetrisRows() *fyne.Container {
	board := container.New(layout.NewGridLayout(b.colums))

	for row := range b.blocks {
		for col := range b.blocks[row] {
			square := TetrisRow(b.blocksize)
			board.Add(square)
			b.blocks[row][col] = square
			// fmt.Printf("row: %d  colm: %d element: %v\n", row, col, square)
		}
	}

	return board

}

func TetrisRow(size int) *canvas.Rectangle {
	square := canvas.NewRectangle(Gray)
	square.SetMinSize(fyne.NewSquareSize(float32(size)))
	return square
}
