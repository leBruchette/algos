package types

import (
	"golang.org/x/exp/constraints"
)

type HeapType int

const (
	MinHeap HeapType = iota
	MaxHeap
)

type Heap[T constraints.Ordered] struct {
	Data    []T
	compare func(a, b T) bool
}

// NewHeap creates a new min-heap from the provided elements.
func NewHeap[T constraints.Ordered](heapType HeapType, elements []T) *Heap[T] {
	comparator := &DefaultComparator[T]{}
	var comparison func(a, b T) bool
	switch heapType {
	case MinHeap:
		comparison = comparator.LessThan
	case MaxHeap:
		comparison = comparator.GreaterThan
	}
	heap := &Heap[T]{Data: elements, compare: comparison}
	heap.Heapify()

	return heap
}

// Insert adds a new item to the heap and maintains the heap property.
func (h *Heap[T]) Insert(item T) {
	h.Data = append(h.Data, item)
	h.bubbleUp(h.size() - 1)
}

// Delete removes the element at index i from the heap and maintains the heap property.
func (h *Heap[T]) Delete(i int) {
	// edge case: handle invalid (too large or negative) index inputs
	if i > h.size()-1 || i < 0 {
		return
	}
	// edge case: handle deleting from a single-element heap
	if h.size() == 1 && i == 0 {
		h.Data = []T{}
		return
	}
	// the element to delete is already at the end, simply resize
	if i == h.size()-1 {
		h.Data = h.Data[:h.size()-1]
		return
	} else {
		// otherwise move the element to delete to the end, resize and bubbleDown
		h.Data[i], h.Data[h.size()-1] = h.Data[h.size()-1], h.Data[i]
		h.Data = h.Data[:h.size()-1]
		h.bubbleDown(i)
	}
}

// Heapify builds the heap from the initial Data slice.
func (h *Heap[T]) Heapify() {
	for i := h.size() / 2; i >= 0; i-- {
		h.bubbleDown(i)
	}
}

// HeapifyToIndex builds the heap from the initial Data slice, considering elements up to maxIdx.
func (h *Heap[T]) HeapifyToIndex(maxIdx int) {
	for i := maxIdx; i >= 0; i-- {
		h.bubbleDownToIndex(i, maxIdx)
	}
}

// bubbleUp restores the heap property by moving the element at index i up.
func (h *Heap[T]) bubbleUp(i int) {
	if i < 1 {
		return
	}
	if h.compare(h.Data[i], h.Data[i/2]) || h.Data[i] == h.Data[i/2] {
		h.Data[i], h.Data[i/2] = h.Data[i/2], h.Data[i]
		h.bubbleUp(i / 2)
	}
}

// bubbleDown restores the heap property by moving the element at index i down.
func (h *Heap[T]) bubbleDown(i int) {
	smallest := i

	if left(i) < h.size() && h.compare(h.Data[left(i)], h.Data[smallest]) {
		smallest = left(i)
	}

	if right(i) < h.size() && h.compare(h.Data[right(i)], h.Data[smallest]) {
		smallest = right(i)
	}

	if smallest != i {
		h.Data[i], h.Data[smallest] = h.Data[smallest], h.Data[i]
		h.bubbleDown(smallest)
	}
}

// bubbleDownToIndex restores the heap property by moving the element at index i down, considering elements up to maxIdx.
func (h *Heap[T]) bubbleDownToIndex(i int, maxIdx int) {
	smallest := i

	if left(i) < maxIdx && h.compare(h.Data[left(i)], h.Data[smallest]) {
		smallest = left(i)
	}

	if right(i) < maxIdx && h.compare(h.Data[right(i)], h.Data[smallest]) {
		smallest = right(i)
	}

	if smallest != i {
		h.Data[i], h.Data[smallest] = h.Data[smallest], h.Data[i]
		h.bubbleDownToIndex(smallest, maxIdx)
	}
}

// isValidHeap checks if the heap property is maintained for the entire heap.
func (h *Heap[T]) IsValidHeap() bool {
	for i := 0; i < h.size(); i++ {
		leftIdx := left(i)
		rightIdx := right(i)
		if leftIdx < h.size() && h.compare(h.Data[leftIdx], h.Data[i]) {
			return false
		}
		if rightIdx < h.size() && h.compare(h.Data[rightIdx], h.Data[i]) {
			return false
		}
	}
	return true
}

// returns the number of elements in the heap.
func (h *Heap[T]) size() int {
	return len(h.Data)
}

// returns left child for a given index, converting from 0 to 1 based indexing
func left(i int) int {
	return 2*i + 1
}

// returns right child for a given index, converting from 0 to 1 based indexing
func right(i int) int {
	return left(i) + 1
}
