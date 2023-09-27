package models

import (
	"fmt"

	"fyne.io/fyne/v2/widget"
)

type DataGame struct {
	scoreLabel *widget.Label
	score      int
	level      string
	lines      int
	gameOver   bool
}

func NewDataGame() *DataGame {
	sl := widget.NewLabel(fmt.Sprintf("%d", 0))
	return &DataGame{
		score:      0,
		level:      "infinite",
		lines:      Rows,
		scoreLabel: sl,
		gameOver:   false,
	}
}

func (d *DataGame) GetScore() int {
	return d.score
}

func (d *DataGame) GetScoreLabel() *widget.Label {
	return d.scoreLabel
}

func (d *DataGame) GetLevel() string {
	return d.level
}
func (d *DataGame) GetLines() int {
	return d.lines
}

func (d *DataGame) UpdateScore(row int) {
	switch row {
	case 1:
		d.score = d.score + 100
	case 2:
		d.score = d.score + 300
	case 3:
		d.score = d.score + 500
	case 4:
		d.score = d.score + 800
	default:
	}
	d.scoreLabel.SetText(fmt.Sprintf("%d", d.score))
}
