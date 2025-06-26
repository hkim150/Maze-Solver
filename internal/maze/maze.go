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

	// For animation
	Highlight
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
	Width      int
	Height     int
	Cells      [][]CellType
	PrevCells  [][]CellType // Track previous state for incremental updates
	StartPos   Pos
	EndPos     Pos
	CursorHome bool // Track if cursor is at home position
}

func NewMaze(width, height int) *Maze {
	cells := make([][]CellType, height)
	prevCells := make([][]CellType, height)
	for row := range cells {
		cells[row] = make([]CellType, width)
		prevCells[row] = make([]CellType, width)
	}

	return &Maze{
		Width:      width,
		Height:     height,
		Cells:      cells,
		PrevCells:  prevCells,
		StartPos:   Pos{1, 1},
		EndPos:     Pos{height - 2, width - 2},
		CursorHome: false,
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
				fmt.Print(Yellow + "██" + Reset)
			case Visited:
				fmt.Print(Gray + "██" + Reset)
			case Visiting:
				fmt.Print(Blue + "██" + Reset)
			case Highlight:
				fmt.Print(Red + "██" + Reset)
			default:
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

// reset up erases non-wall cells as empty and sets the start and end cells
func (m *Maze) Reset() {
	for r := 1; r < m.Height-1; r++ {
		for c := 1; c < m.Width-1; c++ {
			if m.Cells[r][c] != Wall {
				m.Cells[r][c] = Empty
			}
		}
	}

	m.Cells[m.StartPos[0]][m.StartPos[1]] = Start
	m.Cells[m.EndPos[0]][m.EndPos[1]] = End
	
	// Reset animation state to ensure proper display update
	m.CursorHome = false
}

// clean up erases all visiting and visited cells as empty
func (m *Maze) CleanUp() {
	for r := 1; r < m.Height-1; r++ {
		for c := 1; c < m.Width-1; c++ {
			if m.Cells[r][c] == Visiting || m.Cells[r][c] == Visited {
				m.Cells[r][c] = Empty
			}
		}
	}
}

// ClearScreen clears the terminal screen using ANSI escape codes
func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

// HideCursor hides the terminal cursor
func HideCursor() {
	fmt.Print("\033[?25l")
}

// ShowCursor shows the terminal cursor
func ShowCursor() {
	fmt.Print("\033[?25h")
}

// MoveCursor moves the cursor to the specified row and column (1-indexed)
func MoveCursor(row, col int) {
	fmt.Printf("\033[%d;%dH", row+1, col*2+1)
}

// PrintCell prints a single cell at its current cursor position
func (m *Maze) PrintCell(cellType CellType) {
	switch cellType {
	case Wall:
		fmt.Print("██")
	case Start:
		fmt.Print(Green + "██" + Reset)
	case End:
		fmt.Print(Yellow + "██" + Reset)
	case Visited:
		fmt.Print(Gray + "██" + Reset)
	case Visiting:
		fmt.Print(Blue + "██" + Reset)
	case Highlight:
		fmt.Print(Red + "██" + Reset)
	default:
		fmt.Print("  ")
	}
}

// UpdateChangedCells updates only the cells that have changed since the last frame
func (m *Maze) UpdateChangedCells() {
	for i := 0; i < m.Height; i++ {
		for j := 0; j < m.Width; j++ {
			if m.Cells[i][j] != m.PrevCells[i][j] {
				MoveCursor(i, j)
				m.PrintCell(m.Cells[i][j])
				m.PrevCells[i][j] = m.Cells[i][j]
			}
		}
	}
	// Move cursor to bottom of maze to avoid interfering with display
	MoveCursor(m.Height, 0)
}

// InitializeAnimation sets up the initial display for animation
func (m *Maze) InitializeAnimation() {
	ClearScreen()
	HideCursor()
	m.Print()
	// Copy current state to previous state
	for i := 0; i < m.Height; i++ {
		copy(m.PrevCells[i], m.Cells[i])
	}
	m.CursorHome = true
}

// PrintForAnimation efficiently updates only changed cells for animation
func (m *Maze) PrintForAnimation(delay time.Duration) {
	if !m.CursorHome {
		// First time - do full screen setup
		m.InitializeAnimation()
	} else {
		// Subsequent frames - only update changed cells
		m.UpdateChangedCells()
	}
	time.Sleep(delay)
}
