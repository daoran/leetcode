// Source : https://oj.leetcode.com/problems/add-two-numbers/
// Author : daoran
// Date   : 2016-04-10

/********************************************************************************** 
* 
* You are given two linked lists representing two non-negative numbers. 
* The digits are stored in reverse order and each of their nodes contain a single digit. 
* Add the two numbers and return it as a linked list.
* 
* Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
* Output: 7 -> 0 -> 8
*               
**********************************************************************************/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */ 
 
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
   dummyHead := new(ListNode) 
   p := l1
   q := l2 
   curr := dummyHead
   
   carry := 0
   for p != nil || q != nil {
       x := 0
       y := 0
       if p != nil {
           x = p.Val
       } 
       if q != nil {
           y = q.Val
       }
       
       digit := carry + x + y
       carry = digit / 10
       curr.Next = &ListNode{Val:digit%10}
       curr = curr.Next
       if p != nil {
           p = p.Next
       }
       if q != nil {
           q = q.Next
       }
   }
   
   if carry > 0 {
       curr.Next = &ListNode{Val:carry}
   }
   
   return dummyHead.Next
}