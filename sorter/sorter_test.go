package sorter

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"

	"github.com/dimitarkovachev/Programming-with-Go/sorter/resources"
)

var testSliceLenght int = 10000

type sortFunc func([]int)

func TestSortSlice(t *testing.T) {
	testSliceSorted(t, SortSlice)
}

func TestSortSliceEventuallyBetter(t *testing.T) {
	testSliceSorted(t, SortSliceEventuallyBetter)
}

func BenchmarkSortSlice(b *testing.B) {
	benchmarkSliceSorted(b, SortSlice)
}

func BenchmarkSortSliceEventuallyBetter(b *testing.B) {
	benchmarkSliceSorted(b, SortSliceEventuallyBetter)
}

func testSliceSorted(t *testing.T, sortingFunc sortFunc) {
	slice := resources.GenerateSlice(testSliceLenght)

	sortingFunc(slice)

	for i := range slice {
		if i+1 < len(slice) && slice[i] > slice[i+1] {
			t.Errorf("slice[i] = %v; slice[i+1] = %v\n", slice[i], slice[i+1])
			t.FailNow()
		}
	}

	fmt.Printf("test %v passed\n", runtime.FuncForPC(reflect.ValueOf(sortingFunc).Pointer()).Name())
}

func benchmarkSliceSorted(b *testing.B, sortingFunc sortFunc) {
	slice := resources.GenerateSlice(testSliceLenght)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sortingFunc(slice)
	}
}
