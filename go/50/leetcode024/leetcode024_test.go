package leetcode024

import "testing"

// Given 1->2->3->4, you should return the list as 2->1->4->3.
func Test024(t *testing.T) {
	head1 := &ListNode{1, nil}
	addList(head1, 2)
	addList(head1, 3)
	addList(head1, 4)

	printList(swapPairs(head1))
}
