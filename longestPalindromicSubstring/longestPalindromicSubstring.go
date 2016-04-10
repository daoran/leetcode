// Source : https://oj.leetcode.com/problems/longest-palindromic-substring/
// Author : daoran
// Date   : 2016-04-10

/**********************************************************************************
*
* Given a string S, find the longest palindromic substring in S.
* You may assume that the maximum length of S is 1000,
* and there exists one unique longest palindromic substring.
*
**********************************************************************************/

package main

import "fmt"

func longestPalindrome(s string) string {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		len := max(len1, len2)

		if len > end-start {
			start = i - (len-1)/2
			end = i + len/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	l := left
	r := right
	for l >= 0 && r < len(s) && string(s[l]) == string(s[r]) {
		l--
		r++
	}

	return r - l - 1
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

func main() {
	fmt.Println(longestPalindrome("abbacde"))
}
