// Source : https://leetcode.com/problems/palindrome-linked-list/
// Author : daoran
// Date   : 2016-04-13

/**********************************************************************************
 *
 * Given a singly linked list, determine if it is a palindrome.
 *
 * Follow up:
 * Could you do it in O(n) time and O(1) space?
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

func findMiddle(head *ListNode) *ListNode {
	p1 := head
	p2 := head

	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}

	return p1
}

func reverseList(head *ListNode) *ListNode {
	current := new(ListNode)
	prev := new(ListNode)
	next := new(ListNode)

	current = head
	prev = nil

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

func isPalindrome(head *ListNode) bool {
	pMid := findMiddle(head)
	pRev := reverseList(pMid)

	for ; head != pMid; head, pRev = head.Next, pRev.Next {
		if head.Val != pRev.Val {
			return false
		}
	}

	return true
}

func main() {

}
