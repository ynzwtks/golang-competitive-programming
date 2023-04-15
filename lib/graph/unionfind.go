package main

type UnionFind struct {
	root  []int
	size  []int
	group int
}

func NewUnionFind(n int) *UnionFind {
	root := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		root[i] = i
		size[i] = 1
	}
	uf := &UnionFind{root: root, size: size, group: n}
	return uf
}
func (uf *UnionFind) Union(p int, q int) {
	qRoot := uf.Root(q)
	pRoot := uf.Root(p)
	if qRoot == pRoot {
		return
	}
	uf.group--
	if uf.size[qRoot] < uf.size[pRoot] {
		uf.root[qRoot] = uf.root[pRoot]
		uf.size[pRoot] += uf.size[qRoot]
	} else {
		uf.root[pRoot] = uf.root[qRoot]
		uf.size[qRoot] += uf.size[pRoot]
	}
}
func (uf *UnionFind) Root(p int) int {
	if p > len(uf.root)-1 {
		return -1
	}
	for uf.root[p] != p {
		uf.root[p] = uf.root[uf.root[p]]
		p = uf.root[p]
	}
	return p
}
func (uf *UnionFind) find(p int) int {
	return uf.Root(p)
}
func (uf *UnionFind) Connected(p int, q int) bool {
	return uf.Root(p) == uf.Root(q)
}
func (uf *UnionFind) Groups() map[int]int {
	cm := make(map[int]int)
	for i := 0; i < len(uf.root); i++ {
		t := uf.find(uf.root[i])
		cm[t]++
	}
	return cm
}
