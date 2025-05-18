package generator

import "math/rand"

func AldousBroder(width, height int) (Maze, error) {
	maze, err := baseMaze(width, height)
	if err != nil {
		return maze, err
	}

	// notVisited is a map to keep track of the unvisited cells for end condition
	notVisited := make(map[[2]int]bool)
	for row := 1; row < maze.Height-1; row += 2 {
		for col := 1; col < maze.Width-1; col += 2 {
			notVisited[[2]int{row, col}] = true
		}
	}

	// Randomly select a starting cell from the isolated empty cells.
	startRow := rand.Intn((maze.Height-1)/2)*2 + 1
	startCol := rand.Intn((maze.Width-1)/2)*2 + 1
	maze.Cells[startRow][startCol] = Visited
	delete(notVisited, [2]int{startRow, startCol})

	for len(notVisited) > 0 {	
		for _, dir := range randomDirections() {
			neighRow := startRow + dir[0]*2
			neighCol := startCol + dir[1]*2
			if neighRow >= 1 && neighRow < maze.Height-1 && neighCol >= 1 && neighCol < maze.Width-1 {
				if maze.Cells[neighRow][neighCol] != Visited {
					// remove the wall between the current cell and the neighbor
					maze.Cells[startRow+dir[0]][startCol+dir[1]] = Visited
					maze.Cells[neighRow][neighCol] = Visited
					delete(notVisited, [2]int{neighRow, neighCol})
				}
				startRow, startCol = neighRow, neighCol
				break
			}
		}
	}

	return maze, nil
}
