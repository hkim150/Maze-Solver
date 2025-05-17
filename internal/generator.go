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

	maze.Cells[1][1] = Start
	maze.Cells[maze.Height-2][maze.Width-2] = End

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

	maze.Cells[1][1] = Start
	maze.Cells[maze.Height-2][maze.Width-2] = End

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
		elem, _:= rs.GetRandom()
		rs.Remove(elem)
		row, col, rowDir, colDir := elem[0], elem[1], elem[2], elem[3]
		maze.Cells[row+rowDir][col+colDir] = Empty
		dirs := directionsToUnvisitedNeighbors(row+rowDir*2, col+colDir*2)
		for _, dir := range dirs {
			rs.Add([4]int{row + rowDir*2, col + colDir*2, dir[0], dir[1]})
		}
	}

	maze.Cells[1][1] = Start
	maze.Cells[maze.Height-2][maze.Width-2] = End

	return maze, nil
}

func WilsonsAlgorithm(width, height int) (Maze, error) {
	maze, err := baseMaze(width, height)
	if err != nil {
		return maze, err
	}

	notVisited := NewRandomizedSet[[2]int]()
	for i := 1; i < maze.Height-1; i += 2 {
		for j := 1; j < maze.Width-1; j += 2 {
			notVisited.Add([2]int{i, j})
		}
	}

	initial, _ := notVisited.GetRandom()
	row, col := initial[0], initial[1]
	notVisited.Remove(initial)
	maze.Cells[row][col] = Visited

	for !notVisited.IsEmpty() {
		start, _ := notVisited.GetRandom()
		row, col = start[0], start[1]
		stack := [][2]int{{row, col}}
		passedWalls := [][2]int{}
		visiting := map[[2]int]bool{{row, col}: true}
		maze.Print()
		for {
			cell := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			row, col = cell[0], cell[1]
			if maze.Cells[row][col] == Visited {
				for _, cell := range passedWalls {
					maze.Cells[cell[0]][cell[1]] = Visited
				}

				for _, cell := range stack {
					maze.Cells[cell[0]][cell[1]] = Visited
					notVisited.Remove(cell)
				}
				break
			}

			directions := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
			rand.Shuffle(len(directions), func(i, j int) {
				directions[i], directions[j] = directions[j], directions[i]
			})
			neighCount := 0
			for _, dir := range directions {
				neighRow := row + dir[0]*2
				neighCol := col + dir[1]*2
				if neighRow >= 1 && neighRow < maze.Height-1 && neighCol >= 1 && neighCol < maze.Width-1 {
					if !visiting[[2]int{neighRow, neighCol}] {
						stack = append(stack, [2]int{neighRow, neighCol})
						passedWalls = append(passedWalls, [2]int{row + dir[0], col + dir[1]})
						visiting[[2]int{neighRow, neighCol}] = true
						neighCount++
					}
				}
			}
			if neighCount == 0 {
				visiting[[2]int{row, col}] = false
				passedWalls = passedWalls[:len(passedWalls)-1]
			}
		}
	}

	maze.Cells[1][1] = Start
	maze.Cells[maze.Height-2][maze.Width-2] = End
	return maze, nil
}
