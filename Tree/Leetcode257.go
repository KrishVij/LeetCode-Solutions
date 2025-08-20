package main

func Dfs(Node *TreeNode, path_From_Root_To_Leaf_Node []string) []atring{

	if Node.Left == nil && Node.Right == nil {

		return
	}

	path_From_Root_To_Leaf_Node = append(path_From_Root_To_Leaf_Node, Node.Value)
	Dfs(Node.Left)
	Dfs(Node.Right)

	return path_From_Root_To_Leaf_Node

}

func binaryTreePaths(root *TreeNode) []string {
    
	if root == nil {

		return []string{""}
	}

	path_Frome_root_To_Leaf_Node := []string{}
	path_Frome_root_To_Leaf_Node = Dfs(root, path_Frome_root_To_Leaf_Node)

	return path_Frome_root_To_Leaf_Node
}