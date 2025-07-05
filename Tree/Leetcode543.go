/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func max(a,b int) int {

    if b > a {

        return b
    }

    return a
}

func diameterOfBinaryTree(root *TreeNode) int {
    
    if root == nil {

        return 0
    }

    maxDiameter := 0
   
    var Dfs func(Node *TreeNode) int
    Dfs = func(Node *TreeNode) int {

        if Node == nil {

            return 0
        }

        leftLength := Dfs(Node.Left)
        rightLength := Dfs(Node.Right)

        if leftLength + rightLength > maxDiameter {

            maxDiameter = leftLength + rightLength
        }

        return 1 + max(leftLength,rightLength)
    }

    Dfs(root)

    return maxDiameter
}