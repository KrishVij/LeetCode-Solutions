package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepthHelper(node *TreeNode,arr []int,Count int) []int{

    if node == nil {

        return arr
    }

    arr = append(arr,Count)
    arr = maxDepthHelper(node.Left,arr,Count + 1)
    arr = maxDepthHelper(node.Right,arr,Count + 1)

    return arr
    
}

func CalculateMaxDepth(arr []int) int {

    maxDepth := slices.Max(arr)

    return maxDepth
}

func maxDepth(root *TreeNode) int {

    if root == nil {

        return 0
    }

    result := maxDepthHelper(root,[]int{},1)
    maxDepth := CalculateMaxDepth(result)

    return maxDepth
}