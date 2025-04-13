package main

import "fmt"

type Node_go struct {
	data int
	next *Node
}

type LinkedList_go struct {
	head *Node_go
}

func (list *LinkedList_go) Append(val int) *Node_go {

	newNode := &Node_go{data: val}

	if list.head == nil {

		list.head = newNode
		return list.head
	}

	current := list.head
	for current.next != nil {

		current = current.next
	}

	current.next = newNode
	newNode.next = nil

	return list.head
}

func (list *LinkedList) reverseBetween(head *ListNode, left int, right int) *ListNode {

	if head == nil || head.next == nil {

		return head
	}

	var prev *Node
	current := list.head
	for i := 0; i < left-1; i++ {

		prev = current.next
		current = current.next.next
	}

	for left <= right && current != nil {

		nextNode := current.next
		current.next = prev
		prev = current
		current = nextNode
		left += 2

	}
}
