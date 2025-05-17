package internal

import "fmt"

// baseMaze generates a maze with grid like structure
// where all cells are walls except for the odd row and odd column cells are empty
// Because the outer edges are walls, the width and height must be at least 5
// Also the width and height will be round down to the nearest odd number
func baseMaze(width, height int) (Maze, error) {
	if width < 5 || height < 5 {
		return Maze{}, fmt.Errorf("width and height must be at least 5")
	}

	if width%2 == 0 {
		width--
	}
	if height%2 == 0 {
		height--
	}

	cells := make([][]CellType, height)
	for i := range cells {
		cells[i] = make([]CellType, width)
		for j := range cells[i] {
			if i%2 == 1 && j%2 == 1 {
				cells[i][j] = Empty
			} else {
				cells[i][j] = Wall
			}
		}
	}

	maze := Maze{
		Width:  width,
		Height: height,
		Cells:  cells,
		StartCell: [2]int{1, 1},
		EndCell:   [2]int{height - 2, width - 2},
	}

	return maze, nil
}
