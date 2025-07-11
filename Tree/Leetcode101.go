package tree


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSymmetricHelper(Node1, Node2 *TreeNode) bool {

	if Node1 == nil && Node2 == nil {

		return true
	}

	if Node1 == nil || Node2 == nil {

		return false
	}

	if Node1.Val != Node2.Val {
		return false
	}

	return isSymmetricHelper(Node1.Left, Node2.Right) && isSymmetricHelper(Node1.Right, Node2.Left)

	
}

func isSymmetric(root *TreeNode) bool {

	if root == nil {

		return false
	}

	return isSymmetricHelper(root.Left, root.Right)
}