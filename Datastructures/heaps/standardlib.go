package heaps

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func runContainerHeap() {
	h := &IntHeap{}

	heap.Init(h)

	heap.Push(h, 1)
	heap.Push(h, 2)
	heap.Push(h, 3)
	heap.Push(h, 4)
	heap.Push(h, 5)
	heap.Push(h, 6)

	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}

	fmt.Println(h)
}
