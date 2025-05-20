package generator

import (
	"math/rand"
	dataStructure "maze-solver/internal/data_structure"
)

const (
	sidewaysMergeProb = 0.5
	downwardMergeProb = 0.2
)

func EllersAlgorithm(width, height int) (Maze, error) {
	maze, err := baseMaze(width, height)
	if err != nil {
		return maze, err
	}

	// union find for checking connectivity between two cells in O(1) time
	uf := dataStructure.NewUnionFind[[2]int]()

	for r := 1; r < maze.Height-3; r += 2 {
		// random horizonal cell merge in the row
		for c := 1; c < maze.Width-3; c += 2 {
			cell1 := [2]int{r, c}
			cell2 := [2]int{r, c + 2}
			if !uf.IsConnected(cell1, cell2) {
				if rand.Float32() < sidewaysMergeProb {
					maze.Cells[r][c+1] = Empty
					uf.Union(cell1, cell2)
				}
			}
		}

		// random downward cell merge. each set must at at least 1 merge with the row below
		// cellsInSet groups columns in current row that are in the same set after the horizontal cell merge
		cellsInSet := make(map[[2]int][]int)
		for c:=1; c<maze.Width; c+=2 {
			root := uf.Root([2]int{r, c})
			cellsInSet[root] = append(cellsInSet[root], c)
		}

		// for the remaining columns, randomly decide whether to create a downward connection
		for _, cols := range cellsInSet {
			rand.Shuffle(len(cols), func(i, j int) {
				cols[i], cols[j] = cols[j], cols[i]
			})
			
			for i:=0; i<len(cols); i++ {
				// make sure that a set has at least 1 connection downwards
				if i == 0 || rand.Float64() < downwardMergeProb {
					maze.Cells[r+1][cols[i]] = Empty
					uf.Union([2]int{r, cols[i]}, [2]int{r+2, cols[i]})
				}
			}
		}
	}

	// for bottom row, horizontally connect all sets that are not connected
	r := maze.Height-2
	for c := 1; c < maze.Width-3; c += 2 {
		cell1 := [2]int{r, c}
		cell2 := [2]int{r, c + 2}
		if !uf.IsConnected(cell1, cell2) {
			maze.Cells[r][c+1] = Empty
			uf.Union(cell1, cell2)
		}
	}

	return maze, nil
}
