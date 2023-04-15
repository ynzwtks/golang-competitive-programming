package main //トポロジカルソート
//戻り値：ソート結果、Gの有向パスのうち、最長のものの長さ
func topologicalSort(G [][]int) ([]int, int) {
	n := len(G) - 1
	indegrees := make([]int, n+1)
	for i := 0; i < len(G); i++ {
		for j := 0; j < len(G[i]); j++ {
			indegrees[G[i][j]]++
		}
	}
	l := []int{}
	que := []int{}
	cnt := 0
	for i := 1; i <= n; i++ {
		if indegrees[i] == 0 {
			que = append(que, i)
			l = append(l, i)
			indegrees[i] = -1
		}
	}
	for true {
		nq := []int{}
		for _, v := range que {
			for _, v2 := range G[v] {
				indegrees[v2]--
				if indegrees[v2] == 0 {
					nq = append(nq, v2)
					l = append(l, v2)
				}
			}
		}
		if len(nq) == 0 {
			break
		}
		que = nq
		cnt++
	}
	return l, cnt
}
