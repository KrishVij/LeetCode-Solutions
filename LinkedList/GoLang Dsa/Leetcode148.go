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

func (list *LinkedList) MiddleOfList(head *Node) *Node {

	slow, fast := head, head

	var prev *Node
	for fast != nil && fast.next == nil {

		prev = slow
		slow = slow.next
		fast = fast.next.next
	}

	return prev
}

func (list *LinkedList) mergeLinkedList(l1, l2 *Node) *Node {

	dummyNode := &Node{data: -1}
	current := dummyNode

	for l1 != nil && l2 != nil {

		if l1.data < l2.data {

			current.next = l1
			current = current.next
		} else {

			current.next = l2
			current = current.next
		}

	}
	if l1 != nil {

		current.next = l1
	}
	if l2 != nil {

		current.next = nil
	}

	return dummyNode.next
}

func (list *LinkedList) mergeSortForLinkedList(head *Node) *Node {

	if head == nil || head.next == nil {

		return head
	}

	mid := list.MiddleOfList(head)
	rightHalf := mid.next
	mid.next = nil

	leftSorted := list.mergeSortForLinkedList(head)
	rightSorted := list.mergeSortForLinkedList(rightHalf)

	return list.mergeLinkedList(leftSorted, rightSorted)
}
