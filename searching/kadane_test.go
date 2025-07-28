// kadane_test.go
package searching

import "testing"

func TestKadaneSearch(t *testing.T) {
	tests := []struct {
		numbers  []int
		expected int
	}{
		{[]int{-5, 4, -1, 2, 1, -3, 4, -1}, 7},
		{[]int{2, -1, 2, 3, 4, -5}, 10},
		{[]int{-1, 3, -2, 5, -3, 2}, 6},
		{[]int{8, -19, 5, -4, 20}, 21},
		{[]int{-2, -1, -3, -4, -1, -2, -1, -5, -4}, 0},
		{[]int{1, -2, 3, 4, -1, 2, 1, -5, 4}, 9},
		{[]int{-2, -3, 4, -1, -2, 1, 5, -3}, 7},
		{[]int{-1, -2, -3, -4}, 0},
		{[]int{5, 4, -1, 7, 8}, 23},
		{[]int{}, 0},
		{[]int{0, 0, 0}, 0},
		{[]int{1, 2, 3, 4}, 10},
		{[]int{-2, 1}, 1},
	}

	for _, tt := range tests {
		result := KadaneSearch(tt.numbers)
		if result != tt.expected {
			t.Errorf("KadaneSearch(%v) = %d; want %d", tt.numbers, result, tt.expected)
		}
	}
}
