package main

import (
	"container/list"
	"fmt"
)

func main() {
	// double linked list
	ll := list.New()
	ll.PushBack("2")
	ll.PushBack("3")
	ll.PushBack("4")
	ll.PushFront("1")

	// head -> tail
	for node := ll.Front(); node != nil; node = node.Next() {
		fmt.Println(node.Value)
	}

	// tail -> head
	for node := ll.Back(); node != nil; node = node.Prev() {
		fmt.Println(node.Value)
	}
}
