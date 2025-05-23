package algorithm

import (
	"math/rand"
	"maze-solver/internal/maze"
	"time"
)

func AldousBroder(width, height int, animate bool) (*maze.Maze, error) {
	m, err := initialMaze(width, height)
	if err != nil {
		return m, err
	}

	// notVisited is a map to keep track of the unvisited cells for end condition
	notVisited := make(map[maze.Pos]bool)
	for row := 1; row < m.Height-1; row += 2 {
		for col := 1; col < m.Width-1; col += 2 {
			notVisited[maze.Pos{row, col}] = true
		}
	}

	// Randomly select a starting cell from the isolated empty cells.
	r := rand.Intn((m.Height-1)/2)*2 + 1
	c := rand.Intn((m.Width-1)/2)*2 + 1
	m.Cells[r][c] = maze.Visited
	delete(notVisited, maze.Pos{r, c})

	var delay time.Duration
	if animate {
		delay = 40 * time.Millisecond
		m.PrintForAnimation(delay)
	}

	for len(notVisited) > 0 {
		for _, dir := range randomDirections() {
			neighRow := r + dir[0]*2
			neighCol := c + dir[1]*2
			if neighRow >= 1 && neighRow < m.Height-1 && neighCol >= 1 && neighCol < m.Width-1 {
				if m.Cells[neighRow][neighCol] != maze.Visited {
					// remove the wall between the current cell and the neighbor
					m.Cells[r+dir[0]][c+dir[1]] = maze.Visited

					// mark the next cell as visiting to highlight in the animation
					if animate {
						m.Cells[neighRow][neighCol] = maze.Visiting
						m.PrintForAnimation(delay)
					}

					m.Cells[neighRow][neighCol] = maze.Visited
					delete(notVisited, maze.Pos{neighRow, neighCol})
				}

				r, c = neighRow, neighCol
				break
			}
		}
	}

	return m, nil
}
