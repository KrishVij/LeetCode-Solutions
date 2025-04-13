package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

type DoublyLinkedList struct {
	head *Node
}

func (list *DoublyLinkedList) Append(val int) {

	newNode := &Node{data: val}

	if list.head == nil {

		list.head = newNode
		return
	}

	current := list.head

	for current.next != nil {

		current = current.next
	}

	current.next = newNode
	newNode.prev = current

}

func (list *DoublyLinkedList) Delete(val int) {

	if list.head == nil {

		fmt.Println("Doesn't exist")
		return
	}

	current := list.head

	if current.data == val {

		list.head = current.next

		if list.head != nil {

			list.head.prev = nil
		}
	}

	for current != nil && current.data != val {

		current = current.next
	}

	if current == nil {

		fmt.Println("Value doesnt exist")
		return
	}

	if current.next != nil {

		current.next.prev = current.next
	} else {

		list.head = current.next
	}

	if current.prev != nil {

		current.prev.next = current.next
	}

}

func (list *DoublyLinkedList) deleteAllOccurrences(k int) *Node {

	if list.head == nil {

		return list.head
	}

	current := list.head
	nextNode := list.head.next
	prevNode := list.head.prev

	if list.head != nil && list.head.data == k {

		list.head = list.head.next

		if list.head != nil {

			list.head.prev = nil
		}

	}

	for current != nil {

		if current.data == k {

			if current.next != nil {

				current.next.prev = current.prev
			}
			if current.prev != nil {

				current.prev.next = current.next
			}
		}

		current = current.next
	}

	return list.head

}
