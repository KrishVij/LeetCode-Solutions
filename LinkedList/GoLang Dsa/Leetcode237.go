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

func (list *LinkedList) giveAccesstoNode(val int) *Node {

	current := list.head

	for current != nil && current.data != val {

		current = current.next
	}

	return current
}

func (list *LinkedList) deleteNode(node *Node) *Node {

	if node == nil || node == nil {

		return nil
	}

	node.data = node.next.data
	node.next = node.next.next

	return node

}
