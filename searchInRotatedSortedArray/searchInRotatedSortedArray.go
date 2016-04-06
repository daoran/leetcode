// Source : https://oj.leetcode.com/problems/search-in-rotated-sorted-array/
// Author : daoran
// Date   : 2016-04-06

/**********************************************************************************
*
* Suppose a sorted array is rotated at some pivot unknown to you beforehand.
*
* (i.e., 0 1 2 4 5 6 7 might become 4 5 6 7 0 1 2).
*
* You are given a target value to search. If found in the array return its index, otherwise return -1.
*
* You may assume no duplicate exists in the array.
*
**********************************************************************************/
// NOTE: f: first, l: last, m:mid
package main

import "fmt"

func search(nums []int, target int) int {
	f, l := 0, len(nums)-1
	for f <= l {
		m := (l-f)/2 + f
		if nums[m] == target {
			return m
		}
		if nums[f] <= nums[m] {
			if nums[f] <= target && target < nums[m] {
				l = m - 1
			} else {
				f = m + 1
			}
		} else {
			if nums[m] < target && target <= nums[l] {
				f = m + 1
			} else {
				l = m - 1
			}
		}
	}

	return -1
}

func main() {
	test := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(test, 0))
	test2 := []int{3, 1}
	fmt.Println(search(test2, 1))
}
