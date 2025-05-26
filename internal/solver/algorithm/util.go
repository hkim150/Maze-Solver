package algorithm

import "maze-solver/internal/maze"

// neighbors returns a slice of neighboring cells that is in boundary and not a wall
func neighbors(m *maze.Maze, pos maze.Pos) []maze.Pos {
	neighPos := []maze.Pos{}
	for _, dir := range [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
		row := pos[0] + dir[0]
		col := pos[1] + dir[1]
		if row >= m.StartPos[0] && row <= m.EndPos[0] && col >= m.StartPos[1] && col <= m.EndPos[1] && m.Cells[row][col] != maze.Wall {
			neighPos = append(neighPos, maze.Pos{row, col})
		}
	}

	return neighPos
}
