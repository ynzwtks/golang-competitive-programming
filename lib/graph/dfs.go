package main

// DFSは深さ優先探索を実行するための構造体です。
type DFS struct {
	g            [][]int // グラフ（隣接リスト形式）
	visited      []bool  // 各頂点の訪問状態
	path         []int   // 探索パス
	shortestPath []int   // 最短パス
}

// NewDFSは、新しいDFS構造体を作成し初期化します。
func NewDFS(g [][]int) *DFS {
	d := &DFS{}
	d.g = g
	d.visited = make([]bool, len(g))
	d.path = make([]int, 0)
	d.shortestPath = make([]int, len(g)+1)
	return d
}

// Searchは、深さ優先探索を実行し、指定した終了ノードに到達した場合はtrueを返します。
func (d *DFS) Search(s int, t int) bool {
	pos := s
	d.visited[pos] = true
	d.path = append(d.path, pos)

	if s == t {
		if len(d.path) < len(d.shortestPath) {
			d.shortestPath = make([]int, len(d.path))
			copy(d.shortestPath, d.path)
		}
	} else {
		for i := 0; i < len(d.g[pos]); i++ {
			next := d.g[pos][i]
			if d.visited[next] == false {
				d.Search(next, t)
			}
		}
	}
	d.visited[pos] = false
	d.path = d.path[:len(d.path)-1]
	return len(d.shortestPath) <= len(d.g)
}
