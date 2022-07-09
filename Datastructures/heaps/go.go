package heaps

import (
	"fmt"
	"time"
)

// Go has an inbuilt heap

// A heap is a tree with the property that each node is the minimum-valued node in its subtree.

type MinHeap struct {
	Heap []int
}

func (*MinHeap) Parent(index int) int {
	return (index - 1) / 2
}

func (h *MinHeap) Swap(first, second int) {
	temp := h.Heap[first]
	h.Heap[first] = h.Heap[second]
	h.Heap[second] = temp
}

func (h *MinHeap) Insert(value int) {
	h.Heap = append(h.Heap, value)
	h.upHeapify(len(h.Heap) - 1)
}

func (h *MinHeap) upHeapify(index int) {
	parentIdx := h.Parent(index)

	for h.Heap[index] < h.Heap[parentIdx] {
		h.Swap(index, parentIdx)
	}
}

func main() {
	start := time.Now()

	fmt.Println("Running main")

	mh := MinHeap{}

	mh.Insert(8)
	mh.Insert(7)
	mh.Insert(6)
	mh.Insert(5)
	mh.Insert(3)
	mh.Insert(2)

	fmt.Println(mh.Heap)

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
