package sorting

import (
	"golang.org/x/exp/constraints"
)

func QuickSort[T constraints.Ordered](items []T) {
	copy(items, quickSort(items, DefaultComparator[T]{}))
}

func QuickSortWithComparator[T any](items []T, comparator Comparator[T]) {
	copy(items, quickSort(items, comparator))
}

func quickSort[T any](items []T, comparator Comparator[T]) []T {
	if len(items) <= 1 {
		return items
	}

	var left, middle, right []T
	pivot := len(items) / 2

	for i := range items {
		switch {
		case comparator.LessThan(items[i], items[pivot]):
			left = append(left, items[i])
		case comparator.EqualTo(items[i], items[pivot]):
			middle = append(middle, items[i])
		case comparator.GreaterThan(items[i], items[pivot]):
			right = append(right, items[i])
		}
	}

	return append(append(quickSort[T](left, comparator), middle...), quickSort[T](right, comparator)...)

}
