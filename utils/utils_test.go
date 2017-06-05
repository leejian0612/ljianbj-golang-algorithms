package utils

import (
	"fmt"
	"testing"
)

func TestRandomArray(t *testing.T) {
	a1 := RandomArray(10)
	a2 := RandomArray(10)
	fmt.Printf("TestRandomArray\na1: %v\na2: %v\n", a1, a2)
}

func TestUniqueArray(t *testing.T) {
	a1 := UniqueArray(10)
	a2 := UniqueArray(10)
	fmt.Printf("TestUniqueArray\na1: %v\na2: %v\n", a1, a2)
}
