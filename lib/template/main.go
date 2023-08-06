package main

import (
	"bufio"
	"container/list"
	"fmt"
	"math"
	"math/big"
	"math/bits"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

const (
	BUFSIZE                = 10000000
	MOD2305843009213693951 = 2305843009213693951
	MOD1000000007          = 1000000007
	MOD998244353           = 998244353
	MOD                    = MOD998244353
	INF                    = 1 << 60
)

var rdr *bufio.Scanner
var wtr *bufio.Writer
var p10 []int

func solve() {

}
func flush() { wtr.Flush() }
func main() {
	defer flush()
	rdr = bufio.NewScanner(os.Stdin)
	rdr.Buffer(make([]byte, 4096), math.MaxInt64)
	rdr.Split(bufio.ScanWords)
	wtr = bufio.NewWriterSize(os.Stdout, BUFSIZE)
	p10 = []int{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000, 1000000000, 10000000000, 100000000000, 1000000000000, 10000000000000, 100000000000000, 1000000000000000, 10000000000000000, 100000000000000000, 1000000000000000000}
	solve()

}

// 構造体定義
type Data struct{ start, end, val int }
type Pair struct{ a, b int }
type Item struct{ key, num int }
type Point struct{ y, x float64 }
type Graph struct {
	n     int
	edges [][]Edge
}
type Edge struct {
	from, to, w int
}
type FenwickTree struct {
	data []int
	len  int
}
type Bit2D struct {
	data       [][]int
	lenX, lenY int
}
type UnionFind struct {
	root  []int
	size  []int
	group int
}
type Deque struct {
	deq *list.List
	cur *list.Element
	sum int
}
type DataDeque struct {
	deq *list.List
	cur *list.Element
}
type IntPQ struct {
	d                []int
	prioritySmallest bool
}
type ItemPQ struct {
	d                []Item
	prioritySmallest bool
}
type Frac struct {
	top    int
	bottom int
}
type SegInt struct {
	data []int
	e    int
	op   func(int, int) int
	len  int
}
type BitSets []big.Int

// 出力
func out(a ...interface{}) {
	str := fmt.Sprintln(a...)
	_, err := wtr.WriteString(str)
	if err != nil {
		return
	}
	return
}
func outf(f float64) {
	outfmt("%0.9f\n", f)
}
func outn(x int) {
	outfmt("%d ", x)
}
func outfn(f float64) {
	outfmt("%0.9f ", f)
}
func outfmt(format string, a ...interface{}) (int, error) {
	str := fmt.Sprintf(format, a...)
	return wtr.WriteString(str)
}
func outListX(l []int) {
	s := make([]string, 0, len(l))
	for i := 0; i < len(l); i++ {
		s = append(s, strconv.Itoa(l[i]))
	}
	write(strings.Join(s, " "))
}
func outListY(l []int) {
	s := make([]string, 0, len(l))
	for i := 0; i < len(l); i++ {
		s = append(s, strconv.Itoa(l[i]))
	}
	write(strings.Join(s, "\n"))
}
func outYN(x bool) {
	if x == true {
		out("Yes")
	} else {
		out("No")
	}
}
func write(s string) {
	_, err := wtr.WriteString(s)
	if err != nil {
		return
	}
	_, err = wtr.WriteString("\n")
	if err != nil {
		return
	}
	return
}
func outGrid(g [][]byte) {
	for i := 0; i < len(g); i++ {
		out(string(g[i]))
	}
}

// 入力
func rs() string { rdr.Scan(); return rdr.Text() }
func ri() int {
	rdr.Scan()
	i, e := strconv.Atoi(rdr.Text())
	if e != nil {
		panic(e)
	}
	return i
}
func ri2() (int, int) {
	a := ri()
	b := ri()
	return a, b
}
func ri3() (int, int, int) {
	a := ri()
	b := ri()
	c := ri()
	return a, b, c
}
func ri4() (int, int, int, int) {
	a := ri()
	b := ri()
	c := ri()
	d := ri()
	return a, b, c, d
}
func ris(n int) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = ri()
	}
	return res
}
func riy(n int) []int {
	ret := make([]int, 0)
	for i := 0; i < n; i++ {
		ret = append(ret, ri())
	}
	return ret
}
func risapp(a []int, n int) []int {
	for i := 0; i < n; i++ {
		a = append(a, ri())
	}
	return a
}
func ris2(n int) ([]int, []int) {
	res := make([]int, n)
	res2 := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = ri()
		res2[i] = ri()
	}
	return res, res2
}
func rf() float64 {
	f, e := strconv.ParseFloat(rs(), 64)
	if e != nil {
		panic(e)
	}
	return f
}
func rfs(n int) []float64 {
	res := make([]float64, n)
	for i := 0; i < n; i++ {
		res[i] = rf()
	}
	return res
}
func rss(n int) []string {
	res := make([]string, n)
	for i := 0; i < n; i++ {
		res[i] = rs()
	}
	return res
}
func readIntLines(h, w int) [][]int {
	l := make([][]int, h)
	for i := 0; i < h; i++ {
		l[i] = ris(w)
	}
	return l
}
func readGrid(h int) [][]byte {
	g := make([][]byte, h)
	for i := 0; i < h; i++ {
		g[i] = []byte(rs())
	}
	return g
}
func readEdges(m int, readWeight bool) []Edge {
	ret := make([]Edge, m)
	for i := 0; i < m; i++ {
		a, b := ri2()
		ret[i].from = a
		ret[i].to = b
		ret[i].w = 1
		if readWeight == true {
			ret[i].w = ri()
		}
	}
	return ret
}
func readPairs(m int) []Pair {
	ret := make([]Pair, m)
	for i := 0; i < m; i++ {
		a, b := ri2()
		ret[i] = Pair{a, b}
	}
	return ret
}

// 汎用関数
func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}
func maxf(i float64, j float64) float64 {
	if i > j {
		return i
	}
	return j
}
func min(i int, j int) int {
	if i < j {
		return i
	}
	return j
}
func minf(i float64, j float64) float64 {
	if i < j {
		return i
	}
	return j
}
func chmin(a *int, b int) {
	if *a > b {
		*a = b
	}
}
func chmax(a *int, b int) {
	if *a < b {
		*a = b
	}
}
func isInRange(x, low, high int) bool {
	return low <= x && x <= high
}
func isOutRange(x, low, high int) bool {
	return !isInRange(x, low, high)
}
func cs(a []int) []int {
	ret := make([]int, len(a))
	copy(ret, a)
	return ret
}
func maxSlice(l []int) int {
	ret := math.MinInt64
	for i := 0; i < len(l); i++ {
		if ret < l[i] {
			ret = l[i]
		}
	}
	return ret
}
func minSlice(l []int) int {
	ret := math.MaxInt64
	for i := 0; i < len(l); i++ {
		if ret > l[i] {
			ret = l[i]
		}
	}
	return ret
}
func maxSliceIdx(l []int) int {
	cur := math.MinInt64
	idx := 0
	for i := 0; i < len(l); i++ {
		if cur < l[i] {
			cur = l[i]
			idx = i
		}
	}
	return idx
}
func minSliceIdx(l []int) int {
	cur := math.MaxInt64
	idx := 0
	for i := 0; i < len(l); i++ {
		if cur > l[i] {
			cur = l[i]
			idx = i
		}
	}
	return idx
}
func maxMapKey(m map[int]int) int {
	tm, key := -INF, 0
	for i, v := range m {
		if v > tm {
			tm = v
			key = i
		}
	}
	return key
}
func minMapKey(m map[int]int) int {
	tm, key := INF, 0
	for i, v := range m {
		if v < tm {
			tm = v
			key = i
		}
	}
	return key
}
func maxMapVal(m map[int]int) int {
	t := -INF
	for _, v := range m {
		chmax(&t, v)
	}
	return t
}
func minMapVal(m map[int]int) int {
	t := INF
	for _, v := range m {
		chmin(&t, v)
	}
	return t
}
func maxvar(x ...int) int { return maxSlice(x) }
func minvar(x ...int) int { return minSlice(x) }
func condInt(a bool, b, c int) int {
	if a == true {
		return b
	} else {
		return c
	}
}
func condBool(a, b, c bool) bool {
	if a == true {
		return b
	} else {
		return c
	}
}
func condStr(a bool, b, c string) string {
	if a == true {
		return b
	} else {
		return c
	}
}
func divCeil(a, b int) int { return (a + b - 1) / b }

// Grid操作
func gridRangeChecker(aMin, aMax, bMin, bMax int) func(Pair) bool {
	f := func(p Pair) bool {
		return isInRange(p.a, aMin, aMax) && isInRange(p.b, bMin, bMax)
	}
	return f
}
func createGrid(h, w int, ch byte) [][]byte {
	ret := make([][]byte, h)
	for i := 0; i < h; i++ {
		ret[i] = make([]byte, w)
		if ch == 0 {
			continue
		}
		for j := 0; j < w; j++ {
			ret[i][j] = ch
		}
	}
	return ret
}
func rotateGrid(g [][]byte, rotation int, flipY, flipX bool) [][]byte {
	h := len(g)
	w := len(g[0])
	var ret [][]byte
	if rotation == 1 || rotation == 3 {
		ret = createGrid(w, h, 0)
	} else {
		ret = createGrid(h, w, 0)
	}
	conv := gridPosConv(h, w, rotation, flipY, flipX)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			ch := g[i][j]
			ni, nj := conv(i, j)
			ret[ni][nj] = ch
		}
	}
	return ret
}
func gridPosConv(maxY, maxX, rotation int, flipY, flipX bool) func(int, int) (int, int) {
	return func(y, x int) (int, int) {
		if flipY {
			y = maxY - 1 - y
		}
		if flipX {
			x = maxX - 1 - x
		}
		switch rotation % 4 {
		case 1:
			y, x = x, maxY-1-y
		case 2:
			y, x = maxY-1-y, maxX-1-x
		case 3:
			y, x = maxX-1-x, y
		}
		return y, x
	}
}
func clipGrid(grid [][]byte, y, x, height, width int) [][]byte {
	if y < 0 || x < 0 || y+height > len(grid) || x+width > len(grid[0]) {
		panic("The provided range is out of grid bounds.")
	}
	clipped := make([][]byte, height)
	for i := range clipped {
		clipped[i] = make([]byte, width)
	}
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			clipped[i][j] = grid[y+i][x+j]
		}
	}
	return clipped
}
func pasteGrid(original, edited [][]byte, y, x int) {
	if y < 0 || x < 0 || y+len(edited) > len(original) || x+len(edited[0]) > len(original[0]) {
		panic("The provided range is out of original grid bounds.")
	}
	for i := range edited {
		for j := range edited[i] {
			original[y+i][x+j] = edited[i][j]
		}
	}
}
func flipSubGrid(original [][]byte, y, x, height, width int, flipY, flipX bool) {
	extracted := clipGrid(original, y, x, height, width)
	flipped := rotateGrid(extracted, 0, flipY, flipX)
	pasteGrid(original, flipped, y, x)
}
func isSubGridMatch(grid [][]byte, y, x int, subgrid [][]byte, eval func(byte, byte) bool) bool {
	//eval例:
	//func(a,b byte)bool{return  a == b || a == '?' || b == '?'}
	if y+len(subgrid) > len(grid) || x+len(subgrid[0]) > len(grid[0]) {
		return false
	}
	for i := range subgrid {
		for j := range subgrid[i] {
			if !eval(grid[y+i][x+j], subgrid[i][j]) {
				return false
			}
		}
	}
	return true
}
func strToGrid(s []string) [][]byte {
	ret := make([][]byte, len(s))
	for i := 0; i < len(s); i++ {
		ret[i] = []byte(s[i])
	}
	return ret
}
func addGridBanpei(g [][]byte, c byte) [][]byte {
	head := make([]byte, len(g[0])+2)
	tail := make([]byte, len(g[0])+2)
	for i := 0; i < len(g[0])+2; i++ {
		head[i] = c
		tail[i] = c
	}
	ret := make([][]byte, len(g)+2)
	ret[0] = head
	ret[len(g)+1] = tail
	for i := 1; i <= len(g); i++ {
		t := g[i-1]
		t = append([]byte{c}, t...)
		t = append(t, c)
		ret[i] = t
	}
	return ret
}

// Mod
func modAdd(a, b int) int { return (a + b) % MOD }
func modSub(a, b int) int { return (((a - b) % MOD) + MOD) % MOD }
func modMul(a, b int) int { return (a * b) % MOD }
func modInv(x int) int {
	a, _ := exgcd(x, MOD)
	return (a + MOD) % MOD
}
func sumMod(a []int) int {
	ret := 0
	for i := 0; i < len(a); i++ {
		ret = modAdd(ret, a[i])
	}
	return ret
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
func rangeCount(a []int, nmin, nmax int) int {
	return UpperBound(a, nmax) - LowerBound(a, nmin)
}
func searchNearestNum(a []int, x int) int {
	k := LowerBound(a, x)
	if k == len(a) {
		return a[k-1]
	}
	if k == 0 {
		return a[0]
	}
	if abs(a[k]-x) > abs(a[k-1]-x) {
		return a[k-1]
	} else {
		return a[k]
	}
}
func countIntervalLen(a []int, x, lowerLimit, higherLimit int) int {
	idx := UpperBound(a, x)
	low := lowerLimit
	high := higherLimit
	if idx != 0 {
		low = a[idx-1]
	}
	if idx != len(a) {
		high = a[idx]
	}
	return high - low - 1
} // 昇順ソート済みの[]intについて指定した数値の前後要素の区間を求める
func searchPairsRange(p []Pair, l, r int) (int, int) {
	ret := sort.Search(len(p), func(i int) bool {
		return l < p[i].a
	})
	if ret != 0 && l < p[ret-1].b {
		ret--
	}
	ret2 := sort.Search(len(p), func(i int) bool {
		return r < p[i].b
	})
	if ret2 != len(p) && p[ret2].a > r {
		ret2--
	}
	chmax(&ret, 0)
	chmax(&ret2, 0)
	chmin(&ret, len(p)-1)
	chmin(&ret2, len(p)-1)
	chmin(&ret, ret2)
	return ret, ret2
} // 昇順の重複のない区間([]Pair)について区間[l,r]が含まれるindexの範囲を返す

// スライス・map操作系
func sumIntSlice(l []int) int {
	s := 0
	for i := 0; i < len(l); i++ {
		s += l[i]
	}
	return s
}
func slice2map(l []int) map[int]int {
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
	ret := make([]int, 0, len(x))
	for i := len(x) - 1; i >= 0; i-- {
		ret = append(ret, x[i])
	}
	return ret
}
func reverseBytes(x []byte) []byte {
	ret := make([]byte, 0, len(x))
	for i := len(x) - 1; i >= 0; i-- {
		ret = append(ret, x[i])
	}
	return ret
}
func haveSameKeys(map1, map2 interface{}) bool {
	val1, val2 := reflect.ValueOf(map1), reflect.ValueOf(map2)
	if val1.Kind() != reflect.Map || val2.Kind() != reflect.Map {
		panic("Both parameters must be maps")
	}
	if val1.Len() != val2.Len() {
		return false
	}
	for _, key := range val1.MapKeys() {
		if val2.MapIndex(key).Kind() == reflect.Invalid {
			return false
		}
	}
	return true
}
func intSliceCnt(a []int, maxRange int) []int {
	ret := intSlice(maxRange, 0)
	for i := 0; i < len(a); i++ {
		ret[a[i]]++
	}
	return ret
}
func defaultIntMap(d int) (map[int]int, func(int) int) {
	defaultVal := d
	m := make(map[int]int)
	accessor := func(x int) int {
		val, ok := m[x]
		if ok == true {
			return val
		} else {
			return defaultVal
		}
	}
	return m, accessor
}
func transpose(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return matrix
	}
	rows := len(matrix)
	cols := len(matrix[0])
	transposeMatrix := make([][]int, cols)
	for i := 0; i < cols; i++ {
		transposeMatrix[i] = make([]int, rows)
		for j := 0; j < rows; j++ {
			transposeMatrix[i][j] = matrix[j][i]
		}
	}
	return transposeMatrix
}
func removeElements(s []int, start, end int) []int {
	if start < 0 || start >= len(s) || end < 0 || end >= len(s) || start > end {
		return nil
	}
	return append(s[:start], s[end+1:]...)
} // スライスの[begin,end]の区間を削除する
func insertSliceAt(index int, a []int, b []int) []int {
	if index < 0 || index > len(a) {
		out("Error: index out of range")
		return nil
	}
	a = append(a, make([]int, len(b))...)
	copy(a[index+len(b):], a[index:])
	copy(a[index:], b)
	return a
} // スライスaのindexの位置にbを挿入する
func intSlice(a, value int) []int {
	ret := make([]int, a)
	if value == 0 {
		return ret
	}
	for i := 0; i < a; i++ {
		ret[i] = value
	}
	return ret
}
func intSlice2(a, b, value int) [][]int {
	ret := make([][]int, a)
	for i := 0; i < a; i++ {
		ret[i] = make([]int, b)
		if value == 0 {
			continue
		}
		for j := 0; j < b; j++ {
			ret[i][j] = value
		}
	}
	return ret
}
func intSlice3(a, b, c, value int) [][][]int {
	ret := make([][][]int, a)
	for i := 0; i < a; i++ {
		ret[i] = make([][]int, b)
		for j := 0; j < b; j++ {
			ret[i][j] = make([]int, c)
			if value == 0 {
				continue
			}
			for k := 0; k < c; k++ {
				ret[i][j][k] = value
			}
		}
	}
	return ret
}
func floatSlice(a int, value float64) []float64 {
	ret := make([]float64, a)
	if value == 0 {
		return ret
	}
	for i := 0; i < a; i++ {
		ret[i] = value
	}
	return ret
}
func floatSlice2(a, b int, value float64) [][]float64 {
	ret := make([][]float64, a)
	for i := 0; i < a; i++ {
		ret[i] = make([]float64, b)
		if value == 0 {
			continue
		}
		for j := 0; j < b; j++ {
			ret[i][j] = value
		}
	}
	return ret
}
func boolSlice(a int) []bool {
	ret := make([]bool, a)
	return ret
}
func boolSlice2(a, b int) [][]bool {
	ret := make([][]bool, a)
	for i := 0; i < a; i++ {
		ret[i] = make([]bool, b)
	}
	return ret
}
func rangeSlice(begin, end int) []int {
	ret := make([]int, 0, end-begin+1)
	for i := begin; i <= end; i++ {
		ret = append(ret, i)
	}
	return ret
}
func rotate90(a [][]int) [][]int {
	w, h := len(a), len(a[0])
	res := make([][]int, h)
	for i := 0; i < h; i++ {
		res[i] = make([]int, w)
		for j := 0; j < w; j++ {
			res[i][j] = a[j][h-i-1]
		}
	}
	return res
}
func intsToBytes(a []int, add int) []byte {
	ret := make([]byte, len(a))
	for i := 0; i < len(a); i++ {
		ret[i] = byte(a[i] + add)
	}
	return ret
}
func bytesToInts(a []byte, add int) []int {
	ret := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		ret[i] = int(a[i]) + add
	}
	return ret
}
func stringIndexer() (func(string) int, func(int) string) {
	cur := 0
	m := make(map[string]int)
	r := make(map[int]string)
	f := func(s string) int {
		v := m[s]
		if v == 0 {
			cur++
			m[s] = cur
			r[cur] = s
			return cur
		}
		return v
	}
	f2 := func(x int) string {
		return r[x]
	}
	return f, f2
} // 文字列をindexに変換
func compressCoordinate(arr []int) ([]int, map[int]int, map[int]int) {
	arrCopy := append([]int(nil), arr...)
	sort.Ints(arrCopy)
	m := make(map[int]int)
	rev := make(map[int]int)
	rank := 0
	for _, a := range arrCopy {
		if _, ok := m[a]; !ok {
			m[a] = rank
			rank++
		}
	}
	result := make([]int, len(arr))
	for i, a := range arr {
		result[i] = m[a]
		rev[i] = a
	}
	return result, m, rev
} // 座標圧縮

// Pair操作
func dedupPair(p []Pair, KeyIsA bool, f func(int, int) int) []Pair {
	t := make(map[int]int)
	ret := make([]Pair, 0)
	if f == nil {
		for k := range pairs2map(p) {
			ret = append(ret, k)
		}
		return ret
	}
	for i := 0; i < len(p); i++ {
		if KeyIsA == true {
			val, ok := t[p[i].a]
			if ok == true {
				t[p[i].a] = f(val, p[i].b)
			} else {
				t[p[i].a] = p[i].b
			}
		} else {
			val, ok := t[p[i].b]
			if ok == true {
				t[p[i].b] = f(val, p[i].a)
			} else {
				t[p[i].b] = p[i].a
			}
		}
	}
	for k, v := range t {
		if KeyIsA == true {
			ret = append(ret, Pair{k, v})
		} else {
			ret = append(ret, Pair{v, k})
		}
	}
	return ret
}
func sortPairs(p []Pair, sortByA, prioritySmallestKey1, prioritySmallestKey2 bool) {
	f := []func(int, int) bool{func(a, b int) bool { return a < b }, func(a, b int) bool { return a > b }}
	f1 := condInt(prioritySmallestKey1, 0, 1)
	f2 := condInt(prioritySmallestKey2, 0, 1)
	if sortByA == true {
		sort.Slice(p, func(i, j int) bool { return condBool(p[i].a == p[j].a, f[f2](p[i].b, p[j].b), f[f1](p[i].a, p[j].a)) })
	} else {
		sort.Slice(p, func(i, j int) bool { return condBool(p[i].b == p[j].b, f[f2](p[i].a, p[j].a), f[f1](p[i].b, p[j].b)) })
	}
}
func pairs2map(p []Pair) map[Pair]int {
	ret := make(map[Pair]int)
	for i := 0; i < len(p); i++ {
		ret[p[i]]++
	}
	return ret
}
func map2pair(m map[int]int) []Pair {
	t := make([]Pair, 0, len(m))
	for i, v := range m {
		t = append(t, Pair{i, v})
	}
	return t
}
func splitPair(p []Pair) ([]int, []int) {
	res := make([]int, len(p))
	res2 := make([]int, len(p))
	for i, v := range p {
		res[i] = v.a
		res2[i] = v.b
	}
	return res, res2
}
func slice2Pairs(a, b []int) []Pair {
	if len(a) != len(b) {
		return nil
	}
	ret := make([]Pair, 0, len(a))
	for i := 0; i < len(a); i++ {
		ret = append(ret, Pair{a[i], b[i]})
	}
	return ret
}

// 文字列
func findChar(s string, p byte) ([]int, []int) {
	ret := make([]int, 0)
	ret2 := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == p {
			ret = append(ret, i)
			ret2 = append(ret2, len(s)-i-1)
		}
	}
	return ret, ret2
} // 左と右から文字が出現する位置を返す
func repeat(s string, n int) string {
	return strings.Repeat(s, n)
}

// グラフ
func NewGraph(n int) *Graph {
	g := &Graph{}
	g.n = n
	g.edges = make([][]Edge, g.n)
	return g
}
func (g *Graph) AddEdge(a, b, c int) {
	var t Edge
	t.from, t.to, t.w = a, b, c
	g.edges[a] = append(g.edges[a], t)
}
func (g *Graph) ReadSimpleUndirectedGraph(m int) {
	for i := 0; i < m; i++ {
		a, b := ri2()
		g.AddEdge(a, b, 1)
		g.AddEdge(b, a, 1)
	}
}
func (g *Graph) ReadWightedUndirectedGraph(m int) {
	for i := 0; i < m; i++ {
		a, b, c := ri3()
		g.AddEdge(a, b, c)
		g.AddEdge(b, a, c)
	}
}
func (g *Graph) ReadSimpleDirectedGraph(m int) {
	for i := 0; i < m; i++ {
		a, b := ri2()
		g.AddEdge(a, b, 1)
	}
}
func (g *Graph) ReadWightedDirectedGraph(m int) {
	for i := 0; i < m; i++ {
		a, b, c := ri3()
		g.AddEdge(a, b, c)
	}
}
func (g *Graph) PairToEdge(p []Pair, undirected bool) {
	for i := 0; i < len(p); i++ {
		g.AddEdge(p[i].a, p[i].b, 1)
		if undirected == true {
			g.AddEdge(p[i].b, p[i].a, 1)
		}
	}
}
func (g *Graph) SortEdgeByNode(prioritySmallest bool) {
	for v := 0; v < len(g.edges); v++ {
		if prioritySmallest == true {
			sort.Slice(g.edges[v], func(i, j int) bool {
				return g.edges[v][i].to < g.edges[v][j].to
			})
		} else {
			sort.Slice(g.edges[v], func(i, j int) bool {
				return g.edges[v][i].to > g.edges[v][j].to
			})
		}
	}
}
func readMatrixEdges(n, m int, readCost, directed bool) [][]int {
	ret := intSlice2(n, n, 0)
	a, b, c := 0, 0, 1
	for i := 0; i < m; i++ {
		if readCost == true {
			a, b, c = ri3()
		} else {
			a, b = ri2()
		}
		ret[a][b] = c
		if directed == false {
			ret[b][a] = c
		}
	}
	return ret
} // 隣接行列の読み込み
func restorePath(p []int, begin, end int) []int {
	ret := make([]int, 0)
	cur := end
	ret = append(ret, cur)
	for cur != -1 && cur != begin {
		ret = append(ret, p[cur])
		cur = p[cur]
	}
	ret = reverseList(ret)
	return ret
} // 最短パス復元(BFS,ダイクストラ,ベルマンフォード共通)
func searchGrid(g [][]byte, s string) [][]Pair {
	ret := make([][]Pair, len(s))
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[0]); j++ {
			for k := 0; k < len(s); k++ {
				if g[i][j] == s[k] {
					ret[k] = append(ret[k], Pair{i, j})
				}
			}
		}
	}
	return ret
} //Grid上の対象の座標を探索

// BIT
func NewFenwickTree(n int) *FenwickTree {
	ret := &FenwickTree{}
	ret.data = make([]int, n)
	ret.len = n
	return ret
}
func (f *FenwickTree) Add(p, x int) {
	if p < 0 || p >= f.len {
		panic("worong range")
		return
	}
	p += 1
	for p <= f.len {
		f.data[p-1] += x
		p += p & -p
	}
}
func (f *FenwickTree) Replace(p, x int) {
	if p < 0 || p >= f.len {
		panic("wrong range")
		return
	}
	oldValue := f.Sum(p, p+1)
	diff := x - oldValue
	f.Add(p, diff)
}
func (f *FenwickTree) Sum(l, r int) int {
	if l < 0 || l >= f.len || r < 0 || r > f.len || l > r {
		return 0
	}
	sum := func(r int) int {
		s := 0
		for r > 0 {
			s += f.data[r-1]
			r -= r & -r
		}
		return s
	}
	return sum(r) - sum(l)
} // 区間[l,r)の合計
func (f *FenwickTree) ModSum(l, r, mod int) int {
	if l < 0 || l >= f.len || r < 0 || r > f.len || l > r {
		return 0
	}
	sum := func(r int) int {
		s := 0
		for r > 0 {
			s += f.data[r-1]
			s %= mod
			r -= r & -r
		}
		return s % mod
	}
	return modSub(sum(r), sum(l))
} // 区間[l,r)の合計(MOD)
func (f *FenwickTree) Set(l []int) {
	n := len(l)
	if n != f.len {
		panic("worong length")
		return
	}
	for i, v := range l {
		f.Add(i, v)
	}
}
func (f *FenwickTree) LowerBound(l, k int) (int, int, bool) {
	if f.Sum(l, f.len) < k {
		return f.len, 0, false
	}
	b := f.Sum(0, l)
	idx := sort.Search(f.len, func(i int) bool {
		t := f.Sum(0, i)
		return t-b >= k
	})
	return idx, f.Sum(l, idx), true
}
func (f *FenwickTree) RangeCount(idx, nmax int) int {
	ret1, _, ok := f.LowerBound(idx, nmax+1)
	ret2, _, _ := f.LowerBound(idx, 0)
	if ok == true {
		return ret1 - ret2 - 1
	} else {
		return ret1 - ret2
	}
}
func countNumberOfFall(a []int) int {
	ret := 0
	t, _, _ := compressCoordinate(a)
	f := NewFenwickTree(len(t))
	for i := len(t) - 1; i >= 0; i-- {
		f.Add(t[i], 1)
		ret += f.Sum(0, t[i])
	}
	return ret
} // 座圧して転倒数を求める
func rangeBIT(n int) (func(int, int, int), func(int, int) int) {
	f0 := NewFenwickTree(n + 1)
	f1 := NewFenwickTree(n + 1)
	rangeAddFn := func(l, r, val int) {
		f0.Add(l, -val*l)
		f0.Add(r, val*r)
		f1.Add(l, val)
		f1.Add(r, -val)
	}
	sumFn := func(l, r int) int {
		ret := f0.Sum(0, r) + f1.Sum(0, r)*r - f0.Sum(0, l) + f1.Sum(0, l)*l
		return ret
	}
	return rangeAddFn, sumFn
} // [l,r)についてvalを区間加算、[l,r)の区間和を返すクロージャーを返す

// 2次元BIT
func NewBit2D(x, y int) *Bit2D {
	ret := &Bit2D{}
	ret.data = make([][]int, x+1)
	for i := range ret.data {
		ret.data[i] = make([]int, y+1)
	}
	ret.lenX = x + 1
	ret.lenY = y + 1
	return ret
}
func (f *Bit2D) Add(x, y, v int) {
	x += 1
	y += 1
	if x <= 0 || x >= f.lenX || y <= 0 || y >= f.lenY {
		panic("wrong range")
		return
	}
	for i := x; i < f.lenX; i += i & -i {
		for j := y; j < f.lenY; j += j & -j {
			f.data[i][j] += v
		}
	}
}
func (f *Bit2D) Replace(x, y, v int) {
	oldValue := f.Sum(x, y, x+1, y+1)
	diff := v - oldValue
	f.Add(x, y, diff)
}
func (f *Bit2D) Sum(x1, y1, x2, y2 int) int {
	sum := func(x, y int) int {
		s := 0
		for i := x; i > 0; i -= i & -i {
			for j := y; j > 0; j -= j & -j {
				s += f.data[i][j]
			}
		}
		return s
	}
	return sum(x2, y2) - sum(x1, y2) - sum(x2, y1) + sum(x1, y1)
}

// 2次元累積和
func imos2D(h, w int) (func(int, int, int), func(int, int, int, int) int) {
	table := intSlice2(h, w, 0)
	doneCompute := false
	add := func(y, x, val int) {
		if x < 0 || y < 0 || x >= w || y >= h {
			panic("Out of range")
		}
		table[y][x] += val
	}
	compute := func() {
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				if i > 0 {
					table[i][j] += table[i-1][j]
				}
				if j > 0 {
					table[i][j] += table[i][j-1]
				}
				if i > 0 && j > 0 {
					table[i][j] -= table[i-1][j-1]
				}
			}
		}
	}
	rangeSum := func(y1, x1, y2, x2 int) int {
		if doneCompute == false {
			compute()
			doneCompute = true
		}
		if y1 >= y2 || x1 >= x2 {
			return 0
		}
		d1, d2, d3 := 0, 0, 0
		if y1 != 0 && x1 != 0 {
			d1 = table[y1-1][x1-1]
		}
		if y1 != 0 {
			d2 = table[y1-1][x2-1]
		}
		if x1 != 0 {
			d3 = table[y2-1][x1-1]
		}
		return table[y2-1][x2-1] + d1 - d2 - d3
	}
	return add, rangeSum
} // 2次元累積和のadd(y,x)とrangeSum(y,xの閉開区間)のクロージャーを返す

// セグメント木(int)
func NewSegInt(a []int, e int, op func(int, int) int) *SegInt {
	ret := &SegInt{}
	ret.e = e
	ret.op = op
	ret.len = 1
	for ret.len < len(a) {
		ret.len *= 2
	}
	ret.data = make([]int, ret.len*2-1)

	for i := 0; i < len(a); i++ {
		ret.data[i] = a[i]
	}
	for i := len(a); i < ret.len; i++ {
		ret.data[i] = ret.e
	}
	for i := 0; i < ret.len*2-2; i += 2 {
		ret.data[i/2+ret.len] = op(ret.data[i], ret.data[i+1])
	}
	return ret
}
func (seg *SegInt) Update(idx int, val int) {
	cur := idx
	seg.data[cur] = val
	for cur < seg.len*2-2 {
		if cur%2 == 1 {
			seg.data[cur/2+seg.len] = seg.op(seg.data[cur-1], seg.data[cur])
		} else {
			seg.data[cur/2+seg.len] = seg.op(seg.data[cur], seg.data[cur+1])
		}
		cur = cur/2 + seg.len
	}
}
func (seg *SegInt) Query(l, r int) int {
	ret := seg.e
	for l < r {
		if l%2 == 1 {
			ret = seg.op(ret, seg.data[l])
			l++
		}
		if r%2 == 1 {
			ret = seg.op(ret, seg.data[r-1])
			r--
		}
		if l == r {
			return ret
		}
		l = l/2 + seg.len
		r = r/2 + seg.len
	}
	ret = seg.op(ret, seg.data[l])
	return ret
} // 半開区間[l,r)のクエリ
func sliceOrderChecker(a []int, asc bool) (func(int, int), func(int, int) bool) {
	dat := make([]int, len(a))
	copy(dat, a)
	var fn func(int, int) int
	if asc == true {
		fn = func(u, v int) int { return condInt(u <= v, 1, 0) }
	} else {
		fn = func(u, v int) int { return condInt(u >= v, 1, 0) }
	}
	d := make([]int, len(a))
	d[len(a)-1] = 1
	for i := 0; i < len(a)-1; i++ {
		d[i] = fn(a[i], a[i+1])
	}
	seg := NewSegInt(d, 1, func(u, v int) int { return condInt(u == 1 && v == 1, 1, 0) })
	update := func(idx, val int) {
		dat[idx] = val
		if len(dat) == 1 {
			return
		}
		if idx != len(dat)-1 {
			seg.Update(idx, fn(dat[idx], dat[idx+1]))
		}
		if idx != 0 {
			seg.Update(idx-1, fn(dat[idx-1], dat[idx]))
		}
	}
	query := func(l, r int) bool {
		if r-l == 0 || seg.Query(l, r-1) == 1 {
			return true
		}
		return false
	}
	return update, query
} // スライスの[l,r)について昇順(or 降順)の場合にtrueとなるクロージャーを返す
func charSet(a []int) (func(int, int), func(int, int) ([]int, int, int, []int)) {
	const maxCharTypes = 26
	dat := make([]int, len(a))
	copy(dat, a)
	count := intSlice(maxCharTypes, 0)
	seg := make([]*SegInt, maxCharTypes)
	for i := 0; i < maxCharTypes; i++ {
		seg[i] = NewSegInt(intSlice(len(a), 0), 0, func(a, b int) int { return a + b })
	}
	for i := 0; i < len(a); i++ {
		idx := a[i]
		seg[idx].Update(i, 1)
		count[idx]++
	}
	update := func(idx, val int) {
		old := dat[idx]
		next := val
		dat[idx] = next
		seg[old].Update(idx, 0)
		seg[next].Update(idx, 1)
		count[old]--
		count[next]++
	}
	query := func(l, r int) ([]int, int, int, []int) {
		ret := intSlice(maxCharTypes, 0)
		minIdx, maxIdx := -1, -1
		for i := 0; i < maxCharTypes; i++ {
			t := seg[i].Query(l, r)
			if minIdx == -1 && t != 0 {
				minIdx = i
			}
			if t != 0 {
				maxIdx = i
			}
			ret[i] += t
		}
		return ret, minIdx, maxIdx, count
	}
	return update, query
} // 区間[l,r)に含まれる文字カウント、辞書順の最小文字、最大文字、全体の文字カウントを求めるクロージャーを返す

// UnionFind
func NewUnionFind(n int) *UnionFind {
	root := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		root[i] = i
		size[i] = 1
	}
	uf := &UnionFind{root: root, size: size, group: n}
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

// Int型Deque
func NewDeque() *Deque {
	deq := &Deque{}
	deq.deq = list.New()
	return deq
}
func (deq *Deque) PushBack(n int) {
	deq.cur = deq.deq.PushBack(n)
	deq.sum += n
}
func (deq *Deque) PushFront(n int) {
	deq.cur = deq.deq.PushFront(n)
	deq.sum += n
}
func (deq *Deque) InsertAfterCurrent(n int) {
	deq.cur = deq.deq.InsertAfter(n, deq.cur)
	deq.sum += n
}
func (deq *Deque) InsertBeforeCurrent(n int) {
	deq.cur = deq.deq.InsertBefore(n, deq.cur)
	deq.sum += n
}
func (deq *Deque) PopBack() int {
	ret := deq.deq.Back().Value.(int)
	deq.deq.Remove(deq.deq.Back())
	deq.sum -= ret
	return ret
}
func (deq *Deque) PopFront() int {
	ret := deq.deq.Front().Value.(int)
	deq.deq.Remove(deq.deq.Front())
	deq.sum -= ret
	return ret
}
func (deq *Deque) DumpBack() []int {
	ret := make([]int, 0, deq.Len())
	for e := deq.deq.Back(); e != nil; e = e.Prev() {
		ret = append(ret, e.Value.(int))
	}
	return ret
}
func (deq *Deque) DumpFront() []int {
	ret := make([]int, 0, deq.Len())
	for e := deq.deq.Front(); e != nil; e = e.Next() {
		ret = append(ret, e.Value.(int))
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
func NewIntPQ(prioritySmallest bool) *IntPQ {
	ret := &IntPQ{}
	ret.d = make([]int, 0, 100)
	n := ret.Len()
	for i := n/2 - 1; i >= 0; i-- {
		ret.down(i, n)
	}
	ret.prioritySmallest = prioritySmallest
	return ret
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
func (pq *IntPQ) Push(x int) {
	pq.d = append(pq.d, x)
	pq.up(len(pq.d) - 1)
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
func (pq *IntPQ) PushSlice(a []int) {
	for _, v := range a {
		pq.Push(v)
	}
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
func (pq *IntPQ) PopAndPush(x int) int {
	ret := pq.Peek()
	pq.d[0] = x
	pq.down(0, pq.Len())
	return ret
}
func (pq *IntPQ) PopUniq() int {
	ret := pq.Pop()
	for pq.Len() != 0 && pq.Peek() == ret {
		pq.Pop()
	}
	return ret
}

// 構造体Deque
func NewDataDeque() *DataDeque {
	deq := &DataDeque{}
	deq.deq = list.New()
	return deq
}
func (deq *DataDeque) PushBack(d Data) {
	deq.cur = deq.deq.PushBack(d)
}
func (deq *DataDeque) PushFront(d Data) {
	deq.cur = deq.deq.PushFront(d)
}
func (deq *DataDeque) InsertAfterCurrent(d Data) {
	deq.cur = deq.deq.InsertAfter(d, deq.cur)
}
func (deq *DataDeque) InsertBeforeCurrent(d Data) {
	deq.cur = deq.deq.InsertBefore(d, deq.cur)
}
func (deq *DataDeque) PopBack() Data {
	ret := deq.deq.Back().Value.(Data)
	deq.deq.Remove(deq.deq.Back())
	return ret
}
func (deq *DataDeque) PopFront() Data {
	ret := deq.deq.Front().Value.(Data)
	deq.deq.Remove(deq.deq.Front())
	return ret
}
func (deq *DataDeque) DumpBack() []Data {
	ret := make([]Data, 0, deq.Len())
	for e := deq.deq.Back(); e != nil; e = e.Prev() {
		ret = append(ret, e.Value.(Data))
	}
	return ret
}
func (deq *DataDeque) DumpFront() []Data {
	ret := make([]Data, 0, deq.Len())
	for e := deq.deq.Front(); e != nil; e = e.Next() {
		ret = append(ret, e.Value.(Data))
	}
	return ret
}
func (deq *DataDeque) Back() Data {
	return deq.deq.Back().Value.(Data)
}
func (deq *DataDeque) Front() Data {
	return deq.deq.Front().Value.(Data)
}
func (deq *DataDeque) Len() int {
	return deq.deq.Len()
}

// データ構造体用ヒープ
func NewItemPQ(prioritySmallest bool) *ItemPQ {
	ret := &ItemPQ{}
	ret.d = make([]Item, 0, 100)
	n := ret.Len()
	for i := n/2 - 1; i >= 0; i-- {
		ret.down(i, n)
	}
	ret.prioritySmallest = prioritySmallest
	return ret
}
func (pq *ItemPQ) less(i, j int) bool {
	if pq.prioritySmallest == true {
		return pq.d[i].key < pq.d[j].key
	} else {
		return pq.d[i].key > pq.d[j].key
	}
}
func (pq *ItemPQ) swap(i, j int) { pq.d[i], pq.d[j] = pq.d[j], pq.d[i] }
func (pq *ItemPQ) down(i0, n int) bool {
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
func (pq *ItemPQ) Push(x Item) {
	pq.d = append(pq.d, x)
	pq.up(len(pq.d) - 1)
}
func (pq *ItemPQ) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !pq.less(j, i) {
			break
		}
		pq.swap(i, j)
		j = i
	}
}
func (pq *ItemPQ) Pop() Item {
	n := pq.Len() - 1
	x := pq.d[0]
	pq.swap(0, n)
	pq.down(0, n)
	pq.d = pq.d[0 : pq.Len()-1]
	return x
}
func (pq *ItemPQ) Len() int   { return len(pq.d) }
func (pq *ItemPQ) Peek() Item { return pq.d[0] }
func (pq *ItemPQ) Remove(i int) Item {
	n := pq.Len() - 1
	if n != i {
		pq.swap(i, n)
		if !pq.down(i, n) {
			pq.up(i)
		}
	}
	return pq.Pop()
}
func (pq *ItemPQ) Fix(i int) {
	if !pq.down(i, pq.Len()) {
		pq.up(i)
	}
}

// 整数関連
func lcm(a, b int) int {
	return (a / gcd(a, b)) * b
}
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	c := 1
	for c != 0 {
		c = a % b
		a, b = b, c
	}
	return a
}
func exgcd(a, b int) (int, int) {
	q := 0
	x0, x1, y0, y1 := 1, 0, 0, 1
	for b != 0 {
		q, a, b = a/b, b, a%b
		x0, x1 = x1, x0-q*x1
		y0, y1 = y1, y0-q*y1
	}
	return x0, y0
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
func divisorCnt(maxNum int) func(int) int {
	memo := intSlice(maxNum, -1)
	f := func(n int) int {
		if memo[n] != -1 {
			return memo[n]
		}
		var cnt int
		for i := 1; i*i <= n; i++ {
			if n%i == 0 {
				cnt++
				if i != n/i {
					cnt++
				}
			}
		}
		memo[n] = cnt
		return cnt
	}
	return f
}
func divisorPairs(n int) []Pair {
	var p []Pair
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			p = append(p, Pair{i, n / i})
		}
	}
	return p
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
func fastFacorization(n int) func(x int) map[int]int {
	rp := make([]int, n+1)
	p := primeList(n)
	rp[1] = 1
	for _, v := range p {
		for i := v; i <= n; i += v {
			rp[i] = v
		}
	}
	var f func(int) map[int]int
	f = func(k int) map[int]int {
		t := k
		m := make(map[int]int)
		for t != 1 {
			m[rp[t]]++
			t = t / rp[t]
		}
		return m
	}
	return f
}
func isPrime(n int) bool { return big.NewInt(int64(n)).ProbablyPrime(0) }
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
func powMod(x, k, m int) int {
	if k == 0 {
		return 1
	}
	if x > m {
		x %= m
	}
	if k%2 == 0 {
		return powMod(x*x%m, k/2, m)
	} else {
		return x * powMod(x, k-1, m) % m
	}
}
func pow(a, b int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = res * a
		}
		a = a * a
		b >>= 1
	}
	return res
}
func combination(n, r int) int {
	if n < r {
		return 0
	}
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
func combinationMod(a, b int) int {
	t1, t2 := 1, 1
	for i := 2; i <= b; i++ {
		t2 = (i * t2) % MOD
	}
	inv := modInv(t2)
	for i := a - b + 1; i <= a; i++ {
		t1 = (i * t1) % MOD
	}
	return (t1 * inv) % MOD
}
func sqrt(x int) int                  { return int(math.Sqrt(float64(x))) }
func cbrt(x int) int                  { return int(math.Cbrt(float64(x))) }
func seqSum(x int) int                { return (x*x + x) / 2 }
func seqRangeSum(from, to int) int    { return seqSum(to) - seqSum(from-1) }
func seqSumMod(x int) int             { return modMul(modAdd(modMul(x%MOD, x%MOD), x%MOD), modInv(2)) }
func seqRangeSumMod(from, to int) int { return modSub(seqSum(to), seqSum(from-1)) }
func getArithmeticSeqRange(d, a, maxY, maxX int) (int, int) {
	t := (-a + d - 1) / d
	return max(1, t), min(maxX, (maxY-a)/d)
} // 等差数列(初項a,公差d)について、y[0,maxY],x[0,maxX]の条件でxの範囲を求める
func MulMatrix(m1, m2 [][]int, mod int) [][]int {
	size := len(m1)
	ret := make([][]int, size)
	for i := 0; i < size; i++ {
		ret[i] = make([]int, size)
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			for k := 0; k < size; k++ {
				ret[i][k] += m1[i][j] * m2[j][k]
				ret[i][k] %= mod
			}
		}
	}
	return ret
} // 正方行列の掛け算
func PowMatrix(m [][]int, n int, mod int) [][]int {
	// 行列累乗
	// a_0=s a_1=t
	// a_(i+2) = p*a_(i+1) + q*a_i
	//           ↓
	// |a_(i+2)| = | p q |^n (a_1)
	// |a_(i+1)| = | 1 0 |   (a_0)
	size := len(m)
	tm := make([][]int, size)
	for i := 0; i < size; i++ {
		tm[i] = make([]int, size)
		tm[i][i] = 1
	}
	for n >= 2 {
		if n%2 == 0 {
			m = MulMatrix(m, m, mod)
			n = n / 2
		} else {
			tm = MulMatrix(tm, m, mod)
			m = MulMatrix(m, m, mod)
			n = (n - 1) / 2
		}
	}
	m = MulMatrix(tm, m, mod)
	return m
} // 正方行列の冪乗

// int型操作(1-indexed)
func getDigitLen(x int) int   { return condInt(x <= 0, 0, 1+int(math.Log10(float64(x)))) }
func getDigit(x, idx int) int { return condInt(idx > 17 || idx < 1, 0, (x/p10[idx-1])%10) }
func setDigit(x, idx, v int) int {
	return condInt(idx > 17 || v > 9 || v < 0, 0, x+(v-getDigit(x, idx))*p10[idx-1])
}
func swapDigit(x, idx1, idx2 int) int {
	t1 := getDigit(x, idx1)
	t2 := getDigit(x, idx2)
	x = setDigit(x, idx2, t1)
	x = setDigit(x, idx1, t2)
	return x
}

// ビット操作(0-indexed)
func bitLen(x int) int              { return bits.Len(uint(x)) }                                        // xの2進数での長さを取得する
func getNthBit(x, idx int) int      { return condInt(x&(1<<uint(idx)) == 0, 0, 1) }                     // idx番目のbitを取得する
func setNthBit(x, idx, bit int) int { return condInt(bit == 0, x & ^(1<<uint(idx)), x|(1<<uint(idx))) } // idx番目のbitを変更する
func toggleNthBit(x, idx int) int   { return x ^ (1 << uint(idx)) }                                     // idx番目のbitを反転させる
func itobstr(x, digits int) string  { return fmt.Sprintf("%0*b", digits, x) }                           // 数値をバイナリ表記したstringに変換する(デバッグ用)
func btoi(b []byte) int {
	ret, _ := strconv.ParseInt(string(b), 2, 64)
	return int(ret)
}                                // バイナリ表記[]byteをintに変換する
func bitRightmostOne(n int) int  { return bits.TrailingZeros(uint(n)) }  // bitが1になっている一番右側の位置を返す
func bitRightmostZero(n int) int { return bits.TrailingZeros(^uint(n)) } // bitが0になっている一番右側の位置を返す
func getBitPosConv(mask int) []int {
	t := make([]int, bits.OnesCount(uint(mask)))
	k := bits.Len(uint(mask))
	cnt := 0
	for i := 0; i < k; i++ {
		if getNthBit(mask, i) == 1 {
			t[cnt] = i
			cnt++
		}
	}
	return t
} // bitが1になっている位置の一覧を返す
func nbitsNth(bit0Sum, bit1Sum, kth int) int {
	ret := 0
	t := 0
	a, b := bit0Sum, bit1Sum
	s := bit0Sum + bit1Sum
	for i := 0; i < s; i++ {
		if a <= 0 {
			ret = setNthBit(ret, s-i-1, 1)
		} else if b <= 0 {
			ret = setNthBit(ret, s-i-1, 0)
		} else {
			n := combination(a+b-1, a-1)
			if n+t < kth {
				t += n
				ret = setNthBit(ret, s-i-1, 1)
				b--
			} else {
				ret = setNthBit(ret, s-i-1, 0)
				a--
			}
		}
	}
	return ret
} // 2進数でbitの0と1を指定した数分含む値ついてkth番目の値を求める
func bitApply(bitLen, num, bit0, bit1 int) int {
	ret, t := 0, 0
	for i := 0; i < bitLen; i++ {
		k := getNthBit(num, i)
		if k == 0 {
			t = getNthBit(bit0, i)
		} else {
			t = getNthBit(bit1, i)
		}
		ret = setNthBit(ret, i, t)
	}
	return ret
} // numを2進数に変換した各桁について、0の場合はbit0,1の場合はbit1で対応する桁の値に置き換える
func bitIndexSum(a []int, x int) int {
	ret := 0
	for i := 0; i < bitLen(x); i++ {
		if getNthBit(x, i) == 1 {
			ret += a[i]
		}
	}
	return ret
} // aについてxのbitが1のposの和を取る
func countRangeBit(x, from, to int) int {
	t := (1<<(to-from+1) - 1) << from
	return bits.OnesCount(uint(t & x))
} // [l,r]に含まれるbit1の数をカウントする(0-indexed)

// Pair操作
func subPair(p, t Pair) Pair     { return Pair{p.a - t.a, p.b - t.b} }
func addPair(p, t Pair) Pair     { return Pair{p.a + t.a, p.b + t.b} }
func mulPair(p Pair, d int) Pair { return Pair{p.a * d, p.b * d} }
func detPair(p, t Pair) int      { return p.a*t.b - p.b*t.a }
func dotPair(p, t Pair) int      { return p.a*t.a + p.b*t.b }
func distPair(p, t Pair) int     { return dotPair(subPair(p, t), subPair(p, t)) }
func distIntMatrix(p []Pair) [][]int {
	n := len(p)
	ret := intSlice2(n, n, INF)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			t := distPair(p[i], p[j])
			ret[i][j] = t
			ret[j][i] = t
		}
	}
	return ret
}
func distFloatMatrix(p []Pair) [][]float64 {
	n := len(p)
	ret := make([][]float64, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			ret[i][j] = math.MaxFloat32
		}
	}
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			t := math.Sqrt(float64(distPair(p[i], p[j])))
			ret[i][j] = t
			ret[j][i] = t
		}
	}
	return ret
}
func intersectSum(p1, p2 Pair) int {
	return max(0, min(p1.b, p2.b)-max(p1.a, p2.a))
} // [a1,b1),[a2,b2)の重複区間のカウント
func normalizePairs(p []Pair) []Pair {
	ma, mb := INF, INF
	for i := 0; i < len(p); i++ {
		ma = min(ma, p[i].a)
		mb = min(mb, p[i].b)
	}
	ret := make([]Pair, len(p))
	for i := 0; i < len(p); i++ {
		ret[i].a = p[i].a - ma
		ret[i].b = p[i].b - mb
	}
	return ret
}
func addPairs(p []Pair, d Pair) []Pair {
	ret := make([]Pair, len(p))
	for i := 0; i < len(p); i++ {
		ret[i] = addPair(p[i], d)
	}
	return ret
}
func sumProductOfPairs(p []Pair) int {
	ret := 0
	for _, v := range p {
		ret += v.a * v.b
	}
	return ret
}
func mergeIntervals(p []Pair) []Pair {
	sortPairs(p, true, true, true)
	pa, pb := p[0].a, p[0].b
	ret := make([]Pair, 0)
	for i := 0; i < len(p); i++ {
		if pb >= p[i].a {
			chmax(&pb, p[i].b)
		} else {
			ret = append(ret, Pair{pa, pb})
			pa, pb = p[i].a, p[i].b
		}
		if i == len(p)-1 {
			ret = append(ret, Pair{pa, pb})
		}
	}
	return ret
} // 複数の[a,b)の区間について重複する区間をマージする

// グルーピング
func splitGroupIndexRange(a []int) []Pair {
	d := 1 //差分がd以下だと同じグループ
	if len(a) == 0 {
		return nil
	}
	if len(a) == 1 {
		return []Pair{{0, 0}}
	}
	begin := 0
	ret := make([]Pair, 0)
	for i := 1; i < len(a); i++ {
		if a[i]-a[i-1] <= d {
			continue
		}
		ret = append(ret, Pair{begin, i - 1})
		begin = i
	}
	ret = append(ret, Pair{begin, len(a) - 1})
	return ret
} // 単調増加のスライスで隣合う要素の差がd以下のものをグルーピングし開始/終了のindexを返す
func divideConnectedGroup(n int, p []Pair) [][]int {
	uf := NewUnionFind(n + 1)
	for i := 0; i < len(p); i++ {
		uf.Union(p[i].a, p[i].b)
	}
	nm := make(map[int][]int)
	for i := 1; i <= n; i++ {
		k := uf.find(i)
		nm[k] = append(nm[k], i)
	}
	ret := make([][]int, 0)
	for _, v := range nm {
		ret = append(ret, v)
	}
	return ret
} // 無向グラフを連結成分毎にグルーピングする (1-indexedでノード0は無視)

// 数えあげ
func findSumPattern(a []int, mod int, maxCnt int) [][][]int {
	m := make([][][]int, mod)
	cnt := 0
	var f func(int, int, []int)
	f = func(idx, val int, t []int) {
		if idx == len(a) || (cnt >= maxCnt && maxCnt != -1) {
			return
		}
		nv := (val + a[idx]) % mod
		np := make([]int, len(t))
		copy(np, t)
		np = append(np, idx+1)
		m[nv] = append(m[nv], np)
		cnt++
		f(idx+1, val, t)
		f(idx+1, nv, np)
		return
	}
	f(0, 0, nil)
	return m
} // スライスの部分列の和(mod)について、組み合わせを最大maxCnt回記録する(maxCnt=-1の場合は全列挙)

// 分数
func compareFrac(a, b Frac) int {
	// -1(a>b), 1(a<b), 0(a==b)
	t1 := a.top * b.bottom
	t2 := a.bottom * b.top
	if t1 == t2 {
		return 0
	} else if t1 < t2 {
		return 1
	} else {
		return -1
	}
}
func AddFrac(a, b Frac) Frac {
	top := a.top*b.bottom + a.bottom*b.top
	bottom := a.bottom * b.bottom
	g := gcd(top, bottom)
	return Frac{top / g, bottom / g}
}
func SubFrac(a, b Frac) Frac {
	top := a.top*b.bottom - a.bottom*b.top
	bottom := a.bottom * b.bottom
	g := gcd(top, bottom)
	return Frac{top / g, bottom / g}
}
func MulFrac(a, b Frac) Frac {
	top := a.top * b.top
	bottom := a.bottom * b.bottom
	g := gcd(top, bottom)
	return Frac{top / g, bottom / g}
}

// ベクトル
func degreeToRadian(d float64) float64 { return d * math.Pi / 180 }
func rotatePoint(p Point, rad float64) Point {
	cos := math.Cos(rad)
	sin := math.Sin(rad)
	return Point{sin*p.x + cos*p.y, cos*p.x - sin*p.y}
}
func addPoint(a, b Point) Point         { return Point{a.y + b.y, a.x + b.x} }
func subPoint(a, b Point) Point         { return Point{a.y - b.y, a.x - b.x} }
func mulPoint(a Point, b float64) Point { return Point{a.y * b, a.x * b} }
func distance(a, b Point) float64 {
	t := subPoint(a, b)
	return math.Hypot(t.x, t.y)
}

// 図形判定
func isConnectedCircle(p1, p2 Pair, r1, r2 int) bool {
	d := distPair(p1, p2)
	if d > pow(r1+r2, 2) || d < pow(r1-r2, 2) {
		return false
	}
	return true
}

// 順列
func sumPermutation(s, n int) [][]int {
	var results [][]int
	t := make([]int, n)
	var dfs func(int, int)
	dfs = func(sum, index int) {
		if index == n {
			if sum == s {
				result := make([]int, n)
				copy(result, t)
				results = append(results, result)
			}
			return
		}
		for i := 0; i <= s; i++ {
			t[index] = i
			dfs(sum+i, index+1)
		}
	}
	dfs(0, 0)
	return results
} // 合計がsとなるn個の要素のスライスの組み合わせを返す
func sumPermutations(s, n int) [][]int {
	perm := make([][]int, 0)
	for i := 0; i <= s; i++ {
		t := sumPermutation(i, n)
		perm = append(perm, t...)
	}
	return perm
} // 合計がs以下となるn個の要素のスライスの組み合わせを返す
func nextPermIterator(orderdSlice []int) func() []int {
	x := orderdSlice
	first := true
	f := func(x sort.Interface) bool {
		n := x.Len() - 1
		if n < 1 {
			return false
		}
		j := n - 1
		for ; !x.Less(j, j+1); j-- {
			if j == 0 {
				return false
			}
		}
		l := n
		for !x.Less(j, l) {
			l--
		}
		x.Swap(j, l)
		for k, l := j+1, n; k < l; {
			x.Swap(k, l)
			k++
			l--
		}
		return true
	}
	f2 := func() []int {
		if first == true {
			first = false
			return x
		} else {
			ret := f(sort.IntSlice(x))
			if ret == false {
				return nil
			}
			return x
		}
	}
	return f2
} // nextPermutation()のIterator
func ascPerm(size, low, high int) [][]int {
	p := intSlice(size, low)
	ret := make([][]int, 0)
	var f func(int, int)
	f = func(idx, cur int) {
		if idx == size {
			ret = append(ret, cs(p))
			return
		}
		for i := cur; i <= high; i++ {
			p[idx] = i
			f(idx+1, i)
		}
	}
	f(0, low)
	return ret
} // low〜highの要素が昇順(ai <= aj)となるような組み合わせを返す

// bit set
func newBitSets(size int) *BitSets {
	bs := make(BitSets, size)
	return &bs
}
func (tb *BitSets) SetBit(i int, bitNum int, bitVal uint) {
	(*tb)[i].SetBit(&(*tb)[i], bitNum, bitVal)
}
func (tb *BitSets) Mex(u int) int {
	temp := new(big.Int)
	temp.Not(&(*tb)[u])
	return int(temp.TrailingZeroBits())
} // 集合のmexを求める
func (tb *BitSets) Mex2(u int, v int) int {
	temp := new(big.Int)
	temp.Or(&(*tb)[u], &(*tb)[v])
	temp.Not(temp)
	return int(temp.TrailingZeroBits())
} //和集合のmexを求める
