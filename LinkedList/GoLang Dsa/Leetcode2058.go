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

func findMinMaxDistance(countStore []int) []int {

	if len(countStore) == 2 {

		return []int{-1, -1}
	}

	minDis := countStore[1] - countStore[0]

	for i := 1; i < len(countStore); i++ {

		diff := countStore[i] - countStore[i-1]
		if diff < minDis {

			minDis = diff
		}
	}

	maxDis := countStore[len(countStore)-1] - countStore[0]

	return []int{minDis, maxDis}
}

func (list *LinkedList) nodesBetweenCriticalPoints(head *Node) []int {

	if head == nil || head.next == nil || head.next.next == nil {

		return []int{-1, -1}
	}

	current := list.head
	count := 0
	var countStore []int

	var prevNode *Node

	for current.next != nil {

		nextNode := current.next
		if prevNode != nil && nextNode != nil {

			if current.data > prevNode.data && current.data > nextNode.data {

				countStore = append(countStore, count)

			} else if current.data < prevNode.data && current.data < nextNode.data {

				countStore = append(countStore, count)

			}
		}

		count++
		prevNode = current
		current = current.next
	}

	return findMinMaxDistance(countStore)
}
