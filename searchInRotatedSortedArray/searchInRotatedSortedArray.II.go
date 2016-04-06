// Source : https://oj.leetcode.com/problems/search-in-rotated-sorted-array-ii/
// Author : Hao Chen
// Date   : 2014-06-29

/**********************************************************************************
*
* Follow up for "Search in Rotated Sorted Array":
* What if duplicates are allowed?
*
* Would this affect the run-time complexity? How and why?
*
* Write a function to determine if a given target is in the array.
*
**********************************************************************************/

// NOTE: f: first, l: last, m:mid
package main

import "fmt"

func search(nums []int, target int) bool {
	f, l := 0, len(nums)-1
	for f <= l {
		m := (l-f)/2 + f
		if nums[m] == target {
			return true
		}
		if nums[f] < nums[m] {
			if nums[f] <= target && target < nums[m] {
				l = m - 1
			} else {
				f = m + 1
			}
		} else if nums[f] > nums[m] {
			if nums[m] < target && target <= nums[l] {
				f = m + 1
			} else {
				l = m - 1
			}
		} else {
			// skip duplicate one
			f++
		}
	}

	return false
}

func main() {
	test := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(test, 0))
	test2 := []int{3, 1}
	fmt.Println(search(test2, 1))
	test3 := []int{1, 3, 1, 1, 1}
	fmt.Println(search(test3, 1))
}
