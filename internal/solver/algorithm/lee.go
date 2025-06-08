package algorithm

import (
	"fmt"
	"math"
	"maze-solver/internal/maze"
	"time"
)

func Lee(m *maze.Maze, animate bool) error {
	// copy the maze for setting the dist value
	distMaze := make([][]int, m.Height)
	for i := range distMaze {
		distMaze[i] = make([]int, m.Width)
		for j := range distMaze[i] {
			// initialize the distance to each cell as infinity
			distMaze[i][j] = math.MaxInt
		}
	}

	pos := m.StartPos
	q := []maze.Pos{pos}
	distMaze[pos[0]][pos[1]] = 0
	m.Cells[pos[0]][pos[1]] = maze.Visited

	delay := 20 * time.Millisecond

	for len(q) > 0 {
		pos = q[0]
		q = q[1:]

		if animate {
			m.PrintForAnimation(delay)
		}

		if pos == m.EndPos {
			break
		}

		for _, neighbor := range neighbors(m, pos) {
			if m.Cells[neighbor[0]][neighbor[1]] == maze.Empty {
				distMaze[neighbor[0]][neighbor[1]] = distMaze[pos[0]][pos[1]] + 1
				m.Cells[neighbor[0]][neighbor[1]] = maze.Visited
				q = append(q, neighbor)
			}
		}
	}

	if pos != m.EndPos {
		return fmt.Errorf("Could not find a solution for the maze")
	}

	for {
		m.Cells[pos[0]][pos[1]] = maze.Highlight

		if animate {
			m.PrintForAnimation(delay)
		}

		if pos == m.StartPos {
			break
		}

		minDist := math.MaxInt
		nextPos := pos

		for _, neighbor := range neighbors(m, pos) {
			if distMaze[neighbor[0]][neighbor[1]] < minDist {
				minDist = distMaze[neighbor[0]][neighbor[1]]
				nextPos = neighbor
			}
		}

		pos = nextPos
	}

	return nil
}
