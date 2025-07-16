package sorting

import (
	"github.com/lebruchette/algos/types"
	"reflect"
	"testing"
	"time"
)

func TestMergeSortWithInts(t *testing.T) {
	data := []int{5, 2, 9, 1, 5, 6}
	expected := []int{1, 2, 5, 5, 6, 9}

	MergeSort(data)

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v, but got %v", expected, data)
	}
}

func TestMergeSortWithSingleInt(t *testing.T) {
	data := []int{5}
	expected := []int{5}

	MergeSort(data)

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v, but got %v", expected, data)
	}
}

func TestMergeSortWithEmptyInt(t *testing.T) {
	var data []int
	var expected []int

	MergeSort(data)

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v, but got %v", expected, data)
	}
}

func TestMergeSortWithDuplicates(t *testing.T) {
	data := []int{4, 2, 2, 8, 3, 3, 1}
	expected := []int{1, 2, 2, 3, 3, 4, 8}

	MergeSort(data)

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v, but got %v", expected, data)
	}
}

func TestMergeSortWithAlreadySortedData(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6}
	expected := []int{1, 2, 3, 4, 5, 6}

	MergeSort(data)

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v, but got %v", expected, data)
	}
}

func TestMergeSortWithReverseSortedData(t *testing.T) {
	data := []int{6, 5, 4, 3, 2, 1}
	expected := []int{1, 2, 3, 4, 5, 6}

	MergeSort(data)

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v, but got %v", expected, data)
	}
}

func TestMergeSortWithMixedIntegers(t *testing.T) {
	data := []int{-3, 5, -1, 0, 2, -2}
	expected := []int{-3, -2, -1, 0, 2, 5}

	MergeSort(data)

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v, but got %v", expected, data)
	}
}

func TestMergeSortWithComparator(t *testing.T) {
	bob := types.Person{Name: "Bob", Dob: time.Date(1985, time.March, 15, 0, 0, 0, 0, time.UTC)}
	frank := types.Person{Name: "Frank", Dob: time.Date(1992, time.September, 25, 0, 0, 0, 0, time.UTC)}
	alice := types.Person{Name: "Alice", Dob: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)}
	diana := types.Person{Name: "Diana", Dob: time.Date(1995, time.December, 5, 0, 0, 0, 0, time.UTC)}
	charlie := types.Person{Name: "Charlie", Dob: time.Date(2000, time.July, 20, 0, 0, 0, 0, time.UTC)}
	eve := types.Person{Name: "Eve", Dob: time.Date(1988, time.April, 10, 0, 0, 0, 0, time.UTC)}

	data := []types.Person{bob, frank, alice, diana, charlie, eve}
	expected := []types.Person{bob, eve, alice, frank, diana, charlie}

	MergeSortWithComparator(data, types.PersonComparator{})

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v, but got %v", expected, data)
	}
}
