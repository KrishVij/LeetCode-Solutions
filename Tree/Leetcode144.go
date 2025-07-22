package main

func preorderTraversalHelper(Node *TreeNode,arr []int) []int {

    if Node == nil {

        return arr
    }

    arr = append(arr,Node.Val)
   arr =  preorderTraversalHelper(Node.Left,arr)
   arr =  preorderTraversalHelper(Node.Right,arr)

   return arr
}

func preorderTraversal(root *TreeNode) []int {
    
    if root == nil {

        return []int{}
    }

    arr := []int{}
    return preorderTraversalHelper(root,arr)
}