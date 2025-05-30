package algorithm

import (
	"fmt"
	"math/rand"
	"maze-solver/internal/maze"
	"time"
)

func DeadEndFilling(m *maze.Maze, animate bool) error {
	var delay time.Duration
	if animate {
		delay = 30 * time.Millisecond
	}

	// identify initial dead ends
	candidates := make([]maze.Pos, 0)
	for row := m.StartPos[0]; row <= m.EndPos[0]; row++ {
		for col := m.StartPos[1]; col <= m.EndPos[1]; col++ {
			pos := maze.Pos{row, col}
			// skip the start and end positions so the solution is not filled
			if pos == m.StartPos || pos == m.EndPos {
				continue
			}

			if ok, _ := checkDeadEnd(m, pos); ok {
				candidates = append(candidates, pos)
			}
		}
	}

	// shuffle the candidates to randomize the order of filling dead ends
	rand.Shuffle(len(candidates), func(i, j int) {
		candidates[i], candidates[j] = candidates[j], candidates[i]
	})

	// fill the dead ends until there are no dead ends left except the start and end positions
	for len(candidates) > 0 {
		candidate := candidates[len(candidates)-1]
		candidates = candidates[:len(candidates)-1]

		// skip the start and end positions so the solution is not filled
		if candidate == m.StartPos || candidate == m.EndPos {
			continue
		}

		if ok, emptyPos := checkDeadEnd(m, candidate); ok {
			if animate {
				m.Cells[candidate[0]][candidate[1]] = maze.Highlight
				m.PrintForAnimation(delay)
			}
			m.Cells[candidate[0]][candidate[1]] = maze.Visited
			candidates = append(candidates, emptyPos)
		}
	}

	// there should be no dead ends left, so we can walk from start to end
	curr := m.StartPos
	for {
		ok, next := checkDeadEnd(m, curr)
		if !ok {
			return fmt.Errorf("Could not find a solution for the maze")
		}
		
		m.Cells[curr[0]][curr[1]] = maze.Highlight
		curr = next
		
		if curr == m.EndPos {
			m.Cells[curr[0]][curr[1]] = maze.Highlight
			break
		}

		if animate {
			m.PrintForAnimation(delay)
		}
	}

	// clean up the maze by removing the visited cells
	for row := m.StartPos[0]; row <= m.EndPos[0]; row++ {
		for col := m.StartPos[1]; col <= m.EndPos[1]; col++ {
			if m.Cells[row][col] == maze.Visited {
				m.Cells[row][col] = maze.Empty
			}
		}
	}

	return nil
}

// checkDeadEnd checks if the given position is a dead end in the maze and returns a neighboring empty position if it is a dead end.
func checkDeadEnd(m *maze.Maze, pos maze.Pos) (bool, maze.Pos) {
	if m.Cells[pos[0]][pos[1]] != maze.Empty {
		return false, maze.Pos{}
	}

	emptyCount := 0
	var emptyPos maze.Pos
	for _, nPos := range neighbors(m, pos) {
		if m.Cells[nPos[0]][nPos[1]] == maze.Empty {
			emptyCount++
			if emptyCount > 1 {
				break
			}
			emptyPos = nPos
		}
	}

	if emptyCount == 1 {
		return true, emptyPos
	}

	return false, maze.Pos{}
}
