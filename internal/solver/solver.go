package solver

import (
	"fmt"
	"maze-solver/internal/maze"
	"maze-solver/internal/solver/algorithm"
	"strings"
)

var solvers = map[string]solverFunc{
	"dfs": algorithm.DFS,
	"bfs": algorithm.BFS,
	"a-star": algorithm.AStar,
}

type solverFunc func(m *maze.Maze, animate bool) error

func (f solverFunc) Solve(m *maze.Maze, animate bool) error {
	return f(m, animate)
}

func Solve(m *maze.Maze, algorithm string, animate bool) error {
	// set the stant and end position as empty for solving
	m.Cells[m.StartPos[0]][m.StartPos[1]] = maze.Empty
	m.Cells[m.EndPos[0]][m.EndPos[1]] = maze.Empty

	solveFunc, ok := solvers[algorithm]
	if !ok {
		keys := make([]string, 0, len(solvers))
		for key := range solvers {
			keys = append(keys, key)
		}
		algorithms := strings.Join(keys, ", ")

		return fmt.Errorf("Unknown solver algorithm; Choose one from: %v\n", algorithms)
	}

	return solveFunc.Solve(m, animate)
}
