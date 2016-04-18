// Source : https://oj.leetcode.com/problems/sort-list/
// Author : doaran
// Date   : 2016-04-18

/**********************************************************************************
*
* Sort a linked list in O(n log n) time using constant space complexity.
*
**********************************************************************************/

package main

// ListNode represents linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// Find the middle place
	p1 := head
	p2 := head.Next

	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}

	p2 = p1.Next
	p1.Next = nil

	return mergeTwoLists(sortList(head), sortList(p2))
}

func mergeTwoLists(head1 *ListNode, head2 *ListNode) *ListNode {
	p1 := head1
	p2 := head2

	dummyHead := new(ListNode)
	tail := dummyHead

	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			tail.Next = p1
			p1 = p1.Next
		} else {
			tail.Next = p2
			p2 = p2.Next
		}
		tail = tail.Next
	}

	if p1 != nil {
		tail.Next = p1
	}
	if p2 != nil {
		tail.Next = p2
	}

	return dummyHead.Next
}

func main() {

}
