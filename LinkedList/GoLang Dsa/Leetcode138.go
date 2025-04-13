package main

import "fmt"

type NodeQ struct {
	data   int
	next   *NodeQ
	random *NodeQ
}

type LinkedList struct {
	head *NodeQ
}

func (list *LinkedList) copyRandomList(head *NodeQ) *NodeQ {

	if head == nil || head.next == nil {

		return head
	}

	dummyNode := &NodeQ{data: -1}
	temp1 := head
	temp2 := dummyNode
	nodeMap := make(map[*NodeQ]*NodeQ)

	for temp1 != nil {

		nodeMap[temp1] = &NodeQ{data: temp1.data}
		temp1 = temp1.next
	}

	temp1 = head
	for temp1 != nil {

		temp2.next = nodeMap[temp1]
		temp2 = temp2.next
		if temp1.random != nil {

			temp2.random = nodeMap[temp1.random]

		}

		temp1 = temp1.next
	}

	return dummyNode.next

}
