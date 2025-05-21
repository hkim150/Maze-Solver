package generator

import "math/rand"

func BinaryTreeAlgorithm(width, height int) (Maze, error) {
	maze, err := baseMaze(width, height)
	if err != nil {
		return maze, err
	}

	// for each cell, if can merge to right and down, choose randomly
	// if only right or down, merge the only way
	// if neither, do nothing
	for r := 1; r < maze.Height-1; r += 2 {
		for c := 1; c < maze.Width-1; c += 2 {
			right := c + 2 < maze.Width-1
			down := r + 2 < maze.Height-1
			if right && down {
				right = rand.Intn(2) == 0
			}

			if right {
				maze.Cells[r][c+1] = Visited
			} else if down {
				maze.Cells[r+1][c] = Visited
			}
		}
	}

	return maze, nil
}
