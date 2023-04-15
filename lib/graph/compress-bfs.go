package main

import "container/list"

// CompressBFSは、座標圧縮版の幅優先探索 (BFS) を実行する関数です。
// 入力としてグラフg（map[int][]int）と始点beginを受け取り、
// 各ノードに到達できるかどうかを表すvisitedマップと、経路情報を格納したprevマップを返します。
func compressBFS(g map[int][]int, begin int) (map[int]int, map[int]int) {
	visited := make(map[int]int)
	prev := make(map[int]int)

	deq := list.New()
	deq.PushBack(begin)
	visited[begin] = 0
	prev[begin] = -1

	for deq.Len() != 0 {
		pos := deq.Front().Value.(int)
		deq.Remove(deq.Front())
		for _, v := range g[pos] {
			_, ok := visited[v]
			if ok == false {
				visited[v] = visited[pos] + 1
				deq.PushBack(v)
				prev[v] = pos
			}
		}
	}
	return visited, prev
}

// restoreCompressedBFSPathは、CompressBFSで得られた経路情報からパスを復元する関数です。
// 入力として経路情報マップp、始点begin、終点endを受け取り、
// 復元されたパス（[]int）を返します。
func restoreCompressedBFSPath(p map[int]int, begin, end int) []int {
	ret := make([]int, 0)
	cur := end
	ret = append(ret, cur)

	for cur != -1 && cur != begin {
		ret = append(ret, p[cur])
		cur = p[cur]
	}
	ret = reverseList(ret)
	return ret
}
