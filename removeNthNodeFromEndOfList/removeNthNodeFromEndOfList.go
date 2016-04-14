// Source : https://oj.leetcode.com/problems/remove-nth-node-from-end-of-list/
// Author : daoran
// Date   : 2016-04-14

/**********************************************************************************
*
* Given a linked list, remove the nth node from the end of list and return its head.
*
* For example,
*
*    Given linked list: 1->2->3->4->5, and n = 2.
*
*    After removing the second node from the end, the linked list becomes 1->2->3->5.
*
* Note:
* Given n will always be valid.
* Try to do this in one pass.
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

import "fmt"

// ListNode represents Linked List
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n <= 0 {
		return head
	}

	dummyHead := new(ListNode)
	dummyHead.Next = head
	head = dummyHead

	p1 := head
	p2 := head

	for i := 0; i < n; i++ {
		if p2 == nil {
			return nil
		}
		p2 = p2.Next
	}

	for p2.Next != nil {
		p2 = p2.Next
		p1 = p1.Next
	}

	p1.Next = p1.Next.Next

	return dummyHead.Next
}

func main() {
	list := &ListNode{Val: 1, Next: nil}
	fmt.Println(removeNthFromEnd(list, 1))
}
