package internal

import "fmt"

type CellType int

const (
	Empty CellType = iota
	Wall
	Start
	End
	Visited
)

type Maze struct {
	Width  int
	Height int
	Cells  [][]CellType
}

func (m *Maze) Print() {
	for i := 0; i < m.Height; i++ {
		for j := 0; j < m.Width; j++ {
			switch m.Cells[i][j] {
			case Visited:
				fallthrough
			case Empty:
				fmt.Print(" ")
			case Wall:
				fmt.Print("#")
			case Start:
				fmt.Print("S")
			case End:
				fmt.Print("E")
			}
		}
		fmt.Println()
	}
}
