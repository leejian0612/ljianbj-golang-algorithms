package sorting

import (
	"math/rand"
)

func QuickSort(array []int) {
	quickSort(array, 0, len(array)-1)
}

func quickSort(array []int, start int, end int) {
	if start < end {
		index := start + rand.Intn(end-start)
		array[index], array[end] = array[end], array[index]
		sorted := start
		for i := start; i < end; i++ {
			if array[i] <= array[end] {
				array[i], array[sorted] = array[sorted], array[i]
				sorted++
			}
		}
		array[end], array[sorted] = array[sorted], array[end]
		quickSort(array, start, sorted-1)
		quickSort(array, sorted+1, end)
	}
}
