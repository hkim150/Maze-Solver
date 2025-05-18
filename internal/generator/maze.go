package generator

import (
	"fmt"
)

type CellType int

const (
	Empty CellType = iota
	Wall
	Start
	End

	// For pathfinding
	Visited
	Visiting
)

// Define ANSI color codes
const (
	Reset  = "\033[0m" // Reset to default color
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	White  = "\033[37m"
	Gray   = "\033[90m"
)

type Maze struct {
	Width     int
	Height    int
	Cells     [][]CellType
	StartCell [2]int
	EndCell   [2]int
}

func (m *Maze) Print() {
	for i := 0; i < m.Height; i++ {
		for j := 0; j < m.Width; j++ {
			switch m.Cells[i][j] {
			case Wall:
				fmt.Print("██")
			case Start:
				fmt.Print(Green + "██" + Reset)
			case End:
				fmt.Print(Red + "██" + Reset)
			default:
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}
