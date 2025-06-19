package sorting

import (
	"golang.org/x/exp/constraints"
	"time"
)

// defaultComparator provides a default comparator for ordered types.
func defaultComparator[T constraints.Ordered]() func(a, b T) bool {
	return func(a, b T) bool {
		return a > b
	}
}

// placing test dependencies here for now.
type Person struct {
	Name string
	Dob  time.Time
}
