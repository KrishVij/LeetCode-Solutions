package tree


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func midOfASLL(Node *ListNode) (*ListNode,*ListNode) {

	if Node == nil {

		return nil,nil
	}

    var prev *ListNode

	slow := Node
	fast := Node

	for fast != nil && fast.Next != nil {

        prev = slow
        slow = slow.Next
		fast = fast.Next.Next
	}

	return prev,slow
}
func sortedListToBST(head *ListNode) *TreeNode {

	if head == nil {

		return nil
	}

	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}
    
    prev, mid := midOfASLL(head)
      
    if prev != nil {
        prev.Next = nil
    }

    root := &TreeNode{Val: mid.Val}
	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(mid.Next)

	return root

}