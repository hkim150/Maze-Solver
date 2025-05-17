package internal

import (
	"math/rand"
)

type MazeGenerator interface {
	Generate(width, height int) Maze
}

type mazeGeneratorFunc func(width, height int) Maze

func (f mazeGeneratorFunc) Generate(width, height int) Maze {
	return f(width, height)
}

func RandomizedDFS(width, height int) (Maze, error) {
	maze, err := baseMaze(width, height)
	if err != nil {
		return maze, err
	}

	startRow := rand.Intn((maze.Height-1)/2)*2 + 1
	startCol := rand.Intn((maze.Width-1)/2)*2 + 1

	maze.Cells[startRow][startCol] = Visited
	stack := [][2]int{{startRow, startCol}}

	for len(stack) > 0 {
		row, col := stack[len(stack)-1][0], stack[len(stack)-1][1]
		stack = stack[:len(stack)-1]

		directions := [][2]int{{-2, 0}, {0, -2}, {2, 0}, {0, 2}}
		rand.Shuffle(len(directions), func(i, j int) {
			directions[i], directions[j] = directions[j], directions[i]
		})
		for _, dir := range directions {
			neighRow := row + dir[0]
			neighCol := col + dir[1]
			if neighRow >= 1 && neighRow < maze.Height-1 && neighCol >= 1 && neighCol < maze.Width-1 && maze.Cells[neighRow][neighCol] != Visited {
				maze.Cells[row+dir[0]/2][col+dir[1]/2] = Visited
				maze.Cells[neighRow][neighCol] = Visited
				stack = append(stack, [2]int{row, col})
				stack = append(stack, [2]int{neighRow, neighCol})
				break
			}
		}
	}

	return maze, nil
}

func RandomizedKruskal(width, height int) (Maze, error) {
	maze, err := baseMaze(width, height)
	if err != nil {
		return maze, err
	}

	walls := make([][2]int, 0)
	for row := 1; row < maze.Height-1; row++ {
		for col := row%2 + 1; col < maze.Width-1; col += 2 {
			walls = append(walls, [2]int{row, col})
		}
	}

	to1D := func(row, col int) int {
		return (row-1)/2*(maze.Width-1)/2 + (col-1)/2
	}

	uf := NewUnionFind((maze.Width - 1) / 2 * (maze.Height - 1) / 2)

	rand.Shuffle(len(walls), func(i, j int) {
		walls[i], walls[j] = walls[j], walls[i]
	})

	for _, wall := range walls {
		row, col := wall[0], wall[1]
		var cell_1, cell_2 int
		if row%2 == 1 {
			cell_1 = to1D(row, col-1)
			cell_2 = to1D(row, col+1)
		} else {
			cell_1 = to1D(row-1, col)
			cell_2 = to1D(row+1, col)
		}

		if !uf.IsConnected(cell_1, cell_2) {
			maze.Cells[row][col] = Empty
			uf.Union(cell_1, cell_2)
		}
	}

	return maze, nil
}

func RandomizedPrim(width, height int) (Maze, error) {
	maze, err := baseMaze(width, height)
	if err != nil {
		return maze, err
	}

	// directionsToUnvisitedNeighbors returns the directions to the unvisited neighbor
	directionsToUnvisitedNeighbors := func(row, col int) [][2]int {
		directions := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
		unvisitedDir := make([][2]int, 0)
		for _, dir := range directions {
			neighRow := row + dir[0]*2
			neighCol := col + dir[1]*2
			if neighRow >= 1 && neighRow < maze.Height-1 && neighCol >= 1 && neighCol < maze.Width-1 && maze.Cells[neighRow][neighCol] != Visited {
				maze.Cells[neighRow][neighCol] = Visited
				unvisitedDir = append(unvisitedDir, dir)
			}
		}

		return unvisitedDir
	}

	row := rand.Intn(maze.Height/2)*2 + 1
	col := rand.Intn(maze.Width/2)*2 + 1
	maze.Cells[row][col] = Visited

	rs := NewRandomizedSet[[4]int]()
	dirs := directionsToUnvisitedNeighbors(row, col)
	for _, dir := range dirs {
		rs.Add([4]int{row, col, dir[0], dir[1]})
	}

	for !rs.IsEmpty() {
		elem, _ := rs.GetRandom()
		rs.Remove(elem)
		row, col, rowDir, colDir := elem[0], elem[1], elem[2], elem[3]
		maze.Cells[row+rowDir][col+colDir] = Empty
		dirs := directionsToUnvisitedNeighbors(row+rowDir*2, col+colDir*2)
		for _, dir := range dirs {
			rs.Add([4]int{row + rowDir*2, col + colDir*2, dir[0], dir[1]})
		}
	}

	return maze, nil
}

func WilsonsAlgorithm(width, height int) (Maze, error) {
	maze, err := baseMaze(width, height)
	if err != nil {
		return maze, err
	}

	// notVisited is a set of all cells that are not visited
	// It will be used to randomly select a starting point for the maze generation
	notVisited := NewRandomizedSet[[2]int]()
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
		visiting := NewHashStack[[2]int]()
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
			directions := [][2]int{{-2, 0}, {0, -2}, {2, 0}, {0, 2}}
			rand.Shuffle(len(directions), func(i, j int) {
				directions[i], directions[j] = directions[j], directions[i]
			})
			for _, dir := range directions {
				neighRow := currRow + dir[0]
				neighCol := currCol + dir[1]
				if neighRow >= 1 && neighRow < maze.Height-1 && neighCol >= 1 && neighCol < maze.Width-1 && !visiting.Contains([2]int{neighRow, neighCol}) {
					stack = append(stack, [4]int{currRow, currCol, dir[0], dir[1]})
				}
			}
		}
	}

	return maze, nil
}
