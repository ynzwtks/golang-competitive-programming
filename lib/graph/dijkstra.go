package main

import (
	"math"
)

// dijkstra関数は、グラフGとコストCの行列、始点beginを入力として、
// 始点から各頂点への最小コストの配列と最短パスを復元するためのprev配列を返します。
func dijkstra(G, C [][]int, begin int) ([]int, []int) {
	n := len(G)
	dist := make([]int, n)
	visited := make([]bool, n)
	prev := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt64 // 最初はすべての頂点へのコストを無限大に設定
	}
	dist[begin] = 0 // 始点のコストを0に設定
	h := NewStructPQ(true)
	h.Push(PQData{0, begin})

	// 優先度付きキューが空になるまで繰り返す
	for h.Len() > 0 {
		t := h.Pop()
		id, _ := t.num, t.key
		if visited[id] == true { // 既に訪問済みの場合はスキップ
			continue
		}
		visited[id] = true
		for i, v := range G[id] {
			// 未訪問で、コストが更新される場合
			if visited[v] == false && dist[v] > dist[id]+C[id][i] {
				dist[v] = dist[id] + C[id][i]
				h.Push(PQData{dist[v], v})
				prev[v] = id
			}
		}
	}
	return dist, prev // 最小コスト配列と経路復元用配列を返す
}

// restorePath関数は、終点tと最短パス復元用の配列pを入力として、
// 最短パス上の頂点の配列を返します。
func restoreDijkstraPath(t int, p []int) []int {
	path := make([]int, 0)
	path = append(path, t)
	for p[t] != t {
		path = append(path, p[t])
		t = p[t]
	}
	return reverseList(path)
}

//データ構造体用ヒープ
type PQData struct {
	key int //並び替え用キー
	num int
}

func NewStructPQ(prioritySmallest bool) *StructPQ {
	ret := &StructPQ{}
	ret.d = make([]PQData, 0, 100)
	n := ret.Len()
	for i := n/2 - 1; i >= 0; i-- {
		ret.down(i, n)
	}
	ret.prioritySmallest = prioritySmallest
	return ret
}

type StructPQ struct {
	d                []PQData
	prioritySmallest bool
}

func (pq *StructPQ) less(i, j int) bool {
	if pq.prioritySmallest == true {
		return pq.d[i].key < pq.d[j].key
	} else {
		return pq.d[i].key > pq.d[j].key
	}
}
func (pq *StructPQ) swap(i, j int) { pq.d[i], pq.d[j] = pq.d[j], pq.d[i] }
func (pq *StructPQ) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && pq.less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !pq.less(j, i) {
			break
		}
		pq.swap(i, j)
		i = j
	}
	return i > i0
}
func (pq *StructPQ) Push(x PQData) {
	pq.d = append(pq.d, x)
	pq.up(len(pq.d) - 1)
}
func (pq *StructPQ) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !pq.less(j, i) {
			break
		}
		pq.swap(i, j)
		j = i
	}
}
func (pq *StructPQ) Pop() PQData {
	n := pq.Len() - 1
	x := pq.d[0]
	pq.swap(0, n)
	pq.down(0, n)
	pq.d = pq.d[0 : pq.Len()-1]
	return x
}
func (pq *StructPQ) Len() int {
	return len(pq.d)
}
func (pq *StructPQ) Peek() PQData {
	return pq.d[0]
}
func (pq *StructPQ) Remove(i int) PQData {
	n := pq.Len() - 1
	if n != i {
		pq.swap(i, n)
		if !pq.down(i, n) {
			pq.up(i)
		}
	}
	return pq.Pop()
}
func (pq *StructPQ) Fix(i int) {
	if !pq.down(i, pq.Len()) {
		pq.up(i)
	}
}
