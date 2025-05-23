package maze

import (
	"fmt"
	"time"
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

type Pos [2]int

type Maze struct {
	Width    int
	Height   int
	Cells    [][]CellType
	StartPos Pos
	EndPos   Pos
}

func NewMaze(width, height int) *Maze {
	cells := make([][]CellType, height)
	for row := range cells {
		cells[row] = make([]CellType, width)
	}

	return &Maze{
		Width:    width,
		Height:   height,
		Cells:    cells,
		StartPos: Pos{1, 1},
		EndPos:   Pos{height - 2, width - 2},
	}
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
			case Visited:
				fmt.Print("  ")
			case Visiting:
				fmt.Print(Blue + "██" + Reset)
			default:
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

// ClearScreen clears the terminal screen using ANSI escape codes
func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

// PrintForAnimation clears the screen and prints the current state of the maze
// with a delay for animation effect
func (m *Maze) PrintForAnimation(delay time.Duration) {
	ClearScreen()
	m.Print()
	time.Sleep(delay)
}
