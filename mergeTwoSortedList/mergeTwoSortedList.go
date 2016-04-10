// Source : https://oj.leetcode.com/problems/merge-two-sorted-lists/
// Author : daoran
// Date   : 2016-04-10

/**********************************************************************************
*
* Merge two sorted linked lists and return it as a new list. The new list should be
* made by splicing together the nodes of the first two lists.
*
**********************************************************************************/

package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := new(ListNode)
	p := dummyHead

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}

	if l1 != nil {
		p.Next = l1
	}
	if l2 != nil {
		p.Next = l2
	}

	return dummyHead.Next
}

func main() {

}
