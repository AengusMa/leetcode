package leetcode019

// ListNode 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	current := head

	l := 0
	for {
		if current == nil {
			break
		}
		current = current.Next
		l++
	}
	if l == 1 {
		if n == 1 {
			head = nil
			return head
		} else {
			return head
		}
	}
	if n == l {
		head = head.Next
		return head
	}
	current = head
	for i := 0; i < l-n-1; i++ {
		current = current.Next
	}
	current.Next = current.Next.Next
	return head
}

func addList(head *ListNode, val int) {
	current := head
	for {
		if current.Next == nil {
			break
		}
		current = current.Next
	}
	node := &ListNode{val, nil}
	current.Next = node
}
func printList(head *ListNode) {
	current := head
	for {
		if current == nil {
			break
		}
		print(current.Val, "\t")
		current = current.Next
	}
	println()
}
