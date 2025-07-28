// kmp_search_test.go
package searching

import (
	"reflect"
	"testing"
)

func TestKmpSearch(t *testing.T) {
	tests := []struct {
		text      string
		subString string
		expected  []int
	}{
		{"abcxabcdabxabcdabcdabcy", "abcdabcy", []int{15}},
		{"aaaaa", "aa", []int{0, 1, 2, 3}},
		{"abcabcabc", "abc", []int{0, 3, 6}},
		{"abcabcabc", "abcd", []int{}},
		{"", "a", []int{}},
		{"a", "", []int{}},
		{"", "", []int{}},
		{"abc", "abc", []int{0}},
		{"abc", "bc", []int{1}},
		{"abc", "d", []int{}},
		{"mississippi", "issi", []int{1, 4}},
		{"abababab", "abab", []int{0, 2, 4}},
		{"abcde", "e", []int{4}},
		{"abcde", "a", []int{0}},
		{"abcde", "abcde", []int{0}},
		{"abcabcabcabc", "cab", []int{2, 5, 8}},
		{"xyzxyzxyz", "yzx", []int{1, 4}},
		{"aabaaabaaac", "aabaaac", []int{4}},
	}

	for _, tt := range tests {
		result := KmpSearch(tt.text, tt.subString)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("KmpSearch(%q, %q) = %v; want %v", tt.text, tt.subString, result, tt.expected)
		}
	}
}
