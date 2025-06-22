package sorting

import (
	"reflect"
	"testing"
	"time"
)

func TestSelectionSortWithInts(t *testing.T) {
	data := []int{5, 2, 9, 1, 5, 6}
	expected := []int{1, 2, 5, 5, 6, 9}

	SelectionSort(data)

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v, but got %v", expected, data)
	}
}

func TestSelectionSortWithSingleInt(t *testing.T) {
	data := []int{5}
	expected := []int{5}

	SelectionSort(data)

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v, but got %v", expected, data)
	}
}

func TestSelectionSortWithEmptySlice(t *testing.T) {
	var data []int
	var expected []int

	SelectionSort(data)

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v, but got %v", expected, data)
	}
}

func TestSelectionSortWithDuplicates(t *testing.T) {
	data := []int{4, 2, 2, 8, 3, 3, 1}
	expected := []int{1, 2, 2, 3, 3, 4, 8}

	SelectionSort(data)

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v, but got %v", expected, data)
	}
}

func TestSelectionSortWithAlreadySortedData(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6}
	expected := []int{1, 2, 3, 4, 5, 6}

	SelectionSort(data)

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v, but got %v", expected, data)
	}
}

func TestSelectionSortWithReverseSortedData(t *testing.T) {
	data := []int{6, 5, 4, 3, 2, 1}
	expected := []int{1, 2, 3, 4, 5, 6}

	SelectionSort(data)

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v, but got %v", expected, data)
	}
}

func TestSelectionSortWithMixedIntegers(t *testing.T) {
	data := []int{-3, 5, -1, 0, 2, -2}
	expected := []int{-3, -2, -1, 0, 2, 5}

	SelectionSort(data)

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v, but got %v", expected, data)
	}
}

func TestSelectionSortWithCustomComparator(t *testing.T) {
	alice := Person{Name: "Alice", Dob: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)}
	bob := Person{Name: "Bob", Dob: time.Date(1985, time.March, 15, 0, 0, 0, 0, time.UTC)}
	charlie := Person{Name: "Charlie", Dob: time.Date(2000, time.July, 20, 0, 0, 0, 0, time.UTC)}

	data := []Person{charlie, alice, bob}
	expected := []Person{bob, alice, charlie}

	SelectionSortWithComparator(data, &PersonComparator{})

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v, but got %v", expected, data)
	}
}
