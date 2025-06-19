# `sorting` Package

The `sorting` package provides various implementations for sorting slices of data in Go. More specifically: 
- **Generic Sorting**: Sort slices of any type that satisfies the `constraints.Ordered` interface (e.g., integers, floats, strings).
- **Custom Comparator Support**: Sort slices of any type using a user-defined comparator function.
- **Stable Sorting**: Maintains the relative order of equal elements.

### Functions
### `InsertionSort`
Sorts a slice of ordered types in ascending order using insertion sort.

```go
func InsertionSort[T constraints.Ordered](data []T)
```

#### Parameters:
- `data`: A slice of any type that satisfies `constraints.Ordered`.

#### Example:
```go
data := []int{5, 2, 9, 1, 5, 6}
sorting.InsertionSort(data)
// data is now: []int{1, 2, 5, 5, 6, 9}
```

---

### `InsertionSortWithComparator`
Sorts a slice of any type using a custom comparator type.  This type must implement the interface `Comparator`

```go
func InsertionSortWithComparator[T any](data []T, comparator Comparator[T])
```

#### Parameters:
- `data`: A slice of any type.
- `comparator`: A function that defines comparison logic via `Comparator`. 

#### Example:
```go
type Person struct {
    Name string
    Dob  time.Time
}

data := []Person{
    {Name: "Bob", Dob: time.Date(1985, time.March, 15, 0, 0, 0, 0, time.UTC)}
    {Name: "Frank", Dob: time.Date(1992, time.September, 25, 0, 0, 0, 0, time.UTC)}
    {Name: "Alice", Dob: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)}
}

// from sorting/types.go
type Comparator[T any] interface {
    GreaterThan(a, b T) bool
    LessThan(a, b T) bool
    EqualTo(a, b T) bool
}

// implementing Comparator for Person, comparing dates of birth
type AgeComparator struct{}
func (ac AgeComparator) GreaterThan(a, b Person) bool { return a.Dob.After(b.Dob) }
func (ac AgeComparator) LessThan(a, b Person) bool    { return a.Dob.Before(b.Dob) }
func (ac AgeComparator) EqualTo(a, b Person) bool     { return a.Dob.Equal(b.Dob) }

// data is now sorted by age in ascending order
sorting.InsertionSortWithComparator(data, AgeComparator{})

```

---

### `MergeSort`
Sorts a slice of ordered types in ascending order using merge sort.

```go
func MergeSort[T constraints.Ordered](data []T)
```

#### Parameters:
- `data`: A slice of any type that satisfies `constraints.Ordered`.

#### Example:
```go
data := []int{5, 2, 9, 1, 5, 6}
sorting.MergeSort(data)
// data is now: []int{1, 2, 5, 5, 6, 9}
```

---

### `MergeSortWithComparator`
Sorts a slice of any type using a custom comparator type.  This type must implement the interface `Comparator`

```go
func MergeSortWithComparator[T any](data []T, comparator Comparator[T])
```

#### Parameters:
- `data`: A slice of any type.
- `comparator`: A function that defines comparison logic via `Comparator`.

#### Example:
```go
type Person struct {
    Name string
    Dob  time.Time
}

data := []Person{
    {Name: "Bob", Dob: time.Date(1985, time.March, 15, 0, 0, 0, 0, time.UTC)}
    {Name: "Frank", Dob: time.Date(1992, time.September, 25, 0, 0, 0, 0, time.UTC)}
    {Name: "Alice", Dob: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)}
}

// from sorting/types.go
type Comparator[T any] interface {
    GreaterThan(a, b T) bool
    LessThan(a, b T) bool
    EqualTo(a, b T) bool
}

// implementing Comparator for Person, comparing dates of birth
type AgeComparator struct{}
func (ac AgeComparator) GreaterThan(a, b Person) bool { return a.Dob.After(b.Dob) }
func (ac AgeComparator) LessThan(a, b Person) bool    { return a.Dob.Before(b.Dob) }
func (ac AgeComparator) EqualTo(a, b Person) bool     { return a.Dob.Equal(b.Dob) }

// data is now sorted by age in ascending order
sorting.MergeSortWithComparator(data, AgeComparator{})

```

---

### `QuickSort`
Sorts a slice of ordered types in ascending order using insertion sort.

```go
func QuickSort[T constraints.Ordered](data []T)
```

#### Parameters:
- `data`: A slice of any type that satisfies `constraints.Ordered`.

#### Example:
```go
data := []int{5, 2, 9, 1, 5, 6}
sorting.InsertionSort(data)
// data is now: []int{1, 2, 5, 5, 6, 9}
```

---

### `QuickSortWithComparator`
Sorts a slice of any type using a custom comparator type.  This type must implement the interface `Comparator`

```go
func QuickSortWithComparator[T any](data []T, comparator Comparator[T])
```

#### Parameters:
- `data`: A slice of any type.
- `comparator`: A function that defines comparison logic via `Comparator`.

#### Example:
```go
type Person struct {
    Name string
    Dob  time.Time
}

data := []Person{
    {Name: "Bob", Dob: time.Date(1985, time.March, 15, 0, 0, 0, 0, time.UTC)}
    {Name: "Frank", Dob: time.Date(1992, time.September, 25, 0, 0, 0, 0, time.UTC)}
    {Name: "Alice", Dob: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)}
}

// from sorting/types.go
type Comparator[T any] interface {
    GreaterThan(a, b T) bool
    LessThan(a, b T) bool
    EqualTo(a, b T) bool
}

// implementing Comparator for Person, comparing dates of birth
type AgeComparator struct{}
func (ac AgeComparator) GreaterThan(a, b Person) bool { return a.Dob.After(b.Dob) }
func (ac AgeComparator) LessThan(a, b Person) bool    { return a.Dob.Before(b.Dob) }
func (ac AgeComparator) EqualTo(a, b Person) bool     { return a.Dob.Equal(b.Dob) }

// data is now sorted by age in ascending order
sorting.QuickSortWithComparator(data, AgeComparator{})

```

---

## Installation

To use the `algos/sorting` package, add it to your Go project by including it in your `go.mod` file:

```bash
go get github.com/lebruchette/algos
```

Ensure you have Go version `1.23.3` or later.

## Testing

The package includes comprehensive unit tests to verify its functionality. To run the tests, use the following command:

```bash
go test ./sorting
```

## License

This package is licensed under the MIT License. See the LICENSE file for details.