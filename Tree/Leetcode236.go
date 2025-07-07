package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func Dfs(Node, p, q *TreeNode, arr []*TreeNode) []*TreeNode {

	if Node == nil {

		return arr
	}

	if Node == p {

		arr[0] = p
		
	}

	if Node == q {

		arr[1] = q
		
	}

	Dfs(Node.Left, p, q, arr)
	Dfs(Node.Right, p, q, arr)

	return arr

}

func LCAHelper(p, q, Node *TreeNode) *TreeNode {

	if Node == nil {

		return nil
	}

	if Node == p || Node == q {
		return Node
	}

	left := LCAHelper(p, q, Node.Left)
	right := LCAHelper(p, q, Node.Right)

	if left != nil && right != nil {
		return Node
	}

	if left != nil {
		return left
	}

	return right

}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if root == nil || p == nil || q == nil {

		return nil
	}
    
	arr := make([]*TreeNode, 2)
	arr = Dfs(root, p, q, arr)

	Result := LCAHelper(arr[0], arr[1], root)

	return Result

}