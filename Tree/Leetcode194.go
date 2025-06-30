package main

func inorderHelper(node *TreeNode,arr *[]int) {

    if node == nil {

        return
    }

    inorderHelper(node.Left,arr)
    *arr = append(*arr,node.Val)
    inorderHelper(node.Right,arr)

    
}

func inorderTraversal(root *TreeNode) []int {

    if root == nil {

        return []int{}
    }
    
    inorderArray := []int{}
    inorderHelper(root,&inorderArray)

    return inorderArray

}