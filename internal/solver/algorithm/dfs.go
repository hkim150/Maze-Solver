package algorithm

import (
	"fmt"
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

	solved := false
	for len(stack) > 0 {
		pos := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if pos == m.EndPos {
			solved = true
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

	if !solved {
		return fmt.Errorf("Could not find a solution for the maze")
	}

	// Reconstruct the path
	m.Reset()
	for p := m.EndPos; p != m.StartPos; p = parent[p] {
		m.Cells[p[0]][p[1]] = maze.Highlight
	}
	m.Cells[m.StartPos[0]][m.StartPos[1]] = maze.Highlight

	return nil
}
