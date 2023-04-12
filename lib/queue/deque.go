package main

import (
	"container/list"
	"fmt"
)

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

func main() {
	deq := NewDeque()
	deq.PushBack(1)
	deq.PushBack(2)

	fmt.Println(deq.DumpFront())
}
