package sorter

import (
	"fmt"
	"reflect"
	"runtime"
	"sorter/resources"
	"testing"
)

var testSliceLenght int = 100000

type sortFunc func([]int)

func TestSortSlice(t *testing.T) {
	testSliceSorted(t, SortSlice)
}

func TestSortSliceEventuallyBetter(t *testing.T) {
	testSliceSorted(t, SortSliceEventuallyBetter)
}

func testSliceSorted(t *testing.T, sortingFunc sortFunc) {
	slice := resources.GenerateSlice(testSliceLenght)

	sortingFunc(slice)

	failed := false

	for i := range slice {
		if i+1 < len(slice) && slice[i] > slice[i+1] {
			t.Errorf("slice[i] = %v; slice[i+1] = %v\n", slice[i], slice[i+1])
			t.Fail()
			failed = true
			break
		}
	}
	if !failed {
		fmt.Printf("test %v passed\n", runtime.FuncForPC(reflect.ValueOf(sortingFunc).Pointer()).Name())
	}
}
