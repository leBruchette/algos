# `sorting` Package

The `sorting` package provides various implementations for sorting slices of data in Go. More specifically: 
- **Generic Sorting**: Sort slices of any type that satisfies the `constraints.Ordered` interface (e.g., integers, floats, strings).
- **Custom Comparator Support**: Sort slices of any type using a user-defined comparator function.
- **Stable Sorting**: Maintains the relative order of equal elements.

### Functions
### `InsertionSort`
Sorts a slice of ordered types in ascending order using the insertion sort algorithm.

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
Sorts a slice of any type using a custom comparator function.

```go
func InsertionSortWithComparator[T any](data []T, comparator func(a, b T) bool)
```

#### Parameters:
- `data`: A slice of any type.
- `comparator`: A function that defines the sorting logic. It should return `true` if the first argument is "greater than" the second argument.

#### Example:
```go
type Person struct {
    Name string
    Age  int
}

data := []Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
    {Name: "Charlie", Age: 35},
}

ageComparator := func(a, b Person) bool {
    return a.Age > b.Age
}

sorting.InsertionSortWithComparator(data, ageComparator)
// data is now sorted by age in descending order
```

---

### `MergeSort`
Sorts a slice of ordered types in ascending order using a merge sort algorithm.

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
Sorts a slice of any type using a custom comparator function.

```go
func MergeSortWithComparator[T any](data []T, comparator func(a, b T) bool)
```

#### Parameters:
- `data`: A slice of any type.
- `comparator`: A function that defines the sorting logic. It should return `true` if the first argument is "greater than" the second argument.

#### Example:
```go
type Person struct {
    Name string
    Age  int
}

data := []Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
    {Name: "Charlie", Age: 35},
}

ageComparator := func(a, b Person) bool {
    return a.Age > b.Age
}

sorting.MergeSortWithComparator(data, ageComparator)
// data is now sorted by age in descending order
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