package algorithm

import (
	"math/rand"
	dataStructure "maze-solver/internal/data_structure"
	"maze-solver/internal/maze"
)

const (
	sidewaysMergeProb = 0.5
	downwardMergeProb = 0.2
)

func Ellers(width, height int) (*maze.Maze, error) {
	m, err := initialMaze(width, height)
	if err != nil {
		return m, err
	}

	// union find for checking connectivity between two cells in O(1) time
	uf := dataStructure.NewUnionFind[maze.Pos]()

	for r := 1; r < m.Height-3; r += 2 {
		// random horizonal cell merge in the row
		for c := 1; c < m.Width-3; c += 2 {
			cell1 := maze.Pos{r, c}
			cell2 := maze.Pos{r, c + 2}
			if !uf.IsConnected(cell1, cell2) {
				if rand.Float32() < sidewaysMergeProb {
					m.Cells[r][c+1] = maze.Empty
					uf.Union(cell1, cell2)
				}
			}
		}

		// random downward cell merge. each set must at at least 1 merge with the row below
		// cellsInSet groups columns in current row that are in the same set after the horizontal cell merge
		cellsInSet := make(map[maze.Pos][]int)
		for c := 1; c < m.Width; c += 2 {
			root := uf.Root(maze.Pos{r, c})
			cellsInSet[root] = append(cellsInSet[root], c)
		}

		// for the remaining columns, randomly decide whether to create a downward connection
		for _, cols := range cellsInSet {
			rand.Shuffle(len(cols), func(i, j int) {
				cols[i], cols[j] = cols[j], cols[i]
			})

			for i := 0; i < len(cols); i++ {
				// make sure that a set has at least 1 connection downwards
				if i == 0 || rand.Float64() < downwardMergeProb {
					m.Cells[r+1][cols[i]] = maze.Empty
					uf.Union(maze.Pos{r, cols[i]}, maze.Pos{r + 2, cols[i]})
				}
			}
		}
	}

	// for bottom row, horizontally connect all sets that are not connected
	r := m.Height - 2
	for c := 1; c < m.Width-3; c += 2 {
		cell1 := maze.Pos{r, c}
		cell2 := maze.Pos{r, c + 2}
		if !uf.IsConnected(cell1, cell2) {
			m.Cells[r][c+1] = maze.Empty
			uf.Union(cell1, cell2)
		}
	}

	return m, nil
}
