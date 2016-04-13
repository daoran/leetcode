// Source : https://leetcode.com/problems/contains-duplicate-ii/
// Author : daoran
// Date   : 2016-04-13

/**********************************************************************************
 *
 * Given an array of integers and an integer k, find out whether there there are
 * two distinct indices i and j in the array such that nums[i] = nums[j] and
 * the difference between i and j is at most k.
 *
 *
 **********************************************************************************/

package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		n := nums[i]
		if _, ok := m[n]; ok && i-m[n] <= k {
			return true
		}
		m[n] = i
	}

	return false
}

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 4, 5, 2}, 5))
}
