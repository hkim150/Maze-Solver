package algorithm

import (
	"fmt"
	"math/rand"
	"maze-solver/internal/maze"
	"time"
)

func RandomMouse(m *maze.Maze, animate bool) error {
	var delay time.Duration
	if animate {
		delay = 7 * time.Millisecond
	}

	visited := make(map[maze.Pos]bool)
	deadEnds := make(map[maze.Pos]bool)
	junctions := make(map[maze.Pos]bool)

	prev := m.StartPos
	curr := m.StartPos
	for curr != m.EndPos {
		if animate {
			m.Cells[curr[0]][curr[1]] = maze.Highlight
			m.PrintForAnimation(delay)
			m.Cells[curr[0]][curr[1]] = maze.Empty
		}

		n := neighbors(m, curr)
		var next maze.Pos
		if len(n) == 0 {
			return fmt.Errorf("No neighbors found for position %v, cannot continue.\n", curr)
		} else if len(n) == 1 {
			next = n[0]
			deadEnds[curr] = true
		} else if len(n) >= 1 {
			if len(n) >= 3 {
				junctions[curr] = true
			}

			i := rand.Intn(len(n))
			if n[i] == prev {
				i = (i + 1) % len(n)
			}
			next = n[i]
		}

		// toggle visited state so that reverse path is not highlighted and only the solution path is highlighted
		visited[next] = !visited[next]
		prev, curr = curr, next
	}

	if animate {
		m.Cells[curr[0]][curr[1]] = maze.Highlight
		m.PrintForAnimation(delay)
	}

	m.Reset()

	// the deadends are always toggled true when they cannot be part of the solution path
	for cell, ok := range visited {
		if ok && !deadEnds[cell] {
			m.Cells[cell[0]][cell[1]] = maze.Highlight
		}
	}

	m.Cells[m.StartPos[0]][m.StartPos[1]] = maze.Highlight

	// junctions of the solution path may not be toggled on. Highlight junction if neighbor is highlighted
	for cell := range junctions {
		m.Cells[cell[0]][cell[1]] = maze.Empty
		for _, n := range neighbors(m, cell) {
			if m.Cells[n[0]][n[1]] == maze.Highlight {
				m.Cells[cell[0]][cell[1]] = maze.Highlight
				break
			}
		}
	}

	return nil
}
