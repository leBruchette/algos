package types

import (
	"golang.org/x/exp/constraints"
	"time"
)

type Comparator[T any] interface {
	GreaterThan(a, b T) bool
	LessThan(a, b T) bool
	EqualTo(a, b T) bool
}

type DefaultComparator[T constraints.Ordered] struct{}

func (d DefaultComparator[T]) GreaterThan(a, b T) bool { return a > b }
func (d DefaultComparator[T]) LessThan(a, b T) bool    { return a < b }
func (d DefaultComparator[T]) EqualTo(a, b T) bool     { return a == b }

// ===========================================================================
// 						===== Test Dependencies =====
// ===========================================================================

type Person struct {
	Name string
	Dob  time.Time
}

type PersonComparator struct{}

func (d PersonComparator) GreaterThan(a, b Person) bool { return a.Dob.After(b.Dob) }
func (d PersonComparator) LessThan(a, b Person) bool    { return a.Dob.Before(b.Dob) }
func (d PersonComparator) EqualTo(a, b Person) bool     { return a.Dob.Equal(b.Dob) }
