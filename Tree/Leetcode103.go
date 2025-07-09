package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {

	if root == nil {

		return [][]int{}
	}

	queue := []*TreeNode{root}
	zigZagTraversal := [][]int{}
	level := 1

	for len(queue) > 0 {

        size := len(queue)
        arr := []int{}

        for i := 0;i < size; i++ {

            Current := queue[0]
			queue = queue[1:]

            if level & 1 == 1 {

                arr = append(arr,Current.Val)
            }else {

                arr = append([]int{ Current.Val}, arr...)
            }

            if Current.Left != nil {

                queue = append(queue, Current.Left)
            }

            if Current.Right != nil {

                queue = append(queue, Current.Right)
            }
        }

        zigZagTraversal = append(zigZagTraversal, arr)
        level++
    }

    return zigZagTraversal

}