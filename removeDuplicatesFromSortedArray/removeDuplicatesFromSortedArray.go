// Source : https://oj.leetcode.com/problems/remove-duplicates-from-sorted-array/
// Author : daoran
// Date   : 2016-04-6

/********************************************************************************** 
* 
* Given a sorted array, remove the duplicates in place such that each element appear 
* only once and return the new length.
* 
* Do not allocate extra space for another array, you must do this in place with constant memory.
* 
* For example,
* Given input array A = [1,1,2],
* 
* Your function should return length = 2, and A is now [1,2].
* 
*               
**********************************************************************************/
// Note: when the slice value is passed into the function as an parameter, 
// the value is copied, although we can modified this slice content,
// because it contains a pointer to the underlying array, we cannot modify original 
// slice's header, since we do not have the header's address.
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var index int
	for i := 1; i < len(nums); i++ {
		if nums[index] != nums[i] {
			index++
			nums[index] = nums[i]
		}
	}

	return index + 1
}