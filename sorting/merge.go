package sorting

import (
	"golang.org/x/exp/constraints"
)

// MergeSort sorts a slice of ordered types in ascending order using the merge sort algorithm.
// It uses the default comparator for types that satisfy constraints.Ordered.
func MergeSort[T constraints.Ordered](items []T) {
	splitAndSort(items, 0, len(items)-1, DefaultComparator[T]{})
}

// MergeSortWithComparator sorts a slice of any type using the merge sort algorithm.
// It takes a custom comparator function to define the sorting logic for types that do not satisfy constraints.Ordered.
// The comparator should return true if the first argument is "greater than" the second argument
// (or whatever custom logic is required for sorting).
func MergeSortWithComparator[T any](items []T, comparator Comparator[T]) {
	splitAndSort(items, 0, len(items)-1, comparator)
}

// splitAndSort recursively divides the slice into halves, sorts each half, and merges them back together.
// It uses the provided comparator to determine the sorting order.
func splitAndSort[T any](items []T, left, right int, comparator Comparator[T]) {
	if left == right || len(items) < 1 {
		//single element, so return
		return
	} else if left+1 == right {
		// two elements...simple swap after comparison
		if comparator.GreaterThan(items[left], items[right]) {
			items[right], items[left] = items[left], items[right]
			return
		}
	} else {
		// keep halving the slice, recursively sorting left and right halves
		mid := (left + right) / 2
		splitAndSort[T](items, left, mid, comparator)
		splitAndSort[T](items, mid+1, right, comparator)

		// and merge each iteration
		merge(items, left, mid, right, comparator)
	}
}

// merge combines two sorted sub-slices into a single sorted slice.
// The left sub-slice is defined by indices [left, mid], and the right sub-slice is defined by indices [mid+1, right].
// It uses the provided comparator to determine the sorting order.
func merge[T any](items []T, left, mid, right int, comparator Comparator[T]) {
	lPtr, rPtr := left, mid+1

	tempSlice := make([]T, 0)
	for lPtr <= mid || rPtr <= right {
		// Both slices still have elements to compare, append the smaller item at the given pointers
		if lPtr <= mid && rPtr <= right {
			if comparator.LessThan(items[lPtr], items[rPtr]) {
				tempSlice = append(tempSlice, items[lPtr])
				lPtr++
			} else {
				tempSlice = append(tempSlice, items[rPtr])
				rPtr++
			}
		} else if lPtr > mid {
			// left side is fully traversed, append the rest of right side
			tempSlice = append(tempSlice, items[rPtr])
			rPtr++
		} else {
			// right is fully traversed, append the rest of left
			tempSlice = append(tempSlice, items[lPtr])
			lPtr++
		}
	}

	// overwrite original, unsorted indices with a sorted slice of the same elements
	copy(items[left:right+1], tempSlice)
}
