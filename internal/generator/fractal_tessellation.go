package generator

import (
	"math/rand"
)

func FractalTessellation(width, height int) (Maze, error) {
	// the algorithm generates squares of certain sizes
	// clip the length of the maze to the largest square within the width and height
	length := 3
	maxLength := min(width, height)
	for {
		if 2*length-1 > maxLength {
			break
		}
		
		length = 2*length - 1
	}

	maze, err := baseMaze(length, length)
	if err != nil {
		return maze, err
	}

	// set the top left start cell as empty so that the start cell doesn't get copied over
	maze.Cells[1][1] = Empty

	length = 1
	for 2*length-1 < maze.Width {
		// copy the current square fractal three times - right, below, diagonal
		startCells := [3][2]int{{1, length + 2}, {length + 2, 1}, {length + 2, length + 2}}
		for _, startCell := range startCells {
			row, col := startCell[0], startCell[1]
			for i := range length {
				for j := range length {
					maze.Cells[row+i][col+j] = maze.Cells[i+1][j+1]
				}
			}
		}

		// make holes on three of the four walls that divide each copied section
		r1 := 1 + rand.Intn((length+1)/2)*2
		c1 := 1 + rand.Intn((length+1)/2)*2
		r2 := length + 2 + rand.Intn((length+1)/2)*2
		c2 := length + 2 + rand.Intn((length+1)/2)*2

		holeWalls := [4][2]int{{r1, length + 1}, {length + 1, c1}, {r2, length + 1}, {length + 1, c2}}
		rand.Shuffle(4, func(i, j int) {
			holeWalls[i], holeWalls[j] = holeWalls[j], holeWalls[i]
		})

		for i := range 3 {
			maze.Cells[holeWalls[i][0]][holeWalls[i][1]] = Empty
		}

		length = 2*length + 1
	}

	maze.Cells[1][1] = Start
	maze.Cells[maze.Height-2][maze.Width-2] = End

	return maze, nil
}
