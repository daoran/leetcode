// Source : https://oj.leetcode.com/problems/longest-substring-without-repeating-characters/
// Author : daoran
// Date   : 2016-04-09

/**********************************************************************************
*
* Given a string, find the length of the longest substring without repeating characters.
* For example, the longest substring without repeating letters for "abcabcbb" is "abc",
* which the length is 3. For "bbbbb" the longest substring is "b", with the length of 1.
*
**********************************************************************************/

package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	exist := make([]bool, 256)

	i := 0
	maxLen := 0
	for j := 0; j < len(s); j++ {
		for exist[s[j]] {
			exist[s[i]] = false
			i++
		}

		exist[s[j]] = true
		maxLen = max(maxLen, j-i+1)
	}

	return maxLen
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcda"))
}
