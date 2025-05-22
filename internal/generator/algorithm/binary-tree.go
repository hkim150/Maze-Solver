package algorithm

import (
	"math/rand"
	"maze-solver/internal/maze"
)

func BinaryTree(width, height int) (*maze.Maze, error) {
	m, err := initialMaze(width, height)
	if err != nil {
		return m, err
	}

	// for each cell, if can merge to right and down, choose randomly
	// if only right or down, merge the only way
	// if neither, do nothing
	for r := 1; r < m.Height-1; r += 2 {
		for c := 1; c < m.Width-1; c += 2 {
			right := c+2 < m.Width-1
			down := r+2 < m.Height-1
			if right && down {
				right = rand.Intn(2) == 0
			}

			if right {
				m.Cells[r][c+1] = maze.Visited
			} else if down {
				m.Cells[r+1][c] = maze.Visited
			}
		}
	}

	return m, nil
}
