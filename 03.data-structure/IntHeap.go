package main

// from https://pkg.go.dev/container/heap@go1.19
import (
	"container/heap"
	"fmt"
	"math/rand"
	"time"
)

// An IntHeap is a min-heap of ints.

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }

func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

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

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	for i := 0; i < random.Intn(20); i++ {
		heap.Push(h, random.Intn(100))
	}
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
