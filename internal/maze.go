package internal

import "fmt"

type CellType int

const (
	Empty CellType = iota
	Wall
	Start
	End

	// For pathfinding
	Visited
)

type Maze struct {
	Width  int
	Height int
	Cells  [][]CellType
	StartCell [2]int
	EndCell   [2]int
}

func (m *Maze) Print() {
	for i := 0; i < m.Height; i++ {
		for j := 0; j < m.Width; j++ {
			switch m.Cells[i][j] {
			case Wall:
				fmt.Print("#")
			default:
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
