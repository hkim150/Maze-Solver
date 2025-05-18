package generator

import (
	"math/rand"
	dataStructure "maze-solver/internal/data_structure"
)

func RandomizedPrim(width, height int) (Maze, error) {
	maze, err := baseMaze(width, height)
	if err != nil {
		return maze, err
	}

	// directionsToUnvisitedNeighbors is a helper function that returns the directions to the unvisited neighbor
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

	// initialize a randomized set to store frontier walls.
    // each wall is represented as [row, col, rowDir, colDir], where rowDir and colDir
    // indicate the direction to the neighboring cell.
	rs := dataStructure.NewRandomizedSet[[4]int]()
	dirs := directionsToUnvisitedNeighbors(row, col)
	for _, dir := range dirs {
		rs.Add([4]int{row, col, dir[0], dir[1]})
	}

	for !rs.IsEmpty() {
		// Randomly select a frontier wall from the pool
		elem, _ := rs.GetRandom()
		rs.Remove(elem)

		// Remove the wall
		row, col, rowDir, colDir := elem[0], elem[1], elem[2], elem[3]
		maze.Cells[row+rowDir][col+colDir] = Visited

		// mark the neighboring cell as visited and add the walls in between to the pool 
		dirs := directionsToUnvisitedNeighbors(row+rowDir*2, col+colDir*2)
		for _, dir := range dirs {
			rs.Add([4]int{row + rowDir*2, col + colDir*2, dir[0], dir[1]})
		}
	}

	maze.Cells[maze.StartCell[0]][maze.StartCell[1]] = Start
	maze.Cells[maze.EndCell[0]][maze.EndCell[1]] = End

	return maze, nil
}
