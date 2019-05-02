package leetcode083

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	curr := head
	res := head
	for head != nil {
		if head.Val != curr.Val {
			curr.Next = head
			curr = head
		} else {
			head, head.Next = head.Next, nil
		}
	}
	return res
}
