package main

//WarshallFloyd アルゴリズムを実行し、すべてのノードペア間の最短経路を見つけます。
type WarshallFloyd struct {
	n     int
	edges [][]int
}

func NewWarshallFloyd(n int) *WarshallFloyd {
	ret := &WarshallFloyd{}
	ret.n = n
	ret.edges = make([][]int, n)
	for i := 0; i < n; i++ {
		ret.edges[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				ret.edges[i][j] = 0
			} else {
				ret.edges[i][j] = INF
			}
		}
	}
	return ret
}
func (g *WarshallFloyd) ReadEdges(m int) {
	for i := 0; i < m; i++ {
		a, b, c := readint3()
		g.AddEdge(a, b, c)
	}
}
func (g *WarshallFloyd) AddEdge(source, destination, weight int) {
	g.edges[source][destination] = weight
}
func (g *WarshallFloyd) Search() {
	for i := 0; i < g.n; i++ {
		for j := 0; j < g.n; j++ {
			for k := 0; k < g.n; k++ {
				if g.edges[j][i] != INF && g.edges[i][k] != INF {
					g.edges[j][k] = min(g.edges[j][k], g.edges[j][i]+g.edges[i][k])
				}
			}
		}
	}
}
func (g *WarshallFloyd) GetResult() [][]int {
	return g.edges
}
