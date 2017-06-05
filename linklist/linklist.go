package linklist

import (
	"fmt"
)

type LinkList struct {
	head *LinkNode
}

type LinkNode struct {
	Value interface{}
	Next  *LinkNode
}

func Insert(value interface{}, list *LinkList) {
	node := &LinkNode{value, nil}
	if list == nil {
		list = &LinkList{}
	}

	//insert first value as head
	if list.head == nil {
		list.head = node
		return
	}

	tail := list.head
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = node
}

func PrintList(list *LinkList) {
	if list == nil {
		fmt.Printf("The link list is not initialized!\n")
		return
	}

	if list.head == nil {
		fmt.Printf("The link list is empty!\n")
		return
	}

	node := list.head
	fmt.Printf("%v\n", node.Value)
	for node.Next != nil {
		node = node.Next
		fmt.Printf("%v\n", node.Value)
	}
}
