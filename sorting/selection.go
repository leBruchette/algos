package sorting

import "golang.org/x/exp/constraints"

func SelectionSort[T constraints.Ordered](items []T) {
	swapSort[T](items, 0, DefaultComparator[T]{})
}

func SelectionSortWithComparator[T any](items []T, comparator Comparator[T]) {
	swapSort[T](items, 0, comparator)
}

func swapSort[T any](items []T, from int, comparator Comparator[T]) {
	if len(items) <= 1 {
		return
	}

	for i := from; i < len(items); i++ {
		if comparator.LessThan(items[i], items[from]) || comparator.EqualTo(items[i], items[from]) {
			items[from], items[i] = items[i], items[from]
			swapSort[T](items, from+1, comparator)
		}
	}

	return
}
