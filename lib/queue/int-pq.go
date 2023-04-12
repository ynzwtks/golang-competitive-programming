package main

import "fmt"

// int型専用ヒープ
// 標準のcontainer/heapではinterface型のオーバヘッドで大量のpush,popだと遅くなるのでint専用で高速化。
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

func main() {
	pq := NewIntPQ(false)
	pq.Push(1)
	pq.Push(3)
	pq.Push(2)
	pq.Push(7)
	fmt.Println(pq.Peek())
	fmt.Println(pq.Pop())
	fmt.Println(pq.Pop())
}
