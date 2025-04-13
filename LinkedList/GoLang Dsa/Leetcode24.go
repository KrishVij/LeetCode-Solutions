package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) Apend(val int) *Node {

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

func (list *LinkedList) swapPairs(head *Node) *Node {

	if head == nil || head.next == nil {

		return head
	}

	dummyNode := &Node{next: head}
	prev := dummyNode
	slow := list.head
	fast := list.head.next

	for slow != nil && fast != nil {

		slow.next = fast.next
		fast.next = slow
		prev.next = fast

		prev = slow
		slow = slow.next
		if slow != nil {

			fast = slow.next
		}

	}

	return dummyNode.next
}
