package main

import (
	"bufio"
	"container/list"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

const (
	BUFSIZE = 10000000
	MOD     = 1000000007
)

var rdr *bufio.Reader

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func max(x ...int) int {
	ret := math.MinInt64
	for i := 0; i < len(x); i++ {
		if ret < x[i] {
			ret = x[i]
		}
	}
	return ret
}
func min(x ...int) int {
	ret := math.MaxInt64
	for i := 0; i < len(x); i++ {
		if ret > x[i] {
			ret = x[i]
		}
	}
	return ret
}

func readline() string {
	buf := make([]byte, 0, 16)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			fmt.Println(e.Error())
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return strings.TrimSpace(string(buf))
}

func readint() int {
	return s2i(readline())
}
func readint2() (int, int) {
	lines := strings.Split(readline(), " ")
	return s2i(lines[0]), s2i(lines[1])
}
func readint3() (int, int, int) {
	lines := strings.Split(readline(), " ")
	return s2i(lines[0]), s2i(lines[1]), s2i(lines[2])
}
func readint4() (int, int, int, int) {
	lines := strings.Split(readline(), " ")
	return s2i(lines[0]), s2i(lines[1]), s2i(lines[2]), s2i(lines[3])
}
func s2i(s string) int {
	v, ok := strconv.Atoi(s)
	if ok != nil {
		panic("Faild : " + s + " can't convert to int")
	}
	return v
}

func reverseList(x []int) []int {
	ret := make([]int, 0)
	for i := len(x) - 1; i >= 0; i-- {
		ret = append(ret, x[i])
	}
	return ret
}
func readGrid(h int) [][]byte {
	g := make([][]byte, h)
	for i := 0; i < h; i++ {
		g[i] = []byte(readline())
	}
	return g
}

// グラフ
type Graph struct {
	n    int
	edge [][]int
	cost [][]int
}

func NewGraph(n int) *Graph {
	g := &Graph{}
	g.n = n + 1 // 要素n(1-indexed)のグラフ
	g.edge = make([][]int, g.n)
	g.cost = make([][]int, g.n)
	return g
}
func (g *Graph) AddEdge(a, b int) {
	g.edge[a] = append(g.edge[a], b)
}
func (g *Graph) AddCost(a, b int) {
	g.cost[a] = append(g.cost[a], b)
}
func (g *Graph) ReadSimpleUndirectedGraph(m int) {
	for i := 0; i < m; i++ {
		a, b := readint2()
		g.AddEdge(a, b)
		g.AddEdge(b, a)
		g.AddCost(a, 1)
		g.AddCost(b, 1)
	}
}
func (g *Graph) ReadWightedUndirectedGraph(m int) {
	for i := 0; i < m; i++ {
		a, b, c := readint3()
		g.AddEdge(a, b)
		g.AddEdge(b, a)
		g.AddCost(a, c)
		g.AddCost(b, c)
	}
}
func (g *Graph) readSimpleDirectedGraph(m int) {
	for i := 0; i < m; i++ {
		a, b := readint2()
		g.AddEdge(a, b)
		g.AddCost(a, 1)
	}
}
func (g *Graph) readWightedDirectedGraph(m int) {
	for i := 0; i < m; i++ {
		a, b, c := readint3()
		g.AddEdge(a, b)
		g.AddCost(a, c)
	}
}

type Edge struct {
	source, destination, weight int
}
type Pair struct {
	a, b int
}

func keys(m map[int]int) []int {
	ret := make([]int, 0, len(m))
	for i := range m {
		ret = append(ret, i)
	}
	return ret
}
func values(m map[int]int) []int {
	ret := make([]int, 0, len(m))
	for _, v := range m {
		ret = append(ret, v)
	}
	return ret
}
func reverseSort(a []int) []int {
	sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })
	return a
}

func intSlice3(a, b, c, value int) [][][]int {
	ret := make([][][]int, a)
	for i := 0; i < a; i++ {
		ret[i] = make([][]int, b)
		for j := 0; j < b; j++ {
			ret[i][j] = make([]int, c)
			for k := 0; k < c; k++ {
				ret[i][j][k] = value
			}
		}
	}
	return ret
}
func intSlice2(a, b, value int) [][]int {
	ret := make([][]int, a)
	for i := 0; i < a; i++ {
		ret[i] = make([]int, b)
		for j := 0; j < b; j++ {
			ret[i][j] = value
		}
	}
	return ret
}
func intSlice(a, value int) []int {
	ret := make([]int, a)
	for i := 0; i < a; i++ {
		ret[i] = value
	}
	return ret
}

// Int型Dequeラッパー
type Deque struct {
	deq *list.List
}

func NewDeque() *Deque {
	deq := &Deque{}
	deq.deq = list.New()
	return deq
}
func (deq *Deque) PushBack(n int) {
	deq.deq.PushBack(n)
}
func (deq *Deque) PushFront(n int) {
	deq.deq.PushFront(n)
}
func (deq *Deque) PopBack() int {
	ret := deq.deq.Back().Value.(int)
	deq.deq.Remove(deq.deq.Back())
	return ret
}
func (deq *Deque) PopFront() int {
	ret := deq.deq.Front().Value.(int)
	deq.deq.Remove(deq.deq.Front())
	return ret
}
func (deq *Deque) DumpBack() []int {
	ret := make([]int, 0)
	for deq.deq.Len() != 0 {
		ret = append(ret, deq.deq.Back().Value.(int))
		deq.deq.Remove(deq.deq.Back())
	}
	return ret
}
func (deq *Deque) DumpFront() []int {
	ret := make([]int, 0)
	for deq.deq.Len() != 0 {
		ret = append(ret, deq.deq.Front().Value.(int))
		deq.deq.Remove(deq.deq.Front())
	}
	return ret
}
func (deq *Deque) Back() int {
	return deq.deq.Back().Value.(int)
}
func (deq *Deque) Front() int {
	return deq.deq.Front().Value.(int)
}
func (deq *Deque) Len() int {
	return deq.deq.Len()
}

// int型ヒープ
func NewPQ(prioritySmallest bool) *IntPQ {
	ret := &IntPQ{}
	ret.d = make([]int, 0, 100)
	n := ret.Len()
	for i := n/2 - 1; i >= 0; i-- {
		ret.down(i, n)
	}
	ret.prioritySmallest = prioritySmallest
	return ret
}

type IntPQ struct {
	d                []int
	prioritySmallest bool
}

func (pq *IntPQ) less(i, j int) bool {
	if pq.prioritySmallest == true {
		return pq.d[i] < pq.d[j]
	} else {
		return pq.d[i] > pq.d[j]
	}
}
func (pq *IntPQ) swap(i, j int) { pq.d[i], pq.d[j] = pq.d[j], pq.d[i] }
func (pq *IntPQ) down(i0, n int) bool {
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
func (pq *IntPQ) Push(x int) {
	pq.d = append(pq.d, x)
	pq.up(len(pq.d) - 1)
}
func (pq *IntPQ) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !pq.less(j, i) {
			break
		}
		pq.swap(i, j)
		j = i
	}
}
func (pq *IntPQ) Pop() int {
	n := pq.Len() - 1
	x := pq.d[0]
	pq.swap(0, n)
	pq.down(0, n)
	pq.d = pq.d[0 : pq.Len()-1]
	return x
}
func (pq *IntPQ) Len() int {
	return len(pq.d)
}
func (pq *IntPQ) Peek() int {
	return pq.d[0]
}
func (pq *IntPQ) Remove(i int) int {
	n := pq.Len() - 1
	if n != i {
		pq.swap(i, n)
		if !pq.down(i, n) {
			pq.up(i)
		}
	}
	return pq.Pop()
}
func (pq *IntPQ) Fix(i int) {
	if !pq.down(i, pq.Len()) {
		pq.up(i)
	}
}
