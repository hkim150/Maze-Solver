package generator

import (
	"math/rand"
	dataStructure "maze-solver/internal/data_structure"
)

func RandomizedKruskal(width, height int) (Maze, error) {
	maze, err := baseMaze(width, height)
	if err != nil {
		return maze, err
	}

	// initialize a list to store all walls in the maze.
    // walls are represented as [row, col] coordinates.
	walls := make([][2]int, 0)
	for row := 1; row < maze.Height-1; row++ {
		for col := row%2 + 1; col < maze.Width-1; col += 2 {
			walls = append(walls, [2]int{row, col})
		}
	}

	uf := dataStructure.NewUnionFind[[2]int]()

	// shuffle the walls randomly to ensure a random maze generation
	rand.Shuffle(len(walls), func(i, j int) {
		walls[i], walls[j] = walls[j], walls[i]
	})

	for _, wall := range walls {
		row, col := wall[0], wall[1]

		// two cells separated by the current wall.
		var cell_1, cell_2 [2]int
		if row%2 == 1 { // Horizontal wall
			cell_1 = [2]int{row, col-1} // Cell to the left of the wall
			cell_2 = [2]int{row, col+1} // Cell to the right of the wall
		} else { // Vertical wall
			cell_1 = [2]int{row-1, col} // Cell above the wall
			cell_2 = [2]int{row+1, col} // Cell below the wall
		}

		// If the two cells are not already connected, remove the wall and connect them.
		if !uf.IsConnected(cell_1, cell_2) {
			maze.Cells[row][col] = Empty
			uf.Union(cell_1, cell_2)
		}
	}

	return maze, nil
}
