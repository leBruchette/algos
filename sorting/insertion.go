package sorting

import "golang.org/x/exp/constraints"

// InsertionSort sorts a slice of ordered types in ascending order using the insertion sort algorithm.
// It uses the default comparator for types that satisfy constraints.Ordered.
func InsertionSort[T constraints.Ordered](data []T) {
	sort[T](data, defaultComparator[T]())
}

// InsertionSortWithComparator sorts a slice of any type using the insertion sort algorithm.
// It takes a custom comparator function to define the sorting logic for types that do not satisfy constraints.Ordered.
// The comparator should return true if the first argument is "greater than" the second argument
// (or whatever custom logic is required for sorting).
func InsertionSortWithComparator[T any](data []T, comparator func(a, b T) bool) {
	sort[T](data, comparator)
}

// sort is the core implementation of the insertion sort algorithm.
// It iterates through the slice, and for each element, it places it in its correct position
// relative to the already sorted portion of the slice, using the provided comparator.
func sort[T any](data []T, comparator func(a T, b T) bool) {
	n := len(data)
	for i := 1; i < n; i++ {
		key := data[i]
		j := i - 1
		// Move elements of the sorted portion that are greater than the key
		// (as determined by the comparator) one position ahead.
		for j >= 0 && comparator(data[j], key) {
			data[j+1] = data[j]
			j--
		}
		// Place the key in its correct position.
		data[j+1] = key
	}
}
