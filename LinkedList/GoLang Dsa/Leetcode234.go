package main

import (
	"fmt"
	"reflect"
)

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

func (list *LinkedList) reverseList(head *Node) *Node {

	if head == nil || head.next == nil {

		return head
	}

	current := head
	var prev *Node

	for current != nil {

		nextNode := current.next
		current.next = prev
		prev = current
		current = nextNode
	}

	return prev

}

func (list *LinkedList) isPallindrome(head *Node) bool {

	if head == nil || head.next == nil {

		return false
	}

	current := head
	var originaListArray []int

	for current != nil {

		originaListArray = append(originaListArray, current.data)
		current = current.next
	}

	newHead := list.reverseList(head)
	var reverseListArray []int

	temp := newHead

	for temp != nil {

		reverseListArray = append(reverseListArray, temp.data)
		temp = temp.next
	}

	if reflect.DeepEqual(originaListArray, reverseListArray) {

		return true
	}

	return false

}
