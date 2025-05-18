package generator

import (
	dataStructure "maze-solver/internal/data_structure"
)

func WilsonsAlgorithm(width, height int) (Maze, error) {
	maze, err := baseMaze(width, height)
	if err != nil {
		return maze, err
	}

	// notVisited is a set of all cells that are not visited
	// It will be used to randomly select a starting point for the maze generation
	notVisited := dataStructure.NewRandomizedSet[[2]int]()
	for i := 1; i < maze.Height-1; i += 2 {
		for j := 1; j < maze.Width-1; j += 2 {
			notVisited.Add([2]int{i, j})
		}
	}

	// Initially, one cell is marked as visited to provide destination for the random walk
	initial, _ := notVisited.GetRandom()
	notVisited.Remove(initial)
	row, col := initial[0], initial[1]
	maze.Cells[row][col] = Visited

	for !notVisited.IsEmpty() {
		// Randomly select a cell from the set of unvisited cells and start a random walk
		start, _ := notVisited.GetRandom()
		notVisited.Remove(start)

		// stack holds the next cells to visit.
		// It is in old cell + direction format so that in case of walking into a dead end,
		// we can easily track and erase the path that was taken
		stack := [][4]int{{start[0], start[1], 0, 0}} // initially the direction is 0,0

		// visiting holds the nodes that have been part of the path so far from the random walk
		// It is used to detect cycles and to mark the cells as visited at the end of the walk
		// hashstack is used so that we can check if a cell is already visited in O(1) time
		visiting := dataStructure.NewHashStack[[2]int]()
		for {
			// curr holds the direction to the curr cell to visit from an old cell we came from
			curr := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			oldRow, oldCol, rowDir, colDir := curr[0], curr[1], curr[2], curr[3]
			currRow := oldRow + rowDir
			currCol := oldCol + colDir

			// if the current cell is not connected to the last cell in the walk path, it means we've reached a dead end in the last cell
			// backtrack the path until we find a cell that is connected to the current cell
			for !visiting.IsEmpty() {
				last, _ := visiting.Peek()
				if last[0] == oldRow && last[1] == oldCol {
					break
				}

				visiting.Pop()
			}
			// add the current cell to the walk path
			visiting.Push([2]int{currRow, currCol})

			// check if we reached a cell that is visited and need to end the random walk
			if maze.Cells[currRow][currCol] == Visited {
				// make the walls in between the visiting cells as visited, so that it creates a path
				// mark the visiting cells as visited and remove them from the notVisited set
				s := visiting.ToSlice()
				for i := 0; i < len(s)-1; i++ {
					r1, c1 := s[i][0], s[i][1]
					r2, c2 := s[i+1][0], s[i+1][1]
					maze.Cells[(r1+r2)/2][(c1+c2)/2] = Visited
					maze.Cells[r1][c1] = Visited
					notVisited.Remove([2]int{r1, c1})
				}
				// mark the last cell as visited and remove it from the notVisited set
				maze.Cells[s[len(s)-1][0]][s[len(s)-1][1]] = Visited
				notVisited.Remove([2]int{s[len(s)-1][0], s[len(s)-1][1]})
				break
			}

			// add the current cell's unvisited neighbors to the stack
			for _, dir := range randomDirections() {
				neighRow := currRow + dir[0]*2
				neighCol := currCol + dir[1]*2
				if neighRow >= 1 && neighRow < maze.Height-1 && neighCol >= 1 && neighCol < maze.Width-1 && !visiting.Contains([2]int{neighRow, neighCol}) {
					stack = append(stack, [4]int{currRow, currCol, dir[0] * 2, dir[1] * 2})
				}
			}
		}
	}

	return maze, nil
}
