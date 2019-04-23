package leetcode023

// ListNode 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var head *ListNode
	var current *ListNode
	hasEle := true
	for hasEle {
		min := int(^uint(0) >> 1)
		minIndex := -1
		hasEle = false
		for i := 0; i < len(lists); i++ {
			if lists[i] != nil && lists[i].Val < min {
				min = lists[i].Val
				minIndex = i
				hasEle = true
			}
		}
		if minIndex != -1 {
			if head == nil {
				head = lists[minIndex]
				current = head
			} else {
				current.Next = lists[minIndex]
				current = current.Next
			}
			lists[minIndex] = lists[minIndex].Next
		}
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
