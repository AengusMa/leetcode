package leetcode019

import "testing"

// Given linked list: 1->2->3->4->5, and n = 2.
// After removing the second node from the end, the linked list becomes 1->2->3->5.
func Test019(t *testing.T) {
	head := &ListNode{1, nil}
	// printList(head)
	removeNthFromEnd(head, 1)
	printList(head)
}
