package main

// オイラーツアー
type EulerTour struct {
	g       [][]int
	inTime  []int
	outTime []int
	route   []int
	depth   []int
	order   map[int]int
	time    int
	root    int
	visited []bool
}

// valTable[inTime[a]:outTime[a]]でノードa配下の値
func NewEulerTour(g [][]int, root int) *EulerTour {
	d := &EulerTour{}
	d.g = g
	d.inTime = make([]int, len(g))
	d.outTime = make([]int, len(g))
	d.depth = make([]int, 0)
	d.order = make(map[int]int)
	d.root = root
	d.time = 0
	d.route = make([]int, 0)
	d.visited = make([]bool, len(g))
	d.dfs(root, 0)
	d.setTimeTable()
	d.visited = make([]bool, len(g))
	return d
}
func (d *EulerTour) setTimeTable() {
	t := make(map[int]bool)
	for i := 0; i < len(d.route); i++ {
		k := d.route[i]
		if t[k] == false {
			t[k] = true
			d.inTime[k] = i
		} else {
			d.outTime[k] = i
		}
	}
}
func (d *EulerTour) dfs(s, depth int) {
	pos := s
	d.visited[pos] = true
	d.route = append(d.route, pos)
	d.depth = append(d.depth, depth)
	d.order[pos] = len(d.depth) - 1
	for i := 0; i < len(d.g[pos]); i++ {
		next := d.g[pos][i]
		if d.visited[next] == false {
			d.dfs(next, depth+1)
		}
	}
	d.route = append(d.route, pos)
}

//ノードを訪れてた時間、別のノードかあら戻ってきた間を返す
func (d *EulerTour) GetRange(x int) (int, int) {
	return d.inTime[x], d.outTime[x]
}

//訪れたノードの順番を返す
func (d *EulerTour) GetRoute() []int {
	return d.route
}

//各ノードの深さを返す
func (d *EulerTour) GetDepth() (map[int]int, []int) {
	return d.order, d.depth
}
