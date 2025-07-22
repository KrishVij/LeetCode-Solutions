package main

func postorderTraversalHelper(Node *TreeNode,arr []int) []int {

    if Node == nil {

        return arr
    }

    arr = postorderTraversalHelper(Node.Left,arr)
    arr = postorderTraversalHelper(Node.Right,arr)
    arr = append(arr,Node.Val)

    return arr
}

func postorderTraversal(root *TreeNode) []int {
    
    if root == nil {

        return []int{}
    }

    return postorderTraversalHelper(root,[]int{})
}