package main

import (
	"fmt"
)

// 削除予約付きInt型PriorityQueue
func NewRemovablePQ(prioritySmallest bool) *RemovableIntPQ {
	ret := &RemovableIntPQ{}
	ret.d = make([]int, 0, 100)
	n := ret.Len()
	for i := n/2 - 1; i >= 0; i-- {
		ret.down(i, n)
	}
	ret.prioritySmallest = prioritySmallest
	ret.m = make(map[int]int)

	return ret
}

type RemovableIntPQ struct {
	d                []int
	prioritySmallest bool
	m                map[int]int
	sum              int
	len              int
}

func (pq *RemovableIntPQ) less(i, j int) bool {
	if pq.prioritySmallest == true {
		return pq.d[i] < pq.d[j]
	} else {
		return pq.d[i] > pq.d[j]
	}
}
func (pq *RemovableIntPQ) swap(i, j int) { pq.d[i], pq.d[j] = pq.d[j], pq.d[i] }
func (pq *RemovableIntPQ) down(i0, n int) bool {
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
func (pq *RemovableIntPQ) Push(x int) {
	pq.len++
	pq.sum += x
	pq.d = append(pq.d, x)
	pq.up(len(pq.d) - 1)
}
func (pq *RemovableIntPQ) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !pq.less(j, i) {
			break
		}
		pq.swap(i, j)
		j = i
	}
}
func (pq *RemovableIntPQ) pop() int {
	n := pq.Len() - 1
	x := pq.d[0]
	pq.swap(0, n)
	pq.down(0, n)
	pq.d = pq.d[0 : pq.Len()-1]
	return x
}
func (pq *RemovableIntPQ) Pop() int {
	t := 0
	for {
		t = pq.pop()
		if pq.m[t] == 0 {
			break
		}
		pq.m[t]--
	}
	pq.sum -= t
	pq.len--
	return t
}
func (pq *RemovableIntPQ) Len() int {
	return len(pq.d)
}
func (pq *RemovableIntPQ) Peek() int {
	return pq.d[0]
}
func (pq *RemovableIntPQ) Remove(x int) {
	pq.len--
	pq.sum -= x
	pq.m[x]++
	for pq.m[pq.d[0]] != 0 {
		pq.m[pq.d[0]]--
		pq.Pop()
	}
}
func (pq *RemovableIntPQ) Fix(i int) {
	if !pq.down(i, pq.Len()) {
		pq.up(i)
	}
}

func main() {
	r := NewRemovablePQ(false)
	for i := 0; i < 10; i++ {
		r.Push(i)
	}
	fmt.Println(r.Pop())
	fmt.Println(r.Pop())
	fmt.Println(r.sum, r.len)
	r.Remove(9)
	fmt.Println(r.sum, r.len)
	r.Remove(0)
	fmt.Println(r.sum, r.len)
	fmt.Println(r.Pop())
	fmt.Println(r.Peek())
	fmt.Println(r.sum, r.len)
}
