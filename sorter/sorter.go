package sorter

func InsertionSort(slice []int) {
	for i := 0; i < len(slice); i++ {
		for j := i - 1; j >= 0 && slice[j] > slice[j+1]; j-- {
			temp := slice[j+1]
			slice[j+1] = slice[j]
			slice[j] = temp
		}
	}
}

func MergeSort(slice []int) []int {
	sliceLength := len(slice)

	if sliceLength > 1 {
		midPoint := sliceLength / 2

		lSlice := []int{}
		rSlice := []int{}

		for i, v := range slice {
			if i < midPoint {
				lSlice = append(lSlice, v)
			} else {
				rSlice = append(rSlice, v)
			}
		}

		lSlice = MergeSort(lSlice)
		rSlice = MergeSort(rSlice)
		return merge(lSlice, rSlice)
	}

	return slice
}

func merge(lSlice, rSlice []int) []int {
	mergedSlice := []int{}

	i, j := 0, 0

	for i < len(lSlice) && j < len(rSlice) {
		if lSlice[i] < rSlice[j] {
			mergedSlice = append(mergedSlice, lSlice[i])
			i++
		} else {
			mergedSlice = append(mergedSlice, rSlice[j])
			j++
		}
	}

	for i < len(lSlice) {
		mergedSlice = append(mergedSlice, lSlice[i])
		i++
	}

	for j < len(rSlice) {
		mergedSlice = append(mergedSlice, rSlice[j])
		j++
	}

	return mergedSlice
}
