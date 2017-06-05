package linklist

import (
	"../utils"
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	array := utils.RandomArray(10)
	fmt.Printf("TestInsert\narray: %v\n", array)
	head := &LinkList{}
	fmt.Printf("linklist: ")
	PrintList(head)
	for _, i := range array {
		Insert(i, head)
	}
	fmt.Printf("linklist: ")
	PrintList(head)
}
