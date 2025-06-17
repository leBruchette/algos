package sorting

import (
	"reflect"
	"testing"
	"time"
)

func TestInsertionSortWithInts(t *testing.T) {
	data := []int{5, 2, 9, 1, 5, 6}
	expected := []int{1, 2, 5, 5, 6, 9}

	InsertionSort(data)

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v, but got %v", expected, data)
	}
}

func TestInsertionSortWithSingleInt(t *testing.T) {
	data := []int{5}
	expected := []int{5}

	InsertionSort(data)

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v, but got %v", expected, data)
	}
}

func TestInsertionSortWithEmptyInt(t *testing.T) {
	var data []int
	var expected []int

	InsertionSort(data)

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v, but got %v", expected, data)
	}
}

func TestInsertionSortWithDuplicates(t *testing.T) {
	data := []int{4, 2, 2, 8, 3, 3, 1}
	expected := []int{1, 2, 2, 3, 3, 4, 8}

	InsertionSort(data)

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v, but got %v", expected, data)
	}
}

func TestInsertionSortWithAlreadySortedData(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6}
	expected := []int{1, 2, 3, 4, 5, 6}

	InsertionSort(data)

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v, but got %v", expected, data)
	}
}

func TestInsertionSortWithReverseSortedData(t *testing.T) {
	data := []int{6, 5, 4, 3, 2, 1}
	expected := []int{1, 2, 3, 4, 5, 6}

	InsertionSort(data)

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v, but got %v", expected, data)
	}
}

func TestInsertionSortWithMixedIntegers(t *testing.T) {
	data := []int{-3, 5, -1, 0, 2, -2}
	expected := []int{-3, -2, -1, 0, 2, 5}

	InsertionSort(data)

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v, but got %v", expected, data)
	}
}

type Person struct {
	Name string
	Dob  time.Time
}

func TestInsertionSortWithComparator(t *testing.T) {
	bob := Person{Name: "Bob", Dob: time.Date(1985, time.March, 15, 0, 0, 0, 0, time.UTC)}
	frank := Person{Name: "Frank", Dob: time.Date(1992, time.September, 25, 0, 0, 0, 0, time.UTC)}
	alice := Person{Name: "Alice", Dob: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)}
	diana := Person{Name: "Diana", Dob: time.Date(1995, time.December, 5, 0, 0, 0, 0, time.UTC)}
	charlie := Person{Name: "Charlie", Dob: time.Date(2000, time.July, 20, 0, 0, 0, 0, time.UTC)}
	eve := Person{Name: "Eve", Dob: time.Date(1988, time.April, 10, 0, 0, 0, 0, time.UTC)}

	dobComparator := func(a, b Person) bool {
		return a.Dob.After(b.Dob)
	}

	data := []Person{bob, frank, alice, diana, charlie, eve}
	expected := []Person{bob, eve, alice, frank, diana, charlie}

	InsertionSortWithComparator(data, dobComparator)

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v, but got %v", expected, data)
	}
}
