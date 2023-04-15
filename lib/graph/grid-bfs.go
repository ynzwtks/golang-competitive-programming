package main

import "container/list"

// gridBfs は、2Dグリッド上で01-BFSを実行する関数です。
// g: 2Dグリッド(マップ)。 '.' は通行可能で、それ以外の文字は壁を表します。
// y, x: 開始位置の座標
// 戻り値: 各座標への最短移動回数を格納した2Dスライス
func gridBfs(g [][]byte, y, x int) [][]int {
	h := len(g)
	w := len(g[0])
	visited := make([][]int, h)
	for i := 0; i < h; i++ {
		visited[i] = make([]int, w)
		for j := 0; j < w; j++ {
			visited[i][j] = -1
		}
	}
	dxy := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	queue := list.New()
	queue.PushBack([2]int{y, x})
	visited[y][x] = 0

	for queue.Len() != 0 {
		cur := queue.Front().Value.([2]int)
		queue.Remove(queue.Front())
		cy, cx := cur[0], cur[1]

		for _, d := range dxy {
			ny, nx := cy+d[0], cx+d[1]
			if nx >= 0 && nx < w && ny >= 0 && ny < h {
				if g[ny][nx] != '#' && visited[ny][nx] == -1 {
					visited[ny][nx] = visited[cy][cx] + 1
					queue.PushBack([2]int{ny, nx})
				}
			}
		}
	}
	return visited
}
