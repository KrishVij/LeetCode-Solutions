/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {

	fast, slow := head, head

	if head == nil || head.Next == nil {

		return nil
	}

	for fast != nil && fast.Next != nil {

		slow = slow.Next
		fast = fast.Next.Next

		if fast == slow {

			break
		}

	}
	if fast == nil || fast.Next == nil {

		return nil
	}
	slow = head
	for slow != fast {

		slow = slow.Next
		fast = fast.Next

	}
	return slow

}
