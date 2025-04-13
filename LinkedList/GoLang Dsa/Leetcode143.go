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

func (list *LinkedList) middleNode(head *Node) *Node {

	fast, slow := list.head, list.head

	for fast != nil && fast.next != nil {

		slow = slow.next
		fast = fast.next.next
	}

	return slow
}

func (list *LinkedList) reverseList(head *Node) *Node {

	prevNode := &Node{}
	current := list.head

	for current != nil {

		front := current.next
		current.next = prevNode
		prevNode = current
		current = front
	}

	return prevNode

}

func (list *LinkedList) reorderList(head *Node) *Node {

	if head == nil && head.next == nil {

		return nil
	}

	mid := list.middleNode(list.head)
	second := list.reverseList(mid.next)
	mid.next = nil

	first := head
	for second != nil {

		tmp1, tmp2 := first.next, second.next
		first.next = second
		second.next = tmp1
		first, second = tmp1, tmp2
	}

	return list.head

}
