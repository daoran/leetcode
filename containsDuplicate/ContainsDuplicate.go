// Source : https://leetcode.com/problems/contains-duplicate/
// Author : daoran
// Date   : 2016-04-08

/**********************************************************************************
 *
 * Given an array of integers, find if the array contains any duplicates.
 * Your function should return true if any value appears at least twice in the array,
 * and it should return false if every element is distinct.
 *
 **********************************************************************************/

package main

import "fmt"

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)

	for _, n := range nums {
		if _, ok := m[n]; ok {
			return true
		}
		m[n] = true
	}

	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 3, 4, 5, 6, 7}))
}
