package main

import (
	"cmp"
	"fmt"
)

// A min heap
type Heap[T cmp.Ordered] struct {
	nums []T
}

func NewHeap[T cmp.Ordered]() *Heap[T] {
	return &Heap[T]{
		nums: make([]T, 0),
	}
}

func (h *Heap[T]) Insert(e T) {
	h.nums = append(h.nums, e)
	idx := len(h.nums) - 1
	parent_idx := (idx - 1) / 2
	for idx > 0 && h.nums[parent_idx] > h.nums[idx] {
		h.nums[parent_idx], h.nums[idx] = h.nums[idx], h.nums[parent_idx]

		idx = parent_idx
		parent_idx = (idx - 1) / 2
	}

}

func (h *Heap[T]) Pop() {
	if len(h.nums) == 0 {
		return
	}
	h.nums[0] = h.nums[len(h.nums)-1]
	h.nums = h.nums[:len(h.nums)-1]
	h.sift_down(0)
}

func (h *Heap[T]) sift_down(index int) {
	// choose the smaller children to exchange, for better maintainence
	for {
		left := 2*index + 1
		right := 2*index + 2
		smaller := index
		if left < len(h.nums) && h.nums[left] < h.nums[smaller] {
			smaller = left
		}
		if right < len(h.nums) && h.nums[right] < h.nums[smaller] {
			smaller = right
		}
		// no need to swap
		if smaller == index {
			break
		}
		h.nums[smaller], h.nums[index] = h.nums[index], h.nums[smaller]
		index = smaller
	}

}

func (h *Heap[T]) Print() {
	fmt.Println(h.nums)
}

func main() {
	heap := NewHeap[int]()
	heap.Insert(3)
	heap.Insert(10)
	heap.Insert(1)
	heap.Insert(0)
	heap.Insert(999)
	heap.Pop()
	heap.Insert(0)
	heap.Pop()
	heap.Print()

}
