package sorting

import (
	"reflect"
	"testing"
)

func TestHeapSortInts(t *testing.T) {
	data := []int{5, 2, 9, 1, 5, 6}
	expected := []int{1, 2, 5, 5, 6, 9}

	HeapSort(data)

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v, got %v", expected, data)
	}
}

func TestHeapSortEmpty(t *testing.T) {
	var data []int
	var expected []int
	HeapSort(data)
	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v, got %v", expected, data)
	}
}

func TestHeapSortSingle(t *testing.T) {
	data := []int{42}
	expected := []int{42}
	HeapSort(data)
	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v, got %v", expected, data)
	}
}

func TestHeapSortAlreadySorted(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 3, 4, 5}
	HeapSort(data)
	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v, got %v", expected, data)
	}
}

func TestHeapSortReverseSorted(t *testing.T) {
	data := []int{5, 4, 3, 2, 1}
	expected := []int{1, 2, 3, 4, 5}
	HeapSort(data)
	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v, got %v", expected, data)
	}
}

func TestHeapSortWithDuplicates(t *testing.T) {
	data := []int{2, 3, 2, 1, 3, 1}
	expected := []int{1, 1, 2, 2, 3, 3}
	HeapSort(data)
	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v, got %v", expected, data)
	}
}

func TestHeapSortNegativeInts(t *testing.T) {
	data := []int{-3, 5, -1, 0, 2, -2}
	expected := []int{-3, -2, -1, 0, 2, 5}
	HeapSort(data)
	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v, got %v", expected, data)
	}
}
