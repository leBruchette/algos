package sorting

import (
	"github.com/lebruchette/algos/types"
	"golang.org/x/exp/constraints"
)

func HeapSort[T constraints.Ordered](data []T) {
	// create a max heap that satisfies properties of a max heap
	heap := types.NewHeap(types.MaxHeap, data)

	// swap the root node to the end and heapify
	for i := len(data) - 1; i > -1; i-- {
		heap.Data[0], heap.Data[i] = heap.Data[i], heap.Data[0]
		heap.HeapifyToIndex(i)
	}
}
