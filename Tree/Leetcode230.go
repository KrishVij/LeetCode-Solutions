package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func Bfs(node *TreeNode,arr []int) []int {

    if node == nil {

        return arr
    }

    queue := []*TreeNode{}
    queue = append(queue,node)

    for len(queue) > 0 {

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

    return arr
}

func kthSmallestHelper(arr []int,k int) int {

    slices.Sort(arr)

    for idx,ele := range arr {

        if idx == k - 1 {

            return ele
        }
    }

    return 0
}

func kthSmallest(root *TreeNode, k int) int {

    if root == nil {

        return 0
    }

    result := Bfs(root,[]int{})
    kthSmallestEle := kthSmallestHelper(result,k)

    return kthSmallestEle


}