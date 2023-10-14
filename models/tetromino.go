package models

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
)

type Tetromino struct {
	shape [][]bool
	color color.Color
	x, y  int
	board *Board
}

func NewTetromino(b *Board) *Tetromino {
	s := randomShape()
	rangR := rand.Intn(255-130) + 130
	rangG := rand.Intn(255-130) + 130
	rangB := rand.Intn(255-130) + 130

	tetromino := &Tetromino{
		shape: s,
		color: color.RGBA{R: uint8(rangR), G: uint8(rangG), B: uint8(rangB), A: 255},
		board: b,
		x:     0,
		y:     0,
	}

	tetromino.DrawTetromino()

	return tetromino
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
			{true, true, false},
			{false, true, true},
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
	for row := range t.shape {
		for col := range t.shape[row] {
			if t.shape[row][col] {
				t.board.blocks[t.y+row][t.x+col].FillColor = t.color
				t.board.blocks[t.y+row][t.x+col].Refresh()
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
	if !t.CheckColision() {
		t.x++
	}
	t.DrawTetromino()
}

func (t *Tetromino) MoveRight() {
	t.EraseTetromino()
	t.x++
	if !t.CheckColision() {
		t.x--
	}
	t.DrawTetromino()
}

func (t *Tetromino) MoveDown() {
	t.EraseTetromino()
	t.y++
	t.DrawTetromino()
}

func (t *Tetromino) CheckColision() bool {
	for row := range t.shape {
		for col := range t.shape[row] {
			if t.shape[row][col] {
				if t.x+col > Columns-1 || t.x+col < 0 || t.board.blocks_state[t.y+row][t.x+col] != 0 {
					return false
				}
			}
		}
	}
	return true
}

func (t *Tetromino) CanMoveDown() bool {
	for row := range t.shape {
		for col := range t.shape[row] {
			if t.shape[row][col] {
				if t.y+len(t.shape) >= Rows || t.board.blocks_state[t.y+row+1][t.x+col] != 0 {
					t.LockShape()
					return false
				}
			}
		}
	}
	return true
}
func (t *Tetromino) LockShape() {
	for row := range t.shape {
		for col := range t.shape[row] {
			if t.shape[row][col] {
				t.board.blocks_state[t.y+row][t.x+col] = 1
			}
		}
	}
}

func (t *Tetromino) RotateShape() {
	rows := len(t.shape)
	cols := len(t.shape[0])
	rotated := make([][]bool, cols)
	for i := range rotated {
		rotated[i] = make([]bool, rows)
	}
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			rotated[col][rows-1-row] = t.shape[row][col]
		}
	}
	if t.CanMoveDown() && t.CanRotate(rotated) {
		t.EraseTetromino()
		t.shape = rotated
		t.CheckWall()
		t.DrawTetromino()
	}
}

func (t *Tetromino) CanRotate(newShape [][]bool) bool {
	if t.x+len(newShape[0]) > Columns || t.y+len(newShape) > Rows {
		return false
	}
	for row := 0; row < len(newShape); row++ {
		for col := 0; col < len(newShape[0]); col++ {
			if newShape[row][col] && t.board.blocks_state[t.y+row][t.x+col] != 0 {
				return false
			}
		}
	}

	return true
}

func (t *Tetromino) CheckWall() {
	for row := range t.shape {
		for col := range t.shape[row] {
			if t.shape[row][col] {
				if t.x+col < 0 {
					t.x++
				} else if t.x+col > Columns-1 {
					t.x--
				}
			}
		}
	}
	t.CanMoveDown()

}

func (t *Tetromino) CheckTop() bool {
	if !t.CanMoveDown() && t.x == 0 {
		return true
	}
	return false
}

func (t *Tetromino) FallShape(window fyne.Window, quit chan int) {
	for {
		select {
		case <-quit:
			fmt.Println("Fall shape is close!")
			return
		default:
			if !t.board.stop {
				if t.CanMoveDown() {
					t.MoveDown()
				} else {
					t = NewTetromino(t.board)
					t.SetKeys(window)
				}
			}
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func (t *Tetromino) SetKeys(window fyne.Window) {
	window.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {
		if keyEvent.Name == fyne.KeyUp {
			t.RotateShape()
		} else if keyEvent.Name == fyne.KeyLeft {
			t.MoveLeft()
		} else if keyEvent.Name == fyne.KeyRight {
			t.MoveRight()
		} else if keyEvent.Name == fyne.KeyDown {
			if !t.board.stop {
				if t.CanMoveDown() {
					t.MoveDown()
				}
			}
		}
	})
}
