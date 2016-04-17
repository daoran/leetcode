// Source : https://oj.leetcode.com/problems/partition-list/
// Author : daoran
// Date   : 2016-04-17

/**********************************************************************************
*
* Given a linked list and a value x, partition it such that all nodes less than x come
* before nodes greater than or equal to x.
*
* You should preserve the original relative order of the nodes in each of the two partitions.
*
* For example,
* Given 1->4->3->2->5->2 and x = 3,
* return 1->2->2->4->3->5.
*
*
**********************************************************************************/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

// ListNode represents linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	dummyHead := new(ListNode)

	dummyHead.Next = head
	head = dummyHead

	pos := new(ListNode)
	pos = nil

	for p := head; p != nil && p.Next != nil; {
		if pos == nil && p.Next.Val >= x {
			pos = p
			p = p.Next
			continue
		}

		if pos != nil && p.Next.Val < x {
			pNext := p.Next
			p.Next = pNext.Next
			pNext.Next = pos.Next
			pos.Next = pNext
			pos = pNext
			continue
		}

		p = p.Next
	}

	return head.Next
}

func main() {

}
