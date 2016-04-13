// Source : https://oj.leetcode.com/problems/merge-sorted-array/
// Author : daoran
// Date   : 2016-04-13

/**********************************************************************************
*
* Given two sorted integer arrays A and B, merge B into A as one sorted array.
*
* Note:
*   You may assume that A has enough space (size that is greater or equal to m + n)
*   to hold additional elements from B. The number of elements initialized in A and B
*   are m and n respectively.
*
**********************************************************************************/

package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	ia := m - 1
	ib := n - 1
	for i := m + n - 1; i >= 0; i-- {
		if ia >= 0 && ib < 0 {
			break
		}

		if ia < 0 && ib >= 0 {
			nums1[i] = nums2[ib]
			ib--
		}

		if ia >= 0 && ib >= 0 {
			if nums1[ia] > nums2[ib] {
				nums1[i] = nums1[ia]
				ia--
			} else {
				nums1[i] = nums2[ib]
				ib--
			}
		}
	}
}

func main() {
	nums1 := []int{1, 3, 4, 5}
	nums2 := []int{2, 6, 7, 8}
	nums3 := make([]int, len(nums1)+len(nums2))
	copy(nums3, nums1)
	fmt.Println(nums3)
	merge(nums3, len(nums1), nums2, len(nums2))
	fmt.Println(nums3)
}
