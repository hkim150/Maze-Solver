package solver

import (
	"fmt"
	"maze-solver/internal/maze"
	"maze-solver/internal/solver/algorithm"
	"strings"
)

var solvers = map[string]solverFunc{
	"dfs": algorithm.DFS,
}

type solverFunc func(maze *maze.Maze, animate bool) error

func (f solverFunc) Solve(maze *maze.Maze, animate bool) error {
	return f(maze, animate)
}

func Solve(maze *maze.Maze, algorithm string, animate bool) error {
	solveFunc, ok := solvers[algorithm]
	if !ok {
		keys := make([]string, 0, len(solvers))
		for key := range solvers {
			keys = append(keys, key)
		}
		algorithms := strings.Join(keys, ", ")

		return fmt.Errorf("Unknown solver algorithm; Choose one from: %v\n", algorithms)
	}

	return solveFunc.Solve(maze, animate)
}
