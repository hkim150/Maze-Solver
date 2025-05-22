package algorithm

import "maze-solver/internal/maze"

func HuntAndKill(width, height int) (*maze.Maze, error) {
	m, err := initialMaze(width, height)
	if err != nil {
		return m, err
	}

	m.Cells[1][1] = maze.Empty
	m.Cells[m.Height-2][m.Width-2] = maze.Empty

	// to quckly find the top left empty cell, the number of empty cells per row is stored
	// this way, top left empty cell can be found in O(height + width) time
	emptyCount := make([]int, m.Height-1)
	for row := 1; row < m.Width-1; row += 2 {
		emptyCount[row] = (m.Width - 1) / 2
	}

	for {
		// find the first empty cell from top left
		var r, c int
		for row := 1; row < m.Height-1; row += 2 {
			if emptyCount[row] > 0 {
				for col := 1; col < m.Width-1; col++ {
					if m.Cells[row][col] == maze.Empty {
						r, c = row, col
						m.Cells[r][c] = maze.Visited
						emptyCount[row]--
						break
					}
				}
				break
			}
		}

		// no empty cells remain and generation is complete
		if r == 0 {
			break
		}

		// connect the starting position to a random nearby visited cell
		for _, dir := range randomDirections() {
			neighRow, neighCol := r+dir[0]*2, c+dir[1]*2
			wallRow, wallCol := r+dir[0], c+dir[1]
			if neighRow >= 1 && neighRow < m.Height-1 && neighCol >= 1 && neighCol < m.Width-1 && m.Cells[neighRow][neighCol] == maze.Visited {
				m.Cells[wallRow][wallCol] = maze.Visited
				break
			}
		}

		// do a random walk until a dead end is reached
		var isDeadEnd bool

		for !isDeadEnd {
			isDeadEnd = true
			for _, dir := range randomDirections() {
				neighRow, neighCol := r+dir[0]*2, c+dir[1]*2
				wallRow, wallCol := r+dir[0], c+dir[1]
				if neighRow >= 1 && neighRow < m.Height-1 && neighCol >= 1 && neighCol < m.Width-1 && m.Cells[neighRow][neighCol] == maze.Empty {
					m.Cells[neighRow][neighCol] = maze.Visited
					m.Cells[wallRow][wallCol] = maze.Visited
					emptyCount[neighRow]--
					r, c = neighRow, neighCol
					isDeadEnd = false
					break
				}
			}
		}
	}

	return m, nil
}
