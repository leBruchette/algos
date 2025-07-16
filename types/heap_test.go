package types

import (
	"reflect"
	"testing"
)

func TestMinHeapWithSingleElement(t *testing.T) {
	heap := NewHeap[int](MinHeap, []int{10})
	if !heap.IsValidHeap() {
		t.Errorf("Heap property violated: %v", heap.Data)
	}

	expected := []int{10}
	if !reflect.DeepEqual(heap.Data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.Data)
	}
}

func TestMinHeapWithTwoElements(t *testing.T) {
	heap := NewHeap[int](MinHeap, []int{10, 5})
	if !heap.IsValidHeap() {
		t.Errorf("Heap property violated: %v", heap.Data)
	}

	expected := []int{5, 10}
	if !reflect.DeepEqual(heap.Data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.Data)
	}
}

func TestMinHeapWithMultipleElements(t *testing.T) {
	heap := NewHeap[int](MinHeap, []int{10, 15, 20, 5})
	if !heap.IsValidHeap() {
		t.Errorf("Heap property violated: %v", heap.Data)
	}

	expected := []int{5, 10, 20, 15}
	if !reflect.DeepEqual(heap.Data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.Data)
	}
}

func TestMinHeapWithAlreadyHeapifiedData(t *testing.T) {
	heap := NewHeap[int](MinHeap, []int{5, 10, 20, 15})
	if !heap.IsValidHeap() {
		t.Errorf("Heap property violated: %v", heap.Data)
	}

	expected := []int{5, 10, 20, 15}
	if !reflect.DeepEqual(heap.Data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.Data)
	}
}

func TestMinHeapWithEmptyHeap(t *testing.T) {
	heap := NewHeap[int](MinHeap, []int{})
	if !heap.IsValidHeap() {
		t.Errorf("Heap property violated: %v", heap.Data)
	}

	expected := []int{}
	if !reflect.DeepEqual(heap.Data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.Data)
	}
}

func TestMinHeapWithMixedIntegers(t *testing.T) {
	heap := NewHeap[int](MinHeap, []int{0, -5, 10, -10})
	if !heap.IsValidHeap() {
		t.Errorf("Heap property violated: %v", heap.Data)
	}

	expected := []int{-10, -5, 10, 0}
	if !reflect.DeepEqual(heap.Data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.Data)
	}
}

func TestMinHeapWithDuplicateValues(t *testing.T) {
	heap := NewHeap[int](MinHeap, []int{10, 10, 5, 5})
	if !heap.IsValidHeap() {
		t.Errorf("Heap property violated: %v", heap.Data)
	}

	expected := []int{5, 10, 5, 10}
	if !reflect.DeepEqual(heap.Data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.Data)
	}
}

func TestMinHeapBubbleDownWithNegativeValues(t *testing.T) {
	heap := NewHeap[int](MinHeap, []int{-10, -20, -5, -15})
	if !heap.IsValidHeap() {
		t.Errorf("Heap property violated: %v", heap.Data)
	}

	expected := []int{-20, -15, -5, -10}
	if !reflect.DeepEqual(heap.Data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.Data)
	}
}

func TestMinHeapBubbleDownWithLargeHeap(t *testing.T) {
	heap := NewHeap[int](MinHeap, []int{50, 40, 30, 20, 10, 5, 15})
	if !heap.IsValidHeap() {
		t.Errorf("Heap property violated: %v", heap.Data)
	}

	expected := []int{5, 10, 15, 20, 40, 30, 50}
	if !reflect.DeepEqual(heap.Data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.Data)
	}
}

func TestMinHeapInsertSingleElement(t *testing.T) {
	heap := NewHeap[int](MinHeap, []int{})
	heap.Insert(10)
	if !heap.IsValidHeap() {
		t.Errorf("Heap property violated: %v", heap.Data)
	}

	expected := []int{10}
	if !reflect.DeepEqual(heap.Data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.Data)
	}
}

func TestMinHeapInsertMultipleElements(t *testing.T) {
	heap := NewHeap[int](MinHeap, []int{})
	heap.Insert(10)
	heap.Insert(5)
	heap.Insert(20)
	if !heap.IsValidHeap() {
		t.Errorf("Heap property violated: %v", heap.Data)
	}

	expected := []int{5, 10, 20}
	if !reflect.DeepEqual(heap.Data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.Data)
	}
}

func TestMinHeapInsertWithExistingHeap(t *testing.T) {
	heap := NewHeap[int](MinHeap, []int{5, 10, 20})
	heap.Insert(15)
	if !heap.IsValidHeap() {
		t.Errorf("Heap property violated: %v", heap.Data)
	}

	expected := []int{5, 10, 20, 15}
	if !reflect.DeepEqual(heap.Data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.Data)
	}
}

func TestMinHeapInsertWithDuplicateValues(t *testing.T) {
	heap := NewHeap[int](MinHeap, []int{5, 10, 5, 10})
	heap.Insert(5)
	if !heap.IsValidHeap() {
		t.Errorf("Heap property violated: %v", heap.Data)
	}

	expected := []int{5, 5, 10, 10, 5}
	if !reflect.DeepEqual(heap.Data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.Data)
	}
}

func TestMinHeapInsertWithNegativeValues(t *testing.T) {
	heap := NewHeap[int](MinHeap, []int{5, 10, 20})
	heap.Insert(-15)
	if !heap.IsValidHeap() {
		t.Errorf("Heap property violated: %v", heap.Data)
	}

	expected := []int{-15, 5, 20, 10}
	if !reflect.DeepEqual(heap.Data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.Data)
	}
}

func TestMinHeapInsertWithLargeHeap(t *testing.T) {
	heap := NewHeap[int](MinHeap, []int{5, 10, 20, 15, 25, 30, 35})
	heap.Insert(40)
	if !heap.IsValidHeap() {
		t.Errorf("Heap property violated: %v", heap.Data)
	}

	expected := []int{5, 10, 20, 15, 25, 30, 35, 40}
	if !reflect.DeepEqual(heap.Data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.Data)
	}
}

func TestMinHeapDeleteSingleElement(t *testing.T) {
	heap := NewHeap[int](MinHeap, []int{10})
	heap.Delete(0)
	if !heap.IsValidHeap() {
		t.Errorf("Heap property violated: %v", heap.Data)
	}

	expected := []int{}
	if !reflect.DeepEqual(heap.Data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.Data)
	}
}

func TestMinHeapDeleteFirstElement(t *testing.T) {
	heap := NewHeap[int](MinHeap, []int{10, 5, 20})
	heap.Delete(0)
	if !heap.IsValidHeap() {
		t.Errorf("Heap property violated: %v", heap.Data)
	}

	expected := []int{10, 20}
	if !reflect.DeepEqual(heap.Data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.Data)
	}
}

func TestMinHeapDeleteMiddleElement(t *testing.T) {
	heap := NewHeap[int](MinHeap, []int{10, 5, 20})
	heap.Delete(1)
	if !heap.IsValidHeap() {
		t.Errorf("Heap property violated: %v", heap.Data)
	}

	expected := []int{5, 20}
	if !reflect.DeepEqual(heap.Data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.Data)
	}
}

func TestMinHeapDeleteLastElement(t *testing.T) {
	heap := NewHeap[int](MinHeap, []int{10, 5, 20})
	heap.Delete(2)
	if !heap.IsValidHeap() {
		t.Errorf("Heap property violated: %v", heap.Data)
	}

	expected := []int{5, 10}
	if !reflect.DeepEqual(heap.Data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.Data)
	}
}

func TestMinHeapDeleteOutOfBoundsIndex(t *testing.T) {
	heap := NewHeap[int](MinHeap, []int{10, 5, 20})
	heap.Delete(3) // Index out of bounds
	if !heap.IsValidHeap() {
		t.Errorf("Heap property violated: %v", heap.Data)
	}

	expected := []int{5, 10, 20}
	if !reflect.DeepEqual(heap.Data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.Data)
	}
}

func TestMinHeapDeleteNegativeIndex(t *testing.T) {
	heap := NewHeap[int](MinHeap, []int{10, 5, 20})
	heap.Delete(-1) // Negative index
	if !heap.IsValidHeap() {
		t.Errorf("Heap property violated: %v", heap.Data)
	}

	expected := []int{5, 10, 20}
	if !reflect.DeepEqual(heap.Data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.Data)
	}
}

func TestMaxHeapWithSingleElement(t *testing.T) {
	heap := NewHeap[int](MaxHeap, []int{10})
	if !heap.IsValidHeap() {
		t.Errorf("Heap property violated: %v", heap.Data)
	}

	expected := []int{10}
	if !reflect.DeepEqual(heap.Data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.Data)
	}
}

func TestMaxHeapWithTwoElements(t *testing.T) {
	heap := NewHeap[int](MaxHeap, []int{10, 15})
	if !heap.IsValidHeap() {
		t.Errorf("Heap property violated: %v", heap.Data)
	}

	expected := []int{15, 10}
	if !reflect.DeepEqual(heap.Data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.Data)
	}
}

func TestMaxHeapWithMultipleElements(t *testing.T) {
	heap := NewHeap[int](MaxHeap, []int{10, 5, 20, 25})
	if !heap.IsValidHeap() {
		t.Errorf("Heap property violated: %v", heap.Data)
	}

	expected := []int{25, 10, 20, 5}
	if !reflect.DeepEqual(heap.Data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.Data)
	}
}

func TestMaxHeapWithAlreadyHeapifiedData(t *testing.T) {
	heap := NewHeap[int](MaxHeap, []int{25, 10, 20, 5})
	if !heap.IsValidHeap() {
		t.Errorf("Heap property violated: %v", heap.Data)
	}

	expected := []int{25, 10, 20, 5}
	if !reflect.DeepEqual(heap.Data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.Data)
	}
}

func TestMaxHeapWithEmptyHeap(t *testing.T) {
	heap := NewHeap[int](MaxHeap, []int{})
	if !heap.IsValidHeap() {
		t.Errorf("Heap property violated: %v", heap.Data)
	}

	expected := []int{}
	if !reflect.DeepEqual(heap.Data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.Data)
	}
}

func TestMaxHeapInsertSingleElement(t *testing.T) {
	heap := NewHeap[int](MaxHeap, []int{})
	heap.Insert(10)
	if !heap.IsValidHeap() {
		t.Errorf("Heap property violated: %v", heap.Data)
	}

	expected := []int{10}
	if !reflect.DeepEqual(heap.Data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.Data)
	}
}

func TestMaxHeapInsertMultipleElements(t *testing.T) {
	heap := NewHeap[int](MaxHeap, []int{})
	heap.Insert(10)
	heap.Insert(15)
	heap.Insert(5)
	if !heap.IsValidHeap() {
		t.Errorf("Heap property violated: %v", heap.Data)
	}

	expected := []int{15, 10, 5}
	if !reflect.DeepEqual(heap.Data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.Data)
	}
}

func TestMaxHeapInsertWithExistingHeap(t *testing.T) {
	heap := NewHeap[int](MaxHeap, []int{20, 10, 5})
	heap.Insert(15)
	if !heap.IsValidHeap() {
		t.Errorf("Heap property violated: %v", heap.Data)
	}

	expected := []int{20, 15, 5, 10}
	if !reflect.DeepEqual(heap.Data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.Data)
	}
}

func TestMaxHeapInsertWithDuplicateValues(t *testing.T) {
	heap := NewHeap[int](MaxHeap, []int{10, 5, 10, 5})
	heap.Insert(10)
	if !heap.IsValidHeap() {
		t.Errorf("Heap property violated: %v", heap.Data)
	}

	expected := []int{10, 10, 5, 5, 10}
	if !reflect.DeepEqual(heap.Data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.Data)
	}
}

func TestMaxHeapInsertWithNegativeValues(t *testing.T) {
	heap := NewHeap[int](MaxHeap, []int{5, 10, 20})
	heap.Insert(-15)
	if !heap.IsValidHeap() {
		t.Errorf("Heap property violated: %v", heap.Data)
	}

	expected := []int{20, 10, 5, -15}
	if !reflect.DeepEqual(heap.Data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.Data)
	}
}

func TestMaxHeapInsertWithLargeHeap(t *testing.T) {
	heap := NewHeap[int](MaxHeap, []int{5, 10, 20, 15, 25, 30, 35})
	heap.Insert(40)
	if !heap.IsValidHeap() {
		t.Errorf("Heap property violated: %v", heap.Data)
	}

	expected := []int{40, 35, 30, 25, 10, 5, 20, 15}
	if !reflect.DeepEqual(heap.Data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.Data)
	}
}

func TestMaxHeapDeleteSingleElement(t *testing.T) {
	heap := NewHeap[int](MaxHeap, []int{10})
	heap.Delete(0)
	if !heap.IsValidHeap() {
		t.Errorf("Heap property violated: %v", heap.Data)
	}

	expected := []int{}
	if !reflect.DeepEqual(heap.Data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.Data)
	}
}

func TestMaxHeapDeleteFirstElementMaxHeap(t *testing.T) {
	heap := NewHeap[int](MaxHeap, []int{20, 10, 5})
	heap.Delete(0)
	if !heap.IsValidHeap() {
		t.Errorf("Heap property violated: %v", heap.Data)
	}

	expected := []int{10, 5}
	if !reflect.DeepEqual(heap.Data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.Data)
	}
}

func TestMaxHeapDeleteMiddleElement(t *testing.T) {
	heap := NewHeap[int](MaxHeap, []int{20, 10, 5})
	heap.Delete(1)
	if !heap.IsValidHeap() {
		t.Errorf("Heap property violated: %v", heap.Data)
	}

	expected := []int{20, 5}
	if !reflect.DeepEqual(heap.Data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.Data)
	}
}

func TestMaxHeapDeleteLastElement(t *testing.T) {
	heap := NewHeap[int](MaxHeap, []int{20, 10, 5})
	heap.Delete(2)
	if !heap.IsValidHeap() {
		t.Errorf("Heap property violated: %v", heap.Data)
	}

	expected := []int{20, 10}
	if !reflect.DeepEqual(heap.Data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.Data)
	}
}

func TestMaxHeapDeleteOutOfBoundsIndex(t *testing.T) {
	heap := NewHeap[int](MaxHeap, []int{20, 10, 5})
	heap.Delete(3) // Index out of bounds
	if !heap.IsValidHeap() {
		t.Errorf("Heap property violated: %v", heap.Data)
	}

	expected := []int{20, 10, 5}
	if !reflect.DeepEqual(heap.Data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.Data)
	}
}

func TestMaxHeapDeleteNegativeIndex(t *testing.T) {
	heap := NewHeap[int](MaxHeap, []int{20, 10, 5})
	heap.Delete(-1) // Negative index
	if !heap.IsValidHeap() {
		t.Errorf("Heap property violated: %v", heap.Data)
	}

	expected := []int{20, 10, 5}
	if !reflect.DeepEqual(heap.Data, expected) {
		t.Errorf("Expected %v, but got %v", expected, heap.Data)
	}
}
