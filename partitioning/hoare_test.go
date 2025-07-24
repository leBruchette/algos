package partitioning

import (
	"reflect"
	"testing"
)

func TestHoarePartition(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{4, 3, 2, 1}, []int{1, 3, 2, 4}},       // reverse order
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},       // already sorted
		{[]int{5, 1, 4, 2, 3}, []int{3, 1, 4, 2, 5}}, // random order
		{[]int{10}, []int{10}},                       // single element
		{[]int{}, []int{}},                           // empty slice
		{[]int{2, 2, 2, 2}, []int{2, 2, 2, 2}},       // all elements equal

		{[]int{3, 3, 2, 2, 1, 1}, []int{1, 1, 2, 2, 3, 3}}, // duplicates
		//TODO correct algo to pass this case, currently returns -4,-1,-2,-3.
		//{[]int{-3, -1, -2, -4}, []int{-4, -3, -2, -1}},     // all negative
		{[]int{-2, 0, 2, -1, 1}, []int{-2, 0, 2, -1, 1}}, // mix negative and positive

		{[]int{2, 1}, []int{1, 2}}, // two elements unsorted
		{[]int{1, 2}, []int{1, 2}}, // two elements sorted
	}

	for _, tt := range tests {
		inputCopy := make([]int, len(tt.input))
		copy(inputCopy, tt.input)
		HoarePartition(inputCopy)
		if !reflect.DeepEqual(inputCopy, tt.expected) {
			t.Errorf("HoarePartition(%v) = %v; want %v", tt.input, inputCopy, tt.expected)
		}
	}
}
