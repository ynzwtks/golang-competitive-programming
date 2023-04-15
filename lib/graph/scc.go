package main

type Scc struct {
	nodes int
	edges [][]int
}

func NewScc(g [][]int) *Scc {
	return &Scc{
		nodes: len(g),
		edges: g,
	}
}

// SCC は、グラフの強連結成分を求め、それらの成分ごとにスライスを返します。
// len(sccs) は強連結成分の数になります。
func (g *Scc) scc() [][]int {
	index := 0
	stack := []int{}
	indices := make([]int, g.nodes)
	lowlink := make([]int, g.nodes)
	onStack := make([]bool, g.nodes)
	var sccs [][]int

	var strongConnect func(int)
	strongConnect = func(v int) {
		indices[v] = index
		lowlink[v] = index
		index++
		stack = append(stack, v)
		onStack[v] = true

		for _, w := range g.edges[v] {
			if indices[w] < 0 {
				strongConnect(w)
				lowlink[v] = min(lowlink[v], lowlink[w])
			} else if onStack[w] {
				lowlink[v] = min(lowlink[v], indices[w])
			}
		}
		if lowlink[v] == indices[v] {
			var scc []int
			for {
				w := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				onStack[w] = false
				scc = append(scc, w)
				if w == v {
					break
				}
			}
			sccs = append(sccs, scc)
		}
	}
	for v := 0; v < g.nodes; v++ {
		indices[v] = -1
	}
	for v := 0; v < g.nodes; v++ {
		if indices[v] < 0 {
			strongConnect(v)
		}
	}
	return sccs
}
