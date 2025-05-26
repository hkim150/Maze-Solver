package algorithm

import (
	"maze-solver/internal/maze"
	"time"
)

func DFS(m *maze.Maze, animate bool) error {
	stack := []maze.Pos{m.StartPos}

	var delay time.Duration
	if animate {
		delay = 40 * time.Millisecond
		m.PrintForAnimation(delay)
	}

	for len(stack) > 0 {
		pos := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if animate {
			m.Cells[pos[0]][pos[1]] = maze.Highlight
			m.PrintForAnimation(delay)
		}
		m.Cells[pos[0]][pos[1]] = maze.Visited

		for _, nPos := range neighbors(m, pos) {
			if nPos == m.EndPos {
				return nil
			}

			if m.Cells[nPos[0]][nPos[1]] == maze.Empty {
				stack = append(stack, nPos)
			}
		}
	}

	m.PrintForAnimation(delay)

	return nil
}
