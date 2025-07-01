package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func CalculateMinDepth(arr []int) int {

    minDepth := slices.Min(arr)

    return minDepth
}

func minDepthHelper(node *TreeNode,arr []int,Count int) []int {

    if node == nil {

        return arr
    }

    if node.Left == nil && node.Right == nil {
        arr = append(arr, Count)
        return arr
    }

    arr = minDepthHelper(node.Left,arr,Count + 1)
    arr = minDepthHelper(node.Right,arr,Count + 1)

    return arr

}

func minDepth(root *TreeNode) int {
    
    if root == nil {

        return 0
    }

    result := minDepthHelper(root,[]int{},1)

    minDepth := CalculateMinDepth(result)

    return minDepth
}