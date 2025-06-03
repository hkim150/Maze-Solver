package algorithm

import (
	"fmt"
	"math/rand"
	"maze-solver/internal/maze"
	"time"
)

func Tremaux(m *maze.Maze, animate bool) error {
    var delay time.Duration
    if animate {
        delay = 30 * time.Millisecond
    }

    parent := make(map[maze.Pos]maze.Pos)

    // initially, do a random walk to set the prev cell
    prev := m.StartPos
    n := neighbors(m, prev)
    if len(n) == 0 {
        return fmt.Errorf("No neighbors found for start position %v, cannot continue.\n", prev)
    }

    i := rand.Intn(len(n))
    curr := n[i]
    parent[curr] = prev

    var lastJunctionNext maze.Pos
	for curr != m.EndPos{
        if animate {
            tmp := m.Cells[curr[0]][curr[1]]
            m.Cells[curr[0]][curr[1]] = maze.Highlight
            m.PrintForAnimation(delay)
            m.Cells[curr[0]][curr[1]] = tmp
        }

        next := curr
		n := neighbors(m, curr)
		if len(n) == 0 {
			return fmt.Errorf("No neighbors found for position %v, cannot continue.\n", curr)
		} else if len(n) == 2 {
            // passage
            if n[0] == prev {
                next = n[1]
            } else {
                next = n[0]
            }
        } else if len(n) == 1 {
            // dead-end
            m.Cells[prev[0]][prev[1]] = maze.Visited
            next = prev
            lastJunctionNext = next
        } else if len(n) >= 3 {
            // or junction
            if m.Cells[prev[0]][prev[1]] == maze.Empty {
                m.Cells[prev[0]][prev[1]] = maze.Visiting
            } else if prev != lastJunctionNext && m.Cells[prev[0]][prev[1]] == maze.Visiting {
                m.Cells[prev[0]][prev[1]] = maze.Visited
            }

            // 1. pick an arbitrary unmarked entrance, if there is one
            // 2. if all entrances are marked once, go back to the entrance it came from
            // 3. pick any entrance with the fewest marks

            visiting := make([]maze.Pos, 0, len(n))
            for _, pos := range n {
                if m.Cells[pos[0]][pos[1]] == maze.Empty {
                    next = pos
                    break
                } else if m.Cells[pos[0]][pos[1]] == maze.Visiting {
                    visiting = append(visiting, pos)
                }
            }

            if next == curr {
                if len(visiting) == len(n) {
                    next = prev
                } else {
                    i := rand.Intn(len(visiting))
                    next = visiting[i]
                }
            }

            lastJunctionNext = next
            if m.Cells[next[0]][next[1]] == maze.Empty {
                m.Cells[next[0]][next[1]] = maze.Visiting
            } else if m.Cells[next[0]][next[1]] == maze.Visiting {
                m.Cells[next[0]][next[1]] = maze.Visited
            }
        }

        parent[next] = curr
        prev, curr = curr, next
	}

    if animate {
        m.Cells[curr[0]][curr[1]] = maze.Highlight
        m.PrintForAnimation(delay)
    }

    // Reconstruct the path
    prev, curr = m.EndPos, prev
    m.Cells[curr[0]][curr[1]] = maze.Highlight

    for curr != m.StartPos {
        found := false
        for _, nPos := range neighbors(m, curr) {
            if nPos != prev && m.Cells[nPos[0]][nPos[1]] == maze.Visiting {
                prev, curr = curr, nPos
                found = true
                break
            }
        }
        
        if !found {
            prev, curr = curr, parent[curr]
        }

        m.Cells[curr[0]][curr[1]] = maze.Highlight
    }

    m.CleanUp()

	return nil
}
