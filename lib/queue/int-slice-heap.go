// This example demonstrates an integer heap built using the heap interface.
package main

import (
	"container/heap"
	"fmt"
)

type IntSliceHeap [][]int

func (h IntSliceHeap) Len() int           { return len(h) }
func (h IntSliceHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h IntSliceHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntSliceHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *IntSliceHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	h := &IntSliceHeap{[]int{1, 1}, {3, 5}, {8, 9}, {4, 4}}
	heap.Init(h)
	heap.Push(h, []int{3, 2})
	fmt.Printf("minimum: %d\n", (*h)[0][0])
	for h.Len() > 0 {
		fmt.Println(heap.Pop(h))
	}
}
