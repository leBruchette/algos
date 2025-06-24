package ds

import (
	"log"
	"reflect"
	"testing"
)

func TestBubbleUpWithSingleElement(t *testing.T) {
	heap := &Heap[int]{data: []int{10}}
	heap.bubbleUp(0)
	expected := []int{10}

	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.data)
	}
}

func TestBubbleUpWithTwoElements(t *testing.T) {
	heap := &Heap[int]{data: []int{10, 5}}
	heap.bubbleUp(1)
	expected := []int{5, 10}

	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.data)
	}
}

func TestBubbleUpWithMultipleElements(t *testing.T) {
	heap := &Heap[int]{data: []int{10, 15, 20, 5}}
	heap.bubbleUp(3)
	expected := []int{5, 10, 20, 15}

	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.data)
	}
}

func TestBubbleUpWithAlreadyHeapifiedData(t *testing.T) {
	heap := &Heap[int]{data: []int{5, 10, 20, 15}}
	heap.bubbleUp(3)
	expected := []int{5, 10, 20, 15}

	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.data)
	}
}

func TestBubbleUpWithEmptyHeap(t *testing.T) {
	heap := &Heap[int]{data: []int{}}
	heap.bubbleUp(0)
	expected := []int{}

	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.data)
	}
}

func TestBubbleDownWithSingleElement(t *testing.T) {
	heap := &Heap[int]{data: []int{10}}
	heap.bubbleDown(0)
	expected := []int{10}

	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.data)
	}
}

func TestBubbleDownWithTwoElements(t *testing.T) {
	heap := &Heap[int]{data: []int{10, 5}}
	log.Printf("the current heap: %v", heap.data)
	heap.bubbleDown(0)
	expected := []int{5, 10}

	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.data)
	}
}

func TestBubbleDownWithAlreadyHeapifiedData(t *testing.T) {
	heap := &Heap[int]{data: []int{5, 10, 20, 15}}
	heap.bubbleDown(0)
	expected := []int{5, 10, 20, 15}

	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.data)
	}
}

func TestBubbleDownWithEmptyHeap(t *testing.T) {
	heap := &Heap[int]{data: []int{}}
	heap.bubbleDown(0)
	expected := []int{}

	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.data)
	}
}

func TestBubbleDownWithMixedIntegers(t *testing.T) {
	heap := &Heap[int]{data: []int{0, -5, 10, -10}}
	heap.bubbleDown(0)
	expected := []int{-5, -10, 10, 0}

	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.data)
	}
}

func TestBubbleDownWithDuplicateValues(t *testing.T) {
	heap := &Heap[int]{data: []int{10, 10, 5, 5}}
	heap.bubbleDown(0)
	expected := []int{5, 10, 10, 5}

	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.data)
	}
}

func TestBubbleDownWithNegativeValues(t *testing.T) {
	heap := &Heap[int]{data: []int{-10, -20, -5, -15}}
	heap.bubbleDown(0)
	expected := []int{-20, -15, -5, -10}

	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.data)
	}
}

func TestBubbleDownWithLargeHeap(t *testing.T) {
	heap := &Heap[int]{data: []int{50, 40, 30, 20, 10, 5, 15}}
	heap.bubbleDown(0)
	expected := []int{30, 40, 5, 20, 10, 50, 15}

	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.data)
	}
}

func TestInsertSingleElement(t *testing.T) {
	heap := &Heap[int]{data: []int{}}
	heap.Insert(10)
	expected := []int{10}

	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.data)
	}
}

func TestInsertMultipleElements(t *testing.T) {
	heap := &Heap[int]{data: []int{}}
	heap.Insert(10)
	heap.Insert(5)
	heap.Insert(20)
	expected := []int{5, 10, 20}

	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.data)
	}
}

func TestInsertWithExistingHeap(t *testing.T) {
	heap := &Heap[int]{data: []int{5, 10, 20}}
	heap.Insert(15)
	expected := []int{5, 10, 20, 15}

	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.data)
	}
}

func TestInsertWithDuplicateValues(t *testing.T) {
	heap := &Heap[int]{data: []int{5, 10, 5, 10}}
	heap.Insert(5)
	expected := []int{5, 10, 5, 10, 5}

	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.data)
	}
}

func TestInsertWithNegativeValues(t *testing.T) {
	heap := &Heap[int]{data: []int{5, 10, 20}}
	heap.Insert(-15)
	expected := []int{-15, 5, 20, 10}

	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.data)
	}
}

func TestInsertWithLargeHeap(t *testing.T) {
	heap := &Heap[int]{data: []int{5, 10, 20, 15, 25, 30, 35}}
	heap.Insert(40)
	expected := []int{5, 10, 20, 15, 25, 30, 35, 40}

	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.data)
	}
}

func TestDeleteSingleElement(t *testing.T) {
	heap := &Heap[int]{data: []int{10}}
	heap.Delete(0)
	expected := []int{}

	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.data)
	}
}

func TestDeleteFirstElement(t *testing.T) {
	heap := &Heap[int]{data: []int{10, 5, 20}}
	heap.Delete(0)
	expected := []int{5, 20}

	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.data)
	}
}

func TestDeleteMiddleElement(t *testing.T) {
	heap := &Heap[int]{data: []int{10, 5, 20}}
	heap.Delete(1)
	expected := []int{10, 20}

	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.data)
	}
}

func TestDeleteLastElement(t *testing.T) {
	heap := &Heap[int]{data: []int{10, 5, 20}}
	heap.Delete(2)
	expected := []int{10, 5}

	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.data)
	}
}

func TestDeleteOutOfBoundsIndex(t *testing.T) {
	heap := &Heap[int]{data: []int{10, 5, 20}}
	heap.Delete(3) // Index out of bounds
	expected := []int{10, 5, 20}

	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.data)
	}
}

func TestDeleteNegativeIndex(t *testing.T) {
	heap := &Heap[int]{data: []int{10, 5, 20}}
	heap.Delete(-1) // Negative index
	expected := []int{10, 5, 20}

	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.data)
	}
}

func TestHeapifyWithUnorderedData(t *testing.T) {
	heap := &Heap[int]{data: []int{20, 15, 10, 5}}
	heap.Heapify()
	expected := []int{5, 15, 10, 20}

	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.data)
	}
}

func TestHeapifyWithAlreadyHeapifiedData(t *testing.T) {
	heap := &Heap[int]{data: []int{5, 10, 20, 15}}
	heap.Heapify()
	expected := []int{5, 10, 20, 15}

	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.data)
	}
}

func TestHeapifyWithEmptyHeap(t *testing.T) {
	heap := &Heap[int]{data: []int{}}
	heap.Heapify()
	expected := []int{}

	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.data)
	}
}

func TestHeapifyWithSingleElement(t *testing.T) {
	heap := &Heap[int]{data: []int{10}}
	heap.Heapify()
	expected := []int{10}

	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.data)
	}
}

func TestHeapifyWithDuplicateValues(t *testing.T) {
	heap := &Heap[int]{data: []int{10, 10, 5, 5}}
	heap.Heapify()
	expected := []int{5, 10, 5, 10}

	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.data)
	}
}

func TestHeapifyWithNegativeValues(t *testing.T) {
	heap := &Heap[int]{data: []int{-10, -20, -5, -15}}
	heap.Heapify()
	expected := []int{-20, -15, -5, -10}

	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.data)
	}
}

func TestHeapifyWithMixedIntegers(t *testing.T) {
	heap := &Heap[int]{data: []int{0, -5, 10, -10}}
	heap.Heapify()
	expected := []int{-10, -5, 10, 0}

	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.data)
	}
}

func TestHeapifyWithLargeHeap(t *testing.T) {
	heap := &Heap[int]{data: []int{50, 40, 30, 20, 10, 5, 15}}
	heap.Heapify()
	expected := []int{5, 10, 15, 20, 40, 30, 50}

	if !reflect.DeepEqual(heap.data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.data)
	}
}
