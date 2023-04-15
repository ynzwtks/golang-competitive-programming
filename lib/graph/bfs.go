package main

import "container/list"

// BFSは、グラフgと始点beginを入力として、各ノードへの最短距離（到達不可の場合は-1）の配列visitedと
// 経路情報の配列prevを返します。
func bfs(g [][]int, begin int) ([]int, []int) {
	n := len(g)
	visited := make([]int, n)
	prev := make([]int, n)

	for i := 0; i < n; i++ {
		visited[i] = -1
	}
	deq := list.New()
	deq.PushBack(begin)
	visited[begin] = 0
	prev[begin] = -1

	for deq.Len() != 0 {
		pos := deq.Front().Value.(int)
		deq.Remove(deq.Front())
		for _, v := range g[pos] {
			if visited[v] == -1 {
				visited[v] = visited[pos] + 1
				deq.PushBack(v)
				prev[v] = pos
			}
		}
	}
	return visited, prev
}

// restorePathは、BFSの経路情報配列p、始点begin、終点endを入力として、
// 最短経路上のノードの配列を返します。
func restoreBfsPath(p []int, begin, end int) []int {
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

// CreateBFSGroupは、グラフgとノード数nを入力として、
// 連結成分ごとに分類されたグループ情報を返します。
// キーがグループIDで、値がグループ内のノード数です。
func createBfsGroup(g [][]int, n int) map[int]int {
	group := make([]int, n+1)
	for i := 0; i <= n; i++ {
		group[i] = -1
	}
	for i := 0; i <= n; i++ {
		if group[i] == -1 {
			ret, _ := bfs(g, i)
			for j := 0; j <= n; j++ {
				if ret[j] != -1 && group[j] == -1 {
					group[j] = i
				}
			}
		}
	}
	gt := make(map[int]int)
	for i := 0; i < len(group); i++ {
		gt[group[i]]++
	}
	return gt
}
