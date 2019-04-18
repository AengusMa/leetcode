package leetcode021

import "testing"

// Input: 1->2->4, 1->3->4
// Output: 1->1->2->3->4->4
func Test021(t *testing.T) {
	head1 := &ListNode{1, nil}
	addList(head1, 2)
	addList(head1, 4)
	head2 := &ListNode{1, nil}
	addList(head2, 3)
	addList(head2, 4)
	printList(mergeTwoLists1(head1, head2))

}
