package main

import (
	"fmt"
)

//データ構造体用ヒープ
type PQData struct {
	key int //並び替え用キー
	num int
}

func NewStructPQ(prioritySmallest bool) *StructPQ {
	ret := &StructPQ{}
	ret.d = make([]PQData, 0, 100)
	n := ret.Len()
	for i := n/2 - 1; i >= 0; i-- {
		ret.down(i, n)
	}
	ret.prioritySmallest = prioritySmallest
	return ret
}

type StructPQ struct {
	d                []PQData
	prioritySmallest bool
}

func (pq *StructPQ) less(i, j int) bool {
	if pq.prioritySmallest == true {
		return pq.d[i].key < pq.d[j].key
	} else {
		return pq.d[i].key > pq.d[j].key
	}
}
func (pq *StructPQ) swap(i, j int) { pq.d[i], pq.d[j] = pq.d[j], pq.d[i] }
func (pq *StructPQ) down(i0, n int) bool {
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
func (pq *StructPQ) Push(x PQData) {
	pq.d = append(pq.d, x)
	pq.up(len(pq.d) - 1)
}
func (pq *StructPQ) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !pq.less(j, i) {
			break
		}
		pq.swap(i, j)
		j = i
	}
}
func (pq *StructPQ) Pop() PQData {
	n := pq.Len() - 1
	x := pq.d[0]
	pq.swap(0, n)
	pq.down(0, n)
	pq.d = pq.d[0 : pq.Len()-1]
	return x
}
func (pq *StructPQ) Len() int {
	return len(pq.d)
}
func (pq *StructPQ) Peek() PQData {
	return pq.d[0]
}
func (pq *StructPQ) Remove(i int) PQData {
	n := pq.Len() - 1
	if n != i {
		pq.swap(i, n)
		if !pq.down(i, n) {
			pq.up(i)
		}
	}
	return pq.Pop()
}
func (pq *StructPQ) Fix(i int) {
	if !pq.down(i, pq.Len()) {
		pq.up(i)
	}
}

func main() {
	pq := NewStructPQ(false)
	pq.Push(PQData{1, 2})
	pq.Push(PQData{5, 1})
	pq.Push(PQData{9, 3})
	pq.Push(PQData{3, 5})
	fmt.Println(pq.Peek())
	fmt.Println(pq.Pop())
	fmt.Println(pq.Pop())
}
