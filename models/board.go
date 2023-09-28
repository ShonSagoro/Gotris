package models

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

var Gray = color.RGBA{R: 30, G: 30, B: 30, A: 255}

const (
	Rows    = 20
	Columns = 10
)

type Board struct {
	rows         int
	colums       int
	blocks       [Rows][Columns]*canvas.Rectangle
	blocks_state [Rows][Columns]int
	dataGame     *DataGame
	stop         bool
}

func NewBoard(data *DataGame) *Board {
	board := &Board{
		rows:     Rows,
		colums:   Columns,
		dataGame: data,
		stop:     false,
	}
	board.MakeBlocks()
	return board

}

func (b *Board) MakeBlocks() {
	for row := range b.blocks {
		for col := range b.blocks[row] {
			square := makeSquare()
			b.blocks[row][col] = square
			b.blocks_state[row][col] = 0
		}
	}
}
func (b *Board) SetStop(val bool) {
	b.stop = val
}
func (b *Board) GetStop() bool {
	return b.stop
}

func (b *Board) DrawBoard() *fyne.Container {
	board := b.TetrisRows()
	return board
}

func (b *Board) TetrisRows() *fyne.Container {
	board := container.New(layout.NewGridLayout(b.colums))
	for row := range b.blocks {
		for col := range b.blocks[row] {
			board.Add(b.blocks[row][col])
		}
	}
	return board
}

func makeSquare() *canvas.Rectangle {
	square := canvas.NewRectangle(Gray)
	square.SetMinSize(fyne.NewSquareSize(float32(30)))
	return square
}

func (b *Board) CheckBoard() int {
	rows := 0
	for row := range b.blocks_state {
		if b.CheckFullRow(row) {
			b.ClearRow(row)
			rows++
		}
	}
	return rows
}

func (b *Board) CheckFullRow(row int) bool {
	for col := range b.blocks_state[row] {
		if b.blocks_state[row][col] != 1 {
			return false
		} else if b.blocks_state[row][col] == 1 && b.blocks[row][col].FillColor == Gray {
			b.blocks_state[row][col] = 0
			b.blocks[row][col].FillColor = Gray
		}
	}
	return true
}

func (b *Board) ClearRow(row int) {
	for col := range b.blocks_state[row] {
		b.blocks_state[row][col] = 0
		b.blocks[row][col].FillColor = Gray
	}
}

func (b *Board) DownPiecesOnCascade() {
	b.stop = true
	for {
		moved := false

		for row := Rows - 1; row >= 0; row-- {
			for col := 0; col < Columns; col++ {
				if b.blocks_state[row][col] == 1 {
					if row+1 < Rows && b.blocks_state[row+1][col] == 0 {
						b.blocks_state[row+1][col] = 1
						b.blocks[row+1][col].FillColor = b.blocks[row][col].FillColor
						b.blocks_state[row][col] = 0
						b.blocks[row][col].FillColor = Gray
						moved = true
					}
				}
			}
		}

		if !moved {
			break
		}
	}
	b.stop = false
}
func (b *Board) CheckAndClear() {
	for {
		if !b.stop {
			rows := b.CheckBoard()
			if rows != 0 {
				b.dataGame.UpdateScore(rows)
				b.DownPiecesOnCascade()
			}
		}
	}
}

func (b *Board) CheckGameOver() bool {
	for col := range b.blocks_state[0] {
		if b.blocks_state[0][col] == 1 {
			return true
		}
	}
	return false
}
func (b *Board) CheckFalseColition() {
	for {
		if !b.stop {
			for row := range b.blocks_state {
				for col := range b.blocks_state[row] {
					if b.blocks_state[row][col] == 1 && b.blocks[row][col].FillColor == Gray {
						b.blocks_state[row][col] = 0
						b.blocks[row][col].FillColor = Gray
					}
				}
			}
		}
	}

}
