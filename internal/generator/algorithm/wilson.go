package algorithm

import (
	dataStructure "maze-solver/internal/data_structure"
	"maze-solver/internal/maze"
	"time"
)

func Wilson(width, height int, animate bool) (*maze.Maze, error) {
	m, err := initialMaze(width, height)
	if err != nil {
		return m, err
	}

	// notVisited is a set of all cells that are not visited
	// it is to randomly select a starting point for the maze generation
	notVisited := dataStructure.NewRandomizedSet[maze.Pos]()
	for i := 1; i < m.Height-1; i += 2 {
		for j := 1; j < m.Width-1; j += 2 {
			notVisited.Add(maze.Pos{i, j})
		}
	}

	var delay time.Duration
	if animate {
		delay = 25 * time.Millisecond
		m.PrintForAnimation(delay)
	}

	// initially, choose a random cell as a visited cell
	initial, _ := notVisited.GetRandom()
	notVisited.Remove(initial)
	m.Cells[initial[0]][initial[1]] = maze.Visited

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

		visiting := dataStructure.NewStack[maze.Pos]()

		// dfs tries all possible ordering of nodes with the same visiting nodes when backtracking
		// in order prevent that and to significantly reduce the search space,
		// keep track of the seen neighbors so that each unvisited cell is only considered as a candidate once
		seen := make(map[maze.Pos]bool)
		seen[maze.Pos{row, col}] = true

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
				if animate {
					m.Cells[cell[0]][cell[1]] = maze.Highlight
					m.PrintForAnimation(delay)
				}
				m.Cells[cell[0]][cell[1]] = maze.Empty
			}

			// add the current cell to the walk path
			visiting.Push(maze.Pos{currRow, currCol})
			if animate {
				m.Cells[currRow][currCol] = maze.Highlight
				m.PrintForAnimation(delay)
			}
			m.Cells[currRow][currCol] = maze.Visiting

			// add the current cell's unvisited neighbors to the stack
			for _, dir := range randomDirections() {
				neighRow := currRow + dir[0]*2
				neighCol := currCol + dir[1]*2
				if neighRow >= 1 && neighRow < m.Height-1 && neighCol >= 1 && neighCol < m.Width-1 {
					if m.Cells[neighRow][neighCol] == maze.Visited {
						visiting.Push(maze.Pos{neighRow, neighCol})
						randomWalking = false
						break
					} else if !seen[maze.Pos{neighRow, neighCol}] {
						seen[maze.Pos{neighRow, neighCol}] = true
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
					if animate {
						m.Cells[(r1+r2)/2][(c1+c2)/2] = maze.Highlight
						m.PrintForAnimation(delay)
					}
					m.Cells[(r1+r2)/2][(c1+c2)/2] = maze.Visited
					m.Cells[r1][c1] = maze.Visited
					notVisited.Remove(maze.Pos{r1, c1})
				}
				// mark the last cell as visited and remove it from the notVisited set
				m.Cells[s[len(s)-1][0]][s[len(s)-1][1]] = maze.Visited
				notVisited.Remove(maze.Pos{s[len(s)-1][0], s[len(s)-1][1]})
			}
		}
	}

	return m, nil
}
