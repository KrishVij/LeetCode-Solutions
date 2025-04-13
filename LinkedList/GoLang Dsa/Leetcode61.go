package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) Append(val int) *Node {

	newNode := &Node{data: val}

	if list.head == nil {

		list.head = newNode
		return list.head
	}

	current := list.head
	for current.next != nil {

		current = current.next
	}

	current.next = newNode

	return list.head
}

func (list *LinkedList) rotateRight(k int) *Node {

	if list.head == nil || list.head.next == nil || k == 0 {

		return list.head
	}

	tail := list.head
	count := 1

	for tail.next == nil {

		count++
		tail = tail.next
	}

	k = k % count
	if k == 0 {

		return list.head
	}

	current := list.head
	var prev *Node

	for i := 0; i < count-k; i++ {

		prev = current
		current = current.next
	}

	prev.next = nil
	tail.next = list.head
	list.head = current

	return list.head

}
