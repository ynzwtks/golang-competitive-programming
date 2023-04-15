package main

import (
	"sort"
)

type MST struct {
	n         int
	edges     []Edge
	usedEdges []Edge
	uf        *UnionFind
	cost      int
}

func NewMST(n int) *MST {
	ret := &MST{}
	ret.n = n
	ret.uf = NewUnionFind(n)
	return ret
}
func (g *MST) ReadEdges(m int) {
	for i := 0; i < m; i++ {
		a, b, c := readint3()
		g.AddEdge(a, b, c)
	}
}
func (g *MST) AddEdge(source, destination, weight int) {
	g.edges = append(g.edges, Edge{source, destination, weight})
}
func (g *MST) Build(isSmallFirst bool) int {
	if isSmallFirst == true {
		sort.Slice(g.edges, func(i, j int) bool { return g.edges[i].weight < g.edges[j].weight })
	} else {
		sort.Slice(g.edges, func(i, j int) bool { return g.edges[i].weight > g.edges[j].weight })
	}
	for _, t := range g.edges {
		if g.uf.Connected(t.source, t.destination) == false {
			g.usedEdges = append(g.usedEdges, t)
			g.uf.Union(t.source, t.destination)
			g.cost += t.weight
		}
	}
	return g.cost
}
