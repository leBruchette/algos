package partitioning

func HoarePartition(elements []int) {
	if len(elements) <= 1 {
		return
	}

	i, j, pivot := 0, len(elements)-1, elements[0]
	// use two pointers to find
	// - first element larger than pivot (from left)
	// - first element smaller than pivot (from right)
	// then swap.  repeat until pointers have overtaken each other
	for {

		for elements[i] < pivot {
			i++
		}
		for elements[j] > pivot {
			j--
		}

		if i >= j {
			break
		}

		elements[i], elements[j] = elements[j], elements[i]
		i++
		j--
	}

}
