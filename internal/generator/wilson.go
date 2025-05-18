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
	// it is to randomly select a starting point for the maze generation
	notVisited := dataStructure.NewRandomizedSet[[2]int]()
	for i := 1; i < maze.Height-1; i += 2 {
		for j := 1; j < maze.Width-1; j += 2 {
			notVisited.Add([2]int{i, j})
		}
	}

	// initially, choose a random cell as a visited cell
	initial, _ := notVisited.GetRandom()
	notVisited.Remove(initial)
	maze.Cells[initial[0]][initial[1]] = Visited

	// choose a random cell from the unvisited cells and start random walk
	for !notVisited.IsEmpty() {
		start, _ := notVisited.GetRandom()
		notVisited.Remove(start)
		row, col := start[0], start[1]

		// toVisit holds the next cells to visit
		// it's format is old cell + direction so that in case of walking into a dead end,
		// we can easily track and erase the path that was taken
		toVisit := dataStructure.NewStack[[4]int]()
		toVisit.Push([4]int{row, col, 0, 0}) // initially the direction is 0,0

		visiting := dataStructure.NewStack[[2]int]()

		// dfs tries all possible ordering of nodes with the same visiting nodes when backtracking
		// in order prevent that and to significantly reduce the search space,
		// keep track of the seen neighbors so that each unvisited cell is only considered as a candidate once
		seen := make(map[[2]int]bool)
		seen[[2]int{row, col}] = true

		// for optimization, instead of waiting until the currently visiting cell is a visited cell
		// we want to exit early when a visited neighboring cell is found,
		// the randomWalking boolean is used to signal the early exit
		randomWalking := true

		for randomWalking {
			elem, _ := toVisit.Pop()
			oldRow, oldCol, rowDir, colDir := elem[0], elem[1], elem[2], elem[3]
			currRow := oldRow + rowDir
			currCol := oldCol + colDir

			// if the current cell is not connected to the last cell in the walk path, it means we've reached a dead end in the last cell
			// backtrack the path until we find a cell that is connected to the current cell
			for !visiting.IsEmpty() {
				last, _ := visiting.Peek()
				if last[0] == oldRow && last[1] == oldCol {
					break
				}
				cell, _ := visiting.Pop()
				maze.Cells[cell[0]][cell[1]] = Empty
			}

			// add the current cell to the walk path
			visiting.Push([2]int{currRow, currCol})
			maze.Cells[currRow][currCol] = Visiting

			// add the current cell's unvisited neighbors to the stack
			for _, dir := range randomDirections() {
				neighRow := currRow + dir[0]*2
				neighCol := currCol + dir[1]*2
				if neighRow >= 1 && neighRow < maze.Height-1 && neighCol >= 1 && neighCol < maze.Width-1 {
					if maze.Cells[neighRow][neighCol] == Visited {
						visiting.Push([2]int{neighRow, neighCol})
						randomWalking = false
						break
					} else if !seen[[2]int{neighRow, neighCol}] {
						seen[[2]int{neighRow, neighCol}] = true
						toVisit.Push([4]int{currRow, currCol, dir[0] * 2, dir[1] * 2})
					}
				}
			}

			if !randomWalking {
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
			}
		}
	}

	return maze, nil
}
