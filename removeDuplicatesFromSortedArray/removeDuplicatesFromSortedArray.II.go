// Source : https://oj.leetcode.com/problems/remove-duplicates-from-sorted-array-ii/
// Author : daoran
// Date   : 2016-04-06

/********************************************************************************** 
* 
* Follow up for "Remove Duplicates":
* What if duplicates are allowed at most twice?
* 
* For example,
* Given sorted array A = [1,1,1,2,2,3],
* 
* Your function should return length = 5, and A is now [1,1,2,2,3].
* 
*               
**********************************************************************************/
package main

import "fmt"


func main() {
    test := []int{1,2,2,2,3,3,3,4,4,5,6}
    fmt.Println("before:", test)
    result := removeDuplicates(test)
    for i := 0; i < result; i++ {
        fmt.Printf("%d ", test[i])
    }
    fmt.Println("after:", test)
}
func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
        return len(nums)
    }

	var index = 2
	for i := 2; i < len(nums); i++ {
		if nums[index-2] != nums[i] {
			nums[index] = nums[i]
            index++
		}
	}

	return index
}