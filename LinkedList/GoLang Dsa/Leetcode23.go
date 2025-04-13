package main

import (
	"fmt"
	"sort"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) array2List(arr []int) *Node {

	length := len(arr)

	if length == 0 {

		return nil
	}

	list.head = &Node{data: arr[0]}
	current := list.head
	for _, val := range arr[1:] {

		current.next = &Node{data: val}
		current = current.next
	}

	return list.head
}

func (list *LinkedList) mergeKLists(arr [][]int) *Node {

	var result []int

	for _, subarr := range arr {

		result = append(result, subarr...)
	}

	sort.Ints(result)

	newHead := list.array2List(result)

	return newHead

}
