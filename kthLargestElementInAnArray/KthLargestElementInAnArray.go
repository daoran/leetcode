// Source : https://leetcode.com/problems/kth-largest-element-in-an-array/
// Author : daoran
// Date   : 2016-04-16

/**********************************************************************************
 *
 * Find the kth largest element in an unsorted array.
 * Note that it is the kth largest element in the sorted order, not the kth distinct element.
 *
 * For example,
 * Given [3,2,1,5,6,4] and k = 2, return 5.
 *
 * Note:
 * You may assume k is always valid, 1 ≤ k ≤ array's length.
 *
 * Credits:Special thanks to @mithmatt for adding this problem and creating all test cases.
 *
 **********************************************************************************/

package main

import "fmt"

func partition(nums []int, left int, right int) int {
	pivot := nums[left]

	l := left + 1
	r := right

	for l <= r {
		if nums[l] < pivot && nums[r] > pivot {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}

		if nums[l] >= pivot {
			l++
		}

		if nums[r] <= pivot {
			r--
		}
	}

	nums[left], nums[r] = nums[r], nums[left]
	return r
}

func findKthLargest(nums []int, k int) int {
	left := 0
	right := len(nums) - 1

	pos := 0
	for true {
		pos = partition(nums, left, right)
		if pos == k-1 {
			break
		}

		if pos > k-1 {
			right = pos - 1
		} else {
			left = pos + 1
		}
	}

	return nums[pos]
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}
