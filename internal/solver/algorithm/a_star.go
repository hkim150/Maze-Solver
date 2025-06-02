package algorithm

import (
	"container/heap"
	"fmt"
	"maze-solver/internal/maze"
	"time"
)

func AStar(m *maze.Maze, animate bool) error {
	pq := make(priorityQueue, 0)
	heap.Init(&pq)

	// nodes is a map to access the node in priority queue in O(1) time
	nodes := make(map[maze.Pos]*node)

	// parent position to reconstruct the path
	parent := make(map[maze.Pos]maze.Pos)

	startNode := &node{pos: m.StartPos, gScore: 0, fScore: 0}
	nodes[m.StartPos] = startNode
	heap.Push(&pq, startNode)

	var delay time.Duration
	if animate {
		delay = 30 * time.Millisecond
	}

	solved := false
	for len(pq) > 0 {
		curr := heap.Pop(&pq).(*node)

		if curr.pos == m.EndPos {
			solved = true
			break
		}

		if animate {
			m.Cells[curr.pos[0]][curr.pos[1]] = maze.Highlight
			m.PrintForAnimation(delay)
			m.Cells[curr.pos[0]][curr.pos[1]] = maze.Visited
		}

		for _, nPos := range neighbors(m, curr.pos) {
			// use manhattan distance as heuristic
			g := nodes[curr.pos].gScore + 1
			f := g + manhattanDistance(nPos, m.EndPos)

			if neighbor, ok := nodes[nPos]; !ok {
				neighbor = &node{pos: nPos, gScore: g, fScore: f}
				nodes[nPos] = neighbor
				heap.Push(&pq, neighbor)
				parent[nPos] = curr.pos
			} else if f < neighbor.fScore {
				pq.update(neighbor, nPos, g, f)
				parent[nPos] = curr.pos
			}
		}
	}

	if !solved {
		return fmt.Errorf("Could not find a solution for the maze")
	}

	// Reconstruct the path
	m.Reset()
	for p := m.EndPos; p != m.StartPos; p = parent[p] {
		m.Cells[p[0]][p[1]] = maze.Highlight
	}
	m.Cells[m.StartPos[0]][m.StartPos[1]] = maze.Highlight

	return nil
}

type node struct {
	pos    maze.Pos
	gScore int
	fScore int
	index  int // index in the priority queue
}

type priorityQueue []*node

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].fScore < pq[j].fScore
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x any) {
	item := x.(*node)
	item.index = len(*pq)
	*pq = append(*pq, x.(*node))
}

func (pq *priorityQueue) Pop() any {
	old := *pq
	l := len(old)
	x := old[l-1]
	old[l-1] = nil // don't stop the GC from reclaiming the memory
	x.index = -1   // for safety
	*pq = old[:l-1]
	return x
}

func (pq *priorityQueue) update(n *node, pos maze.Pos, g, f int) {
	n.pos = pos
	n.gScore = g
	n.fScore = f
	heap.Fix(pq, n.index)
}

func manhattanDistance(a, b maze.Pos) int {
	return abs(a[0]-b[0]) + abs(a[1]-b[1])
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
