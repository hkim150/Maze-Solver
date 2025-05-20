package dataStructure

type UnionFind[T comparable] struct {
	root map[T]T
	rank map[T]int
}

func NewUnionFind[T comparable]() *UnionFind[T] {
	return &UnionFind[T]{
		root: make(map[T]T),
		rank: make(map[T]int),
	}
}

func (uf *UnionFind[T]) Root(x T) T {
	if _, ok := uf.root[x]; !ok {
		uf.root[x] = x
		return x
	}

	if uf.root[x] != x {
		uf.root[x] = uf.Root(uf.root[x])
	}

	return uf.root[x]
}

func (uf *UnionFind[T]) Union(x, y T) {
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

func (uf *UnionFind[T]) IsConnected(x, y T) bool {
	return uf.Root(x) == uf.Root(y)
}
