package algorithm

import (
	"fmt"
	"maze-solver/internal/maze"
	"time"
)

const (
	left = iota
	up
	right
	down
)

func HandOnWall(m *maze.Maze, animate bool) error {
	curr := m.StartPos
	dirs := [][2]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}} // left-hand rule directions (left, up, right, down)
	
	var delay time.Duration
	if animate {
		delay = 20 * time.Millisecond
		m.Cells[curr[0]][curr[1]] = maze.Highlight
		m.PrintForAnimation(delay)
	}

	// find the initial direction to start with
	d := left
	for i, dir := range dirs {
		row := curr[0] + dir[0]
		col := curr[1] + dir[1]
		if row >= m.StartPos[0] && row <= m.EndPos[0] && col >= m.StartPos[1] && col <= m.EndPos[1] && m.Cells[row][col] != maze.Wall {
			curr = maze.Pos{row, col}
			d = i
			break
		}
	}

	// map to keep track of junctions visit count for highlighting
	junction := make(map[maze.Pos]int)
	
	for curr != m.EndPos {
		if animate {
			if junction[curr] == 1 {
				m.Cells[curr[0]][curr[1]] = maze.Highlight
			} else if m.Cells[curr[0]][curr[1]] == maze.Empty {
				m.Cells[curr[0]][curr[1]] = maze.Highlight
			} else if m.Cells[curr[0]][curr[1]] == maze.Highlight {
				m.Cells[curr[0]][curr[1]] = maze.Empty
			}
			m.PrintForAnimation(delay)
		}

		next := curr
		// turn left of the current direction
		d = (d + len(dirs) - 1) % len(dirs)
		nextD := d
		

		// check if the neighboring cell is empty, turn right if not
		neighborCount := 0
		for range len(dirs) {
			nRow := curr[0] + dirs[d][0]
			nCol := curr[1] + dirs[d][1]

			if nRow >= m.StartPos[0] && nRow <= m.EndPos[0] && nCol >= m.StartPos[1] && nCol <= m.EndPos[1] && m.Cells[nRow][nCol] != maze.Wall {
				if neighborCount == 0 {
					next = maze.Pos{nRow, nCol}
					nextD = d
				}
				neighborCount++
			}

			// turn right
			d = (d + 1) % len(dirs)
		}

		if neighborCount == 0 {
			return fmt.Errorf("Could not find a solution for the maze, stuck at position %v", curr)
		} else if neighborCount == 1 {
			// deadend, erase the path
			m.Cells[curr[0]][curr[1]] = maze.Empty
		} else if neighborCount >= 3{
			junction[curr]++
		}

		curr = next
		d = nextD
	}

	m.Cells[curr[0]][curr[1]] = maze.Highlight

	return nil
}
