package algorithm

import (
	"maze-solver/internal/maze"
	"time"
)

func DFS(m *maze.Maze, animate bool) error {
	stack := []maze.Pos{m.StartPos}
	parent := make(map[maze.Pos]maze.Pos)

	var delay time.Duration
	if animate {
		delay = 30 * time.Millisecond
		m.PrintForAnimation(delay)
	}

	for len(stack) > 0 {
		pos := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if pos == m.EndPos {
			break
		}

		if animate {
			m.Cells[pos[0]][pos[1]] = maze.Highlight
			m.PrintForAnimation(delay)
		}
		m.Cells[pos[0]][pos[1]] = maze.Visited

		for _, nPos := range neighbors(m, pos) {
			if m.Cells[nPos[0]][nPos[1]] == maze.Empty {
				parent[nPos] = pos
				stack = append(stack, nPos)
			}
		}
	}
	m.PrintForAnimation(delay)

	m.CleanUp()
	pos := m.EndPos
	m.Cells[pos[0]][pos[1]] = maze.Highlight
	for pos != m.StartPos {
		pos, _ = parent[pos]
		m.Cells[pos[0]][pos[1]] = maze.Highlight
	}

	return nil
}
