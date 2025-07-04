package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sortedArrayToBST(nums []int) *TreeNode {

	if len(nums) == 0 {

		return nil
	}

	mid := len(nums) / 2

	rootNode := &TreeNode{Val: nums[mid]}
    rootNode.Left = sortedArrayToBST(nums[:mid])
    rootNode.Right = sortedArrayToBST(nums[mid + 1:])

    return rootNode
}