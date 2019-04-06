package leetcode023

import "testing"

// Input:
// [
//   1->4->5,
//   1->3->4,
//   2->6
// ]
// Output: 1->1->2->3->4->4->5->6
func Test023(t *testing.T) {
	head1 := &ListNode{1, nil}
	addList(head1, 4)
	addList(head1, 5)
	head2 := &ListNode{1, nil}
	addList(head2, 3)
	addList(head2, 4)
	head3 := &ListNode{2, nil}
	addList(head3, 6)
	var lists []*ListNode
	lists = append(lists, head1, head2, head3)
	printList(mergeKLists(lists))
}
