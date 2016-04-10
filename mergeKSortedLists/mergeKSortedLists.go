// Source : https://oj.leetcode.com/problems/merge-k-sorted-lists/
// Author : daoran
// Date   : 2016-04-10

/********************************************************************************** 
* 
* Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.
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
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }
    
    end := len(lists) - 1
    for end > 0 {
        begin := 0
        for begin < end {
            lists[begin] = mergeTwoLists(lists[begin], lists[end])
            
            begin++
            end--
        }
    }
    
    return lists[0]
}

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