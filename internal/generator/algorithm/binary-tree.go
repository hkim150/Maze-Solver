package algorithm

import (
	"math/rand"
	"maze-solver/internal/maze"
	"time"
)

func BinaryTree(width, height int, animate bool) (*maze.Maze, error) {
	m, err := initialMaze(width, height)
	if err != nil {
		return m, err
	}

	var delay time.Duration
	if animate {
		delay = 40 * time.Millisecond
		m.PrintForAnimation(delay)
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
				if animate {
					m.Cells[r][c+1] = maze.Highlight
					m.PrintForAnimation(delay)
				}
				m.Cells[r][c+1] = maze.Empty
			} else if down {
				if animate {
					m.Cells[r+1][c] = maze.Highlight
					m.PrintForAnimation(delay)
				}
				m.Cells[r+1][c] = maze.Empty
			}
		}
	}

	return m, nil
}
