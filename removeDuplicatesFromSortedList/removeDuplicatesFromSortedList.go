// Source : https://oj.leetcode.com/problems/remove-duplicates-from-sorted-list/
// Author : Hao Chen
// Date   : 2014-06-21

/********************************************************************************** 
* 
* Given a sorted linked list, delete all duplicates such that each element appear only once.
* 
* For example,
* Given 1->1->2, return 1->2.
* Given 1->1->2->3->3, return 1->2->3.
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
func deleteDuplicates(head *ListNode) *ListNode {
    for p := head; p != nil && p.Next != nil; {
        if p.Val == p.Next.Val {
            p.Next = p.Next.Next
            continue
        }
        p = p.Next
    }
    
    return head
}


