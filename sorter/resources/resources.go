package resources

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateSlice(length int) []int {
	if length <= 0 {
		fmt.Println("please pass positive length")

		return nil
	}

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	slice := []int{}

	for ; length != 0; length-- {
		v := rnd.Intn(5000)
		slice = append(slice, v)
	}

	return slice
}
