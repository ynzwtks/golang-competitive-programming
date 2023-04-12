package main

import (
	"fmt"
)

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

type TwowayHeap struct {
	ih  *IntPQ
	rih *IntPQ
	m   map[int]int
	rm  map[int]int
	sum int
	len int
}

func NewTwowayHeap() *TwowayHeap {
	r := &TwowayHeap{}
	r.ih = NewIntPQ(true)
	r.rih = NewIntPQ(false)
	r.m = make(map[int]int)
	r.rm = make(map[int]int)
	return r
}
func (r *TwowayHeap) Push(x int) {
	r.len++
	r.sum += x
	r.ih.Push(x)
	r.rih.Push(x)
}
func (r *TwowayHeap) Remove(x int) {
	r.len--
	r.sum -= x
	r.m[x]++
	r.rm[x]++
}
func (r *TwowayHeap) Left() int {
	for r.m[r.ih.d[0]] != 0 {
		r.m[r.ih.d[0]]--
		r.ih.Pop()
	}
	return r.ih.d[0]
}
func (r *TwowayHeap) Right() int {
	for r.rm[r.rih.d[0]] != 0 {
		r.rm[r.rih.d[0]]--
		r.rih.Pop()
	}
	return r.rih.d[0]
}
func (r *TwowayHeap) PopLeft() int {
	t := 0
	for {
		t = r.ih.Pop()
		if r.m[t] == 0 {
			break
		}
		r.m[t]--
	}
	r.sum -= t
	r.len--
	r.rm[t]++
	return t
}
func (r *TwowayHeap) PopRight() int {
	t := 0
	for {
		t = r.rih.Pop()
		if r.rm[t] == 0 {
			break
		}
		r.rm[t]--
	}
	r.sum -= t
	r.len--
	r.m[t]++
	return t
}

func main() {
	r := NewTwowayHeap()
	for i := 0; i < 10; i++ {
		r.Push(i)
	}

	r.Remove(9)
	r.Remove(0)
	//fmt.Println(r.popLeft())
	fmt.Println(r.PopRight())
	fmt.Println(r.Right(), r.Left())
	fmt.Println(r.PopRight())
	fmt.Println(r.PopLeft())
	fmt.Println(r.Right(), r.Left())
	fmt.Println(r.ih, r.rih, r.sum, r.len)
}
