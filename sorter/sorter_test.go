package sorter

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"

	"github.com/dimitarkovachev/Programming-with-Go/sorter/resources"
)

var testSliceLenght int = 50

type sortFunctionInPlace func([]int)
type sortFunction func([]int) []int

func TestInsertionSort(t *testing.T) {
	testSliceSorted(t, InsertionSort)
}

func TestMergeSort(t *testing.T) {
	testSliceSorted(t, MergeSort)
}

func BenchmarkInsertionSort(b *testing.B) {
	benchmarkSliceSorted(b, InsertionSort)
}

func BenchmarkMergeSort(b *testing.B) {
	benchmarkSliceSorted(b, MergeSort)
}

func testSliceSorted(t *testing.T, sortingFunc interface{}) {
	slice := resources.GenerateSlice(testSliceLenght)

	switch sort := sortingFunc.(type) {
	case func([]int):
		sort(slice)
	case func([]int) []int:
		slice = sort(slice)
	case sortFunctionInPlace:
		// Can't match function type
	}

	for i := range slice {
		if i+1 < len(slice) && slice[i] > slice[i+1] {
			t.Errorf("slice[%v] = %v; slice[%v] = %v\n", i, slice[i], i+1, slice[i+1])
			t.FailNow()
		}
	}

	fmt.Printf("test %v passed\n", runtime.FuncForPC(reflect.ValueOf(sortingFunc).Pointer()).Name())
}

func benchmarkSliceSorted(b *testing.B, sortingFunc interface{}) {
	slice := resources.GenerateSlice(testSliceLenght)

	switch sort := sortingFunc.(type) {
	case func([]int):
		for i := 0; i < b.N; i++ {
			sort(slice)
		}
	case func([]int) []int:
		for i := 0; i < b.N; i++ {
			sort(slice)
		}
	}
}
