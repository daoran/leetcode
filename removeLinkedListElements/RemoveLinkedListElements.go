// Source : https://leetcode.com/problems/remove-linked-list-elements/
// Author : daoran
// Date   : 2016-04-14

/********************************************************************************** 
 * 
 * Remove all elements from a linked list of integers that have value val.
 * 
 * Example
 * Given: 1 --> 2 --> 6 --> 3 --> 4 --> 5 --> 6,  val = 6
 * Return: 1 --> 2 --> 3 --> 4 --> 5
 * 
 * Credits:Special thanks to @mithmatt for adding this problem and creating all test cases.
 *               
 **********************************************************************************/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    dummyHead := new(ListNode)
    
    dummyHead.Next = head
    p := dummyHead 
    
    for p.Next != nil {
        if p.Next.Val == val {
            p.Next = p.Next.Next
        } else {
            p = p.Next
        }
    }
    
    return dummyHead.Next
}