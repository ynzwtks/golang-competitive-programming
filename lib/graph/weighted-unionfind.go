package main

import "fmt"

type WeightedUnionFind struct {
	parents []int
	weights []int
}

func NewWeightedUnionFind(n int) *WeightedUnionFind {
	uf := &WeightedUnionFind{
		parents: make([]int, n),
		weights: make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.parents[i] = i
		uf.weights[i] = 0
	}
	return uf
}

func (uf *WeightedUnionFind) Root(x int) int {
	if uf.parents[x] == x {
		return x
	}
	root := uf.Root(uf.parents[x])
	uf.weights[x] += uf.weights[uf.parents[x]]
	uf.parents[x] = root
	return root
}

func (uf *WeightedUnionFind) Weight(x int) int {
	uf.Root(x)
	return uf.weights[x]
}

func (uf *WeightedUnionFind) Unite(x, y int, w int) bool {
	rx := uf.Root(x)
	ry := uf.Root(y)
	if rx == ry {
		return false
	}
	w += uf.Weight(x)
	w -= uf.Weight(y)
	uf.parents[rx] = ry
	uf.weights[rx] = -w
	return true
}

func (uf *WeightedUnionFind) Same(x, y int) bool {
	return uf.Root(x) == uf.Root(y)
}

func (uf *WeightedUnionFind) Diff(x, y int) int {
	return uf.Weight(y) - uf.Weight(x)
}

func main() {
	wuf := NewWeightedUnionFind(10)
	wuf.Unite(1, 5, 2)
	wuf.Unite(5, 6, 2)
	fmt.Println(wuf.Diff(1, 6))
}
