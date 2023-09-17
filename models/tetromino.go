package models

import (
	"image/color"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Tetromino struct {
	shape [][]bool
	color color.Color
	x, y  int
	board *Board
}

type Block struct {
	pos       []int
	tetromino Tetromino
}

func NewTetromino(b *Board) Tetromino {
	shapes := [][][]bool{
		{{true, true, true, true}},
		{{true, true}, {true, true}},
		{{true, true, true}, {false, true, false}},
		{{true, true, true}, {true, false, false}},
		{{true, true, true}, {false, false, true}},
	}

	s := shapes[rand.Intn(len(shapes))]

	return Tetromino{
		shape: s,
		color: color.RGBA{R: uint8(rand.Uint32()), G: uint8(rand.Uint32()), B: uint8(rand.Uint32()), A: 255},
		x:     b.width/2 - 2,
		y:     0,
		board: b,
	}
}

func NewBlock(t Tetromino, position []int) *Block {
	return (&Block{
		pos:       position,
		tetromino: t,
	})
}

func (t Tetromino) DrawTetromino() *canvas.Rectangle {
	// for y := 0; y < t.y; y++ {
	// 	for x := 0; x < t.x; x++ {
	// 		if t.shape[x][y] {

	// 		}
	// 	}
	// }
	pos := []int{0, 0}
	block := NewBlock(t, pos)
	return block.DrawBlock()
}

func (b Block) DrawBlock() *canvas.Rectangle {
	square := canvas.NewRectangle(b.tetromino.color)
	square.SetMinSize(fyne.NewSquareSize(float32(b.tetromino.board.cell)))
	square.Move(fyne.NewPos(1, 2))
	return square
}
