package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Solution struct {
	head *Node
}

func Constructor(head *Node) Solution {

	return Solution{head: head}
}

func (this *Solution) GetRandom() int {

}
