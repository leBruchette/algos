package sorting

import (
	"math/rand"
	golangSort "sort"
	"testing"
)

type testCase struct {
	name string
	n    int
}

var iteration = []testCase{
	{"100", 100},
	{"1000", 1000},
	{"10000", 10000},
	{"100000", 100000},
	//{"400000", 400000},
}

func BenchmarkGolangSort(b *testing.B) {
	runSortBenchmark(b, golangSort.Ints)
}

func BenchmarkInsertionSort(b *testing.B) {
	runSortBenchmark(b, InsertionSort)
}

func BenchmarkMergeSort(b *testing.B) {
	runSortBenchmark(b, MergeSort)
}

func BenchmarkQuickSort(b *testing.B) {
	runSortBenchmark(b, QuickSort)
}

func BenchmarkHeapSort(b *testing.B) {
	runSortBenchmark(b, HeapSort)
}

func runSortBenchmark(b *testing.B, sortFunc func([]int)) {
	for _, tc := range iteration {
		b.Run(tc.name, func(b *testing.B) {
			startBenchmarkForSortFunc(b, tc.n, sortFunc)
		})
	}
}

func startBenchmarkForSortFunc(b *testing.B, n int, sortFunc func([]int)) {
	arr := make([]int, n)
	for i := range arr {
		arr[i] = rand.Intn(n)
	}

	//start := time.Now()
	for i := 0; i < b.N; i++ {
		sortFunc(arr)
	}
	//fmt.Printf("\tDuration for %v items: %d\n", n, time.Since(start))

}
