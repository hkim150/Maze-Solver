package algorithm

import (
	"math/rand"
	dataStructure "maze-solver/internal/data_structure"
	"maze-solver/internal/maze"
	"time"
)

func Kruskal(width, height int, animate bool) (*maze.Maze, error) {
	m, err := initialMaze(width, height)
	if err != nil {
		return m, err
	}

	// initialize a list to store all walls in the maze.
	// walls are represented as [row, col] coordinates.
	walls := make([]maze.Pos, 0)
	for row := 1; row < m.Height-1; row++ {
		for col := row%2 + 1; col < m.Width-1; col += 2 {
			walls = append(walls, maze.Pos{row, col})
		}
	}

	uf := dataStructure.NewUnionFind[maze.Pos]()

	// shuffle the walls randomly to ensure a random maze generation
	rand.Shuffle(len(walls), func(i, j int) {
		walls[i], walls[j] = walls[j], walls[i]
	})

	var delay time.Duration
	if animate {
		delay = 40 * time.Millisecond
		m.PrintForAnimation(delay)
	}

	for _, wall := range walls {
		row, col := wall[0], wall[1]

		// two cells separated by the current wall.
		var cell_1, cell_2 maze.Pos
		if row%2 == 1 { // Horizontal wall
			cell_1 = maze.Pos{row, col - 1} // Cell to the left of the wall
			cell_2 = maze.Pos{row, col + 1} // Cell to the right of the wall
		} else { // Vertical wall
			cell_1 = maze.Pos{row - 1, col} // Cell above the wall
			cell_2 = maze.Pos{row + 1, col} // Cell below the wall
		}

		// If the two cells are not already connected, remove the wall and connect them.
		if !uf.IsConnected(cell_1, cell_2) {
			if animate {
				m.Cells[row][col] = maze.Visiting
				m.PrintForAnimation(delay)
			}

			m.Cells[row][col] = maze.Empty
			uf.Union(cell_1, cell_2)
		}
	}

	return m, nil
}
