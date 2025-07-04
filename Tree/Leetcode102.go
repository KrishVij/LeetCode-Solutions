package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {

	if root == nil {
        
        return [][]int{}
	}

	levelMatrix := [][]int{}

	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) > 0 {

        levelSize := len(queue) 
        arr := []int{}

        for i := 0;i < levelSize;i++ {

            Current := queue[0]
            queue = queue[1:]

            arr = append(arr,Current.Val)

            if Current.Left != nil {

                queue = append(queue,Current.Left)
            }

            if Current.Right != nil {

                queue = append(queue,Current.Right)
            }

            
        }

        levelMatrix = append(levelMatrix,arr)
    }

    return levelMatrix
}