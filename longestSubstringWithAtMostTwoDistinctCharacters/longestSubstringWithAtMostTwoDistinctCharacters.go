// Source : https://oj.leetcode.com/problems/longest-substring-with-at-most-two-distinct-characters/
// Author : daoran
// Date   : 2016-04-09

/*
 * Given a string, find the length of the longest substring T that contains at most 2 distinct characters.
 *
 * For example, Given s = “eceba”,
 *
 * T is "ece" which its length is 3.
 */

package main

import "fmt"

func lengthOfLongestSubstringTwoDistinct(s string) int {
	i := 0
	j := -1
	maxLen := 0
	for k := 1; k < len(s); k++ {
		if string(s[k]) == string(s[k-1]) {
			continue
		}

		if j >= 0 && string(s[j]) != string(s[k]) {
			maxLen = max(k-i, maxLen)
			i = j + 1
		}

		j = k - 1
	}
	return max(len(s)-i, maxLen)
}

func lengthOfLongestSubstringTwoDistinct2(s string) int {
	count := make([]int, 256)
	i := 0
	numDistinct := 0
	maxLen := 0
	for j := 0; j < len(s); j++ {
		if count[s[j]] == 0 {
			numDistinct++
		}

		count[s[j]]++
		for numDistinct > 2 {
			count[s[i]]--
			if count[s[i]] == 0 {
				numDistinct--
			}
			i++
		}
		maxLen = max(j-i+1, maxLen)
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
	fmt.Println(lengthOfLongestSubstringTwoDistinct("abaac"))
	fmt.Println(lengthOfLongestSubstringTwoDistinct2("abaac"))
}
