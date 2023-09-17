package models

import (
	"fmt"
	"image/color"
	"math/rand"
)

type Tetromino struct {
	shape [][]bool
	color color.Color
	x, y  int
	board *Board
}

func NewTetromino(b *Board) Tetromino {
	s := randomShape()
	fmt.Print("NEW")
	return Tetromino{
		shape: s,
		color: color.RGBA{R: uint8(rand.Uint32()), G: uint8(rand.Uint32()), B: uint8(rand.Uint32()), A: 255},
		board: b,
		x:     0,
		y:     0,
	}
}

func randomShape() [][]bool {
	shapes := [][][]bool{
		{
			{true, true, true, true},
		},
		{
			{true, true},
			{true, true},
		},
		{
			{true, true, true},
			{false, true, false},
		},

		{
			{true, true, true},
			{true, false, false}},

		{
			{true, true, true},
			{false, false, true},
		},
	}

	return shapes[rand.Intn(len(shapes))]

}

func (t *Tetromino) DrawTetromino() {
	// fmt.Printf("Array: %v \n", len(t.board.blocks))
	for row := range t.shape {
		for col := range t.shape[row] {
			if t.shape[row][col] {
				t.board.blocks[t.y+row][t.x+col].FillColor = t.color
				t.board.blocks[t.y+row][t.x+col].Refresh()
				fmt.Printf("row: %d  colm: %d \n", t.y+row, t.x+col)
				// fmt.Printf("position: %v \n", t.board.blocks[t.y+row][t.x+col].Position())
			}
		}
	}
}

func (t *Tetromino) EraseTetromino() {
	for row := range t.shape {
		for col := range t.shape[row] {
			if t.shape[row][col] {
				t.board.blocks[t.y+row][t.x+col].FillColor = Gray
				t.board.blocks[t.y+row][t.x+col].Refresh()
			}
		}
	}
}

func (t *Tetromino) MoveLeft() {
	t.EraseTetromino()
	t.x--
	if t.x <= -1 {
		t.x = 0
	}
	t.DrawTetromino()
}

func (t *Tetromino) MoveRight() {
	t.EraseTetromino()
	t.x++
	fmt.Printf("x despues de +1: %v \n", t.x)
	if t.x >= Columns {
		t.x = Columns - 1
	}
	fmt.Printf("x ya checado contra la rabia, %v\n", t.x)
	t.DrawTetromino()
}

func (t *Tetromino) MoveDown() {
	t.EraseTetromino()
	t.y++
	t.DrawTetromino()
}

func (t *Tetromino) CheckColision() bool {
	for row := range t.shape {
		if t.shape[row][0] {
			if t.x <= -1 {
				return false
			}
		} else if t.shape[row][len(t.shape[row])-1] {
			return false
		}
	}
	return true
}

func (t *Tetromino) CanMoveDown() bool {
	for col := range t.shape[len(t.shape)-1] {
		if t.shape[len(t.shape)-1][col] {
			// fmt.Printf("position: (%d,%d), x: %d, y:%d, value1: %v, value2: %v \n", len(t.shape), col, t.x, t.y, t.y+len(t.shape) >= Rows, t.board.blocks[t.y+len(t.shape)][t.x+col].FillColor != Gray)
			// fmt.Printf("y:%d, ROWS: %v, value1: %v, result: %v\n", t.y, Rows, t.y+len(t.shape), t.y+len(t.shape) >= Rows)
			if t.y+len(t.shape) >= Rows || t.board.blocks[t.y+len(t.shape)][t.x+col].FillColor != Gray {
				return false
			}

		}
	}
	return true
}
