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

		list.head == newNode
		return head
	}

	current := list.head

	for current.next != nil {

		current = current.next
	}

	current.next = newNode
	return list.head
}
func (list *LinkedList) getkthNode(current *Node, k int) *Node {

	k -= 1
	for current != nil && k > 0 {

		current = current.next
	}

	return current
}

func (list *LinkedList) revrseList(temp *Node) *Node {

	if temp == nil || temp.next == nil {

		return temp
	}

	current := temp
	var prevNode, nextNode *Node

	for current != nil {

		nextNode = current.next
		current.next = prevNode
		prevNode = current
		current = nextNode
	}

	return prevNode
}
func (list *LinkedList) reverseKgroup(head *Node, k int) *Node {

	if head == nil || head.next == nil {

		return head
	}

	current := head
	var prevNode *Node
	var newHead *Node

	for current != nil {

		kthNode := list.getkthNode(current, k)
		if kthNode == nil {

			if prevNode != nil {

				prevNode.next = current
			}

			break
		}

		nextNode := kthNode.next
		kthNode.next = nil
		reversedHead := list.reverselist(current)

		if newHead == nil {

			newHead = reversedHead
		}

		if prevNode != nil {

			prevNode.next = reversedHead
		}

		prevNode = current
		current = nextNode

	}

	return head

}
