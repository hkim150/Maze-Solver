package generator

import (
	"fmt"
	"math/rand"
)

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
		Width:     width,
		Height:    height,
		Cells:     cells,
		StartCell: [2]int{1, 1},
		EndCell:   [2]int{height - 2, width - 2},
	}

	return maze, nil
}

// randomDirections returns a random order of directions in the maze (up, left, down, right)
func randomDirections() [][2]int {
	directions := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	rand.Shuffle(len(directions), func(i, j int) {
		directions[i], directions[j] = directions[j], directions[i]
	})
	return directions
}
