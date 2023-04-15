package main

import "math"

const (
	INF = math.MaxInt64
)

// BellmanFordは、Bellman-Fordアルゴリズムを実行するための構造体です。
type BellmanFord struct {
	n     int    // グラフの頂点数
	edges []Edge // グラフのエッジリスト
}

// NewBellmanFordは、新しいBellmanFord構造体を作成し初期化します。
func NewBellmanFord(n int) *BellmanFord {
	ret := &BellmanFord{}
	ret.n = n
	return ret
}

// AddEdgeは、グラフにエッジを追加します。
func (g *BellmanFord) AddEdge(source, destination, weight int) {
	g.edges = append(g.edges, Edge{source, destination, weight})
}

// Searchは、指定した始点からの最短経路をBellman-Fordアルゴリズムで探索し、最短距離のリストを返します。
// グラフに負の閉路が存在する場合、nilを返します。
func (g *BellmanFord) Search(s int) []int {
	dist := make([]int, g.n) // 各頂点への最短距離を格納する配列
	for i := 0; i < g.n; i++ {
		dist[i] = INF
	}
	dist[s] = 0

	// 各頂点に対してエッジを緩和する
	for i := 0; i < len(dist); i++ {
		update := false
		for _, edge := range g.edges {
			// 現在の始点からの距離とエッジの重みを加えたものが、現在の終点への距離より短い場合、距離を更新する
			if dist[edge.source] != INF && dist[edge.source]+edge.weight < dist[edge.destination] {
				dist[edge.destination] = dist[edge.source] + edge.weight
				update = true
			}
		}
		// 更新がない場合、ループを抜ける
		if update == false {
			break
		}
	}

	// 負の閉路の検出
	for _, edge := range g.edges {
		if dist[edge.source] != INF && dist[edge.source]+edge.weight < dist[edge.destination] {
			return nil
		}
	}

	return dist
}
