// Source : https://oj.leetcode.com/problems/swap-nodes-in-pairs/
// Author : daoran
// Date   : 2016-04-10

/********************************************************************************** 
* 
* Given a linked list, swap every two adjacent nodes and return its head.
* 
* For example,
* Given 1->2->3->4, you should return the list as 2->1->4->3.
* 
* Your algorithm should use only constant space. You may not modify the values in the list, 
* only nodes itself can be changed.
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
func swapPairs(head *ListNode) *ListNode {
    dummy := new(ListNode)
    dummy.Next = head
    p := head
    prev := dummy
    
    for p != nil && p.Next != nil {
        q := p.Next
        r := p.Next.Next
        
        prev.Next = q
        q.Next = p
        p.Next = r
        
        
        prev = p
        p = r 
    }
    
    return dummy.Next
}


