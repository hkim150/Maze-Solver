package algorithm

import (
	"fmt"
	"maze-solver/internal/maze"
	"time"
)

const delay = 30 * time.Millisecond

func Recursive(m *maze.Maze, animate bool) error {
	if !recurseFunc(m, m.StartPos, animate) {
		return fmt.Errorf("Could not find a solution for the maze")
	}

	if animate {
		m.CleanUp()
		m.PrintForAnimation(delay)
	}

	return nil
}

func recurseFunc(m *maze.Maze, pos maze.Pos, animate bool) bool {
	if pos == m.EndPos {
		m.Cells[pos[0]][pos[1]] = maze.Highlight
		if animate {
			m.PrintForAnimation(delay)
		}
		return true
	}

	if m.Cells[pos[0]][pos[1]] != maze.Empty {
		return false
	}

	m.Cells[pos[0]][pos[1]] = maze.Visited
	if animate {
		m.PrintForAnimation(delay)
	}

	if pos[0] < m.EndPos[0] && recurseFunc(m, maze.Pos{pos[0] + 1, pos[1]}, animate) {
		m.Cells[pos[0]][pos[1]] = maze.Highlight
		if animate {
			m.PrintForAnimation(delay)
		}
		return true
	}

	if pos[1] < m.EndPos[1] && recurseFunc(m, maze.Pos{pos[0], pos[1] + 1}, animate) {
		m.Cells[pos[0]][pos[1]] = maze.Highlight
		if animate {
			m.PrintForAnimation(delay)
		}
		return true
	}

	if pos[0] > m.StartPos[0] && recurseFunc(m, maze.Pos{pos[0] - 1, pos[1]}, animate) {
		m.Cells[pos[0]][pos[1]] = maze.Highlight
		if animate {
			m.PrintForAnimation(delay)
		}
		return true
	}

	if pos[1] > m.StartPos[1] && recurseFunc(m, maze.Pos{pos[0], pos[1] - 1}, animate) {
		m.Cells[pos[0]][pos[1]] = maze.Highlight
		if animate {
			m.PrintForAnimation(delay)
		}
		return true
	}

	return false
}
