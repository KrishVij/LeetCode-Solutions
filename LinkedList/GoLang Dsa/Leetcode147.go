package main

type Node struct {

	Value int
	Next *Node
}

type LinkedList struct {

	Head *Node
}

func (ll *LinkedList) Append(value int ) {

	newNode := &Node{Value : value} 
	if ll.Head == nil {

		ll.Head = newNode
		return 
	}

	curr := ll.head
	for curr.Next != nil {

		curr = curr.Next
	}

	curr.Next = newNode

}

func (ll *LinkedList) insertionSortList() *Node {

	if ll.Head == nil || ll.Head.Next == nil{

		return ll.Head
	}

	dummyNode := &Node{}
	dummyNode.Next = ll.Head
	prev := nil
	curr := ll.head
	next := curr.next

	for next != nil {

		if next.Value < curr.Value {

			curr.Next = next.next
			next.Next = curr
		}

		if next.Value < curr.Value && next.Value < prev.Value {

			curr.Next = next.next
			next.Next = prev
		}

		if next.Value < curr.Value && next.Value > prev.Value {

			prev.Next = next
			curr.Next = next.Next 
			next.Next = curr
		}

		if next.Value < dummyNode.Next.Value {

			current.Next = next.next
			next.Next = dummyNode.next
		}

		prev = current
		current = current.Next
		next = current.Next
	}

	return dummy.Next

}