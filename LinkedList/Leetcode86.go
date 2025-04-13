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

	if list.head == nil || list.head.next == nil {

		return list.head
	}

	newNode := &Node{data: val}

	current := list.head
	for current.next != nil {

		current = current.next
	}

	current.next = newNode

	return list.head
}

func partition(head *Node, x int) *Node {

	if head == nil || head.next == nil {

		return head
	}

	smallDummy, largeDummy := &Node{}, &Node{}
	smallTemp, largeTemp := smallDummy, largeDummy

	current := head

	for current != nil {

		if current.data < x {

			smallTemp.next = current
			smallTemp = smallTemp.next
		} else {

			largeTemp.next = current
			largeTemp = largeTemp.next
		}

		current = current.next
	}

	largeTemp.next = nil
	smallTemp.next = largeDummy.next

	return smallDummy.next

}
