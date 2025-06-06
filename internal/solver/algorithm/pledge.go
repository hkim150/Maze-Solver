package algorithm

import (
	"maze-solver/internal/maze"
	"time"
)

func Pledge(m *maze.Maze, animate bool) error {
	curr := m.StartPos

	// right, down, left, up
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	
	// initially facing right
	d := 0

	// angle counter to keep track of the number of turns
	// it is used to determine when to stop following the wall
	angleCounter := 0

	// map to keep track of junctions visit count for highlighting
	junction := make(map[maze.Pos]int)

	var delay time.Duration
	if animate {
		delay = 30 * time.Millisecond
		m.Cells[m.StartPos[0]][m.StartPos[1]] = maze.Highlight
		m.PrintForAnimation(delay)
		m.Cells[m.StartPos[0]][m.StartPos[1]] = maze.Visited
	}

outer:
	for {
		// walk straight until it hits a wall
		for {
			if curr == m.EndPos {
				break outer
			}

			dir := dirs[d%len(dirs)]
			next := maze.Pos{curr[0] + dir[0], curr[1] + dir[1]}
			if m.Cells[next[0]][next[1]] == maze.Wall {
				break
			}

			curr = next

			if animate {
				if junction[curr] == 1 {
					m.Cells[curr[0]][curr[1]] = maze.Highlight
				} else if m.Cells[curr[0]][curr[1]] == maze.Empty {
					m.Cells[curr[0]][curr[1]] = maze.Highlight
				} else if m.Cells[curr[0]][curr[1]] == maze.Highlight {
					m.Cells[curr[0]][curr[1]] = maze.Empty
				}
			}
		}

		// turn right
		d = (d + 1) % 4
		angleCounter++

		// follow the wall on the left until angle counter is 0
		// there are three scenarios on wall following:
		// 1. left cell is not a wall - turn left and go forward one cell
		// 2. left cell is a wall and front cell is empty - go forward one cell
		// 3. left cell is a wall and front cell is a wall - turn right
		for angleCounter != 0 {
			// for animation, highlight the current cell if moved
			moved := false

			if curr == m.EndPos {
				break outer
			}

			leftDir := (d + len(dirs) - 1) % len(dirs)
			leftCell := maze.Pos{curr[0] + dirs[leftDir][0], curr[1] + dirs[leftDir][1]}
			// case 1
			if m.Cells[leftCell[0]][leftCell[1]] != maze.Wall {
				d = leftDir
				angleCounter--
				curr = leftCell
				moved = true
			} else {
				frontCell := maze.Pos{curr[0] + dirs[d][0], curr[1] + dirs[d][1]}
				// case 2
				if m.Cells[frontCell[0]][frontCell[1]] != maze.Wall {
					curr = frontCell
					moved = true
				// case 3
				} else {
					d = (d + 1) % len(dirs)
					angleCounter++
				}
			}

			if animate && moved {
				if junction[curr] == 1 {
					m.Cells[curr[0]][curr[1]] = maze.Highlight
				} else if m.Cells[curr[0]][curr[1]] == maze.Empty {
					m.Cells[curr[0]][curr[1]] = maze.Highlight
				} else if m.Cells[curr[0]][curr[1]] == maze.Highlight {
					m.Cells[curr[0]][curr[1]] = maze.Empty
				}
			}
		}
	}

	return nil
}

func isJunction(m *maze.Maze, pos maze.Pos) bool {
	return len(neighbors(m, pos)) >= 3
}

func isDeadEnd(m *maze.Maze, pos maze.Pos) bool {
	return len(neighbors(m, pos)) == 1
}