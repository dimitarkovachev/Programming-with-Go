package sorter

func SortSlice(slice []int) {
	for i := 0; i < len(slice); i++ {
		for j := i - 1; j >= 0; j-- {
			if slice[j] > slice[j+1] {
				temp := slice[j+1]
				slice[j+1] = slice[j]
				slice[j] = temp
			}
		}
	}
}

func SortSliceEventuallyBetter(slice []int) {
	for i := 0; i < len(slice); i++ {
		for j := i - 1; j >= 0 && slice[j] > slice[j+1]; j-- {
			temp := slice[j+1]
			slice[j+1] = slice[j]
			slice[j] = temp
		}
	}
}
