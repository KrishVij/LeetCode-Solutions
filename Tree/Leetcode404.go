package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func Dfs(Node *TreeNode, Sum int, isLeft bool) int {

	if Node == nil {

		return Sum
	}

	if Node.Left == nil && Node.Right == nil && isLeft {

		Sum += Node.Val
	}

	Sum = Dfs(Node.Left, Sum, true)
	Sum = Dfs(Node.Right, Sum, false)

	return Sum
}
func sumOfLeftLeaves(root *TreeNode) int {

	if root == nil {

		return 0
	}

	if root.Left == nil && root.Right == nil {

		return 0
	}

	result := Dfs(root, 0, true)

	return result
}