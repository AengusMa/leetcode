package leetcode024

// ListNode 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	var current *ListNode
	var last *ListNode
	current = head
	for current != nil && current.Next != nil {
		firstNode := current
		secondNode := current.Next
		firstNode.Next = secondNode.Next
		secondNode.Next = firstNode
		if current == head {
			head = secondNode
		}
		if last != nil {
			last.Next = secondNode
		}
		last = firstNode
		current = firstNode.Next
	}
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
