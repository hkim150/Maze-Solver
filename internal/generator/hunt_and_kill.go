package generator

func HuntAndKill(width, height int) (Maze, error) {
	maze, err := baseMaze(width, height)
	if err != nil {
		return maze, err
	}

	maze.Cells[1][1] = Empty
	maze.Cells[maze.Height-2][maze.Width-2] = Empty

	// to quckly find the top left empty cell, the number of empty cells per row is stored
	// this way, top left empty cell can be found in O(height + width) time
	emptyCount := make([]int, maze.Height-1)
	for row := 1; row < maze.Width-1; row += 2 {
		emptyCount[row] = (maze.Width - 1) / 2
	}

	for {
		// find the first empty cell from top left
		var r, c int
		for row := 1; row < maze.Height-1; row += 2 {
			if emptyCount[row] > 0 {
				for col := 1; col < maze.Width-1; col++ {
					if maze.Cells[row][col] == Empty {
						r, c = row, col
						maze.Cells[r][c] = Visited
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
			if neighRow >= 1 && neighRow < maze.Height-1 && neighCol >= 1 && neighCol < maze.Width-1 && maze.Cells[neighRow][neighCol] == Visited {
				maze.Cells[wallRow][wallCol] = Visited
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
				if neighRow >= 1 && neighRow < maze.Height-1 && neighCol >= 1 && neighCol < maze.Width-1 && maze.Cells[neighRow][neighCol] == Empty {
					maze.Cells[neighRow][neighCol] = Visited
					maze.Cells[wallRow][wallCol] = Visited
					emptyCount[neighRow]--
					r, c = neighRow, neighCol
					isDeadEnd = false
					break
				}
			}
		}
	}

	maze.Cells[1][1] = Start
	maze.Cells[maze.Height-2][maze.Width-2] = End

	return maze, nil
}
