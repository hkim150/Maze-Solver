package dataStructure

type UnionFind struct {
	root []int
	rank []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		root: make([]int, n),
		rank: make([]int, n),
	}

	for i := 0; i < n; i++ {
		uf.root[i] = i
	}

	return uf
}

func (uf *UnionFind) Root(x int) int {
	if uf.root[x] != x {
		uf.root[x] = uf.Root(uf.root[x])
	}

	return uf.root[x]
}

func (uf *UnionFind) Union(x, y int) {
	if uf.Root(x) == uf.Root(y) {
		return
	}

	if uf.rank[uf.root[x]] < uf.rank[uf.root[y]] {
		uf.root[uf.root[x]] = uf.root[y]
	} else if uf.rank[uf.root[x]] > uf.rank[uf.root[y]] {
		uf.root[uf.root[y]] = uf.root[x]
	} else {
		uf.root[uf.root[y]] = uf.root[x]
		uf.rank[uf.root[x]]++
	}
}

func (uf *UnionFind) IsConnected(x, y int) bool {
	return uf.Root(x) == uf.Root(y)
}
