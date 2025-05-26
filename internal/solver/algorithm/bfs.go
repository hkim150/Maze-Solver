package algorithm

import (
	"maze-solver/internal/maze"
	"time"
)

func BFS(m *maze.Maze, animate bool) error {
	var delay time.Duration
	if animate {
		delay = 20 * time.Millisecond
		m.Cells[m.StartPos[0]][m.StartPos[1]] = maze.Highlight
		m.PrintForAnimation(delay)
	}

	m.Cells[m.StartPos[0]][m.StartPos[1]] = maze.Visited
	queue := []maze.Pos{m.StartPos}
	parent := make(map[maze.Pos]maze.Pos)


outer:
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]

		for _, nPos := range neighbors(m, pos) {
			if m.Cells[nPos[0]][nPos[1]] == maze.Empty {
				if animate {
					m.Cells[nPos[0]][nPos[1]] = maze.Highlight
					m.PrintForAnimation(delay)
				}

				m.Cells[nPos[0]][nPos[1]] = maze.Visited
				queue = append(queue, nPos)
				parent[nPos] = pos

				if nPos == m.EndPos {
					break outer
				}
			}
		}
	}

	m.CleanUp()
	pos := m.EndPos
	m.Cells[pos[0]][pos[1]] = maze.Highlight
	for pos != m.StartPos {
		pos = parent[pos]
		m.Cells[pos[0]][pos[1]] = maze.Highlight
	}

	return nil
}
