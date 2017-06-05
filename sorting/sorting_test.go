package sorting

import (
	"../utils"
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	array := utils.UniqueArray(10)
	fmt.Printf("TestQuickSort\norgin: %v\n", array)
	QuickSort(array)
	fmt.Printf("sorted: %v\n", array)
}
