package main

import (
	"bufio"
	"container/list"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	BUFSIZE = 10000000
	MOD     = 1000000007
)

var rdr *bufio.Reader

func main() {
	rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
	solve()
}

func solve() {

}

// 汎用関数
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
func printList(l []int) {
	s := make([]string, 0)
	for i := 0; i < len(l); i++ {
		s = append(s, strconv.Itoa(l[i]))
	}
	fmt.Println(strings.Join(s, " "))
}
func sumList(l []int) int {
	sum := 0
	for i := 0; i < len(l); i++ {
		sum += l[i]
	}
	return sum
}

// 標準入出力
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
func readIntSliceBanpei() []int {
	slice := make([]int, 0)
	slice = append(slice, 0)
	lines := strings.Split(readline(), " ")
	for _, v := range lines {
		slice = append(slice, s2i(v))
	}
	slice = append(slice, math.MaxInt64)
	return slice
}
func readIntSlice() []int {
	slice := make([]int, 0)
	lines := strings.Split(readline(), " ")
	for _, v := range lines {
		slice = append(slice, s2i(v))
	}
	return slice
}
func readFloatSlice() []float64 {
	slice := make([]float64, 0)
	lines := strings.Split(readline(), " ")
	for _, v := range lines {
		slice = append(slice, s2f(v))
	}
	return slice
}
func readint() int {
	return s2i(readline())
}
func readfloat() float64 {
	return s2f(readline())
}
func readint2() (int, int) {
	lines := strings.Split(readline(), " ")
	return s2i(lines[0]), s2i(lines[1])
}
func readfloat2() (float64, float64) {
	lines := strings.Split(readline(), " ")
	return s2f(lines[0]), s2f(lines[1])
}
func readint3() (int, int, int) {
	lines := strings.Split(readline(), " ")
	return s2i(lines[0]), s2i(lines[1]), s2i(lines[2])
}
func readfloat3() (float64, float64, float64) {
	lines := strings.Split(readline(), " ")
	return s2f(lines[0]), s2f(lines[1]), s2f(lines[2])
}
func readint4() (int, int, int, int) {
	lines := strings.Split(readline(), " ")
	return s2i(lines[0]), s2i(lines[1]), s2i(lines[2]), s2i(lines[3])
}
func readfloat4() (float64, float64, float64, float64) {
	lines := strings.Split(readline(), " ")
	return s2f(lines[0]), s2f(lines[1]), s2f(lines[2]), s2f(lines[3])
}
func s2i(s string) int {
	v, ok := strconv.Atoi(s)
	if ok != nil {
		panic("Faild : " + s + " can't convert to int")
	}
	return v
}
func s2f(s string) float64 {
	v, ok := strconv.ParseFloat(s, 64)
	if ok != nil {
		panic("Faild : " + s + " can't convert to float")
	}
	return v
}
func readSimpleUndirectedGraph(n, m int) [][]int {
	g := make([][]int, n+1)
	for i := 0; i < m; i++ {
		a, b := readint2()
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	return g
}
func readWightedUndirectedGraph(n, m int) ([][]int, [][]int) {
	g := make([][]int, n+1)
	w := make([][]int, n+1)
	for i := 0; i < m; i++ {
		a, b, c := readint3()
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
		w[a] = append(w[a], c)
		w[b] = append(w[b], c)

	}
	return g, w
}
func readSimpleDirectedGraph(n, m int) [][]int {
	g := make([][]int, n+1)
	for i := 0; i < m; i++ {
		a, b := readint2()
		g[a] = append(g[a], b)
	}
	return g
}
func readWightedDirectedGraph(n, m int) ([][]int, [][]int) {
	g := make([][]int, n+1)
	w := make([][]int, n+1)
	for i := 0; i < m; i++ {
		a, b, c := readint3()
		g[a] = append(g[a], b)
		w[a] = append(w[a], c)
	}
	return g, w
}
func readGrid(h int) [][]byte {
	g := make([][]byte, h)
	for i := 0; i < h; i++ {
		g[i] = []byte(readline())
	}
	return g
}

// 2分探索
func UpperBound(array []int, target int) int {
	low, high, mid := 0, len(array)-1, 0
	for low <= high {
		mid = (low + high) / 2
		if array[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}
func LowerBound(array []int, target int) int {
	low, high, mid := 0, len(array)-1, 0
	for low <= high {
		mid = (low + high) / 2
		if array[mid] >= target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

// 変換
//連続している要素を始点・終点のリストに変換
func splitGroup(a []int) [][]int {
	if len(a) == 0 {
		return nil
	}
	if len(a) == 1 {
		return [][]int{{a[0], a[0]}}
	}
	sort.Ints(a)
	start := a[0]
	end := a[0]
	ret := make([][]int, 0)
	for i := 1; i < len(a); i++ {
		if a[i]-a[i-1] > 1 {
			ret = append(ret, []int{start, end})
			start = a[i]
			end = a[i]
		} else {
			end = a[i]
		}
	}
	ret = append(ret, []int{start, end})
	return ret
}

// スライスの中身をユニークな要素に圧縮
func uniqList(a []int) []int {
	m := make(map[int]int)
	for i := 0; i < len(a); i++ {
		m[a[i]] = 1
	}
	return keys(m)
}

//スライスの指定した始点から終点の順序を入れ替え
func rangeReverse(start, end int, l *[]int) {
	for i := 0; i < (end-start+1)/2; i++ {
		(*l)[start+i], (*l)[end-i] = (*l)[end-i], (*l)[start+i]
	}
}
func mtol(m map[int]int) [][]int {
	t := make([][]int, len(m))
	idx := 0
	for i, v := range m {
		t[idx] = []int{i, v}
		idx++
	}
	sort.Slice(t, func(i, j int) bool { return t[i][0] < t[j][0] })
	return t
}
func ltom(l []int) map[int]int {
	m := make(map[int]int)
	for i := 0; i < len(l); i++ {
		m[l[i]]++
	}
	return m
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
func reverseList(x []int) []int {
	ret := make([]int, 0)
	for i := len(x) - 1; i >= 0; i-- {
		ret = append(ret, x[i])
	}
	return ret
}
func reverseBytes(x []byte) []byte {
	ret := make([]byte, 0)
	for i := len(x) - 1; i >= 0; i-- {
		ret = append(ret, x[i])
	}
	return ret
}
func intSlice3(a, b, c int) [][][]int {
	ret := make([][][]int, a)
	for i := 0; i < a; i++ {
		ret[i] = make([][]int, b)
		for j := 0; j < b; j++ {
			ret[i][j] = make([]int, c)
		}
	}
	return ret
}
func intSlice2(a, b int) [][]int {
	ret := make([][]int, a)
	for i := 0; i < a; i++ {
		ret[i] = make([]int, b)
	}
	return ret
}
func lis(a []int) ([]int, []int) {
	lis := make([]int, 0)
	lisLen := make([]int, len(a))
	lis = append(lis, a[0])
	lisLen[0] = len(lis)
	for i := 1; i < len(a); i++ {
		ret := LowerBound(lis, a[i])
		if ret == len(lis) {
			lis = append(lis, a[i])
		} else {
			lis[ret] = a[i]
		}
		lisLen[i] = len(lis)
	}
	return lis, lisLen
}

// 整数関連
func lcm(a, b int) int {
	return (a / gcd(a, b)) * b
}
func gcd(a, b int) int {
	c := 1
	for c != 0 {
		c = a % b
		a, b = b, c
	}
	return a
}

// 拡張ユークリッド互助法
func exgcd(a, b int) (int, int) {
	x1, y1, m1 := 1, 0, a
	x2, y2, m2 := 0, 1, b
	for b != 0 {
		t := a / b
		if m1 >= m2 {
			x1 -= x2 * t
			y1 -= y2 * t
			m1 -= m2 * t
		} else {
			x2 -= x1 * t
			y2 -= y1 * t
			m2 -= m1 * t
		}
		a, b = b, a%b
		if m1 == 1 {
			return x1, y1
		}
	}
	return x2, y2
}
func divisorList(n int) []int {
	var l []int
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			l = append(l, i)
			if i != n/i {
				l = append(l, n/i)
			}
		}
	}
	sort.Slice(l, func(i, j int) bool { return l[i] < l[j] })
	return l
}
func factorization(n int) map[int]int {
	m := make(map[int]int)
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			m[i]++
			n = n / i
		}
	}
	if n != 0 && n != 1 {
		m[n]++
	}
	return m
}
func factorizationLarge(n int) map[int]int {
	var f func(x, a, n *big.Int) *big.Int
	var recursive func(k int)
	var rho func(n *big.Int, g func(*big.Int, *big.Int, *big.Int) *big.Int) (*big.Int, error)
	ret := make(map[int]int)

	f = func(x, a, n *big.Int) *big.Int {
		p := big.NewInt(2)
		x2 := new(big.Int).Exp(x, p, n)
		x2.Add(x2, a)
		x2.Mod(x2, n)
		return x2
	}
	rho = func(n *big.Int, g func(*big.Int, *big.Int, *big.Int) *big.Int) (*big.Int, error) {
		rand.Seed(time.Now().UnixNano())
		one := big.NewInt(1)
		two := big.NewInt(2)
		t := int(n.Int64())
		r1 := int64(rand.Intn(t))
		r2 := 1 + int64(rand.Intn(t-3))

		x, y, a, d := big.NewInt(r1), big.NewInt(r1), big.NewInt(r2), big.NewInt(1)
		if n.Int64()%2 == 0 {
			return two, nil
		}
		for d.Cmp(one) == 0 {
			x = g(x, a, n)
			y = g(g(y, a, n), a, n)
			sub := new(big.Int).Sub(x, y)
			d.GCD(nil, nil, sub.Abs(sub), n)
		}
		if d.Cmp(n) == 0 {
			return rho(n, f)
		}
		return d, nil
	}
	recursive = func(n int) {
		k := big.NewInt(int64(n))
		factor, _ := rho(k, f)
		if factor.ProbablyPrime(0) == true {
			ret[int(factor.Int64())]++
		} else {
			recursive(int(factor.Int64()))
		}
		n = n / int(factor.Int64())
		k = big.NewInt(int64(n))
		if k.ProbablyPrime(0) == true {
			ret[int(k.Int64())]++
		} else {
			recursive(int(k.Int64()))
		}
	}
	k := big.NewInt(int64(n))
	if k.ProbablyPrime(0) == true || n == 1 {
		ret[n]++
		return ret
	}
	recursive(n)
	return ret
}

type FastFactorization struct {
	n  int
	p  []int
	rp []int
}

func NewFastFactorization(n int) *FastFactorization {
	ff := &FastFactorization{}
	ff.n = n
	ff.rp = make([]int, n+1)
	ff.p = primeList(ff.n)
	ff.rp[1] = 1
	for _, v := range ff.p {
		for i := v; i <= ff.n; i += v {
			ff.rp[i] = v
		}
	}
	return ff
}
func (ff *FastFactorization) Factorize(n int) map[int]int {
	t := n
	m := make(map[int]int)
	for t != 1 {
		m[ff.rp[t]]++
		t = t / ff.rp[t]
	}
	return m
}
func isPrime(n int) bool {
	t := big.NewInt(int64(n))
	return t.ProbablyPrime(0)
}
func primeList(n int) []int {
	if n < 2 {
		return nil
	} else if n == 2 {
		return []int{2}
	}
	l := make([]int, n+1)
	primes := make([]int, 0, n)
	for i := 4; i <= n; i += 2 {
		l[i] = 1
	}
	for i := 3; i <= n; i += 2 {
		if l[i] == 1 {
			continue
		}
		for j := i * 2; j <= n; j += i {
			l[j] = 1
		}
	}
	primes = append(primes, 2)
	for i := 3; i <= n; i += 2 {
		if l[i] == 0 {
			primes = append(primes, i)
		}
	}
	return primes
}
func combination(n, r int) int {
	r = min(n-r, r)
	d := make([]int, r)

	for i := 0; i < r; i++ {
		d[i] = n - i
	}
	for i := 2; i <= r; i++ {
		ti := i
		for j := 0; j < r; j++ {
			g := gcd(d[j], ti)
			if g != 1 {
				ti /= g
				d[j] /= g
			}
			if ti == 1 {
				break
			}
		}
	}
	ret := 1
	for i := 0; i < r; i++ {
		ret *= d[i]
	}
	return ret
}
func combinationMod(a, b, mod int) int {
	t1, t2 := 1, 1
	for i := 2; i <= b; i++ {
		t2 = (i * t2) % mod
	}
	gyaku := powMod(t2, mod-2, mod)

	for i := a - b + 1; i <= a; i++ {
		t1 = (i * t1) % mod
	}
	return (t1 * gyaku) % mod
}
func pow(a, b int) int {
	if b == 0 {
		return 1
	}
	p, q := a, b
	t := 1
	for q >= 2 {
		if q%2 == 0 {
			p = p * p
			q = q / 2
		} else {
			t = t * p
			p = p * p
			q = (q - 1) / 2
		}
	}
	return t * p
}
func powMod(a, b, mod int) int {
	if b == 0 {
		return 1
	}
	p, q := a, b
	t := 1
	for q >= 2 {
		if q%2 == 0 {
			p = (p * p) % mod
			q = q / 2
		} else {
			t = (t * p) % mod
			p = (p * p) % mod
			q = (q - 1) / 2
		}
	}
	return (t * p) % mod
}
func decimalToBaseN(s string, base int) string {
	n, _ := strconv.Atoi(s)
	t := strconv.FormatInt(int64(n), base)
	return t
}
func baseNToDecimal(s string, base int) string {
	t, _ := strconv.ParseInt(s, base, 64)
	ret := strconv.Itoa(int(t))
	return ret
}

// ビット操作
func mkBitList(x, keta int) []int {
	l := make([]int, keta)
	idx := 0
	for x != 0 {
		l[idx] = x & 1
		x = x >> 1
		idx++
	}
	return l
}

// Deque Wrapper(Int)
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

// Heap
// int型専用ヒープ
// 標準のcontainer/heapではinterface型のオーバヘッドで大量のpush,popだと遅くなるのでint専用で高速化。
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

// グラフ
type UnionFind struct {
	root  []int
	size  []int
	group int
}

func newUnionFind(n int) UnionFind {
	root := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		root[i] = i
		size[i] = 1
	}
	uf := UnionFind{root: root, size: size, group: n}
	return uf
}
func (uf *UnionFind) Union(p int, q int) {
	qRoot := uf.Root(q)
	pRoot := uf.Root(p)
	if qRoot == pRoot {
		return
	}
	uf.group--
	if uf.size[qRoot] < uf.size[pRoot] {
		uf.root[qRoot] = uf.root[pRoot]
		uf.size[pRoot] += uf.size[qRoot]
	} else {
		uf.root[pRoot] = uf.root[qRoot]
		uf.size[qRoot] += uf.size[pRoot]
	}
}
func (uf *UnionFind) Root(p int) int {
	if p > len(uf.root)-1 {
		return -1
	}
	for uf.root[p] != p {
		uf.root[p] = uf.root[uf.root[p]]
		p = uf.root[p]
	}
	return p
}
func (uf *UnionFind) find(p int) int {
	return uf.Root(p)
}
func (uf *UnionFind) Connected(p int, q int) bool {
	return uf.Root(p) == uf.Root(q)
}
func (uf *UnionFind) Groups() map[int]int {
	cm := make(map[int]int)
	for i := 0; i < len(uf.root); i++ {
		t := uf.find(uf.root[i])
		cm[t]++
	}
	return cm
}

// 数え上げ
//x未満、x超え、x以下、x以上の要素数をカウントする
func CountLessMoreBelowOver(l []int, x int) (int, int, int, int) {
	idx := LowerBound(l, x)
	idx2 := UpperBound(l, x)
	ret1, ret2, ret3, ret4 := 0, 0, 0, 0
	ret1 = idx
	ret2 = len(l) - idx2
	if idx == len(l) {
		return len(l), 0, len(l), 0
	}
	if l[idx] == x {
		ret3 = idx2
		ret4 = len(l) - idx
	} else {
		ret3 = ret1
		ret4 = ret2
	}
	ret4 = len(l) - idx
	return ret1, ret2, ret3, ret4
}

//a[i]の中に連続してx以上の連続する要素が何グループあるか
func countGroup(a []int, x int) int {
	cnt := 0
	t := 0
	for i := 0; i < len(a); i++ {
		if a[i]-x >= 0 {
			t++
			continue
		} else if t != 0 {
			cnt++
			t = 0
		}
	}
	if t != 0 {
		cnt++
	}
	return cnt
}
