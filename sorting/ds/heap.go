package ds

import (
	"github.com/lebruchette/algos/sorting"
	"golang.org/x/exp/constraints"
)

type HeapType int

const (
	MinHeap HeapType = iota
	MaxHeap
)

type Heap[T constraints.Ordered] struct {
	data    []T
	compare func(a, b T) bool
}

// NewHeap creates a new min-heap from the provided elements.
func NewHeap[T constraints.Ordered](heapType HeapType, elements []T) *Heap[T] {
	comparator := &sorting.DefaultComparator[T]{}
	var comparison func(a, b T) bool
	switch heapType {
	case MinHeap:
		comparison = comparator.LessThan
	case MaxHeap:
		comparison = comparator.GreaterThan
	}
	heap := &Heap[T]{data: elements, compare: comparison}
	heap.Heapify()

	return heap
}

// Insert adds a new item to the heap and maintains the heap property.
func (h *Heap[T]) Insert(item T) {
	h.data = append(h.data, item)
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
		h.data = []T{}
		return
	}
	// the element to delete is already at the end, simply resize
	if i == h.size()-1 {
		h.data = h.data[:h.size()-1]
		return
	} else {
		// otherwise move the element to delete to the end, resize and bubbleDown
		h.data[i], h.data[h.size()-1] = h.data[h.size()-1], h.data[i]
		h.data = h.data[:h.size()-1]
		h.bubbleDown(i)
	}
}

// Heapify builds the heap from the initial data slice.
func (h *Heap[T]) Heapify() {
	for i := h.size() / 2; i >= 0; i-- {
		h.bubbleDown(i)
	}
}

// bubbleUp restores the heap property by moving the element at index i up.
func (h *Heap[T]) bubbleUp(i int) {
	if i < 1 {
		return
	}
	if h.compare(h.data[i], h.data[i/2]) || h.data[i] == h.data[i/2] {
		h.data[i], h.data[i/2] = h.data[i/2], h.data[i]
		h.bubbleUp(i / 2)
	}
}

// bubbleDown restores the heap property by moving the element at index i down.
func (h *Heap[T]) bubbleDown(i int) {
	smallest := i

	if left(i) < h.size() && h.compare(h.data[left(i)], h.data[smallest]) {
		smallest = left(i)
	}

	if right(i) < h.size() && h.compare(h.data[right(i)], h.data[smallest]) {
		smallest = right(i)
	}
	if smallest != i {
		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		h.bubbleDown(smallest)
	}
}

// isValidHeap checks if the heap property is maintained for the entire heap.
func (h *Heap[T]) IsValidHeap() bool {
	for i := 0; i < h.size(); i++ {
		leftIdx := left(i)
		rightIdx := right(i)
		if leftIdx < h.size() && h.compare(h.data[leftIdx], h.data[i]) {
			return false
		}
		if rightIdx < h.size() && h.compare(h.data[rightIdx], h.data[i]) {
			return false
		}
	}
	return true
}

// returns the number of elements in the heap.
func (h *Heap[T]) size() int {
	return len(h.data)
}

// returns left child for a given index, converting from 0 to 1 based indexing
func left(i int) int {
	return 2*i + 1
}

// returns right child for a given index, converting from 0 to 1 based indexing
func right(i int) int {
	return left(i) + 1
}
