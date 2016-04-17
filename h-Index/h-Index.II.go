// Source : https://leetcode.com/problems/h-index-ii/
// Author : daoran
// Date   : 2016-04-17

/***************************************************************************************
 *
 * Follow up for H-Index: What if the citations array is sorted in ascending order?
 * Could you optimize your algorithm?
 *
 ***************************************************************************************/

package main

import "fmt"

func hIndex(citations []int) int {
	n := len(citations)

	low := 0
	high := n - 1

	for low <= high {
		mid := low + (high-low)/2
		if citations[mid] == n-mid {
			return n - mid
		} else if citations[mid] > n-mid {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return n - low
}

func main() {
	fmt.Println(hIndex([]int{0, 1, 3, 5, 6}))
}
