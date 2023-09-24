package models

type DataGame struct {
	score int
	level string
	lines int
}

func NewDataGame() *DataGame {
	return &DataGame{
		score: 0,
		level: "infinite",
		lines: Rows,
	}
}
func (d *DataGame) GetScore() int {
	return d.score
}
func (d *DataGame) GetLevel() string {
	return d.level
}
func (d *DataGame) GetLines() int {
	return d.lines
}
