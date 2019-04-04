package leetcode021

// ListNode 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

// Input: 1->2->4, 1->3->4
// Output: 1->1->2->3->4->4
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode
	if l1 == nil && l2 == nil {
		return result
	}
	var current *ListNode
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			node := &ListNode{l1.Val, nil}
			if result == nil {
				result = node
				current = node
			} else {
				current.Next = node
				current = current.Next
			}
			l1 = l1.Next
		} else {
			node := &ListNode{l2.Val, nil}
			if result == nil {
				result = node
				current = node
			} else {
				current.Next = node
				current = current.Next
			}
			l2 = l2.Next
		}
	}
	if l1 != nil {
		if result == nil {
			result = l1
		} else {
			current.Next = l1
		}
	}
	if l2 != nil {
		if result == nil {
			result = l2
		} else {
			current.Next = l2
		}
	}
	return result
}

func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	minEle := func() *ListNode {
		if l1 == nil && l2 == nil {
			return nil
		}
		var node *ListNode
		if l1 == nil {
			node = l2
			l2 = l2.Next
		} else if l2 == nil {
			node = l1
			l1 = l1.Next
		} else if l1.Val < l2.Val {
			node = l1
			l1 = l1.Next
		} else {
			node = l2
			l2 = l2.Next
		}
		return node
	}
	head := minEle()
	current := head
	for current != nil {
		current.Next = minEle()
		current = current.Next
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
