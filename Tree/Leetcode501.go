/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type map_count_to_value struct {
	
    count     int
	value     int
    max_count int

}

func Dfs_501(Node *TreeNode, tree_node *map_count_to_value, modes []int) []int {

	if Node == nil {

		return modes
	}

    modes = Dfs_501(Node.Left, tree_node, modes)

    if Node.Val == tree_node.value {

		tree_node.count = tree_node.count + 1
	}

	if Node.Val != tree_node.value {

		tree_node.value = Node.Val
		tree_node.count = 1
	}

    if tree_node.count > tree_node.max_count {

        tree_node.max_count = tree_node.count
        modes = []int{Node.Val}

    }else if tree_node.count == tree_node.max_count {

        modes = append(modes, Node.Val)
    }

	modes = Dfs_501(Node.Right, tree_node, modes)

	return modes

}

func findMode(root *TreeNode) []int {

	if root == nil {

		return []int{}
	}

    if root.Left == nil && root.Right == nil {

        return []int{root.Val}
    }

	tree_node := &map_count_to_value{

		count: 0,
		value: root.Val,
	}

	return Dfs_501(root, tree_node, []int{})
}