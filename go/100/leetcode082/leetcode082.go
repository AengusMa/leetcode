package leetcode082

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	pre := dummy
	cur := head
	for cur != nil {
		if cur.Next == nil || cur.Val != cur.Next.Val {
			if pre.Next == cur {
				pre = cur
			} else {
				pre.Next = cur.Next
			}
		}
		cur = cur.Next
	}
	return dummy.Next
}
