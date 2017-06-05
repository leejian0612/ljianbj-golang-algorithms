package linklist

import (
	"../utils"
	"testing"
)

func TestInsert(t *testing.T) {
	array := utils.RandomArray(10)
	head := &LinkList{}
	PrintList(head)
	for _, i := range array {
		Insert(i, head)
	}
	PrintList(head)
}
