package main

import "fmt"

type Node struct {

	Value int
	Children []*Node
}

type Tree struct {

	Root *Node
	Count int
}

func (t *Tree) AddChild(Data int,NumOfNodes int) {

	if t.Count >= NumOfNodes {

		return
	}
 
	if t.Root == nil {

		t.Root = &Node{Value: Data}
		t.Count = 1
		return
	}
	
	queue := []*Node{}
	queue = append(queue, t.Root)
	
	for len(queue) > 0 {

		Current := queue[0]
		queue = queue[1:]
		

		if len(Current.Children) < 2  {

			Current.Children = append(Current.Children, &Node{Value: Data})
			t.Count = t.Count + 1

			return
		}

		queue = append(queue, Current.Children...)
	}
}

func (t *Tree) Traversal() {

	if t.Root == nil {

		return 
	}

	queue := []*Node{}
	queue = append(queue, t.Root)

	for len(queue) > 0  {

		Current := queue[0]
		queue = queue[1:]

		fmt.Println(Current.Value)

		for _, child := range Current.Children {
			queue = append(queue, child)
		}
	}
}


func (t *Tree) Dfs(node *Node) {

	if node == nil {

		return
	}

	fmt.Print(node.Value)

	for _,child := range node.Children {

		t.Dfs(child)
	}

}

func (t *Tree) PostOrder(node *Node) {
	
	if node == nil {
		
		return
	}
	
	for _,child := range node.Children {

		t.PostOrder(child)

	}

	fmt.Print(node.Value)

	

}

func (t *Tree) PreOrder(node *Node) {
	
	if node == nil {
		
		return
	}

	fmt.Print(node.Value)
	for _,child := range node.Children {

		t.PreOrder(child)

	}

}

func (t *Tree) InOrder(node *Node) {


	if node == nil {

		return
	}

	mid := len(node.Children)/2

	for i := 0;i < mid;i++ {

		t.InOrder(node.Children[i])
	}

	fmt.Print(node.Value)

	for i := mid; i < len(node.Children);i++ {

		t.InOrder(node.Children[i])
	}
}



func main() {

	tree := &Tree{}
	
	numOfNodes := 7
	
	for i := 1; i <= numOfNodes; i++ {
		tree.AddChild(i, numOfNodes)
	}

	fmt.Println("BFS Traversal:")
	tree.Traversal()

	fmt.Println("DFS Traversal:")
	tree.Dfs(tree.Root)
	fmt.Println()

	fmt.Println("PostOrder Traversal: ")
	tree.PostOrder(tree.Root)

	fmt.Println("\nPreOrder Traversal: ")
	tree.PreOrder(tree.Root)

	fmt.Println("\nInOrder Traversal: ")
	tree.InOrder(tree.Root)


}