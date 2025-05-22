package algorithm

import (
	"math/rand"
	dataStructure "maze-solver/internal/data_structure"
	"maze-solver/internal/maze"
)

const mergeProb = 0.5

func Sidewinder(width, height int) (*maze.Maze, error) {
	m, err := initialMaze(width, height)
	if err != nil {
		return m, err
	}

	// connect all columns in the top row
	for c := 2; c < m.Width-2; c += 2 {
		m.Cells[1][c] = maze.Visited
	}

	// for each cell in a row, randomly decide whether to connect horizontal cells
	// if not connecting, randomly choose one cell from the set and connect it with the row above, then clear the set
	set := dataStructure.NewRandomizedSet[int]()
	for r := 3; r < m.Height-1; r += 2 {
		set.Add(1)
		for c := 3; c < m.Width-1; c += 2 {
			if rand.Float64() < mergeProb {
				m.Cells[r][c-1] = maze.Visited
			} else {
				randCol, _ := set.GetRandom()
				m.Cells[r-1][randCol] = maze.Visited
				set.Clear()
			}
			set.Add(c)
		}

		// flush the remainig set
		randCol, _ := set.GetRandom()
		m.Cells[r-1][randCol] = maze.Visited
		set.Clear()
	}

	return m, nil
}
