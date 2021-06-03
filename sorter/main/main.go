package main

import (
	"fmt"
	"sort"

	"github.com/dimitarkovachev/Programming-with-Go/sorter/resources"
)

type Square struct {
	Length int
}

type ByLength []Square

func (a ByLength) Len() int           { return len(a) }
func (a ByLength) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLength) Less(i, j int) bool { return a[i].Length < a[j].Length }

func main() {
	squares := []Square{}

	ints := resources.GenerateSlice(10)

	for _, v := range ints {
		squares = append(squares, Square{Length: v})
	}

	fmt.Println(squares)

	sort.Sort(ByLength(squares))

	sort.Slice(squares, func(i, j int) bool {
		return squares[i].Length < squares[j].Length
	})

	fmt.Println(squares)
}
