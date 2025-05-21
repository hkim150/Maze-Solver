package generator

import (
	"math/rand"
	dataStructure "maze-solver/internal/data_structure"
)

const mergeProb = 0.5

func SidewinderAlgorithm(width, height int) (Maze, error) {
	maze, err := baseMaze(width, height)
	if err != nil {
		return maze, err
	}

	// connect all columns in the top row
	for c := 2; c < maze.Width-2; c += 2 {
		maze.Cells[1][c] = Visited
	}

	// for each cell in a row, randomly decide whether to connect horizontal cells
	// if not connecting, randomly choose one cell from the set and connect it with the row above, then clear the set
	set := dataStructure.NewRandomizedSet[int]()
	for r:=3; r<maze.Height-1; r+=2 {
		set.Add(1)
		for c:=3; c<maze.Width-1; c+=2 {
			if rand.Float64() < mergeProb {
				maze.Cells[r][c-1] = Visited
			} else {
				randCol, _ := set.GetRandom()
				maze.Cells[r-1][randCol] = Visited
				set.Clear()
			}
			set.Add(c)
		}
		
		// flush the remainig set
		randCol, _ := set.GetRandom()
		maze.Cells[r-1][randCol] = Visited
		set.Clear()
	}

	return maze, nil
}
